import Vue from "vue";

Vue.mixin({
  methods: {
    logout: function () {
      window.backend.WalletApplication.LogOut().then(txFinishedState => {
        if (txFinishedState) {
          let darkMode = this.$store.state.wallet.darkMode
          this.$store.dispatch('transaction/reset').then(() => {
            this.$store.dispatch('addressBook/reset').then(() => {
              this.$store.dispatch('wallet/reset').then(() => {
                this.$store.dispatch('app/reset').then(() => {
                  this.$router.push({
                    name: 'login',
                    params: {
                      message: "Please enter your credentials below to access your Molly Wallet.",
                      darkMode: darkMode
                    }
                  });
                  return;
                })
              })
            })
          })
        }
      }), (this.random = "1");
      return;
    }
  }
});