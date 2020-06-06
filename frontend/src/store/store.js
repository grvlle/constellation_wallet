import Vue from 'vue';
import Vuex from 'vuex';
import Wallet from './modules/wallet';
import App from './modules/app';
import Transactions from './modules/transations';

Vue.use(Vuex);

export const store = new Vuex.Store({
    modules: {
        app: App,
        walletInfo: Wallet,
        txInfo: Transactions
    },
    state: {
        errorMessage: "",
        warningMessage: "",
        successMessage: "",
        loginErrorMsg: "",
        newRelease: "",
        network: "MAINNET",
        displayLoginError: false,
        downloading: {
          filename: "",
          size: ""
        },
        counters: {
            blockCounter: 5,
            tokenCounter: 30,
            valueCounter: 60,
            nodesOnlineCounter: 24
        },
        toggleDashboard: {
            showNodesOnline: false,
            showTransactions: true,
            showThroughput: true,
        },
        chartData: {
            nodesOnline: {
                labels: [], // ChartData.nodes_online.labels,
                series: []  // ChartData.nodes_online.series
            },
            transactions: {
                labels:    [], //[ChartData.transactions.labels],
                series: [] //[ChartData.transactions.series_one, ChartData.transactions.series_two]
            },
            throughput: {
                labels:    [], //[ChartData.throughput.labels],
                series: [] //[ChartData.throughput.series_one, ChartData.throughput.series_two]
            }
        }
    },
    getters: {
        runningOnWindows (state) {
            return state.OS.windows
        },
        runningOnLinux (state) {
            return state.OS.linux
        },
        runningOnMacOS (state) {
            return state.OS.macOS
        }
    }
})

export default store;

