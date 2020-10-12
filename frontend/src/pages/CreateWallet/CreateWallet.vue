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
                        :disabled="!valid"
                        @click.native="moveToRecoveryPhraseInfo()"
                      >
                        <span style="display: block"> CREATE NEW ACCOUNT</span>
                      </p-button>
                    </div>
                  </div>
                  <div class="row">
                    <div class="col">
                      <p-button
                        type="secondary"
                        block
                        @click.native="moveToImportWallet()"
                      >
                        <span style="display: block"> RESTORE ACCOUNT</span>
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
import Swal from "sweetalert2/dist/sweetalert2";

export default {
  components: {},
  name: "create-wallet",
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
  mounted() {
    this.migrateNotification();
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
    moveToRecoveryPhraseInfo: function() {
      if (this.valid === false) return;
      Swal.close();
      this.$store.dispatch("wallet/reset").then(() => {
        this.$router.push({
          name: "recovery phrase info",
          params: {
            message: "Let's first create our recovery phrase!",
            title: "Recovery Phrase",
            darkMode: this.$route.params.darkMode,
          },
        });
      });
    },
    moveToImportWallet: function() {
      Swal.close();
      this.$store.dispatch("wallet/reset").then(() => {
        this.$router.push({
          name: "import wallet",
          params: {
            message: "Please select how you would like to import your wallet:",
            title: "Import an existing wallet",
            darkMode: this.$route.params.darkMode,
          },
        });
      });
    },
    migrateNotification: function() {
      let timerInterval;
      Swal.fire({
        title:
          "<p style='text-align: left; color: white; margin: auto;'>Note</p>",
        html: `<br><p style='text-align: left; color: white;'>This is the password you will use to login to your wallet each session. If you already have an existing wallet, please select the <b>restore</b> option.</p>`,
        width: 300,
        padding: 20,
        backdrop: false,
        toast: true,
        background: "#2654C0",
        position: "top-end",
        showConfirmButton: false,
        allowOutsideClick: false,
        showCloseButton: true,
        timer: 20000,
        timerProgressBar: true,
        willOpen: () => {
          Swal.showLoading();
          timerInterval = setInterval(() => {
            const content = Swal.getContent();
            if (content) {
              const b = content.querySelector("b");
              if (b) {
                b.textContent = Swal.getTimerLeft();
              }
            }
          }, 100);
        },
        onClose: () => {
          clearInterval(timerInterval);
        },
      });
    },
    completeMigration: function() {
      this.$router.push({
        name: "password migration complete",
        params: {
          message:
            "Congratulations! You have completed the Molly Wallet password migration!",
        },
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
