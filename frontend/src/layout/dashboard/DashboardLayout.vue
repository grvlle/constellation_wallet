<template>
  <div id="app" class="wrapper">
    <side-bar>
      <template slot="links">
        <sidebar-link to="dashboard" name="Dashboard" icon="fa fa-columns" />
        <sidebar-link to="wallet-info" name="Wallet Info" icon="fa fa-wallet" />
        <sidebar-link to="submit-transaction" name="Transactions" icon="fa fa-paper-plane" />
        <sidebar-link to="about" name="About" icon="fa fa-question-circle" />
        <sidebar-link to="address-book" name="Address Book" icon="fa fa-address-book" />
        <p class="nav-item">
          <a @click="notifyVue2('top', 'right')" class="nav-link" style="cursor: pointer; color: #6c757d;">
            <i class="fa fa-trophy"></i>Team
          </a>
        </p>
        <p class="nav-item">
          <a @click="notifyVue2('top', 'right')" class="nav-link" style="cursor: pointer; color: #6c757d;">
            <i class="fa fa-gavel"></i>Governance
          </a>
        </p>
      </template>
      <ul class="nav text-center" style="width: 100%; bottom: 0;position: absolute;">
        <hr class="divider"/>
        <li class="nav-item">
          <p class="nav-item" style="font-size: 0.6rem;">
            <a href="javascript:void(0)" @click="toTermsofService" class="nav-link" style="cursor: pointer;">
              Terms of Service
            </a>
          </p>
        </li>
      </ul>
      <mobile-menu>
        <li class="nav-item">
          <a class="nav-link" @click="logout" style="cursor: pointer;">
            <i class="ti-lock"></i>
            <p>LOGOUT</p>
          </a>
        </li>
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
      <dashboard-content @click.native="toggleSidebar"></dashboard-content>
    </div>
  </div>
</template>

<style scoped lang="scss">
</style>

<script>
import TopNavbar from "./TopNavbar.vue";
import DashboardContent from "./Content.vue";
import MobileMenu from "./MobileMenu";
import PathBlockedNotification from "../../pages/Notifications/PathBlocked.vue";

export default {
  components: {
    TopNavbar,
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
    },
    toTermsofService() {
      this.$router.push({
        name: 'terms of service'
      });
    }
  }
};
</script>
