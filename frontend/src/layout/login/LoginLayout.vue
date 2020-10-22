<template>
  <div id="app" class="login-bg vertical-center">
    <link
      href="//maxcdn.bootstrapcdn.com/font-awesome/4.1.0/css/font-awesome.min.css"
      rel="stylesheet"
    />
    <link
      href="https://fonts.googleapis.com/css?family=Lato"
      rel="stylesheet"
      type="text/css"
    />
    <div class="container">
      <i class="fa fa-angle-left go-back-btn" @click="$router.go(-1)"></i>
      <div class="row">
        <div class="header-box">
          <p
            class="header-title"
            v-if="this.$route.params.title && !isMigrationWizard"
          >
            {{ this.$route.params.title }}
          </p>

          <p
            class="header-title"
            style="color: #dd8d74"
            v-else-if="isMigrationWizard"
          >
            {{ this.$route.params.title }}
          </p>
          <p class="header-title" v-else>Welcome to Molly Wallet 2.0</p>

          <p class="sub-title" v-if="this.$route.params.message">
            {{ this.$route.params.message }}
          </p>
          <p class="sub-title" v-else>
            Downloading $DAG wallet dependencies...
          </p>
          <div class="page-error-box text-danger" v-if="displayLoginError">
            <p>{{ loginErrorMsg }}</p>
          </div>
        </div>
      </div>
      <div class="row">
        <transition :name="transitionName" mode="out-in">
          <router-view></router-view>
        </transition>
      </div>
    </div>
    <div class="logo">
      <img
        class="img-fluid"
        v-if="isDarkMode"
        src="~@/assets/img/Constellation-Logo-White.png"
        style="max-height: 2rem"
      />
      <img class="img-fluid" v-else src="~@/assets/img/logo.png" />
    </div>
    <div class="version">
      <p class="version">
        Connected to:
        {{ network }}<br />
        Molly wallet version: {{ uiVersion }}
      </p>
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";
import BrightBG from "../../assets/img/nodes2.jpg";
import DarkBG from "../../assets/img/nodes2_dark.jpg";

export default {
  data: () => ({
    transitionName: "",
  }),
  watch: {
    $route(from, to) {
      if (
        to.name === "download" ||
        to.name === "new wallet" ||
        (from.name === "new wallet" &&
          (to.name === "login" || to.name === "login single password"))
      ) {
        this.transitionName = "";
      } else {
        this.transitionName = "fade";
      }
    },
  },
  computed: {
    isMigrationWizard() {
      return (
        this.$route.name === "keystore migrate" ||
        this.$route.name === "keystore migration complete" ||
        this.$route.name === "password migration"
      );
    },
    themeBG: function() {
      if (this.isDarkMode) {
        return DarkBG;
      } else {
        return BrightBG;
      }
    },
    isDarkMode: function() {
      if (this.darkMode || this.$route.params.darkMode) {
        return true;
      } else {
        return false;
      }
    },
    ...mapState("app", [
      "displayLoginError",
      "loginErrorMsg",
      "network",
      "uiVersion",
    ]),
    ...mapState("wallet", ["darkMode"]),
  },
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
  opacity: 0;
}

.version {
  width: 10rem;
  position: fixed;
  bottom: 0;
  right: 0;
  font-size: 0.7rem;
  text-align: right;
  font-family: Poppins;
  margin-right: 24px;
  margin-bottom: 24px;
}

.go-back-btn {
  width: 10rem;
  position: fixed;
  color: #b9b9b9;
  top: 0;
  left: 0;
  font-size: 2.7rem;
  display: flex;
  align-items: flex-end;
  margin-left: 0.5em;
  margin-top: 0.5em;
}

.go-back-btn:hover {
  color: #ce9483;
  cursor: pointer;
}

.header-title {
  color: #2d9cdb;
  font-family: Poppins;
  font-weight: 500;
  font-size: 28px;
  line-height: 24px;
  margin-bottom: 48px;
}

.sub-title {
  color: #666666;
  font-family: Poppins;
  font-weight: 500;
  font-size: 14px;
  line-height: 24px;

  &:last-child {
    margin-bottom: 48px;
  }
}

.header-box {
  margin: auto;
  height: fit-content;
  max-width: 27rem;
  min-width: 27rem;
}

.logo {
  width: 10rem;
  position: fixed;
  bottom: 0;
  right: 0;
  font-size: 0.7rem;
  display: flex;
  bottom: 0;
  left: 0;
  margin: 1.8em;
}

.page-error-box {
  height: 1.875em;
  padding-bottom: 0.625rem;
}

.page-error-box p {
  font-size: 0.75rem;
}
</style>
