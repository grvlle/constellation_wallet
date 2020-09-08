<template>
  <div id="app" class="login-bg vertical-center" v-bind:style="{ backgroundImage: 'url(' + themeBG + ')' }" >
    <div class="container">
      <div class="row">
        <div class="col mx-auto text-center header">
          <div>
            <img class="img-fluid" v-if="isDarkMode" src="~@/assets/img/Constellation-Logo-White.png" style="max-height: 5.8rem;" />
            <img class="img-fluid" v-else src="~@/assets/img/Constellation-Logo-Black.png" style="max-height: 5.8rem;" />
            <p v-if="this.$route.params.message">{{this.$route.params.message}}</p>
            <p v-else>Downloading $DAG wallet dependencies...</p>
          </div>
          <div class="page-error-box text-danger" v-if="displayLoginError">
            <p>{{loginErrorMsg}}</p>
          </div>
          <div class="page-error-box text-danger" v-else></div>
        </div>
      </div>
      <div class="row" style="min-height: 32rem;">
        <transition :name="transitionName" mode="out-in">
          <router-view></router-view>
        </transition>
      </div>
    </div>
    <div class="version">
      <p class="version">Connected to: {{network}}<br />
      Molly Wallet version: {{uiVersion}}</p>
    </div>
  </div>
</template>

<script>
import {mapState} from 'vuex'
import BrightBG from '../../assets/img/nodes2.jpg';
import DarkBG from '../../assets/img/nodes2_dark.jpg';

export default {
  data: () => ({
    transitionName: ""
  }),
  watch: {
    '$route' (from, to) {
      if (
        to.name == "download" ||
        to.name == "new wallet" ||
        from.name == "new wallet" && to.name == "login"
      ) {
        this.transitionName = ""
      } else {
        this.transitionName = "fade"
      }
    }
  },
  computed: {
    themeBG: function () {
      if (this.isDarkMode) {
          return DarkBG;
        } else {
          return BrightBG;
        }
    },
    isDarkMode: function () {
      if (this.darkMode || this.$route.params.darkMode) {
        return true
      } else {
        return false
      }
    },
    ...mapState('app', ['displayLoginError', 'loginErrorMsg', 'network', 'uiVersion']),
    ...mapState('wallet', ['darkMode'])
  }
};
</script>

<style scoped lang="scss">
.login-bg {
  /* Full height */
  height: 100%;
  position: absolute;
  width: 100%;
  overflow: hidden;

  /* Center and scale the image nicely */
  background-position: center;
  background-repeat: no-repeat;
  background-size: cover;
}

.vertical-center {
  min-height: 100%; /* Fallback for browsers do NOT support vh unit */
  min-height: 100vh; /* These two lines are counted as one :-)       */
  display: flex;
  align-items: center;
}

.fade-enter-active,
.fade-leave-active {
  transition-duration: 1s;
  transition-property: opacity;
  transition-timing-function: ease;
}

.fade-enter,
.fade-leave-active {
  opacity: 0
}

.version {
  width: 10rem;
  position: fixed;
  bottom:0;
  right:0;
  font-size: 0.7rem;
  display: flex;
  align-items: bottom;
  margin-right: 1.8em;
}

.page-error-box {
  height: 1.875em;
  padding-bottom: 0.625rem;
}

.page-error-box p {
    font-size: 0.75rem;
}
</style>
