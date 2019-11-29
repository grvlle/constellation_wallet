<template>
  <card class="card">
    <div>
      <form @submit.prevent>
        <div class="row">
          <div class="col-md-12">
            <fg-input type="text"
                      label="Wallet Address"
                      :disabled="true"
                      placeholder="0x161D1B0bca85e29dF546AFba1360eEc6Ab4aA7Ee"
                      v-model="wallet.address">
            </fg-input>
          </div>
        </div>
    <div class="row">
          <div class="col-md-12">
            <fg-input type="text"
                      label="Available Balance"
                      :disabled="true"
                      placeholder="0"
                      v-model="wallet.available_balance">
            </fg-input>
          </div>
          </div>
          <div class="row">
          <div class="col-md-12">
            <fg-input type="text"
                      label="Nonce"
                      :disabled="true"
                      placeholder="0"
                      v-model="wallet.nonce">
            </fg-input>
          </div>
          </div>
          <div class="row">
          <div class="col-md-12">
            <fg-input type="text"
                      label="Total Balance"
                      :disabled="true"
                      placeholder="0"
                      v-model="wallet.total_balance">
            </fg-input>
          </div>
        </div>

        <div class="row">
          <div class="col-md-12">
            <fg-input type="text"
                      label="Wallet Client Version"
                      :disabled="true"
                      placeholder="0"
                      v-model="this.$store.state.walletInfo.version">
            </fg-input>
          </div>
        </div>

        <div class="row">
          <div class="col-md-12">
            <fg-input type="text"
                      label="Wallet UI Version"
                      :disabled="true"
                      placeholder="0"
                      v-model="this.$store.state.walletInfo.uiVersion">
            </fg-input>
          </div>
        </div>
        
        <div class="clearfix"></div>
      </form>
    </div>
  </card>
</template>
<script>
import Wallet from "../../../../JSONdata/wallet.json"
export default {
  data() {
    return {
      wallet: Wallet,
    };
  },
  methods: {
    updateProfile() {
      alert("Your data: " + JSON.stringify(Wallet));
    }
  }
};
</script>
<style>
</style>
