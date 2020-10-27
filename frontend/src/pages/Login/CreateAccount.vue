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
                    >Enter a name for your Private Key file</label
                  >
                  <input
                    type="text"
                    name="fileName"
                    class="inputfile"
                    @input="checkPassword($event.target.value)"
                  />
                  <div class="validate text-danger">
                    <p
                      class="error-message"
                      v-bind:class="{ resolved: fileNameValid }"
                    >
                      alpha-numeric
                    </p>
                  </div>
                </div>
                <div>
                  <password-input
                    v-model="keystorePassword"
                    label="Password"
                    :validate="true"
                    @valid="validatePassword(true)"
                    @invalid="validatePassword(false)"
                  />
                </div>
                <div>
                  <password-input
                    style="margin: 0"
                    v-model="keyPassword"
                    label="Repeat Password"
                    @input="confirmPassword()"
                  />
                  <div class="validate text-danger">
                    <p
                      class="error-message"
                      v-bind:class="{ resolved: confirmed_password }"
                    >
                      password match
                    </p>
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
                        :disabled="!valid"
                        @click.native="createKeyStore()"
                      >
                        <span style="display: block"> CREATE</span>
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
    <page-overlay text="Loading" :isActive="overlay" />
  </div>
</template>

<script>
import Swal from "sweetalert2/dist/sweetalert2";
import { mapState } from "vuex";
import { keyStore } from "@stardust-collective/dag-keystore";

export default {
  components: {},
  name: "create-account",
  data: () => ({
    keystorePassword: "",
    keyPassword: "",
    fileName: "",
    fileNameValid: false,
    valid_password: false,
    confirmed_password: true,
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
    checkPassword: function(value) {
      this.$emit("input", value);
      const format = /^[0-9a-zA-Z_]+$/;

      this.fileNameValid = format.test(value);

      if (this.fileNameValid) {
        this.$emit("valid");
      } else {
        this.$emit("invalid");
      }

      this.fileName = value;

      this.confirmed_password = this.keystorePassword === this.keyPassword;

      this.valid =
        this.valid_password && this.fileNameValid && this.confirmed_password;
    },
    confirmPassword: function() {
      this.confirmed_password = this.keystorePassword === this.keyPassword;
      this.valid =
        this.valid_password && this.fileNameValid && this.confirmed_password;
    },
    validatePassword: function(is_valid) {
      this.valid_password = is_valid;
      this.confirmPassword();
    },
    createKeyStore: async function() {
      var self = this;
      self.$Progress.start();
      self.overlay = true;

      const jsonObj = await keyStore.generateEncryptedPrivateKey(
        this.keyPassword
      );

      return window.backend.WalletApplication.CreateKeyStoreFile(
        this.fileName,
        JSON.stringify(jsonObj)
      ).then((filePath) => {
        if (filePath) {
          self.overlay = false;
          self.$Progress.finish();
          var path = filePath;
          if (filePath[filePath.length - 1] === ".") {
            path = filePath.slice(0, filePath.length - 1);
            Swal.fire(
              "Info: A file already exists using the name, \"" + this.fileName + "\".",
              "The new file has been created as - " + filePath
            );
          }
          this.$store.dispatch("wallet/reset").then(() => {
            this.$router.push({
              name: "create account complete",
              params: {
                message:
                  "Congratulations! You have created a Private Key file for Molly Wallet!",
                title: "Create a Private Key File",
                filePath: path,
                darkMode: this.$route.params.darkMode,
              },
            });
          });
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
.inputfile {
  height: 36px;
  width: 100%;
  border-radius: 0.25rem;
  //font-size: 0.75em;
  font-weight: 600;
  border: 1px solid #dddddd;
  display: block;
  background: #f9f7f7 !important;
  font-family: Poppins;
  color: #666666;
  padding: 6px 16px;
}

.login-box {
  max-width: 29rem;
  min-width: 29rem;
  padding-bottom: 2rem;
}

.input-box > div {
  margin-bottom: 1.875em;
}

.error-message {
  font-family: Poppins;
  font-style: normal;
  font-weight: normal;
  font-size: 10px;
  line-height: 15px;
  letter-spacing: 0.05em;
  color: #eb5757;
  &.resolved {
    color: #219653;
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
