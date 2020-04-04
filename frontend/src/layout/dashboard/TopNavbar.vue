<template>
  <nav class="navbar navbar-expand-lg navbar-light">
    <div class="container-fluid">
      <img class="img-fluid"
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
            <a class="nav-link" @click="test" style="cursor: pointer;">
              <i class="fa fa-check"></i>TxEvent
            </a>
          </li>
          <li class="nav-item">
            <a class="nav-link" @click="test2" style="cursor: pointer;">
              <i class="fa fa-check"></i>SysEvent
            </a>
          </li>
          <drop-down class="nav-item" title-classes="nav-link">
            <template v-slot:title>
              <i class="ti-bell"></i>
              <span>
                Notifications
                <span v-if="(txNotificationCount + systemNotificationCount) > 0" class="badge badge-danger" style="width:0.875rem;">
                  {{(txNotificationCount + systemNotificationCount)}}
                </span>
                <span v-else class="badge invisible" style="width:0.875rem;">0</span>
              </span>
            </template>
            <div v-if="(txNotificationCount + systemNotificationCount) == 0" class="dropdown-item inactive text-center">
              No new notifications available
            </div>
            <li class="dropdown-header" style="display: flex;" v-if="this.$store.state.notificationInfo.txNotifications.length > 0">
              <div style="margin-right: auto;">Transactions</div>
              <div class="dropdown-header-link" @click="markTxNotficationsAsRead">mark as read</div>
              <div class="dropdown-header-link" @click="clearTxNotfications">clear</div>
            </li>
            <div v-for="n in this.$store.state.notificationInfo.txNotifications" v-bind:key="n.id" class="dropdown-item" 
              :class="n.read ? 'inactive' : 'active'" style="display: flex;" @click="markTxNotificationAsRead(n)">
              <div>
                {{n.tx.datetime}} {{n.description}}
                <p class="small">{{n.tx.hash}}</p>
              </div>
              <div style="margin-top: 0.45rem; margin-left: auto;">
                <i v-if="n.tx.status === 'Complete'" class="fa fa-check"></i>
                <i v-if="n.tx.status === 'Pending'" class="ti-timer"></i>
                <i v-if="n.tx.status === 'Error'" class="fa fa-times"></i>
              </div>
            </div>
            <li class="dropdown-header" style="display: flex;" v-if="this.$store.state.notificationInfo.systemNotifications.length > 0">
              <div style="margin-right: auto;">System</div>
              <div class="dropdown-header-link" @click="markSystemNotificationsAsRead">mark as read</div>
              <div class="dropdown-header-link" @click="clearSystemNotfications">clear</div>
            </li>
            <div v-for="n in this.$store.state.notificationInfo.systemNotifications" v-bind:key="n.id" class="dropdown-item" 
              :class="n.read ? 'inactive' : 'active'" style="display: flex;" @click="markSystemNotificationAsRead(n)">
              <div>
                <div>{{n.description}}</div>
                <p class="small">...</p>
              </div>
              <div style="margin-top: 0.45rem; margin-left: auto;">
                <i class="fa fa-exclamation-circle"></i>
              </div>
            </div>
          </drop-down>
          <li class="nav-item">
            <router-link class="nav-link" to="/settings">
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
export default {
  computed: {
    routeName() {
      const { name } = this.$route;
      return this.capitalizeFirstLetter(name);
    },
    txNotificationCount () {
      let unreadTxNotifications = this.$store.state.notificationInfo.txNotifications.filter(
        function (e) {return e.read == false;}
      );
      return unreadTxNotifications.length;
    },
    systemNotificationCount () {
      let unreadSystemNotifications = this.$store.state.notificationInfo.systemNotifications.filter(
        function (e) {return e.read == false;}
      );
      return unreadSystemNotifications.length;
    }
  },
  watch: {
    txNotificationCount (newCount, oldCount) {
      if (newCount > oldCount) {
        let latestItem = this.$store.state.notificationInfo.txNotifications[this.$store.state.notificationInfo.txNotifications.length - 1]
        this.showNotification(latestItem.type, latestItem.description, "info");
      }
    },
    systemNotificationCount (newCount, oldCount) {
      if (newCount > oldCount) {
        let latestItem = this.$store.state.notificationInfo.systemNotifications[this.$store.state.notificationInfo.systemNotifications.length - 1]
        this.showNotification(latestItem.type, latestItem.description, "info");
      }
    }
  },
  data() {
    return {
      activeNotifications: false,
      random: ""
    };
  },
  methods: {
    logout() {
      window.backend.WalletApplication.LogOut().then(txFinishedState => {
        if (txFinishedState) {
          this.$store.state.txInfo.txHistory = [];
          this.$store.state.walletInfo.keystorePath = "";
          this.$store.state.walletInfo.alias = "";
          this.$store.state.walletInfo.keystorePassword = "";
          this.$store.state.walletInfo.KeyPassword = "";
          this.$store.state.walletInfo.email = "";
          this.$store.state.walletInfo.totalValue = 0;
          this.$store.state.walletInfo.tokenAmount = 0;
          this.$store.state.app.isLoading = false;
          this.$store.state.app.isLoggedIn = false;
          this.$store.state.app.register = false;
          this.$store.state.app.import = false;
          this.$store.state.app.login = true;
          return;
        }
      }),
        (this.random = "1");
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
    },
    showNotification(title, message, type) {
      setTimeout(() => {
        this.$notifications.clear();
      }, 3000);
      this.$notify({
        title: title,
        message: message,
        icon: "fa fa-bell",
        horizontalAlign: "right",
        verticalAlign: "bottom",
        type: type,
        onClick: () => {
          this.$notifications.clear();
        }
      });
    },
    markTxNotificationAsRead(notification) {
      notification.read = true;
      this.$store.commit("updateTxNotification", notification);
    },
    markTxNotficationsAsRead() {
      this.$store.commit("updateTxNotificationsAsRead")
    },
    clearTxNotfications() {
      this.$store.commit("deleteTxNotifications")
    },
    markSystemNotificationAsRead(notification) {
      notification.read = true;
      this.$store.commit("updateSystemNotification", notification);
    },
    markSystemNotificationsAsRead() {
      this.$store.commit("updateSystemNotificationsAsRead")
    },
    clearSystemNotfications() {
      this.$store.commit("deleteSystemNotifications")
    },
    test() {
      this.$store.commit("addTxNotification", 
        {
          id: Math.random(),
          type:"transaction",
          description:"Awaiting confirmation from mainnet", 
          read: false,
          tx: {
            datetime: "Mon Jan _2 15:04:05 2006",
            hash: "7ed02e54dd824ce425752c9c03bdd27f11c05ecd5c36741a311d8020759f54e9",
            status: "Pending"
          }
        }
      );
    },
    test2() {
      this.$store.commit("addSystemNotification", 
        {
          id: Math.random(),
          type:"system",
          description:"Some system error has occured.", 
          read: false
        }
      );
    }
  }
};
</script>

<style>
</style>
