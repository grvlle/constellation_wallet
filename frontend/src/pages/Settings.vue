<template>
    <div class="row">
        <div class="col-md-12">
            <card>
                <template slot="header">
                                <h4 class="card-title">Settings</h4>
                                <p class="card-category">Configure the Constellation Wallet to your personal preferences</p>
                                <br>
                  <form @submit.prevent>
                    <div class="row">
                      <div class="col-md-12">
                        <fg-input type="text"
                                  label="Private Key"
                                  :disabled="true"
                                  placeholder="Private Key"
                                  v-model="this.$store.state.walletInfo.privateKey">
                        </fg-input>
                      </div>
                    </div>
                   </form>
                <form @submit.prevent>
                    <div class="row">
                      <div class="col-md-12">
                        <fg-input type="text"
                                  label="Public Key"
                                  :disabled="true"
                                  placeholder="Public Key"
                                  v-model="this.$store.state.walletInfo.publicKey">
                        </fg-input>
                      </div>
                    </div>
                   </form>
</template>
    </card>
  </div>
 </div>
</template>

<script>
export default {
    // computed: {
    //     wallet3() {
    //         return this.$store.state.walletInfo;
    //     }
    // },
    mounted() {
        window.wails.Events.On("wallet_keys", (privateKey, publicKey) => {
            this.$store.state.walletInfo.privateKey = privateKey;
            this.$store.state.walletInfo.publicKey = publicKey;
        });

    },
};
</script>

<style>

</style>
