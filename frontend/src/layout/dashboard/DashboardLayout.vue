<template>
  <div class="wrapper">
    <side-bar>
      <template slot="links">
        <sidebar-link to="/dashboard" name="Dashboard" icon="fa fa-columns" />
        <sidebar-link
          to="/wallet-info"
          name="Wallet Information"
          icon="fa fa-wallet"
        />
        <sidebar-link
          to="/submit-transaction"
          name="Transactions"
          icon="fa fa-paper-plane"
        />
        <!-- <sidebar-link to="/typography" name="Typography" icon="ti-text"/>
        <sidebar-link to="/icons" name="Icons" icon="ti-pencil-alt2"/>
        <sidebar-link to="/maps" name="Map" icon="ti-map"/>-->
        <sidebar-link to="/about" name="About" icon="fa fa-question-circle" />

        <p class="nav-item">
          <a @click="notifyVue2('top', 'right')" class="nav-link"
            ><i class="fa fa-trophy"></i>The League</a
          >
        </p>
        <p class="nav-item">
          <a @click="notifyVue2('top', 'right')" class="nav-link"
            ><i class="fa fa-gavel"></i>Governance</a
          >
        </p>
      </template>
      <mobile-menu>
        <li class="nav-item">
          <a class="nav-link">
            <i class="ti-panel"></i>
            <p>Stats</p>
          </a>
        </li>
        <drop-down
          class="nav-item"
          title="Notifications"
          title-classes="nav-link"
          icon="ti-bell"
        >
          <a class="dropdown-item">Notification 1</a>
          <a class="dropdown-item">Notification 2</a>
          <a class="dropdown-item">Notification 3</a>
          <a class="dropdown-item">Notification 4</a>
          <a class="dropdown-item">Another notification</a>
        </drop-down>
        <li class="nav-item">
          <a class="nav-link">
            <i class="ti-settings"></i>
            <p>Settings</p>
          </a>
        </li>
        <li class="divider"></li>
      </mobile-menu>
    </side-bar>
    <div class="main-panel">
      <top-navbar></top-navbar>

      <dashboard-content @click.native="toggleSidebar"> </dashboard-content>

      <!-- <content-footer></content-footer> -->
    </div>
  </div>
</template>
<style lang="scss"></style>
<script>
import TopNavbar from "./TopNavbar.vue";
// import ContentFooter from "./ContentFooter.vue";
import DashboardContent from "./Content.vue";
import MobileMenu from "./MobileMenu";
import PathBlockedNotification from "../../pages/Notifications/PathBlocked.vue";
export default {
  components: {
    TopNavbar,
    // ContentFooter,
    DashboardContent,
    MobileMenu
  },
  methods: {
    notifyVue2(verticalAlign, horizontalAlign) {
      setTimeout(() => {
        this.$notifications.clear();
      }, 6000);
      this.$notify({
        component: PathBlockedNotification,
        icon: "fa fa-lock",
        horizontalAlign: horizontalAlign,
        verticalAlign: verticalAlign,
        type: "danger",
        onClick: () => {
          this.$notification.close();
        }
      });
    },
    toggleSidebar() {
      if (this.$sidebar.showSidebar) {
        this.$sidebar.displaySidebar(false);
      }
    }
  }
};
</script>
