import Vue from "vue";

Vue.mixin({
  methods: {
    logout: function() {
      window.backend.WalletApplication.LogOut().then(txFinishedState => {
        if (txFinishedState) {
          this.$store.dispatch('resetWalletState');
          this.$store.dispatch('resetAppState');
          this.$store.dispatch('resetTransactionsState');
          this.$router.push({
            name: 'login',
            params: { message: "Please enter your credentials below to access your Molly Wallet." }
          });
          return;
        }
      }), (this.random = "1");
      return;
    }
  }
});