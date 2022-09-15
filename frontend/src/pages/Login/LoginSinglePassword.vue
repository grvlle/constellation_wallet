<template>
  <div class="container">
    <div class="row">
      <div class="col-12">
        <form ref="textareaform" @submit.prevent="form" class="container">
          <div class="row">
            <div class="col mx-auto login-box">
              <div>
                <label class="control-label"
                  >Select your JSON Private Key</label
                >
                <file-selector
                  v-model="keystorePath"
                  action="SelectFile3"
                  @file="fileSelected"
                />
              </div>

              <div class="input-box">
                <div>
                  <password-input
                    v-model="keystorePassword"
                    label="Password"
                    :validate="false"
                    @input="validCheck()"
                    @enter="loginJsonWallet(keystorePath, keystorePassword)"
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
                        :disabled="!valid"
                        @click.native="
                          loginJsonWallet(keystorePath, keystorePassword)
                        "
                      >
                        <span style="display: block">LOGIN</span>
                      </p-button>
                    </div>
                  </div>
                  <div class="row">
                    <div class="col">
                      <p class="text-left poppin-text">
                        Don't have a Wallet?
                        <a class="link-text" @click="createAccount()">
                          Create one here!
                        </a>
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
  import { keyStore } from '@stardust-collective/dag4-keystore';
import { keyStoreFile } from "../../p12";
import { dag4 } from '@stardust-collective/dag4';
import Swal from "sweetalert2/dist/sweetalert2";

