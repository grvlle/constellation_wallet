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
                  />
                  <div class="input-group-append">
                    <button type="button" @click="setMaxDAGs()" class="btn">Max.</button>
                  </div>
                </div>
                <div class="validate text-danger" v-if="$v.txAmount.$invalid && txAmount != 0">
                  <p>Invalid amount. Please verify.</p>
                </div>
                <div class="validate text-danger" v-else></div>
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
                    @keypress="setName($event.target.value)"
                    placeholder="Enter Recipients Wallet Address..."
                  />
                  <div class="input-group-append">
                    <button type="button" class="btn" @click="toggleAddressBook">
                      <i class="fa fa-address-book"></i>
                      Address book
                    </button>
                  </div>
                </div>
                <div class="validate text-danger" v-if="$v.txAddress.$invalid && txAddress != ''">
                  <p>Invalid wallet address. Please verify.</p>
                </div>
                <div class="validate text-danger" v-else-if="txAddress == address">
                  <p>You can not send to your own wallet address.</p>
                </div>
                <div class="validate text-success" v-else-if="txAddress != ''">
                  <p v-if="txAddressInformation">{{txAddressInformation}}</p>
                  <p
                    v-else class="text-muted"
                  >This DAG address is not stored in any of your address book contacts.</p>
                </div>
                <div class="validate text-danger" v-else></div>
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
          </div>
        </card>
      </div>
    </div>
    <div class="row" v-if="showAddressBook">
      <div class="col-md-5" />
      <div class="col-md-5">
        <address-book-search v-model="txAddress" v-on:input="showAddressBook = false"/>
      </div>
    </div>
    <div class="row">
      <div class="col">
        <card class="card" title="Transaction History">
          <div>
            <timeline v-model="txHistoryPage" />
            <pagination :dataset="txHistory" :pageSize="10" v-model="txHistoryPage" />
          </div>
        </card>
      </div>
    </div>
    <page-overlay text="Submitting Transaction..." :isActive="overlay" />
  </div>
</template>

<script>
const verifyPrefix = value =>
  value.substring(0, 3) === "DAG" || value.substring(0, 3) === "";

import { mapState } from "vuex";
import {
  required,
  minLength,
  maxLength,
  between
} from "vuelidate/lib/validators";
import Swal from "sweetalert2/dist/sweetalert2";
import AddressBookSearch from "../components/AddressBookSearch";
import Pagination from "../components/Pagination";
import Timeline from "../components/Timeline";

export default {
  components: {
    AddressBookSearch,
    Pagination,
    Timeline
  },
  created: function() {
    if (this.txAddressParam != "") {
      this.txAddress = this.txAddressParam;
    }

    // TODO: refactor into wallet-init instead calling from indivual pages
    window.backend.WalletApplication.GetAddressBook().then(ab => {
      let addressBook;
      try {
        addressBook = JSON.parse(ab);
      } catch (e) {
        addressBook = [];
      }
      this.$store.commit({ type: "addressBook/setAddressBook", addressBook });
    });    
  },
  computed: {
    tableClass: function() {
      return `table-${this.type}`;
    },
    txInTransit: function() {
      return this.txStatus == "Pending";
    },
    txAddressInformation: function() {
      let addressInfo = "";
      if (!this.$v.txAddress.$invalid) {        
        let contact = this.$store.getters["addressBook/search"](this.txAddress);
        if (contact.length) {
          addressInfo =
            "This DAG address belongs to your address book contact with the name " +
            '"' +
            contact[0].name +
            '".';
        }
      }
      return addressInfo;
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
      txAmount: 0,
      submitStatus: null,
      amountSubmitted: null,
      notifications: {
        topCenter: false
      },
      overlay: false,
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
    txAddressParam: {
      type: String,
      default: ""
    }
  }
};
</script>

<style scoped lang="scss">
.validate {
  height: 1.25em;
}

.icon-point-right {
  @include themed() {
    color: t('successColor');
  }
  font-size: 2.5rem; 
  width:100%;
}
</style>
