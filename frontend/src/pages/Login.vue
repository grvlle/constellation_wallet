<template>
  <div class="container">
    <div class="row">
      <div class="col-12">
        <form ref="textareaform" @submit.prevent="form" class="container">
          <div class="row">
            <div class="col mx-auto login-box">
              <slot name="label">
                <label class="control-label"> Select Network </label>
              </slot>
              <vue-select
                class="select"
                @input="setNetwork"
                :value="network"
                :options="['Main Constellation Network', 'Eros Test Network', 'Ceres Test Network']"
              ></vue-select>
              <br />
              <div class="input-box">
                <div>
                  <label class="control-label"
                    >Select your private key (key.p12)</label
                  >
                  <file-selector
                    v-model="keystorePath"
                    :placeholder="keystorePath"
                    action="SelectFile"
                  />
                </div>
                <div>
                  <fg-input
                    style="margin-bottom: 0.125em"
                    type="text"
                    label="Key Alias"
                    v-model="alias"
                    :placeholder="alias"
                  />
                </div>
                <div>
                  <password-input
                    v-model="keystorePassword"
                    label="Keystore Password"
                    :validate="false"
                  />
                </div>
                <div>
                  <password-input
                    v-model="KeyPassword"
                    label="Key Password"
                    :validate="false"
                  />
                </div>
              </div>
              <div class="button-box">
                <div class="container">
                  <div class="row">
                    <div class="col">
                      <p-button type="primary" block @click.native="login()">
                        <span style="display: block"> LOGIN</span>
                      </p-button>
                    </div>
                  </div>
                  <div class="row">
                    <div class="col">
                      <p class="text-right">
                        Don't have a wallet yet? Create one
                        <a href="javascript:void(0)" @click="newWallet()"
                          >here!</a
                        >
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
import VueSelect from "vue-select";
import { mapState } from "vuex";

export default {
  components: {
    VueSelect,
  },
  name: "login-screen",
  data: () => ({
    keystorePassword: "",
    KeyPassword: "",
    overlay: false,
  }),
  computed: {
    ...mapState("app", [
      "network",
    ]),
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
    setNetwork: function (value) {
      window.backend.WalletApplication.StoreCurrencyStateDB(value).then(
        (result) => {
          if (result) {
            this.$store.commit("app/setNetwork", value);
          }
        }
      );
    },
    newWallet: function () {
      this.$store.dispatch("wallet/reset").then(() => {
        this.$router.push({
          name: "new wallet",
          params: {
            message:
              "Create a new Molly wallet. Please ensure that you backup all information provided below in a safe place.",
            darkMode: this.$route.params.darkMode,
          },
        });
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
