<template>
  <nav class="navbar navbar-expand-lg navbar-light">
    <div class="container-fluid">
      <img class="img-fluid" v-if="darkMode"
        src="~@/assets/img/Constellation-Logo-White.png"
        style="max-height: 6.25rem; max-width: 12.5rem; margin-left: 2rem;" />
      <img class="img-fluid" v-else
        src="~@/assets/img/Constellation-Logo-Black.png"
        style="max-height: 6.25rem; max-width: 12.5rem; margin-left: 2rem;" />
      <button
        class="navbar-toggler navbar-burger"
        type="button"
        @click="toggleSidebar"
        :aria-expanded="$sidebar.showSidebar"
        aria-label="Toggle navigation" >
        <span class="navbar-toggler-bar"></span>
        <span class="navbar-toggler-bar"></span>
        <span class="navbar-toggler-bar"></span>
      </button>
      <div class="collapse navbar-collapse">
        <ul class="navbar-nav ml-auto">
          <li class="nav-item">
            <router-link class="nav-link" to="settings">
              <i class="ti-settings"></i>
              <p class="nav-item">SETTINGS</p>
            </router-link>
          </li>
          <li class="nav-item">
            <a class="nav-link" @click="logout" style="cursor: pointer;">
              <i class="ti-lock"></i>
              <p class="nav-item">LOGOUT</p>
            </a>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script>
import {mapState} from 'vuex'
export default {
  computed: {
    routeName() {
      const { name } = this.$route;
      return this.capitalizeFirstLetter(name);
    },
    ...mapState('wallet', ['darkMode'])
  },
  data() {
    return {
      activeNotifications: false,
      random: ""
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
    }
  }
};
</script>
<style scoped lang="scss">
</style>
