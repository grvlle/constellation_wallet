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
                    v-model.number="txAmount.normalized"
                    @change="setTxAmount($event.target.value)"
                    @keypress="setTxAmount($event.target.value)"
                    pattern="[0-9]+([,\.][0-9]+)?"
                    step="0.01"
                  />
                  <div class="input-group-append">
                    <button type="button" @click="setMaxDAGs()" class="btn">Max.</button>
                  </div>
                </div>
                <div class="validate text-danger" v-if="$v.txAmount.normalized.$invalid && txAmount.normalized != 0">
                  <p>Invalid amount. Please verify.</p>
                </div>
                <div class="validate text-danger" v-else></div>
              </div>
              <div class="col-md-1">
                <i class="fa fa-chevron-circle-right icon-point-right"></i>
                <div class="validate"></div>
              </div>
              <div class="col-md-5" style="padding-right: 0">
                <div class="input-group" style="margin-bottom: 0;">
                  <input
                    type="text"
                    style="border: 1px solid #ced4da;"
                    class="form-control"
                    aria-label="Amount (in DAGs)"
                    v-model.trim="txAddress"
                    @change="setName($event.target.value)"
                    @keypress="setName($event.target.value)"
                    placeholder="Enter Recipients Wallet Address..."
                  />
<!--                  <div class="input-group-append">-->
<!--                    <button type="button" class="btn" @click="toggleAddressBook">-->
<!--                      <i class="fa fa-address-book"></i>-->
<!--                      Address book-->
<!--                    </button>-->
<!--                  </div>-->
                </div>
                <div class="validate text-danger" v-if="$v.txAddress.$invalid && txAddress != ''">
                  <p>Invalid wallet address. Please verify.</p>
                </div>
                <div class="validate text-danger" v-else-if="txAddress == address">
                  <p>You can not send to your own wallet address.</p>
                </div>
                <div class="validate text-success" v-else-if="txAddress != ''">
                  <p v-if="txAddressInformation">{{txAddressInformation}}</p>
<!--                  <p-->
<!--                    v-else class="text-muted"-->
<!--                  >This DAG address is not stored in any of your address book contacts.</p>-->
                </div>
                <div class="validate text-danger" v-else></div>
              </div>
              <div class="col-md-2" style="padding-left: 0; margin-left: 0">
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

import { mapState, mapGetters } from "vuex";
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
import { keyStore } from "@stardust-collective/dag-keystore";
import { dagWalletAccount } from "@stardust-collective/dag-wallet-sdk";

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
    ...mapGetters('wallet', ['normalizedAvailableBalance']),
    ...mapState("transaction", ["txHistory", "txStatus", "txFinished"])
  },
  methods: {
    setTxAmount(value) {
      this.txAmount.normalized = value;
      // this.txAmount.denormalized = this.txAmount.normalized;
      this.$v.txAmount.normalized.$touch();
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
              "<p>You are about to send <b>" + self.txAmount.normalized + "</b> $DAG tokens to " + self.txAddress + "</p>" +
              "<div class='border border-dark'>" + 
                "<p class='mt-3 font-weight-light text-muted font-italic'>(Available balance: " + self.normalizedAvailableBalance + ")</p>" +              
                "<p>Transaction amount: " + self.txAmount.normalized + "</p>" +
                "<p>To: " + self.txAddress + "</p>" +
              "</div>",
            showCancelButton: true,
            confirmButtonColor: "#5FD1FB",
            confirmButtonText: "Yes, please proceed!"
          },
          {
            title: "Set a fee to prioritize your transaction.",
            html: 
              "<div class='border border-dark'>" + 
                "<p class='mt-3 font-weight-light text-muted font-italic'>(Available balance: " + self.normalizedAvailableBalance + ")</p>" +              
                "<p>Transaction amount: " + self.txAmount.normalized + "</p>" +
                "<p>To: " + self.txAddress + "</p>" +
              "</div>" + 
              "<div>" + 
                "</br>" +                 
                "This fee is <b>optional</b>, enter 0 for no fee." +
            "</div>",
            input: "number",
            inputValue: 0,
            confirmButtonText: "Send transaction",
            confirmButtonColor: "#6DECBB",
            showCancelButton: true,
            animation: false,
            inputValidator: value => {
              return new Promise(resolve => {
                if (value + self.txAmount.normalized > self.normalizedAvailableBalance) {
                  resolve("The transaction amount + fee can not exceed your balance");
                } else if (value < 0 ||  value > 3711998690 || isNaN(parseFloat(value))) {
                  resolve("Please enter a transaction fee between 0 and 3711998690");
                }
                else {
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
            let feeResult = result.value;
            const swalPopup = Swal.mixin({
              customClass: {
                container: this.darkMode ? "theme--dark" : "theme--light"
              }
            });

            const amount = parseFloat(self.txAmount.normalized, 10);
            const fee = parseFloat(feeResult[1], 10);

            window.backend.WalletApplication.GetLastAcceptedTransactionRef().then(result => {
              if (!result) {
                self.txFailure(swalPopup);
              }
              else {
                const tokens = result.split(',');
                const lastRef = {ordinal: parseInt(tokens[0]), prevHash: tokens[1]}
                keyStore.generateTransaction(amount, self.txAddress, dagWalletAccount.keyTrio, lastRef, fee).then(tx => {
                  window.backend.WalletApplication.SendTransaction2(
                      JSON.stringify(tx)
                  ).then(success => {
                    if (success) {
                      self.txSuccess(swalPopup);
                    }
                    else {
                      self.txFailure(swalPopup);
                    }
                  });
                }, () => {
                  self.txFailure(swalPopup);
                })
              }
            })
          }
        });
      }
    },
    txSuccess (swalPopup) {
      this.$Progress.finish();
      this.overlay = false;
      this.txAddress = "";
      this.txAmount.normalized = 0

      swalPopup.fire({
        title: "Success!",
        text:
            "You have sent " +
            self.txAmount.normalized +
            " $DAG tokens to address " +
            self.txAddress +
            ".",
        icon: "success"
      });

    },
    txFailure (swalPopup) {
      this.$Progress.finish();
      this.overlay = false;
      this.txAddress = "";
      this.txAmount.normalized = 0

      swalPopup.fire({
        title: "Transaction Failed!",
        text: "Unable to send Transaction",
        type: "error"
      });
    },
    setMaxDAGs() {
      this.txAmount.normalized = this.normalizedAvailableBalance;
      // this.txAmount.denormalized = this.availableBalance;
    }
  },
  data() {
    return {
      txAddress: "",
      txAmount: {
        normalized: 0
      },
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
  validations: {
    txAddress: {
      required,
      minLength: minLength(40),
      maxLength: maxLength(40),
      verifyPrefix
    },
    txAmount: {
      normalized: {
        required,
        inBetween: between(0.00000001, 3711998690)
      }
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
