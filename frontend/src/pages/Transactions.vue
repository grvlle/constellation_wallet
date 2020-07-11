<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col">
        <card title="Transactions" subTitle="Submit a $DAG Transaction">
          <div class="container">
            <div class="row align-items-center">
              <div class="col-md-4">
                <div class="input-group" style="margin-bottom: 0;">
                  <div class="input-group-prepend">
                    <span class="input-group-text">
                      <small>DAG</small>
                    </span>
                  </div>
                  <input
                    type="text"
                    class="form-control"
                    aria-label="Amount (in DAGs)"
                    v-model.number="txAmount"
                    @change="sendAmount($event.target.value)"
                    pattern="[0-9]+([,\.][0-9]+)?"
                    step="0.01"
                    placeholder="0"
                  />
                  <div class="input-group-append">
                    <button type="button" @click="setMaxDAGs()" class="btn">Max.</button>
                  </div>
                </div>
                <div class="validate" v-if="!$v.txAmount.inBetween">
                  <p>Invalid amount. Please verify.</p>
                </div>
                <div class="validate" v-else></div>
              </div>
              <div class="col-md-1">
                <i class="fa fa-chevron-circle-right icon-point-right"></i>
                <div class="validate"></div>
              </div>
              <div class="col-md-5">
                <div class="input-group" style="margin-bottom: 0;">
                  <input
                    type="text"
                    class="form-control"
                    aria-label="Amount (in DAGs)"
                    v-model.trim="txAddress"
                    @change="setName($event.target.value)"
                    placeholder="Enter Recipients Wallet Address..."
                  />
                  <div class="input-group-append">
                    <button type="button" class="btn" @click="toggleAddressBook">
                      <i class="fa fa-address-book"></i>
                      Address book
                    </button>
                  </div>
                </div>
                <div
                  class="validate"
                  v-if="!$v.txAddress.minLength || !$v.txAddress.verifyPrefix || !$v.txAddress.maxLength"
                >
                  <p>Invalid wallet address. Please verify.</p>
                </div>
                <div class="validate" v-else-if="txAddress == address">
                  <p>You can not send to your own wallet.</p>
                </div>
                <div class="validate" v-else></div>
              </div>
              <div class="col-md-2">
                <p-button
                  type="info"
                  block
                  @click.native="submitTransaction"
                  style="max-width: 10rem; margin-left: auto;"
                  :disabled="txInTransit || txAddress == address"
                >
                  <span>
                    <i class="fa fa-paper-plane"></i> SEND
                  </span>
                </p-button>
                <div class="validate"></div>
              </div>
            </div>
            <div class="row" v-if="showAddressBook">
              <div class="col-md-5" />
              <div class="col-md-5">
                <address-book-search v-model="txAddress"/>
              </div>
            </div>
          </div>
        </card>
      </div>
    </div>
    <div class="row">
      <div class="col">
        <card class="card" :title="transactionTable.title" :subTitle="transactionTable.subTitle">
          <div class="table-full-width table-responsive" style="width: 100%;">
            <table class="table" :class="tableClass">
              <thead>
                <slot txAddress="columns">
                  <th v-for="column in transactionTable.columns" v-bind:key="column.id">{{column}}</th>
                </slot>
              </thead>
              <tbody>
                <tr v-for="tx in txHistoryPage" v-bind:key="tx.ID">
                  <slot :row="tx">
                    <td class="columnA">
                      <i
                        style="color: #6DECBB;"
                        v-if="tx.status === 'Complete'"
                        class="fa fa-check"
                      ></i>
                      <center>
                        <spinner :size="15" color="#F9EC31" v-if="tx.status === 'Pending'"></spinner>
                      </center>
                      <i style="color: firebrick;" v-if="tx.status === 'Error'" class="fa fa-times"></i>
                    </td>
                    <td class="columnB">
                      <p class="description" style="font-size: 0.9375rem;">
                        <b>{{(tx.amount / 1e8).toFixed(8).replace(/\.?0+$/,"")}}</b> DAG
                      </p>
                    </td>
                    <td class="columnC">
                      <p class="description" style="font-size: 0.9375rem;">{{tx.receiver}}</p>
                    </td>
                    <td class="columnD">
                      <p class="description" style="font-size: 0.9375rem;">{{tx.fee / 1e8}}</p>
                    </td>
                    <td class="columnE">
                      <a id="txhash">
                        <p style="font-size: 0.9375rem;">{{tx.hash}}</p>
                      </a>
                    </td>
                    <td class="columnF">
                      <p class="description" style="font-size: 0.9375rem;">{{tx.date}}</p>
                    </td>
                  </slot>
                </tr>
              </tbody>
            </table>
            <pagination :dataset="txHistory" :pageSize=10 v-model="txHistoryPage" />
          </div>
        </card>
      </div>
    </div>
    <page-overlay text="Submitting Transaction..." :isActive="overlay" />
  </div>
</template>

<script>
const tableColumns = ["Status", "Amount", "Receiver", "Fee", "Hash", "Date"];
const verifyPrefix = value =>
  value.substring(0, 3) === "DAG" || value.substring(0, 3) === "";

import { mapState } from "vuex";
import Spinner from "vue-spinner-component/src/Spinner.vue";
import {
  required,
  minLength,
  maxLength,
  between
} from "vuelidate/lib/validators";
import Swal from "sweetalert2/dist/sweetalert2";
import AddressBookSearch from "../components/AddressBookSearch";
import Pagination from "../components/Pagination";

