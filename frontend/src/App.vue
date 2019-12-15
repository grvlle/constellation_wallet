<template>
    <div :class="{'nav-open': $sidebar.showSidebar}">
        <loading-screen :isLoading="isLoading" :fadeout="!isLoading"></loading-screen>
        <div v-if="!isLoading">
        <notifications></notifications>
        <router-view></router-view>
        </div>
    </div>
</template>

<script>
import ErrorNotification from './pages/Notifications/ErrorMessage';
import LoadingScreen from './pages/LoadingScreen'

export default {
    components: {
        LoadingScreen
    },
    data() {
        return {
            isLoading: true
        }
    },
    mounted() {
        setTimeout(() => {
            this.isLoading = false
        }, 6000)

        // Backend Errors
        window.wails.Events.On("error_handling", (m, err) => {
            this.$store.state.errorMessage = m + err
             setTimeout(() => {
                    this.$notifications.clear();
            }, 60000)
            this.$notify({
                component: ErrorNotification,
                timeout: 500000,
                icon: "fa fa-times",
                horizontalAlign: 'right',
                verticalAlign: 'top',
                type: 'danger',
                onClick: ()=>{
                    this.$notifications.clear();
                }
            })
        });

        // Transactions.vue sockets
        window.wails.Events.On("update_tx_history", (txHistoryFull) => {
            this.$store.state.txInfo.txHistory = txHistoryFull;
            //this.$store.commit('updateFullTxHistory', txHistoryFull)
        });
        window.wails.Events.On("new_transaction", (txObject) => {
            this.$store.commit('updateTxHistory', txObject)
        });

        // Dashboard.vue sockets
        window.wails.Events.On("token", (amount) => {
            this.$store.state.walletInfo.tokenAmount = amount;
        });
        window.wails.Events.On("blocks", (number) => {
            this.$store.state.walletInfo.blocks = number;
        });
        window.wails.Events.On("price", (currency, dagRate) => {
            this.$store.state.walletInfo.usdValue = currency + " " + dagRate
        });
        window.wails.Events.On("token_counter", (count) => {
            this.$store.state.counters.tokenCounter = count;
        });
        window.wails.Events.On("block_counter", (blockCount) => {
            this.$store.state.counters.blockCounter = blockCount;
        });
        window.wails.Events.On("chart_counter", (pieChartCount) => {
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
        window.wails.Events.On("wallet_keys", (privateKey, publicKey) => {
            this.$store.state.walletInfo.privateKey = privateKey;
            this.$store.state.walletInfo.publicKey = publicKey;
        })
    },
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