import Vue from "vue";

Vue.mixin({
  methods: {
    logout: function() {
      window.backend.WalletApplication.LogOut().then(txFinishedState => {
        if (txFinishedState) {
          let darkMode = this.$store.state.walletInfo.darkMode
          this.$store.dispatch('resetWalletState');
          this.$store.dispatch('resetAppState');
          this.$store.dispatch('resetTransactionsState');
          this.$router.push({
            name: 'login',
            params: {
              message: "Please enter your credentials below to access your Molly Wallet.",
              darkMode: darkMode
            }
          });
          return;
        }
      }), (this.random = "1");
      return;
    }
  }
});