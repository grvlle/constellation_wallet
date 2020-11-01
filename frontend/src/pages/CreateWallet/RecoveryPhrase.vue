<template>
  <div class="container">
    <div class="row">
      <div class="col-12">
        <div class="row">
          <div class="col mx-auto login-box">
            <div class="button-box">
              <div class="container">
                <div class="row">
                  <div class="col">
                    <p class="body-text">
                      <textarea
                        class="form-control"
                        label="Mnemonic Seed"
                        v-model="seed"
                        :placeholder="seed"
                        :readonly="true"
                      />
                    </p>
                    <p-button
                      type="primary"
                      block
                      @click.native="moveToLogin()"
                    >
                      <span style="display: block">
                        I HAVE WRITTEN DOWN MY RECOVERY PHRASE</span
                      >
                    </p-button>
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
import { keyStore } from "@stardust-collective/dag-keystore";

const seed = keyStore.generateSeedPhrase();

export default {
  name: "recovery-phrase",
  data: () => ({ seed, overlay: false }), //"witch collapse practice feed shame open despair creek road again ice least"
  mounted() {
    this.warningNotification();
  },
  methods: {
    warningNotification: function() {
      let timerInterval;
      Swal.fire({
        title:
          "<p style='text-align: left; color: white; margin: auto;'>Warning</p>",
        html: `<br><p style='text-align: left; color: white;'>Never disclose your recovery phrase. Anyone with this phrase can access your funds.</p>`,
        width: 300,
        padding: 20,
        backdrop: false,
        toast: true,
        background: "#DD8D74",
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
    moveToLogin: function() {

      const privateKey = keyStore.getPrivateKeyFromMnemonic(this.seed);
      window.backend.WalletApplication.SavePhraseandPKeyToKeychain(
        this.seed,
        privateKey
      ).then((result) => {
        if (result) {
          Swal.close();
          this.$router.push({
            name: "login",
            params: {
              message:
                "Please enter your credentials below to access your Molly Wallet.",
            },
          });
        }
      });
    }
  }
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
  margin-top: 5.25em;
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
