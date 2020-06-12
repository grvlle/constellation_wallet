<template>
  <div class="container">
    <div class="row">
      <div class="col-md-6 mx-auto text-center">
        <p
          v-if="this.$store.state.app.downloading.filename !== ''"
        >{{this.$store.state.app.downloading.filename}}: {{this.$store.state.app.downloading.size}}</p>
      </div>
    </div>
    <div class="row">
      <div class="col-md-6 mx-auto text-center">
        <loader />
      </div>
    </div>
    <div class="row" style="height: 21rem;" />
  </div>
</template>

<script>
import Loader from "../components/Loader";
export default {
  name: "downloading-screen",
  components: {
    Loader
  },
  created: function() {
    var self = this;
    window.backend.WalletApplication.CheckAndFetchWalletCLI().then(exists => {
      if (exists) {
        self.$router.push({
          name: 'login', 
          params: {message: "Please enter your credentials below to access your Molly Wallet."}
        });
      }
    });
  }
};
</script>

