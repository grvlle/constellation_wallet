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
          <p class="nav-item">
            <drop-down
              class="nav-item"
              title="Notifications"
              title-classes="nav-link"
              icon="ti-bell"
            >
              <a class="dropdown-item">Empty</a>
            </drop-down>
          </p>
          
          <slot>
          <li class="nav-item">
              
                <router-link class="nav-link" to="/settings">
                  <i class="ti-settings"></i>
                  <p class="nav-item">SETTINGS</p>
                </router-link>
              
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
      activeNotifications: false
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
<style>
</style>