export default {
  components: {
    Spinner,
    AddressBookSearch,
    Pagination
  },
  created: function() {
    if (this.txAddressParam != "") {
      this.txAddress = this.txAddressParam;
    }
  },
  computed: {
    tableClass() {
      return `table-${this.type}`;
    },
    txInTransit() {
      return this.txStatus == "Pending";
    },
    ...mapState("wallet", ["address", "availableBalance", "darkMode"]),
    ...mapState("transaction", ["txHistory", "txStatus", "txFinished"])
  },
  methods: {
    isFloat: function(n) {
      return n === +n && n !== (n | 0);
    },
    isInteger: function(n) {
      return n === +n && n === (n | 0);
    },
    sendAmount(value) {
      this.txAmount = value;
      this.$v.txAmount.$touch();
    },
    setName(value) {
      this.txAddress = value;
      this.$v.txAddress.$touch();
    },
    toggleAddressBook() {
      this.showAddressBook = !this.showAddressBook;
    },
    submitTransaction: function() {
      var self = this;
      self.$v.$touch();
      if (self.$v.$invalid) {
        self.submitStatus = "ERROR";
      } else {
        // do your submit logic here

        if (!self.txFinished) {
          self.submitStatus = "PENDING";
        }
        self.submitStatus = "OK";
      }

      if (self.submitStatus === "OK") {
        Swal.mixin({
          progressSteps: ["1", "2"],
          customClass: {
            container: this.darkMode ? "theme--dark" : "theme--light"
          }
        })
          .queue([
            {
              title: "Are you sure?",
              html:
                "You are about to send <b>" +
                self.txAmount +
                "</b> $DAG tokens to " +
                self.txAddress,
              type: "warning",
              showCancelButton: true,
              confirmButtonColor: "#5FD1FB",
              confirmButtonText: "Yes, please proceed!"
            },
            {
              title: "Set a fee to prioritize your transaction.",
              html: "This is <b>optional</b>, enter 0 for no fee.",
              input: "number",
              inputValue: 0,
              confirmButtonText: "Send transaction",
              confirmButtonColor: "#6DECBB",
              showCancelButton: true,
              inputValidator: value => {
                return new Promise(resolve => {
                  if (
                    value < 0 ||
                    value > 3711998690 ||
                    isNaN(parseFloat(value))
                  ) {
                    resolve("Please enter a value between 0 and 3711998690");
                  } else {
                    resolve();
                  }
                });
              }
            }
          ])
          .then(result => {
            if (result.value) {
              self.$Progress.start();
              self.overlay = true;
              let amount = self.txAmount;
              let address = self.txAddress;
              let fee = result.value;
              const swalPopup = Swal.mixin({
                customClass: {
                  container: this.darkMode ? "theme--dark" : "theme--light"
                }
              });
              window.backend.WalletApplication.TriggerTXFromFE(
                parseFloat(amount, 10),
                parseFloat(fee[1], 10),
                address
              ).then(txFailed => {
                if (txFailed) {
                  swalPopup.fire({
                    title: "Transaction Failed!",
                    text: "Unable to send Transaction",
                    type: "error"
                  });
                  self.$Progress.fail();
                  self.overlay = false;
                }
                if (!txFailed) {
                  swalPopup.fire({
                    title: "Success!",
                    text:
                      "You have sent " +
                      self.txAmount +
                      " $DAG tokens to address " +
                      self.txAddress +
                      ".",
                    type: "success"
                  });
                  self.$Progress.finish();
                  self.overlay = false;
                }
              });
            }
          });
      }
    },
    setMaxDAGs() {
      this.txAmount = this.availableBalance;
    }
  },
  data() {
    return {
      txAddress: "",
      txAmount: null,
      submitStatus: null,
      amountSubmitted: null,
      notifications: {
        topCenter: false
      },
      overlay: false,
      transactionTable: {
        title: "Transaction History",
        subTitle: "Table containing all previous transactions",
        columns: [...tableColumns],
      },
      txHistoryPage: [],
      showAddressBook: false
    };
  },
  filters: {
    dropzero: function(value) {
      if (!value || value === null) return "";
      let index;
      value = value.toString().split("");
      index = value.length - 8;
      value = value.splice(0, index);
      return value.join("");
    },
    truncate: function(value) {
      if (!value || value === null) return "";
      if (value.length > 30) {
        value = value.substring(0, 27) + "...";
      }
      return value;
    }
  },
  validations: {
    txAddress: {
      required,
      minLength: minLength(40),
      maxLength: maxLength(40),
      verifyPrefix
    },
    txAmount: {
      required,
      inBetween: between(0.00000001, 3711998690)
    }
  },
  props: {
    columns: Array,
    data: Array,
    type: {
      type: String, // striped | hover
      default: "striped"
    },
    title: {
      type: String,
      default: ""
    },
    subTitle: {
      type: String,
      default: ""
    },
    txAddressParam: {
      type: String,
      default: ""
    },
  }
};
</script>

<style scoped>
td {
  max-width: 0;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
td.columnA {
  width: 3%;
  padding-top: 0px;
  text-align: center;
}
td.columnB {
  width: 15%;
}
td.columnC {
  width: 40%;
}
td.columnD {
  width: 5%;
}
td.columnE {
  width: 22%;
}
td.columnF {
  width: 15%;
}

txhash a {
  color: blue;
}

txhash a:visited {
  color: blue;
}

txhash p {
  font-weight: bold;
}

.validate {
  height: 1.25em;
}
</style>
