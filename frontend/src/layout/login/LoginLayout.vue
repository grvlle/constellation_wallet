<template>
  <div id="app" class="login-bg vertical-center" v-bind:style="{ backgroundImage: 'url(' + this.bgImg + ')' }" >
    <div class="container">
      <div class="row">
        <div class="col mx-auto text-center header">
          <div>
            <img class="img-fluid" v-if="this.$store.state.walletInfo.darkMode" src="~@/assets/img/Constellation-Logo-White.png" style="max-height: 5.8rem;" />
            <img class="img-fluid" v-else src="~@/assets/img/Constellation-Logo-Black.png" style="max-height: 5.8rem;" />
            <p v-if="this.$route.params.message">{{this.$route.params.message}}</p>
            <p v-else>Downloading $DAG wallet dependencies...</p>
          </div>
          <div class="page-error-box" v-if="this.$store.state.displayLoginError">
            <p>{{this.$store.state.loginErrorMsg}}</p>
          </div>
          <div class="page-error-box" v-else></div>
        </div>
      </div>
      <div class="row" style="min-height: 40rem;">
        <transition :name="transitionName" mode="out-in">
          <router-view></router-view>
        </transition>
      </div>
    </div>
  </div>
</template>

<script>
import BrightBG from '../../assets/img/nodes2.jpg';
import DarkBG from '../../assets/img/nodes2_dark.jpg';
export default {
  data: () => ({
    bgImg: DarkBG,
    transitionName: ""
  }),
  mounted() {
    this.themeBG()
  },
  watch: {
    '$route' (from, to) {
      if (
        to.name == "download" ||
        to.name == "new wallet" ||
        from.name == "new wallet"
      ) {
        this.transitionName = ""
      } else {
        this.transitionName = "fade"
      }
    }
  },
  methods: {
    themeBG: function () {
      if (this.$store.state.walletInfo.darkMode) {
          this.bgImg = DarkBG;
        } else {
          this.bgImg = BrightBG;
        }
    },
  }
};
</script>

<style>
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
</style>
