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
                      A recovery phrase is a series of 12 words in a specific
                      order. This word combination is unique to your wallet.
                      Make sure to have pen and paper ready so you can write it
                      down.
                      <br />
                      <br />
                      Keep the recovery phrase stored in a safe place. Losing
                      this phrase will mean you will no longer be able to access
                      your funds.
                      <br />
                      <br />
                    </p>
                    <p-button
                      type="primary"
                      block
                      @click.native="moveToRecoveryPhrase()"
                    >
                      <span style="display: block"> START</span>
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
export default {
  name: "recovery-phrase-info",
  data: () => ({
    overlay: false,
  }),
  methods: {
    moveToRecoveryPhrase: function() {
      this.$store.dispatch("wallet/reset").then(() => {
        this.$router.push({
          name: "recovery phrase",
          params: {
            message: "Please write down your 12 word recovery phrase.",
            title: "Recovery Phrase",
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
