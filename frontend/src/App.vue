<template>
  <div :class="this.$store.state.walletInfo.darkMode ? 'theme--dark' : 'theme--light'">
    <vue-progress-bar></vue-progress-bar>
    <div :class="{'nav-open': $sidebar.showSidebar}">
      <notifications></notifications>
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import ErrorNotification from "./pages/Notifications/ErrorMessage";
import WarningNotification from "./pages/Notifications/Warning";
import SuccessNotification from "./pages/Notifications/Success";

export default {
  components: {
  },
  data() {
    return {};
  },

  mounted() {
    // Backend Errors
    window.wails.Events.On("error_handling", (m, err) => {
      this.$store.state.errorMessage = m + err;
      this.$notifications.clear();
      setTimeout(() => {
        this.$notifications.clear();
      }, 60000);
      this.$notify({
        component: ErrorNotification,
        timeout: 50000,
        icon: "fa fa-times",
        horizontalAlign: "right",
        verticalAlign: "top",
        type: "danger",
        onClick: () => {
          this.$notifications.clear();
        }
      });
    });

    window.wails.Events.On("warning", m => {
      this.$store.state.warningMessage = m;
      this.$notifications.clear();
      setTimeout(() => {
        this.$notifications.clear();
      }, 60000);
      this.$notify({
        component: WarningNotification,
        timeout: 50000,
        icon: "fa fa-times",
        horizontalAlign: "right",
        verticalAlign: "top",
        type: "warning",
        onClick: () => {
          this.$notifications.clear();
        }
      });
    });

    window.wails.Events.On("success", m => {
      this.$store.state.successMessage = m;
      this.$notifications.clear();
      setTimeout(() => {
        this.$notifications.clear();
      }, 60000);
      this.$notify({
        component: SuccessNotification,
        timeout: 50000,
        icon: "fa fa-check",
        horizontalAlign: "right",
        verticalAlign: "top",
        type: "success",
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
      }, 10000);
    });
    // Transactions.vue sockets
    window.wails.Events.On("update_tx_history", txHistoryFull => {
      this.$store.state.txInfo.txHistory = txHistoryFull;
      //this.$store.commit('updateFullTxHistory', txHistoryFull)
    });
    window.wails.Events.On("tx_in_transit", txFinished => {
      this.$store.state.app.txFinished = txFinished;
    });
    window.wails.Events.On("new_transaction", txObject => {
      this.$store.commit("updateTxHistory", txObject);
    });
    window.wails.Events.On("tx_pending", txStatus => {
      this.$store.state.txInfo.txStatus = txStatus;
    });

    // Downloading.vue sockets
    window.wails.Events.On(
      "downloading_dependencies",
      isDownloadingDependencies => {
        this.$store.state.app.isDownloadingDependencies = isDownloadingDependencies;
      }
    );

    window.wails.Events.On("downloading", (filename, size) => {
      if (this.$store.state.downloading.filename !== filename) {
        this.$store.state.downloading.filename = filename;
      }
      this.$store.state.downloading.size = size;
    });

    // Login.vue sockets
    // window.wails.Events.On("registeredLogin", event => {});

    // Dashboard.vue sockets
    window.wails.Events.On("token", (amount, available, total) => {
      this.$store.state.walletInfo.tokenAmount = amount;
      this.$store.state.walletInfo.availableBalance = available;
      this.$store.state.walletInfo.totalBalance = total;
    });
    window.wails.Events.On("blocks", number => {
      this.$store.state.walletInfo.blocks = number;
    });
    window.wails.Events.On("totalValue", (currency, value) => {
      this.$store.state.walletInfo.currency = currency;
      this.$store.state.walletInfo.totalValue = value;
    });
    window.wails.Events.On("token_counter", count => {
      this.$store.state.counters.tokenCounter = count;
    });
    window.wails.Events.On("value_counter", valueCount => {
      this.$store.state.counters.valueCounter = valueCount;
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
    window.wails.Events.On("wallet_keys", address => {
      // this.$store.state.walletInfo.keystorePath = keystorePath;
      this.$store.state.walletInfo.address = address;
    });
  }
};
</script>

<style lang="scss">
.vue-notifyjs.notifications {
  .alert {
    z-index: 10000;
    font-size: 0.875rem;
  }
  .alert[data-notify="container"] {
    width: 21.875rem;
  }
  .alert-icon {
    margin-left: -0.5em;
    margin-top: -0.5em;
  }
  .list-move {
    transition: transform 0.3s, opacity 0.4s;
  }
  .list-item {
    display: inline-block;
    margin-right: 0.625em;
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
