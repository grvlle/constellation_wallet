<template>
  <nav class="navbar navbar-expand-lg navbar-light">
    <div class="container-fluid">
      <img
        src="https://constellationlabs.io/wp-content/uploads/2019/08/Constellation-Logo-1.png"
        style="max-height: 100px; max-width: 200px;"
      />
      <!-- <a class="navbar-brand">{{routeName}}</a> -->
      <button
        class="navbar-toggler navbar-burger"
        type="button"
        @click="toggleSidebar"
        :aria-expanded="$sidebar.showSidebar"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-bar"></span>
        <span class="navbar-toggler-bar"></span>
        <span class="navbar-toggler-bar"></span>
      </button>
      <div class="collapse navbar-collapse">
        <ul class="navbar-nav ml-auto">
          <!-- <p class="nav-item">
            <drop-down
              class="nav-item"
              title="Logout"
              title-classes="nav-link"
              icon="ti-lock"
            >
              <a class="dropdown-item">Empty</a>
            </drop-down>
          </p>-->
          <li class="nav-item">
            <router-link class="nav-link" to="/settings">
              <i class="ti-settings"></i>
              <p class="nav-item">SETTINGS</p>
            </router-link>
          </li>

          <slot>
            <li class="nav-item">
              <a class="nav-link" @click="logout">
                <i class="ti-lock"></i>
                <p class="nav-item">LOGOUT</p>
              </a>
            </li>
          </slot>
        </ul>
      </div>
    </div>
  </nav>
</template>
<script>
export default {
  computed: {
    routeName() {
      const { name } = this.$route;
      return this.capitalizeFirstLetter(name);
    }
  },
  data() {
    return {
      activeNotifications: false,
      random: "",
    };
  },
  methods: {
    logout() {
      window.backend.WalletApplication.LogOut().then(txFinishedState => {
        if (txFinishedState) {
          this.$store.state.txInfo.txHistory = [];
          this.$store.state.app.isLoading = false;
          this.$store.state.app.isLoggedIn = false;
          this.$store.state.app.register = false;
          this.$store.state.app.import = false;
          this.$store.state.app.login = true;
          this.$store.state.app.margin = 70;
          return;
        }
      }),
      this.random = "1"
      return;
        
    },
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
    }
  }
};
</script>
<style>
</style>
