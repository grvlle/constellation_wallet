<template>
  <div class="container">
    <link
      href="//maxcdn.bootstrapcdn.com/font-awesome/4.1.0/css/font-awesome.min.css"
      rel="stylesheet"
    />
    <link
      href="https://fonts.googleapis.com/css?family=Lato"
      rel="stylesheet"
      type="text/css"
    />
    <div class="row">
      <div class="col-12">
        <div class="row">
          <div class="col mx-auto login-box">
            <div class="button-box">
              <div class="container">
                <div class="row">
                  <div class="col">
                    <button
                      class="primary-btn btn-1 primary-btn-sep icon-list"
                      block
                      v-on:click="moveToImportRecoveryPhrase()"
                    >
                      12 WORD RECOVERY PHRASE
                    </button>
                    <button
                      class="primary-btn btn-1 primary-btn-sep icon-file"
                      block
                      v-on:click="moveToImportKeystore()"
                    >
                      KEYSTORE FILE + SINGLE PASSWORD
                    </button>
                    <button
                      class="primary-btn btn-2 primary-btn-sep icon-file2"
                      block
                      v-on:click="moveToMigrate()"
                    >
                      MIGRATE YOUR KEYSTORE FILE + TWO PASSWORDS
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <page-overlay text="Loading..." :isActive="overlay" />
  </div>
</template>

<script>
import Swal from "sweetalert2/dist/sweetalert2";

export default {
  name: "import-wallet",
  data: () => ({ overlay: false }),
  mounted() {
    this.triggerNotification();
  },
  methods: {
    triggerNotification: function() {
      let timerInterval;
      Swal.fire({
        title:
          "<p style='text-align: left; color: white; margin: auto;'>Note!</p>",
        html: `<br><p style='text-align: left; color: white;'>If you have a Keystore file that was setup with two passwords and you used the same password for both fields, you can skip migrate and directly import it using the single password method.</p>`,
        width: 300,
        padding: 20,
        backdrop: false,
        toast: true,
        background: "#2654C0",
        position: "top-end",
        showConfirmButton: false,
        allowOutsideClick: false,
        showCloseButton: true,
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
    moveToImportRecoveryPhrase: function() {
      Swal.close();
      this.$store.dispatch("wallet/reset").then(() => {
        this.$router.push({
          name: "import recovery phrase",
          params: {
            message:
              "Enter your 12 word recovery phrase to import your Molly Wallet:",
            title: "Import recovery phrase",
            darkMode: this.$route.params.darkMode,
          },
        });
      });
    },
    moveToMigrate: function() {
      Swal.close();
      this.$store.dispatch("wallet/reset").then(() => {
        this.$router.push({
          name: "keystore migrate",
          params: {
            message:
              "Enter your information below to migrate your Molly Wallet:",
            title: "Molly Wallet migration wizard",
            darkMode: this.$route.params.darkMode,
          },
        });
      });
    },
    moveToImportKeystore: function() {
      Swal.close();
      this.$store.dispatch("wallet/reset").then(() => {
        this.$router.push({
          name: "import keystore",
          params: {
            message:
              "Enter your information below to migrate your Molly Wallet:",
            title: "Import keystore file",
            darkMode: this.$route.params.darkMode,
          },
        });
      });
    },
  },
};
</script>

<style scoped lang="scss">
.body-text {
  color: #666666;
  font-family: Poppins;

  font-size: 0.875rem;
  font-weight: 500;
}

.login-box {
  max-width: 29rem;
  min-width: 29rem;
  padding-bottom: 2rem;
  margin-top: 0.25em;
}

.primary-btn {
  padding: 1.5625em 5em;
  display: inline-block;
  margin: 0.9375em 0;
  width: 100%;
  height: 5.625em;
  border-radius: 0.1875em;
  background: #1d40b3;
  font-size: 0.875rem;
  color: white;
  outline: none;
  border: none;
  position: relative;
  -webkit-transition: all 0.3s;
  -moz-transition: all 0.3s;
  transition: all 0.3s;
}

.primary-btn:after {
  content: "";
  position: absolute;
  z-index: -1;
  -webkit-transition: all 0.3s;
  -moz-transition: all 0.3s;
  transition: all 0.3s;
}

/* Pseudo elements for icons */
.primary-btn:before {
  font-family: "FontAwesome";
  font-style: normal;
  font-weight: normal;
  font-variant: normal;
  text-transform: none;
  line-height: 1;
  position: relative;
  -webkit-font-smoothing: antialiased;
}

/* Icon separator */
.primary-btn-sep {
  padding: 1.5625em 3.75em 1.5625em 7.5em;
}

.primary-btn-sep:before {
  background: rgba(0, 0, 0, 0.15);
}

/* Button 1 */
.btn-1 {
  background: #1d40b3;
  color: #fff;
}

.btn-1:hover {
  background: #284bbe;
}

.btn-1:active {
  background: #284bbe;
  top: 0.125em;
}

.btn-1:before {
  position: absolute;
  height: 100%;
  left: 0;
  top: 0;
  line-height: 3;
  font-size: 140%;
  width: 3.75em;
}

/* Button 2 */
.btn-2 {
  background: #dd8d74;
  color: #fff;
}

.btn-2:hover {
  background: #ce9483;
}

.btn-2:active {
  background: #ce9483;
  top: 0.125em;
}

.btn-2:before {
  position: absolute;
  height: 100%;
  left: 0;
  top: 0;
  line-height: 3;
  font-size: 140%;
  width: 3.75em;
}

/* Icons */

.icon-list:before {
  content: "\f0ca";
  padding: 0.8em;
}

.icon-file:before {
  content: "\f15b";
  padding: 0.8em;
}

.icon-info:before {
  content: "\f05a";
  padding: 0.8em;
}

.icon-file2:before {
  content: "\f15c";
  padding: 0.8em;
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
