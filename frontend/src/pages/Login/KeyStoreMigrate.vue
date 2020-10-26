<template>
  <div class="container">
    <div class="row">
      <div class="col-12">
        <form ref="textareaform" @submit.prevent="form" class="container">
          <div class="row">
            <div class="col mx-auto login-box">
              <div class="input-box">
                <div>
                  <label class="control-label"
                    >Select your Private Key (P12 file)</label
                  >
                  <file-selector
                    v-model="keystorePath"
                    :placeholder="keystorePath"
                    action="SelectFile2"
                  />
                </div>
                <div>
                  <fg-input
                    style="margin-bottom: 0.125em"
                    type="text"
                    label="Key Alias"
                    v-model="alias"
                    :placeholder="alias"
                    @input="validCheck()"
                  />
                </div>
                <div>
                  <label class="control-label">
                    <span>Keystore Password </span>
                    <span style="color: #db6e44; font-size: 0.875em">
                    (Also use for future logins)
                  </span>
                  </label>

                  <password-input
                    style="margin: 0"
                    v-model="keystorePassword"
                    :validate="false"
                    @input="validCheck()"
                  />

                </div>
                <div>
                  <password-input
                    v-model="KeyPassword"
                    label="Key Password"
                    :validate="false"
                    @input="validCheck()"
                  />
                </div>
              </div>
              <div class="button-box">
                <div class="container">
                  <div class="row">
                    <div class="col">
                      <p-button
                        class="btn-secondary"
                        block
                        :disabled="!valid"
                        @click.native="migrate()"
                      >
                        <span style="display: block"> COMPLETE MIGRATION</span>
                      </p-button>
                    </div>
                  </div>
                  <div class="row">
                    <div class="col">
                      <!-- <p class="text-right">Don't have a wallet yet? Create one <a href="javascript:void(0)" @click="newWallet()">here!</a></p> -->
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
import Swal from "sweetalert2/dist/sweetalert2";

export default {
  name: "keystore-migrate",
  data: () => ({
    keystorePassword: "",
    KeyPassword: "",
    overlay: false,
    valid: false,
  }),
  computed: {
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
  methods: {
    validCheck: function() {
      this.valid =
        this.keystorePassword &&
        this.KeyPassword &&
        this.alias &&
        this.keystorePath;
    },
    completeMigration: function() {
      this.$router.push({
        name: "keystore migration complete",
        params: {
          title: "Molly Wallet migration wizard",
          message:
            "Congratulations! You have completed the Molly Wallet password migration!",
        },
      });
    },
    migrate: function() {
      var self = this;
      if (self.keystorePath) {
        self.$Progress.start();
        self.overlay = true;
        window.backend.WalletApplication.MigrateWallet(
          self.keystorePath,
          self.keystorePassword,
          self.KeyPassword,
          self.alias
        ).then(
          (result) => {
            self.overlay = false;
            self.$Progress.finish();
            if (result) {
              //TODO handle error case?
              this.$router.push({
                name: "keystore migration complete",
                params: {
                  title: "Molly Wallet migration wizard",
                  message:
                    "Congratulations! You have completed the Molly Wallet password migration!",
                },
              });
            }
          },
          (error) => {
            self.overlay = false;
            self.$Progress.fail();
            Swal.fire("Unable to migrate file", error, "error");
          }
        );
      }
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
        } else {
          self.overlay = false;
          self.$Progress.fail();
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

.input-box > div {
  margin-bottom: 1.875em;
}

.btn-secondary {
  background: #db6e44;
  border: 1px solid #db6e44;
  font-family: Poppins;
  font-style: normal;
  font-weight: 500;
  font-size: 12px;
  line-height: 18px;
  text-align: center;
  letter-spacing: 0.1em;
  border-radius: 4px;
  color: #ffffff;

  &:hover {
    background: #af5836;
    border: 1px solid #af5836;
  }

  &:active,
  &:focus {
    background: #db6e44 !important;
    border: 1px solid #db6e44 !important;
  }

  &:active {
    outline-color: #db6e44 !important;
    outline-width: 0px;
  }

  &:disabled {
    background: #e9a88f !important;
    border: 1px solid #e9a88f !important;
  }
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
