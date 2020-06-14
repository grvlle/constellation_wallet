<template>
  <div class="container">
    <div class="row">
      <div class="col-12">
        <form ref="textareaform" @submit.prevent="form" class="container">
          <div class="row">
            <div class="col mx-auto login-box">
              <div>
                <b>Create a new wallet</b>
                <br />This section will let you create a Molly Wallet to store your
                <b>$DAG</b> tokens in. You simply browse to a path where you wish to save your KeyStore file, give it a name and select 'save'.
                <br />
                <br />Once the path is selected, you get to set up a password to protect the key store.
                <br />
                <br />
                <ul>
                  <li>
                    <b>KeyStore File</b>
                    <i>- Select where to save your KeyStore File.</i>
                  </li>
                  <li>
                    <b>Key Alias</b>
                    <i>- The unique name which is used in the keystore to identify this key entry.</i>
                  </li>
                  <li>
                    <b>Keystore Password</b>
                    <i>- This password unlocks the keystore file.</i>
                  </li>
                  <li>
                    <b>Key Password</b>
                    <i>- Extra layer of security. Both passwords will be needed when accessing/restoring a wallet.</i>
                  </li>
                  <li>
                    <b>Wallet Label</b>
                    <i>
                      - This will set the label of your wallet. This is
                      <b>optional</b> and strictly for cosmetic purposes.
                    </i>
                  </li>
                </ul>
                <b>Important!</b> Please backup your Alias, Store Passwords, Key Password and KeyStore File (key.p12) as these will allow you to restore your wallet at any time.
              </div>
            </div>
            <div class="col mx-auto login-box">
              <div class="input-box">
                <div>
                  <label class="control-label">
                    Select a directory to store your private key (key.p12) in
                  </label>
                  <file-selector 
                    v-model="keystorePath"
                    action="SelectSaveFile"
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
                  <div class="validate" v-if="!this.aliasValid">
                    <p v-if="!this.aliasContainsFiveCharacters">
                      Alias has to be atleast 5 characters long.
                    </p>
                  </div>
                  <div class="validate" v-else />
                </div>
                <div>
                  <password-input
                    v-model="keystorePassword"
                    label="Keystore Password"
                    :validate=true
                    v-on:valid="KeystorePasswordValid = true"
                    v-on:invalid="KeystorePasswordValid = false"
                  />
                </div>
                <div>
                  <password-input
                    v-model="KeyPassword"
                    label="Key Password"
                    :validate=true
                    v-on:valid="KeyPasswordValid = true"
                    v-on:invalid="KeyPasswordValid = false"
                  />
                </div>
                <div>
                  <fg-input
                    type="text"
                    v-model="walletLabel"
                    label="Wallet Label (optional)"
                  ></fg-input>
                </div>
              </div>
              <div class="button-box">
                <div class="container">
                  <div class="row">
                    <div class="col-md-6 pr-md-2 mb-3">
                      <p-button
                        type="default"
                        block
                        @click.native="cancel()"
                      >
                        <span style="display: block;">
                          <i class="fa fa-close"></i>
                          CANCEL
                        </span>
                      </p-button>
                    </div>
                    <div class="col-md-6 pl-md-2 mb-3">
                      <p-button
                        type="warning"
                        block
                        :disabled="!this.isValidNewWallet"
                        @click.native="createWallet()"
                      >
                        <span style="display: block;">
                          <i v-if="!this.isValidNewWallet" class="fa fa-lock"></i>
                          <i v-else class="fa fa-unlock"></i>
                          CREATE
                        </span>
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
export default {
  name: "new-wallet-screen",
  data: () => ({
    aliasValid: false,
    aliasLength: 0,
    aliasContainsFiveCharacters: false,
    keystorePassword: "",
    KeystorePasswordValid: false,
    KeyPassword: "",
    KeyPasswordValid: false,
    overlay: false
  }),
  computed: {
    isValidNewWallet: function() {
      if (
        this.aliasValid &&
        this.KeyPasswordValid &&
        this.KeystorePasswordValid &&
        this.keystorePassword !== "" &&
        this.KeyPassword !== "" &&
        !this.overlay
      ) {
        return true;
      } else {
        return false;
      }
    },
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
        if (value.length >= 5) {
          this.aliasContainsFiveCharacters = true;
          this.aliasValid = true;
        } else {
          this.aliasContainsFiveCharacters = false;
          this.aliasValid = false;
        }
        this.$store.commit('walletInfo/setAlias', value)
      }
    },
    walletLabel: {
      get () {
        return this.$store.state.walletInfo.walletLabel
      },
      set (value) {
        this.$store.commit('walletInfo/setLabel', value)
      }
    }
  },
  methods: {
    cancel: function() {
      this.$store.dispatch('walletInfo/reset').then(() => {
        this.$router.go(-1);
      })
    },
    createWallet: function() {
      var self = this;
      self.$Progress.start();
      self.overlay = true;
      window.backend.WalletApplication.CreateWallet(
        self.keystorePath,
        self.keystorePassword,
        self.KeyPassword,
        self.alias,
        self.walletLabel
      ).then(walletCreated => {
        if (walletCreated) {
          window.backend.WalletApplication.Login(
            self.keystorePath,
            self.keystorePassword,
            self.KeyPassword,
            self.alias
          ).then(loggedIn => {
            if (loggedIn) {
              self.overlay = false;
              self.$store.commit('app/setIsLoggedIn', loggedIn);
              self.$router.push({
                name: 'accept terms of service',
                params: {message: "Terms of Service"}
              });
            } else {
              self.overlay = false;
              self.$Progress.fail();
            }
          });
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
