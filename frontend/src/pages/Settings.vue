<template>
    <div class="row">
      <div class="col-md-6">
      <card>
        <h4 class="card-title">Settings</h4>
                    <p class="card-category">Configure the Constellation Wallet to your personal preferences</p>
                    <br>
        <p>asf</p>
        </card>
      </div>
        <div class="col-md-6">
            <card>
                
                    <h4 class="card-title">Backup / Restore</h4>
                    <p class="card-category">Backup your Wallet</p>
                    <br>
                    
                    <form @submit.prevent>
                                            
                        <label class="control-label">
                          <p>Mnemonic Seed</p>
                        </label>
                        <p class="small" style="margin-bottom: -5px; margin-top: -20px;">This is the seed that will be used to backup and restore your wallet.</p>
                      <div class="row">
                        
                        <div class="input-group col-md-12">
                          
                          <input :type=seed class="form-control" :disabled="true" placeholder="Mnemonic Seed" v-model="this.$store.state.walletInfo.seed">
                          <div class="input-group-append">
                            <p-button class="btn" @click.native="showPassword2()" type="danger"><i v-bind:class="btnText"/></p-button>
                          </div>
                          
                      </div>
                    </div>
                        <div class="row">
                          <div class="col-md-12">
                            <fg-input type="text"
                                      label="Public Key / Private Key"
                                      :disabled="true"
                                      placeholder="Public Key"
                                      v-model="this.$store.state.walletInfo.publicKey">
                            </fg-input>
                          </div>
                        </div>
                        <div class="row">
                        <div class="input-group col-md-12">
                          <input :type=type label="Private Key" class="form-control" :disabled="true" placeholder="Private Key" v-model="this.$store.state.walletInfo.privateKey" aria-describedby="basic-addon2">
                          <div class="input-group-append">
                            <p-button class="btn" @click.native="showPassword()" type="danger"><i v-bind:class="btnText"/></p-button>
                          </div>
                      </div>
                    </div>
            </form>
          
    </card>
  </div>

 </div>

</template>

<script>
export default {

    methods: {
        showPassword() {
            if (this.type === 'password') {
                this.type = 'text'
                this.btnText = 'fa fa-eye-slash'
            } else {
                this.type = 'password'
                this.btnText = 'fa fa-eye'
            }
        },
         showPassword2() {
            if (this.seed === 'password') {
                this.seed = 'text'
                this.btnText = 'fa fa-eye-slash'
            } else {
                this.seed = 'password'
                this.btnText = 'fa fa-eye'
            }
        }
    },
    mounted() {
        window.wails.Events.On("wallet_keys", (privateKey, publicKey) => {
            this.$store.state.walletInfo.privateKey = privateKey;
            this.$store.state.walletInfo.publicKey = publicKey;
        });

    },
    data() {
        return {
            type: 'password',
            seed: 'password',
            btnText: 'fa fa-eye'
        }
    }

};
</script>

<style>

</style>
