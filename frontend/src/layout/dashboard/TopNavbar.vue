<template>
  <nav class="navbar navbar-expand-lg navbar-light">
    <div class="container-fluid">
      <!-- <img
        class="img-fluid"
        v-if="darkMode"
        src="~@/assets/img/Constellation-Logo-White.png"
        style="max-height: 6.25rem; max-width: 12.5rem; margin-left: 2rem"
      />
      <img
        class="img-fluid"
        v-else
        src="~@/assets/img/Constellation-Logo-Black.png"
        style="max-height: 6.25rem; max-width: 12.5rem; margin-left: 2rem"
      /> -->
      <div class="logo">
        <img
          style="height: 54.23px;"
          v-if="darkMode"
          src="~@/assets/img/stardust-collective-logo-white.png"
        />
        <img
          style="height: 54.23px;"
          v-else
          src="~@/assets/img/stardust-collective-logo-black.png"
        />
      </div>
      <!--      <button-->
      <!--        class="navbar-toggler navbar-burger"-->
      <!--        type="button"-->
      <!--        @click="toggleSidebar"-->
      <!--        :aria-expanded="$sidebar.showSidebar"-->
      <!--        aria-label="Toggle navigation"-->
      <!--      >-->
      <!--        <span class="navbar-toggler-bar"></span>-->
      <!--        <span class="navbar-toggler-bar"></span>-->
      <!--        <span class="navbar-toggler-bar"></span>-->
      <!--      </button>-->
      <div class="collapse navbar-collapse">
        <ul class="navbar-nav ml-auto">
          <!--          <li class="nav-item">-->
          <!--            <div class="testnet-toggle">-->
          <!--              <toggle-switch v-model="onTestnet" />-->
          <!--              <p class="nav-item nav-link">MAINNET / TESTNET</p>-->
          <!--            </div>-->
          <!--          </li>-->
          <li class="nav-item">
            <router-link class="nav-link" to="settings">
              <settings-icon />
              <p class="nav-item">SETTINGS</p>
            </router-link>
          </li>
          <li class="nav-item">
            <a class="nav-link" @click="logout" style="cursor: pointer">
              <logout-icon />
              <p class="nav-item">LOGOUT</p>
            </a>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script>
import { mapState } from "vuex";
// import ToggleSwitch from "../../components/Inputs/ToggleSwitch";

export default {
  components: {
    // ToggleSwitch,
  },
  computed: {
    routeName() {
      const { name } = this.$route;
      return this.capitalizeFirstLetter(name);
    },
    ...mapState("wallet", ["darkMode"]),
    ...mapState("app", ["onTestnet"]),
    onTestnet: {
      get: function() {
        return this.$store.state.app.onTestnet;
      },
      set: function() {
        this.$store.dispatch("transaction/reset").then(() => {
          this.switch = !this.switch;
          window.backend.WalletApplication.SelectNetwork(this.switch);
          this.$store.commit("app/setOnTestnet", this.switch);
        });
      },
    },
  },
  data() {
    return {
      activeNotifications: false,
      switch: false,
      random: "",
    };
  },
  methods: {
    capitalizeFirstLetter(string) {
      return string.charAt(0).toUpperCase() + string.slice(1);
    },
    toggleNotificationDropDown() {
      this.activeNotifications = !this.activeNotifications;
    },
    closeDropDown() {
      this.activeNotifications = false;
    },
    toggleSidebar() {
      this.$sidebar.displaySidebar(!this.$sidebar.showSidebar);
    },
    hideSidebar() {
      this.$sidebar.displaySidebar(false);
    },
  },
};
</script>
<style scoped lang="scss">
.logo {
  margin-left: 24px;
}

.toggle {
  margin: 0 0 0 0;
}

.testnet-toggle {
  display: flex;
  margin-top: 2%;
  align-items: center;
}

.navbar .nav-link {
  opacity: 1 !important;
  p {
    font-weight: 500;
    font-size: 12px;
    line-height: 12px;
    letter-spacing: 0.1em;
    text-transform: uppercase;
    margin-left: 6px;
  }
  color: #666666 !important;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
