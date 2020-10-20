<template>
  <div class="container">
    <div class="row">
      <div class="col-12">
        <form ref="textareaform" @submit.prevent="form" class="container">
          <div class="row">
            <div class="col mx-auto login-box">
              <div>
                <label class="control-label"
                >Select your Private Key (P12 or JSON file)</label
                >
                <file-selector
                    v-model="keystorePath"
                    :placeholder="keystorePath"
                    action="SelectFile"
                    @file="fileSelected"
                />
              </div>

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
                          @click.native="loadKeyStoreFile(keystoreFile, keystorePassword)"
                      >
                        <span style="display: block">LOGIN</span>
                      </p-button>
                    </div>
                  </div>
                  <div class="row">
                    <div class="col">
                      <p class="text-right">
                        Don't have a Wallet? <b>Create one
                        <a class="link-text" @click="createAccount()">here!</a></b>
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
import { keyStore, keyStoreFile } from "@stardust-collective/dag-keystore";
import {dagWalletAccount} from '@stardust-collective/dag-wallet-sdk';
import Swal from "sweetalert2/dist/sweetalert2";

export default {
  name: "login-single-password",
  data: () => ({
    keystorePassword: "",
    KeyPassword: "",
    overlay: false,
    valid: false,
    keystoreFile: null
  }),
  computed: {
    keystorePath: {
      get() {
        return this.$store.state.wallet.keystorePath;
      },
      set(value) {
        this.$store.commit("wallet/setKeystorePath", value);
      },
    }
  },
  mounted() {
    this.migrateNotification();
  },
  methods: {
    fileSelected: function(value) {
      this.keystoreFile = value;
      this.keystorePath = value.name;
    },
    migrateNotification: function() {
      Swal.fire({
        title:
          "<p style='text-align: left; color: white; margin: auto;'>Important Update</p>",
        html: `<br><p style='text-align: left; color: white;'>If you have previously signed into Molly Wallet using an alias and two different passwords (versions 1.2.x and earlier), you will need to migrate your credentials before logging in.</p>`,
        width: 300,
        padding: 20,
        backdrop: false,
        toast: true,
        borderColor: "#DD8D74",
        background: "#DD8D74",
        position: "top-end",
        showConfirmButton: true,
        confirmButtonColor: '#DD8D74',
        confirmButtonText: '<div style="color: #B53C19;">MIGRATE</div>',
        allowOutsideClick: false,
        showCloseButton: true,
        timerProgressBar: true,
        willOpen: () => {
          Swal.showLoading();
        },
        onClose: () => {

        }
      }).then((result) => {
        /* Read more about isConfirmed, isDenied below */
        if (result.isConfirmed) {
          // Swal.fire('Saved!', '', 'success')
          this.moveToMigrate()
        }
      })
    },
    moveToMigrate: function() {
      Swal.close();
      this.$store.dispatch("wallet/reset").then(() => {
        this.$router.push({
          name: "keystore migrate",
          params: {
            message:
                "Enter the information below to migrate your existing two password credentials to a single password.",
            title: "Molly Wallet Migration",
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
    loadKeyStoreFile: function (filePath, password) {

      if (!filePath ||  !password) {
        // Swal.fire('Invalid credentials', '', 'error')
        this.loginWithKey();
        return
      }

      let reader = new FileReader();

      reader.readAsBinaryString(filePath);

      reader.onload = () => {

        let keyPair;

        try {
          keyPair = keyStoreFile.readP12(reader.result, password);
        }
        catch (e) {
          this.errorMessage = e.message;
        }

        if (keyPair) {
          this.loginWithKey(keyPair.privateKey);
        }
      };

      reader.onerror = () => {
        //TODO - ERROR
        //this.errorMessage = 'Unable to read file';
      };
    },
    loginWithKey: function (key) {

      if (!key) return

      // TODO - save seed and privKey to KeyChain (Alex)
      dagWalletAccount.loginPrivateKey(key);

      var address = keyStore.getDagAddressFromPublicKey(keyStore.getPublicKeyFromPrivate(key));
      // eslint-disable-next-line no-console
      console.log("getDagAddressFromPublicKey: " + address);
      window.backend.WalletApplication.CreateOrInitWalletV2(
          address
      ).then((result) => {
        if (result) {
          Swal.close();
          this.initWallet();
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
