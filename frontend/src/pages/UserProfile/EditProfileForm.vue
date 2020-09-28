<template>
  <card class="card">
    <div>
      <form @submit.prevent>
        <div class="row">
          <div class="col-md-12">
            <fg-input
              type="text"
              label="Wallet Address"
              :disabled="true"
              placeholder="No DAG Address found."
              v-model="address"
            ></fg-input>
          </div>
        </div>
        <div class="row">
          <div class="col-md-12">
            <fg-input
              type="text"
              label="Available Balance"
              :disabled="true"
              placeholder="0"
              v-model="getNormalizedAvailableBalance"
            ></fg-input>
          </div>
        </div>
        <div class="row">
          <div class="col-md-12">
            <fg-input
              type="text"
              label="Wallet Client Version"
              :disabled="true"
              placeholder="0"
              v-model="version"
            ></fg-input>
          </div>
        </div>
        <div class="row">
          <div class="col-md-12">
            <fg-input
              type="text"
              label="Wallet UI Version"
              :disabled="true"
              placeholder="0"
              v-model="uiVersion"
            ></fg-input>
          </div>
        </div>
        <div class="row">
          <div class="col-md-12">
            <fg-input
              type="text"
              label="Alias"
              :disabled="true"
              placeholder="0"
              v-model="alias"
            ></fg-input>
          </div>
        </div>
        <div class="row">
          <div class="col-md-12">
            <fg-input
              type="text"
              label="Path to private key (key.p12)"
              :disabled="true"
              placeholder="0"
              v-model="keystorePath"
            ></fg-input>
          </div>
        </div>
        <div class="row">
          <div class="col-md-12">
            <fg-input
              type="text"
              label="Terms of Service"
              :disabled="true"
              placeholder="0"
              v-model="termsOfService"
            ></fg-input>
          </div>
        </div>
        <div class="clearfix"></div>
      </form>
    </div>
  </card>
</template>

<script>
import { mapState } from "vuex";
export default {
  data() {
    return {
      normAvailableBalance: this.getNormalizedAvailableBalance() 
    };
  },
  computed: {
    getNormalizedAvailableBalance () {
      return this.$store.getters["wallet/getNormalizedAvailableBalance"];
    },
    ...mapState("app", ["version", "uiVersion"]),
    ...mapState("wallet", [
      "address",
      "availableBalance",
      "nonce",
      "totalBalance",
      "alias",
      "keystorePath",
      "termsOfService",
    ]),    
  },
};
</script>
<style scoped lang="scss">
</style>
