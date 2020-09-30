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
              v-model="normalizedAvailableBalance"
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
import { mapState, mapGetters } from "vuex";
export default {
  computed: {
    ...mapState("app", ["version", "uiVersion"]),
    ...mapState("wallet", [
      "address",
      "alias",
      "keystorePath",
      "termsOfService",
    ]),
    ...mapGetters("wallet", ["normalizedAvailableBalance"])
  },
};
</script>
<style scoped lang="scss">
</style>
