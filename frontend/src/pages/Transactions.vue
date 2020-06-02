<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col">
        <card title="Transactions" subTitle="Submit a $DAG Transaction">
          <!-- <p>Last Transaction State: {{this.$store.state.txInfo.txStatus}}</p> -->
          <form @submit.prevent class="container">
            <div class="form-row align-items-center">
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
                    v-model.number="txAmountValidation"
                    @change="sendAmount($event.target.value)"
                    pattern="[0-9]+([,\.][0-9]+)?"
                    step="0.01"
                    placeholder="0"
                  />
                  <div class="input-group-append">
                    <button type="button" @click="setMaxDAGs()" class="btn">Max.</button>
                  </div>
                </div>
                <div class="validate" v-if="!$v.txAmountValidation.inBetween">
                  <p>Invalid amount. Please verify.</p>
                </div>
                <div class="validate" v-else></div>
              </div>
              <div class="col-md-1">
                <i class="fa fa-chevron-circle-right icon-point-right"></i>
                <div class="validate"></div>
              </div>
              <div class="col-md-5">
                <input
                  type="text"
                  class="form-control"
                  aria-label="Amount (in DAGs)"
                  v-model.trim="txAddressValidation"
                  @change="setName($event.target.value)"
                  placeholder="Enter Recipients Wallet Address..."
                />
                <div
                  class="validate"
                  v-if="!$v.txAddressValidation.minLength || !$v.txAddressValidation.verifyPrefix || !$v.txAddressValidation.maxLength"
                >
                  <p>Invalid wallet address. Please verify.</p>
                </div>
                <div
                  class="validate"
                  v-else-if="txAddressValidation == this.$store.state.walletInfo.address"
                >
                  <p>You can not send to your own wallet.</p>
                </div>
                <div class="validate" v-else></div>
              </div>
              <div class="col-md-2">
                <p-button
                  type="info"
                  block
                  @click.native="tx"
                  style="max-width: 10rem; margin-left: auto;"
                  :disabled="txInTransit || txAddressValidation == this.$store.state.walletInfo.address"
                >
                  <span>
                    <i class="fa fa-paper-plane"></i> SEND
                  </span>
                </p-button>
                <div class="validate"></div>
              </div>
            </div>
          </form>
        </card>
      </div>
    </div>
    <div class="row">
      <div class="col">
        <card class="card" :title="transactionTable.title" :subTitle="transactionTable.subTitle">
          <div class="table-full-width table-responsive" style="width: 100%;">
            <table class="table" :class="tableClass">
              <thead>
                <slot txAddressValidation="columns">
                  <th v-for="column in transactionTable.columns" v-bind:key="column.id">{{column}}</th>
                </slot>
              </thead>
              <tbody>
                <tr v-for="tx in paginatedData" v-bind:key="tx.ID">
                  <slot :row="item">
                    <td class="columnA">
                      <i
                        style="color: #6DECBB;"
                        v-if="tx.status === 'Complete'"
                        class="fa fa-check"
                      ></i>
                      <center><spinner :size="15" color="#F9EC31" v-if="tx.status === 'Pending'"></spinner></center>
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
            <ul
              v-if="this.transactionTable.data.length > 0"
              class="pagination justify-content-center"
            >
              <li class="page-item" :class="pageNumber == 0 ? 'disabled' : ''">
                <a class="page-link" style="cursor: pointer;" @click="prevPage">Previous</a>
              </li>
              <li
                class="page-item"
                :class="page == pageNumber + 1 ? 'active' : ''"
                v-for="page in pageCount"
                :key="page"
              >
                <a class="page-link" style="cursor: pointer;" @click="gotoPage(page)">{{page}}</a>
              </li>
              <li class="page-item" :class="pageNumber >= pageCount - 1 ? 'disabled' : ''">
                <a class="page-link" style="cursor: pointer;" @click="nextPage">Next</a>
              </li>
            </ul>
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

import Spinner from 'vue-spinner-component/src/Spinner.vue';
import {
  required,
  minLength,
  maxLength,
  between
} from "vuelidate/lib/validators";
import Swal from "sweetalert2/dist/sweetalert2";

export default {
  components: {
    Spinner
  },
  computed: {
    tableClass() {
      return `table-${this.type}`;
    },
    txInTransit() {
      return this.$store.state.txInfo.txStatus == "Pending";
    },
    pageCount() {
      let l = this.$store.state.txInfo.txHistory.length,
        s = this.size;
      return Math.ceil(l / s);
    },
    paginatedData() {
      const start = this.pageNumber * this.size,
        end = start + this.size;
      return this.$store.state.txInfo.txHistory.slice(start, end);
    }
  },
  methods: {
    isFloat: function(n) {
      return n === +n && n !== (n | 0);
    },
    isInteger: function(n) {
      return n === +n && n === (n | 0);
    },
    sendAmount(value) {
      this.txAmountValidation = value;
      this.$v.txAmountValidation.$touch();
    },
    setName(value) {
      this.txAddressValidation = value;
      this.$v.txAddressValidation.$touch();
    },
    tx: function() {
      var self = this;
      self.$v.$touch();
      if (self.$v.$invalid) {
        self.submitStatus = "ERROR";
      } else {
        // do your submit logic here

        if (!self.$store.state.app.txFinished) {
          self.submitStatus = "PENDING";
        }
        self.submitStatus = "OK";
      }

      if (self.submitStatus === "OK") {
        self.$swal.mixin({
          progressSteps: ["1", "2"],
          customClass: {
            container: this.$store.state.walletInfo.darkMode
              ? "theme--dark"
              : "theme--light"
          }
        })
          .queue([
            {
              title: "Are you sure?",
              html:
                "You are about to send <b>" +
                self.txAmountValidation +
                "</b> $DAG tokens to " +
                self.txAddressValidation,
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
              let amount = self.txAmountValidation;
              let address = self.txAddressValidation;
              let fee = result.value;
              const swalPopup = Swal.mixin({
                customClass: {
                  container: this.$store.state.walletInfo.darkMode
                    ? "theme--dark"
                    : "theme--light"
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
                      self.txAmountValidation +
                      " $DAG tokens to address " +
                      self.txAddressValidation +
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
      this.txAmountValidation = this.$store.state.walletInfo.availableBalance;
    },
    nextPage() {
      this.pageNumber++;
    },
    prevPage() {
      this.pageNumber--;
    },
    gotoPage(page) {
      this.pageNumber = page - 1;
    }
  },
  data() {
    return {
      txAddressValidation: "",
      txAmountValidation: null,
      submitStatus: null,
      amountSubmitted: null,
      txAmount: "",
      txAddress: "",
      notifications: {
        topCenter: false
      },
      overlay: false,
      transactionTable: {
        title: "Transaction History",
        subTitle: "Table containing all previous transactions",
        columns: [...tableColumns],
        data: this.$store.state.txInfo.txHistory
      },
      pageNumber: 0,
      size: 10
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
    txAddressValidation: {
      required,
      minLength: minLength(40),
      maxLength: maxLength(40),
      verifyPrefix
    },
    txAmountValidation: {
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
    }
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
