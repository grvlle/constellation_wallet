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
                    name: 'login single password',
                    params: {
                      message: "Please enter the credentials to your Keystore file.",
                      darkMode: darkMode
                    }
                  });

                })
              })
            })
          })
        }
      }), (this.random = "1");

    }
  }
});
