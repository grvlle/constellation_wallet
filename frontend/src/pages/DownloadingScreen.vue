<template>
  <div class="container">
    <div class="row">
      <div class="col-md-6 mx-auto text-center">
        <p v-if="downloading.filename !== ''" >
          {{downloading.filename}}: {{downloading.size}}
        </p>
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
import {mapState} from 'vuex'
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
          name: 'login single password',
          params: {message: "Please enter the credentials to your Keystore file."}
        });
      }
    });
  },
  computed: {
    ...mapState('app', ['downloading']) 
  },
};
</script>

<style scoped lang="scss">
</style>
