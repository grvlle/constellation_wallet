<template>
  <div :class="darkMode ? 'theme--dark' : 'theme--light'">
    <vue-progress-bar></vue-progress-bar>
    <div :class="{'nav-open': $sidebar.showSidebar}">
      <notifications></notifications>
      <router-view></router-view>
    </div>
    <page-overlay text="Applying Update. Please wait..." :isActive="overlay" />
  </div>
</template>

<script>
import {mapState} from 'vuex'
import ErrorNotification from "./pages/Notifications/ErrorMessage";
import WarningNotification from "./pages/Notifications/Warning";
import SuccessNotification from "./pages/Notifications/Success";
import NewRelease from "./pages/Notifications/NewRelease";
import Swal from "sweetalert2/dist/sweetalert2";

export default {
  components: {
  },
  computed: {
    ...mapState('app', ['isLoggedIn', 'downloading']),
    ...mapState('wallet', ['darkMode'])
  },
  data() {
    return {
      overlay: false
    };
  },
  onIdle() {
    let timerInterval, closeInSeconds = 60
    const swalPopup = Swal.mixin({
      customClass: {
        container: this.darkMode
          ? "theme--dark"
          : "theme--light"
      }
    });
    if (this.isLoggedIn) {
      swalPopup.fire({
        title: "You have been idle for 5 minutes.",
        html: "To keep your Molly wallet safe from unauthorised access it will automatically logout in <b>10</b> seconds",
        showConfirmButton: false,
        showCancelButton: false,
        timer: closeInSeconds * 1000,
        timerProgressBar: true,
        onBeforeOpen: () => {
          Swal.showLoading();
          timerInterval = setInterval(() => {
            closeInSeconds--;
            const content = Swal.getContent();
            if (content) {
              const b = content.querySelector('b');
              if (b) {
                b.textContent = closeInSeconds;
              }
            }
          }, 1000)
        },
        onClose: () => {
          clearInterval(timerInterval)
        }
      }).then(() => {
        if (this.$store.state.idleVue.isIdle) {
          window.backend.WalletApplication.LogOut().then(txFinishedState => {
            if (txFinishedState) {
              this.logout();
            }
          }), (this.random = "1");
        }
      });
    }
  },
  mounted() {
    // Backend Errors
    window.wails.Events.On("error_handling", (m, err) => {
      this.$store.commit('app/setErrorMessage', m + err);
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
      this.$store.commit('app/setWarningMessage', m);
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
      this.$store.commit('app/setSuccessMessage', m);
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

    window.wails.Events.On("new_release", m => {
      this.$store.commit('app/setNewRelease', m);
      var self = this;
      this.$notifications.clear();
      this.$notify({
        component: NewRelease,
        timeout: 500000,
        icon: "fa fa-info",
        horizontalAlign: "right",
        verticalAlign: "bottom",
        type: "info",
        onClick: () => {
          const swalPopup = Swal.mixin({
            customClass: {
              container: this.darkMode
                ? "theme--dark"
                : "theme--light"
            }
            
          });
          
          swalPopup.fire({
            title: "Update Molly Wallet",
            html: "Do you want to update your Molly Wallet? Selecting update will download the latest build and apply the update. <br><br> <b>The application will restart once update is complete. </b>",
            showCloseButton: true,
            showCancelButton: true,
            focusConfirm: false,
            confirmButtonText:
              '<i class="fa fa-thumbs-up"></i> <b>Update</b>',
            confirmButtonAriaLabel: "Text",
            cancelButtonText:
              'Cancel',
            cancelButtonAriaLabel: "Cancel"
          }).then(result => {
            if (result.value) {
              self.$Progress.start(2000);
              self.overlay = true;
              window.backend.WalletApplication.UpdateMolly()
              self.$notifications.clear();
            }
          });      
        }
      });
    });

    window.wails.Events.On("login_error", (m, err) => {
      this.$store.commit('app/setLoginErrorMessage', m);
      this.$store.commit('app/setDisplayLoginError', err);
      setTimeout(() => {
        this.$store.commit('app/setDisplayLoginError', false);
      }, 10000);
    });

    // Transactions.vue sockets
    window.wails.Events.On("update_tx_history", txHistoryFull => {
      if (Object.entries(txHistoryFull).length != 0) {
        this.$store.commit({type: 'transaction/updateFullTxHistory', txHistoryFull});
      }
    });
    window.wails.Events.On("tx_in_transit", txFinished => {
      this.$store.commit('transaction/setTxFinished', txFinished);
    });
    window.wails.Events.On("new_transaction", txObject => {
      this.$store.commit("transaction/updateTxHistory", txObject);
    });
    window.wails.Events.On("tx_pending", txStatus => {
      this.$store.commit("transaction/updateTxStatus", txStatus);
    });

    window.wails.Events.On("downloading", (filename, size) => {
      if (this.downloading.filename !== filename) {
        this.$store.commit('app/setDownloadFileName', filename);
      }
      this.$store.commit('app/setDownloadFileSize', size);
    });

    window.wails.Events.On("token", (amount, available, total) => {
      this.$store.commit('wallet/setTokenAmount', amount);
      this.$store.commit('wallet/setAvailableBalance', available);
      this.$store.commit('wallet/setTotalBalance', total);
    });
    window.wails.Events.On("blocks", number => {
      this.$store.commit('dashboard/setBlocks', number);
    });
    window.wails.Events.On("tokenPrice", tokenPrice => {
      let rates = []
      rates.push({currency: "USD", tokenprice: tokenPrice.DAG.USD});
      rates.push({currency: "EUR", tokenprice: tokenPrice.DAG.EUR});
      rates.push({currency: "BTC", tokenprice: tokenPrice.DAG.BTC});
      this.$store.commit('wallet/setCurrencyRates', rates);
    });
    window.wails.Events.On("token_counter", count => {
      this.$store.commit('dashboard/setTokenCounter', count);
    });
    window.wails.Events.On("value_counter", valueCount => {
      this.$store.commit('dashboard/setValueCounter', valueCount);
    });
    window.wails.Events.On("block_counter", blockCount => {
      this.$store.commit('dashboard/setBlockCounter', blockCount);
    });
    window.wails.Events.On("chart_counter", pieChartCount => {
      this.$store.commit('dashboard/setChartCounter', pieChartCount);
    });
    window.wails.Events.On("node_stats", (series, labels) => {
      if (Object.entries(series).length != 0 && 
          Object.entries(labels).length != 0) {
        this.$store.commit({type: 'dashboard/setNodeOnlineChart', series, labels});
      }
    });
    window.wails.Events.On("tx_stats", (seriesOne, seriesTwo, labels) => {
      if (Object.entries(seriesOne).length != 0 && 
          Object.entries(seriesTwo).length != 0 &&
          Object.entries(labels).length != 0) {
        this.$store.commit({type: 'dashboard/setTransactionStatsChart', seriesOne, seriesTwo, labels});
      }
    });
    window.wails.Events.On("network_stats", (seriesOne, seriesTwo, labels) => {
      if (Object.entries(seriesOne).length != 0 && 
          Object.entries(seriesTwo).length != 0 && 
          Object.entries(labels).length != 0) {
        this.$store.commit({type: 'dashboard/setNetworkStatsChart', seriesOne, seriesTwo, labels});
      }
    });

    window.wails.Events.On("currency", currency => {
      this.$store.commit('wallet/setCurrency', currency);
    });

    window.wails.Events.On("wallet_keys", address => {
      this.$store.commit('wallet/setAddress', address);
    });

    window.wails.Events.On("campaign_status", address => {
      this.$store.commit('wallet/setCampaignStatus', address);
    });

    window.wails.Events.On("campaign_claim", address => {
      this.$store.commit('wallet/setCampaignClaim', address);
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
</style>
