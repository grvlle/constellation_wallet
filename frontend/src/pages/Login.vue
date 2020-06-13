<template>
  <div class="container">
    <div class="row">
      <div class="col-12">
        <form ref="textareaform" @submit.prevent="form" class="container">
          <div class="row">
            <div class="col mx-auto login-box">
              <div class="input-box">
                <div>
                  <label class="control-label">Select your private key (key.p12)</label>
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
                    :validate=false
                  />
                </div>
                <div>
                  <password-input
                    v-model="KeyPassword"
                    label="Key Password"
                    :validate=false
                  />
                </div>
              </div>
              <div class="button-box">
                <div class="container">
                  <div class="row">
                    <div class="col">
                      <p-button type="primary" block @click.native="login()">
                        <span style="display: block;"> LOGIN</span>
                      </p-button>
                    </div>
                  </div>
                  <div class="row">
                    <div class="col">
                      <p class="text-right">Don't have a wallet yet? Create one <a href="javascript:void(0)" @click="newWallet()">here!</a></p>
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
export default {
  name: "login-screen",
  data: () => ({
    keystorePassword: "",
    KeyPassword: "",
    overlay: false
  }),
  computed: {
    keystorePath: {
      get () {
        return this.$store.state.walletInfo.keystorePath
      },
      set (value) {
        this.$store.commit('walletInfo/setKeystorePath', value)
      }
    },
    alias: {
      get () {
        return this.$store.state.walletInfo.alias
      },
      set (value) {
        this.$store.commit('walletInfo/setAlias', value)
      }
    }
  },
  methods: {
    newWallet: function() {
      this.$store.dispatch('resetWalletState');
      this.$store.dispatch('resetAppState');
      this.$router.push({
        name: 'new wallet', 
        params: {
          message: "Create a new Molly wallet. Please ensure that you backup all information provided below in a safe place.",
          darkMode: this.$route.params.darkMode}
      });
    },
    login: function() {
      var self = this;
      self.$Progress.start();
      self.overlay = true;
      window.backend.WalletApplication.Login(
        self.$store.state.walletInfo.keystorePath,
        self.keystorePassword,
        self.KeyPassword,
        self.alias
      ).then(result => {
        if (result) {
          window.backend.WalletApplication.SetUserTheme().then(
            darkMode => (self.$store.commit('walletInfo/setDarkMode', darkMode))
          )
          window.backend.WalletApplication.SetWalletTag().then(
            walletTag => (self.$store.commit('walletInfo/setEmail', walletTag))
          );
          window.backend.WalletApplication.SetImagePath().then(
            imagePath => (self.$store.commit('walletInfo/setImgPath', imagePath))
          );
          self.overlay = false;
          self.$Progress.finish();
          self.$store.commit('app/setIsLoggedIn', true);

          window.backend.WalletApplication.CheckTermsOfService()
          .then (result => {
            self.$store.commit('walletInfo/setTermsOfService', result)
            if (result) {
              self.$router.push({
                name: 'loading', 
                params: {message: "Getting your $DAG Wallet ready..."}
              });
            } else {
              self.$router.push({
                name: 'accept terms of service',
                params: {message: "Terms of Service"}
              });
            }
          })
        } else {
          self.overlay = false;
          self.$Progress.fail();
        }
      });
    }
  }
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
