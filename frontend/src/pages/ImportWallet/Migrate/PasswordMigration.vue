<template>
  <div class="container">
    <div class="row">
      <div class="col-12">
        <form ref="textareaform" @submit.prevent="form" class="container">
          <div class="row">
            <div class="col mx-auto login-box">

              <br />
              <div class="input-box">

                <div>
                  <password-input
                    v-model="keystorePassword"
                    label="New Password"
                    :validate="true"
                  />
                </div>
                                <div>
                  <password-input
                    v-model="keyPassword"
                    label="Repeat New Password"
                    :validate="false"
                  />
                </div>
              </div>
              <div class="button-box">
                <div class="container">
                  <div class="row">
                    <div class="col">
                      <p-button style="background: #dd8d74; border-color: #dd8d74;" block @click.native="completeMigration()">
                        <span style="display: block"> COMPLETE MIGRATION</span>
                      </p-button>
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

export default {
  components: {
  },
  name: "password-migration-screen",
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
  methods: {
    completeMigration: function() {
        this.$router.push({
        name: "password migration complete",
        params: {
          title: "Molly Wallet migration wizard",
          message: "Congratulations! You have completed the Molly Wallet password migration!",
        },
      });
    },
    login: function () {
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
