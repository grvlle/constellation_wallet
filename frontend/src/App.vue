<template>
  <div>
    <login-screen
      v-if="!this.$store.state.app.isLoggedIn"
      :isLoggedIn="!this.$store.state.app.isLoggedIn"
    />
    <loading-screen
      v-if="this.$store.state.app.isLoading && this.$store.state.app.isLoggedIn"
      :isLoading="this.$store.state.app.isLoading"
      :fadeout="!this.$store.state.app.isLoading"
    />

    <div
      v-if="!this.$store.state.app.isLoading && this.$store.state.app.isLoggedIn"
      :class="{'nav-open': $sidebar.showSidebar}"
    >
      <notifications v-if="!this.$store.state.app.isLoading && this.$store.state.app.isLoggedIn"></notifications>
      <router-view v-if="!this.$store.state.app.isLoading && this.$store.state.app.isLoggedIn"></router-view>
    </div>
  </div>
</template>

<script>
import ErrorNotification from "./pages/Notifications/ErrorMessage";
import LoadingScreen from "./pages/LoadingScreen";
import LoginScreen from "./pages/Login";

export default {
  components: {
    LoadingScreen,
    LoginScreen
  },
  data() {
    return {};
  },

  mounted() {
    // Backend Errors
    window.wails.Events.On("error_handling", (m, err) => {
      this.$store.state.errorMessage = m + err;
      setTimeout(() => {
        this.$notifications.clear();
      }, 60000);
      this.$notify({
        component: ErrorNotification,
        timeout: 500000,
        icon: "fa fa-times",
        horizontalAlign: "right",
        verticalAlign: "top",
        type: "danger",
        onClick: () => {
          this.$notifications.clear();
        }
      });
    });

    window.wails.Events.On("login_error", (m, err) => {
      this.$store.state.loginErrorMsg = m;
      this.$store.state.displayLoginError = err;
      setTimeout(() => {
        this.$store.state.displayLoginError = false;
      }, 6000);
    });
    // Transactions.vue sockets
    window.wails.Events.On("update_tx_history", txHistoryFull => {
      this.$store.state.txInfo.txHistory = txHistoryFull;
      //this.$store.commit('updateFullTxHistory', txHistoryFull)
    });
    window.wails.Events.On("new_transaction", txObject => {
      this.$store.commit("updateTxHistory", txObject);
    });

    // Login.vue sockets
    window.wails.Events.On("registeredLogin", event => {});

    // Dashboard.vue sockets
    window.wails.Events.On("token", amount => {
      this.$store.state.walletInfo.tokenAmount = amount;
    });
    window.wails.Events.On("blocks", number => {
      this.$store.state.walletInfo.blocks = number;
    });
    window.wails.Events.On("price", (currency, dagRate) => {
      this.$store.state.walletInfo.usdValue = currency + " " + dagRate;
    });
    window.wails.Events.On("token_counter", count => {
      this.$store.state.counters.tokenCounter = count;
    });
    window.wails.Events.On("block_counter", blockCount => {
      this.$store.state.counters.blockCounter = blockCount;
    });
    window.wails.Events.On("chart_counter", pieChartCount => {
      this.$store.state.counters.nodesOnlineCounter = pieChartCount;
    });
    window.wails.Events.On("node_stats", (series, labels) => {
      this.$store.state.chartData.nodesOnline.series = series;
      this.$store.state.chartData.nodesOnline.labels = labels;
    });
    window.wails.Events.On("tx_stats", (seriesOne, seriesTwo, labels) => {
      this.$store.state.chartData.transactions.series = [seriesOne, seriesTwo];
      this.$store.state.chartData.transactions.labels = labels;
    });
    window.wails.Events.On("network_stats", (seriesOne, seriesTwo, labels) => {
      this.$store.state.chartData.throughput.series = [seriesOne, seriesTwo];
      this.$store.state.chartData.throughput.labels = labels;
    });

    // Settings.vue sockets
    window.wails.Events.On("wallet_keys", (address) => {
        // this.$store.state.walletInfo.keystorePath = keystorePath;
        this.$store.state.walletInfo.address = address;
      }
    );
  }
};
</script>

<style lang="scss">
.vue-notifyjs.notifications {
  .alert {
    z-index: 10000;
  }
  .list-move {
    transition: transform 0.3s, opacity 0.4s;
  }
  .list-item {
    display: inline-block;
    margin-right: 10px;
  }
  .list-enter-active {
    transition: transform 0.2s ease-in, opacity 0.4s ease-in;
  }
  .list-leave-active {
    transition: transform 1s ease-out, opacity 0.4s ease-out;
  }
  .list-enter {
    opacity: 0;
    transform: scale(1.1);
  }
  .list-leave-to {
    opacity: 0;
    transform: scale(1.2, 0.7);
  }
}

.fadeout {
  animation: fadeout 2s backwards;
}

@keyframes fadeout {
  to {
    opacity: 0;
    visibility: hidden;
  }
}
</style>