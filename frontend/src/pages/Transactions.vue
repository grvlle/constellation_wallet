<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col">
        <card :title="table1.title" :subTitle="table1.subTitle">
          <form @submit.prevent class="container">
            <div class="form-row align-items-center">
              <div class="col-md-5">
                <fg-input
                  style="margin-bottom: 0.125em;"
                  v-model.number="txAmountValidation"
                  @change="sendAmount($event.target.value)"
                  pattern="[0-9]+([,\.][0-9]+)?"
                  step="0.01"
                  label="Amount: "
                  placeholder="0"
                ></fg-input>
                <div class="validate" v-if="!$v.txAmountValidation.inBetween">
                  <p>Invalid amount. Please verify.</p>
                </div>
                <div class="validate" v-else></div>
              </div>
              <div class="col-md-2">
                <i
                  class="fa fa-chevron-circle-right"
                  style="color: #6DECBB; font-size: 2.5rem; width:100%; margin-top: 1.5rem;"
                ></i>
                <div class="validate"></div>
              </div>
              <div class="col-md-5">
                <fg-input
                  style="margin-bottom: 0.125em;"
                  v-model.trim="txAddressValidation"
                  @change="setName($event.target.value)"
                  type="text"
                  label="Address: "
                  placeholder="Enter Recipients Wallet Address"
                ></fg-input>
                <div
                  class="validate"
                  v-if="!$v.txAddressValidation.minLength || !$v.txAddressValidation.verifyPrefix || !$v.txAddressValidation.maxLength"
                >
                  <p>Invalid wallet address. Please verify.</p>
                </div>
                <div class="validate" v-else></div>
              </div>
            </div>
            <div class="form-row">
              <div class="col-md-9" />
              <div class="col-md-3">
                <p-button
                  type="info"
                  block
                  @click.native="tx"
                  :disabled="!this.$store.state.app.txFinished"
                >
                  <span>
                    <i class="fa fa-paper-plane"></i> SEND
                  </span>
                </p-button>
              </div>
            </div>
          </form>
        </card>
      </div>
    </div>
    <div class="row">
      <div class="col">
        <card class="card" :title="table2.title" :subTitle="table2.subTitle">
          <div class="table-full-width table-responsive">
            <table class="table" :class="tableClass">
              <thead>
                <slot txAddressValidation="columns">
                  <th v-for="column in table2.columns" v-bind:key="column.id">{{column}}</th>
                </slot>
              </thead>
              <tbody>
                <tr v-for="tx in this.$store.state.txInfo.txHistory" v-bind:key="tx.ID">
                  <slot :row="item">
                    <td class="columnA">
                      <p class="description" style="font-size: 0.9375rem;">
                        <b>{{tx.amount / 1e8}}</b> DAG
                      </p>
                    </td>
                    <td class="columnB">
                      <p class="description" style="font-size: 0.9375rem;">{{tx.sender}}</p>
                    </td>
                    <td class="columnC">
                      <p class="description" style="font-size: 0.9375rem;">{{tx.fee / 1e8}}</p>
                    </td>
                    <td class="columnD">
                      <a id="txhash">
                        <p style="font-size: 0.9375rem;">{{tx.hash}}</p>
                      </a>
                    </td>
                    <td class="columnE">
                      <p class="description" style="font-size: 0.9375rem;">{{tx.date}}</p>
                    </td>
                  </slot>
                </tr>
              </tbody>
            </table>
            <center>
              <jw-pagination :items="table2.data" @changePage="onChangePage"></jw-pagination>
            </center>
          </div>
        </card>
      </div>
    </div>
  </div>
</template>

<script>
const tableColumns = ["Amount", "Sender", "Fee", "Hash", "Date"];
let tableData = [];
const verifyPrefix = value =>
  value.substring(0, 3) === "DAG" || value.substring(0, 3) === "";

import Swal from "sweetalert2";
import {
  required,
  minLength,
  maxLength,
  between
} from "vuelidate/lib/validators";

export default {
  computed: {
    tableClass() {
      return `table-${this.type}`;
    }
  },
  methods: {
    isFloat: function(n) {
      return n === +n && n !== (n | 0);
    },

    isInteger: function(n) {
      return n === +n && n === (n | 0);
    },

    onChangePage(pageOfItems) {
      // update page of items
      this.$store.state.pageOfItems = pageOfItems;
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
        Swal.mixin({
          progressSteps: ["1", "2"]
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
              let amount = self.txAmountValidation;
              let address = self.txAddressValidation;
              let fee = result.value;
              window.backend.WalletApplication.TriggerTXFromFE(
                parseFloat(amount),
                parseFloat(fee[1]),
                address
              ).then(txFailed => {
                if (txFailed) {
                  Swal.fire({
                    title: "Transaction Failed!",
                    text: "Unable to send Transaction",
                    type: "error"
                  });
                  self.$Progress.fail();
                }
                if (!txFailed) {
                  Swal.fire({
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
                }
              });
            }
          });
      }
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

      table1: {
        title: "Transactions",
        subTitle: "Submit a $DAG Transaction",
        columns: [...tableColumns],
        data: [...tableData]
      },
      table2: {
        title: "Transaction History",
        subTitle: "Table containing all previous transactions",
        columns: [...tableColumns],
        data: this.$store.state.txInfo.txHistory
      }
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
    width: 15%;
}
td.columnB {
    width: 40%;
}
td.columnC {
    width: 10%;
}
td.columnD {
    width: 15%;
}
td.columnE {
    width: 20%;
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
  display: flex;
}
.validate > p {
  /*flex: 1;*/
  font-size: 0.625rem;
  color: firebrick;
  margin-top: 0em;
  margin-right: 0.125em;
}
</style>
