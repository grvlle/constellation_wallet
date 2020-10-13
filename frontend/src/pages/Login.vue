<template>
  <div class="container">
    <div class="row">
      <div class="col-12">
        <form ref="textareaform" @submit.prevent="form" class="container">
          <div class="row">
            <div class="col mx-auto login-box">
              <div class="input-box">
                <div>
                  <password-input
                    v-model="keystorePassword"
                    label="Password"
                    :validate="false"
                  />
                </div>
              </div>
              <div class="button-box">
                <div class="container">
                  <div class="row">
                    <div class="col">
                      <p-button
                        type="primary"
                        block
                        @click.native="loginPass()"
                      >
                        <span style="display: block"> LOGIN</span>
                      </p-button>
                    </div>
                  </div>
                  <div class="row">
                    <div class="col">
                      <p class="text-right">
                        Don't have a Wallet V2 yet? Create one
                        <a class="link-text" @click="createWallet()">here!</a>
                      </p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </form>
      </div>
    </div>
    <page-overlay text="Loading..." :isActive="overlay" />
  </div>
</template>

<script>
import { mapState } from "vuex";
import Swal from "sweetalert2/dist/sweetalert2";
import { keyStore } from "@stardust-collective/dag-keystore";

export default {
  name: "login-screen",
  data: () => ({
    keystorePassword: "",
    KeyPassword: "",
    overlay: false,
  }),
  computed: {
    ...mapState("app", ["network"]),
    keystorePath: {
      get() {
        return this.$store.state.wallet.keystorePath;
      },
      set(value) {
        this.$store.commit("wallet/setKeystorePath", value);
      },
    },
    alias: {
      get() {
        return this.$store.state.wallet.alias;
      },
      set(value) {
        this.$store.commit("wallet/setAlias", value);
      },
    },
  },
  mounted() {
    this.migrateNotification();
  },
  methods: {
    migrateNotification: function() {
      let timerInterval;
      Swal.fire({
        title:
          "<p style='text-align: left; color: white; margin: auto;'>Important Update</p>",
        html: `<br><p style='text-align: left; color: white;'>If you have previously signed into Molly Wallet using a file, two passwords and alias (versions 1.x.x), then you will need to create a new Wallet 2.0 first. <br><br> You will be able import your key files later.</p>`,
        width: 300,
        padding: 20,
        backdrop: false,
        toast: true,
        background: "#2654C0",
        position: "top-end",
        showConfirmButton: false,
        allowOutsideClick: false,
        showCloseButton: true,
        timerProgressBar: true,
        willOpen: () => {
          Swal.showLoading();
          timerInterval = setInterval(() => {
            const content = Swal.getContent();
            if (content) {
              const b = content.querySelector("b");
              if (b) {
                b.textContent = Swal.getTimerLeft();
              }
            }
          }, 100);
        },
        onClose: () => {
          clearInterval(timerInterval);
        },
      });
    },
    setNetwork: function(value) {
      window.backend.WalletApplication.SelectNetwork(value).then((result) => {
        if (result) {
          this.$store.commit("app/setNetwork", value);
        }
      });
    },
    createWallet: function() {
      Swal.close();
      this.$store.dispatch("wallet/reset").then(() => {
        this.$router.push({
          name: "create wallet",
          params: {
            message: "Please enter your new password below, then choose whether to create a new DAG account or import an existing one.",
            title: "Create Molly Wallet password",
            darkMode: this.$route.params.darkMode,
          },
        });
      });
    },
    loginPass: function() {
      var self = this;
      self.$Progress.start();
      self.overlay = true;
      window.backend.WalletApplication.LoginKeychain(
        self.keystorePassword
      ).then((pkey) => {
        // eslint-disable-next-line no-console
        console.log("pkey: " + pkey);
        if (pkey && pkey.length === 64) {
          var address = keyStore.getDagAddressFromPublicKey(keyStore.getPublicKeyFromPrivate(pkey));
          // eslint-disable-next-line no-console
          console.log("getDagAddressFromPublicKey: " + address);
          window.backend.WalletApplication.CreateOrInitWalletV2(
              address
          ).then((result) => {
            if (result) {
              Swal.close();
              self.initWallet();
            }
          })

          // this.$store.dispatch("wallet/reset").then(() => {
          //   this.$router.push({
          //     name: "home",
          //   });
          // });
        }
      });
    },
    login: function() {
      var self = this;
      self.$Progress.start();
      self.overlay = true;
      window.backend.WalletApplication.Login(
        self.keystorePath,
        self.keystorePassword,
        self.KeyPassword,
        self.alias
      ).then((result) => {
        if (result) {
          self.initWallet();
        } else {
          self.overlay = false;
          self.$Progress.fail();
        }
      });
    },
    initWallet: function () {
      var self = this;
      window.backend.WalletApplication.GetUserTheme().then((darkMode) =>
          self.$store.commit("wallet/setDarkMode", darkMode)
      );
      window.backend.WalletApplication.GetWalletTag().then((walletTag) =>
          self.$store.commit("wallet/setLabel", walletTag)
      );
      window.backend.WalletApplication.GetImagePath().then((imagePath) =>
          self.$store.commit("wallet/setImgPath", imagePath)
      );
      self.overlay = false;
      self.$Progress.finish();
      self.$store.commit("app/setIsLoggedIn", true);

      window.backend.WalletApplication.CheckTermsOfService().then(
          (result) => {
            self.$store.commit("wallet/setTermsOfService", result);
            if (result) {
              self.$router.push({
                name: "loading",
                params: { message: "Getting your $DAG Wallet ready..." },
              });
            } else {
              self.$router.push({
                name: "accept terms of service",
                params: { message: "Terms of Service" },
              });
            }
          }
      );
    }
  },
};
</script>

<style scoped lang="scss">
.login-box {
  max-width: 29rem;
  min-width: 29rem;
  padding-bottom: 2rem;
  margin-top: 5.25em;
}

.link-text {
  color: #34b4e7;
}

.link-text:hover {
  color: #ce9483;
  cursor: pointer;
}

.input-box > div {
  margin-bottom: 1.875em;
}

.button-box .container {
  margin-left: 0em;
  margin-right: 0em;
  padding-left: 0em;
  padding-right: 0em;
}

.button-box .container .row {
  margin-left: 0em;
  margin-right: 0em;
  padding-left: 0em;
  padding-right: 0em;
  margin-top: 1.25em;
}

.button-box .container .row [class^="col"] {
  margin-left: 0em;
  margin-right: 0em;
  padding-left: 0em;
  padding-right: 0em;
}
</style>
