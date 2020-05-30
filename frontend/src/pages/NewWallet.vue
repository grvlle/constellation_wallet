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
                  <label
                    class="control-label"
                  >Select a directory to store your private key (key.p12) in</label>
                  <file-selector
                    v-model="this.$store.state.walletInfo.saveKeystorePath"
                    :placeholder="this.$store.state.walletInfo.saveKeystorePath"
                    action="SelectSaveFile"
                  />
                </div>
                <div>
                  <fg-input
                    style="margin-bottom: 0.125em"
                    type="text"
                    v-model="alias"
                    @input.native="checkAlias(alias)"
                    :placeholder="this.$store.state.walletInfo.alias"
                    label="Key Alias"
                  />
                  <div class="validate" v-if="!this.aliasValid">
                    <p
                      v-if="!this.aliasContainsFiveCharacters"
                    >Alias has to be atleast 5 characters long.</p>
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
                    v-model="newWalletLabel"
                    :placeholder="this.$store.state.walletInfo.email"
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
    newWalletLabel: "",
    alias: "",
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
        this.alias !== "" &&
        this.keystorePassword !== "" &&
        this.KeyPassword !== "" &&
        !this.overlay
      ) {
        return true;
      } else {
        return false;
      }
    }
  }, 
  methods: {
    checkAlias: function() {
      this.aliasLength = this.alias.length;

      if (this.aliasLength >= 5) {
        this.aliasContainsFiveCharacters = true;
        this.aliasValid = true;
      } else {
        this.aliasContainsFiveCharacters = false;
        this.aliasValid = false;
      }
    },
    cancel: function() {
      this.resetData();
      this.$router.go(-1);
    },
    resetData: function() {
      this.alias = "";
      this.aliasLength = 0;
      this.aliasContainsFiveCharacters = false;
      this.aliasValid = false;
      this.KeyPassword = "";
      this.KeyPasswordValid = false;
      this.keystorePassword = "";
      this.KeystorePasswordValid = false;
      this.$store.state.walletInfo.keystorePath = "";
      this.$store.state.walletInfo.alias = "";
      this.$store.state.walletInfo.keystorePassword = "";
      this.$store.state.walletInfo.KeyPassword = "";
      this.$store.state.displayLoginError = false;
    },
    createWallet: function() {
      var self = this;
      self.$Progress.start();
      self.overlay = true;
      if (self.newWalletLabel !== "") {
        self.$store.state.walletInfo.email = self.newWalletLabel;
        window.backend.WalletApplication.StoreWalletLabelInDB(
          self.newWalletLabel
        );
      }
      window.backend.WalletApplication.CreateWallet(
        self.$store.state.walletInfo.keystorePath,
        self.keystorePassword,
        self.KeyPassword,
        self.alias
      ).then(walletCreated => {
        if (walletCreated) {
          window.backend.WalletApplication.Login(
            self.$store.state.walletInfo.keystorePath,
            self.keystorePassword,
            self.KeyPassword,
            self.alias
          ).then(loggedIn => {
            if (loggedIn) {
              self.overlay = false;
              self.$store.state.app.isLoggedIn = true;
              self.$router.push({
                name: 'loading', 
                params: {message: "Getting your $DAG Wallet ready..."}
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