export default {
  name: "login-single-password",
  data: () => ({
    keystorePassword: "",
    KeyPassword: "",
    overlay: false,
    valid: false,
  }),
  mounted() {
    this.migrateNotification();
  },
  computed: {
    keystorePath: {
      get() {
        return this.$store.state.wallet.keystorePath;
      },
    },
  },
  methods: {
    validCheck: function() {
      this.valid = this.keystorePath !== null && this.keystorePassword !== "";
    },
    fileSelected: function() {
      this.validCheck();
    },
    migrateNotification: function() {
      Swal.mixin({
        customClass: {
          popup: "swal2-popup-login",
          actions: "swal2-actions-login",
          title: "swal2-title-login",
          confirmButton: "btn-migrate-login",
          closeButton: "btn-close-login",
          container: "swal2-container-login",
        },
        buttonsStyling: false,
      })
        .fire({
          title: "Important Update",
          html: `<br><p class="login-content">If you have previously signed in using a P12 file (versions 1.2.x and earlier), you will need to migrate your file to the JSON format.</p>`,
          width: "18.75rem",
          padding: "0.75rem",
          toast: true,
          background: "#DD8D74",
          position: "top-end",
          showConfirmButton: true,
          confirmButtonColor: "#DD8D74",
          confirmButtonText: '<div class="login-button-text">MIGRATE</div>',
          showCloseButton: true,
          timerProgressBar: true,
          // willOpen: () => {
          //   Swal.showLoading();
          // },
          onClose: () => {},
        })
        .then((result) => {
          /* Read more about isConfirmed, isDenied below */
          if (result.isConfirmed) {
            // Swal.fire('Saved!', '', 'success')
            this.moveToMigrate();
          }
        });
    },
    moveToMigrate: function() {
      Swal.close();
      this.$store.dispatch("wallet/reset").then(() => {
        this.$router.push({
          name: "keystore migrate",
          params: {
            message:
              "Enter the information below to migrate your P12 file to a JSON file",
            title: "Molly Wallet migration wizard",
            darkMode: this.$route.params.darkMode,
          },
        });
      });
    },
    createAccount: function() {
      Swal.close();
      this.$store.dispatch("wallet/reset").then(() => {
        this.$router.push({
          name: "create account",
          params: {
            message: "Enter a name and password for your Private Key file",
            title: "Create a Private Key File",
            darkMode: this.$route.params.darkMode,
          },
        });
      });
    },
    loadKeyStoreFile: function(filePath, password) {
      if (!filePath || !password) {
        return;
      }

      Swal.close();
      this.$Progress.start();
      this.overlay = true;

      let reader = new FileReader();

      //Settings and Wallet info shows the name of the file
      this.$store.commit("wallet/setKeystorePath", filePath.name);

      const ext = filePath.name
        .split(".")
        .pop()
        .toLowerCase();
      const isJson = ext === "json";

      if (isJson) {
        reader.readAsText(filePath);
      } else {
        reader.readAsBinaryString(filePath);
      }

      reader.onload = async () => {
        let privateKey;

        try {
          if (isJson) {
            privateKey = await keyStore.decryptPrivateKey(
              JSON.parse(reader.result),
              password
            );
          } else {
            privateKey = keyStoreFile.readP12(reader.result, password)
              .privateKey;
          }
        } catch (e) {
          this.overlay = false;
          this.$Progress.fail();
          Swal.fire("Unable to unlock file", e.message, "error").then(() => {
            this.migrateNotification();
          });
        }

        if (privateKey) {
          this.loginWithKey(privateKey);
        }
      };

      reader.onerror = (e) => {
        this.overlay = false;
        this.$Progress.fail();
        Swal.fire("Unable to read file", "", e.message).then(() => {
          this.migrateNotification();
        });
      };
    },
    loginWithKey: function(key) {
      if (!key) return;

      dag4.account.loginPrivateKey(key);

      var address = keyStore.getDagAddressFromPublicKey(
        keyStore.getPublicKeyFromPrivate(key)
      );

      // try {
      //   //NOTE: safely continue if window.firebase.analytics has failed to load
      //   window.firebase.analytics().logEvent('login', {address: address});
      // }
      // catch(e) {
      //   // Swal.fire("Unable to read window.firebase.analytics", "", e.message);
      // }

      window.backend.WalletApplication.CreateOrInitWalletV2(address).then(
        (result) => {
          if (result) {
            Swal.close();
            this.initWallet();
          }
        }
      );
    },
    // eslint-disable-next-line no-unused-vars
    loginJsonWallet: function(filePath, password) {
      // const key =
      //   "d4ace4d04e13e3441b7a34fb869dc09fa729d9b4fbf9e3377cbae3d88f75f049";
      // var address = keyStore.getDagAddressFromPublicKey(
      //   keyStore.getPublicKeyFromPrivate(key)
      // );
      // window.backend.WalletApplication.CreateOrInitWalletV2(address).then(
      //   (result) => {
      //     if (result) {
      //       this.overlay = false;
      //       this.$Progress.finish();
      //       this.initWallet();
      //     }
      //   }
      // );
      Swal.close();
      this.$Progress.start();
      this.overlay = true;

      window.backend.WalletApplication.LoginJsonWallet(filePath, password).then(
        (key) => {
          if (key) {
            dag4.account.loginPrivateKey(key);

            var address = keyStore.getDagAddressFromPublicKey(
              keyStore.getPublicKeyFromPrivate(key)
            );

            window.backend.WalletApplication.CreateOrInitWalletV2(address).then(
              (result) => {
                if (result) {
                  this.overlay = false;
                  this.$Progress.finish();
                  this.initWallet();
                }
              }
            );
          } else {
            this.overlay = false;
            this.$Progress.fail();
          }
        },
        (error) => {
          this.overlay = false;
          this.$Progress.finish();
          Swal.fire("Unable to login", error, "error");
        }
      );
    },
    initWallet: function() {
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

      window.backend.WalletApplication.CheckTermsOfService().then((result) => {
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
      });
    },
  },
};
</script>

<style scoped lang="scss">
.login-box {
  max-width: 29rem;
  min-width: 29rem;
  padding-bottom: 2rem;
}

.poppin-text {
  font-family: Poppins;
}

.link-text {
  color: #2d9cdb;
  transition: all 0.3s;
}

.link-text:hover {
  color: #db6e44;
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
