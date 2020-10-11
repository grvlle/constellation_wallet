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
                      />
                    </p>
                    <p-button type="primary" block @click.native="moveToCreatePassword()">
                      <span style="display: block">
                        IMPORT</span
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

export default {
  name: "import-recovery-phrase",
  data: () => ({
    seed:
      "",
  }), 

  methods: {
 
    moveToCreatePassword: function () {
      Swal.close()
      this.$router.push({
        name: "import recovery phrase create wallet password",
        params: {
          message:
            "Please enter your new Molly Wallet password below:",
            title: "Import recovery phrase",
        },
      });
    },
  },
};
</script>

<style scoped lang="scss">
.body-text {
  color: #666666;
  font-family: Poppins;

  font-size: 14px;
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
