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
              v-model="this.$store.state.walletInfo.address"
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
              v-model="this.$store.state.walletInfo.totalBalance"
            ></fg-input>
          </div>
        </div>
        <div class="row">
          <div class="col-md-12">
            <fg-input
              type="text"
              label="Nonce"
              :disabled="true"
              placeholder="0"
              v-model="this.$store.state.walletInfo.nonce"
            ></fg-input>
          </div>
        </div>
        <div class="row">
          <div class="col-md-12">
            <fg-input
              type="text"
              label="Total Balance"
              :disabled="true"
              placeholder="0"
              v-model="this.$store.state.walletInfo.totalBalance"
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
              v-model="this.$store.state.walletInfo.version"
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
              v-model="this.$store.state.walletInfo.uiVersion"
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
              v-model="this.$store.state.app.termsOfService"
            ></fg-input>
          </div>
        </div>
        
        <div class="clearfix"></div>
      </form>
    </div>
  </card>
</template>
<script>
export default {
  data() {
    return {};
  }
};
</script>
<style>
</style>
