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
                  <label class="control-label"
                  >Select a folder for your KeyStore file</label
                  >
                  <file-selector
                      v-model="keystorePath"
                      :placeholder="keystorePath"
                      action="SelectFile"
                  />
                </div>
                <div>
                  <password-input
                    v-model="keystorePassword"
                    label="New Password"
                    :validate="true"
                    @valid="validatePassword(true)"
                    @invalid="validatePassword(false)"
                  />
                </div>
                <div>
                  <password-input
                    v-model="keyPassword"
                    label="Repeat New Password"
                    @input="confirmPassword()"
                  />
                  <div
                    class="validate text-danger"
                    v-if="!valid && valid_password"
                  >
                    Need to confirm the password
                  </div>
                </div>
              </div>
              <div class="button-box">
                <div class="container">
                  <div class="row">
                    <div class="col">
                      <p-button
                        type="primary"
                        block
                        :disabled="valid"
                        @click.native="createKeyStore()"
                      >
                        <span style="display: block"> CREATE KEYSTORE FILE</span>
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
  components: {},
  name: "create account",
  data: () => ({
    keystorePassword: "",
    keyPassword: "",
    valid_password: false,
    valid: false,
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
    confirmPassword: function() {
      this.valid =
        this.valid_password && this.keystorePassword === this.keyPassword;
    },
    validatePassword: function(is_valid) {
      this.valid_password = is_valid;
      this.confirmPassword();
    },
    createKeyStore: function () {
      // const fileBin = '';

      // return window.backend.WalletApplication.CreateKeyStoreFile(fileBin).then((result) => {
      //   if (result) {
          this.$store.dispatch("wallet/reset").then(() => {
            this.$router.push({
              name: "create account complete",
              params: {
                message: "Congratulations! You have created a KeyStore file for Molly Wallet!",
                title: "Create a KeyStore File",
                darkMode: this.$route.params.darkMode,
              },
            });
          });
      //   }
      // });
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
