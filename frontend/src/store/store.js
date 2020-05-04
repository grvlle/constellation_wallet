import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);
export const store = new Vuex.Store({
    state: {
        errorMessage: "",
        warningMessage: "",
        successMessage: "",
        loginErrorMsg: "",
        network: "TESTNET",
        displayLoginError: false,
        app: {
            toc: false,
            txFinished: true,
            isLoading: false,
            isLoggedIn: false, // Change to false
            isDownloadingDependencies: true,
            import: false,
            register: false,
            login: true
        },
        downloading: {
          filename: "",
          size: ""
        },
        walletInfo: {
            version: "v2.5.2",
            uiVersion: "v0.1.9-BETA",
            email: "Molly Wallet",
            imgPath: 'faces/face-0.jpg',
            transactions: 0,
            tokenAmount: 0, 
            totalBalance: 0,
            availableBalance: 0,
            nonce: 0,
            currency: "USD",
            totalValue: "NaN",
            blocks: "NaN",
            address: "N/A",
            keystorePath: "",
            saveKeystorePath: "",
            alias: "",
            privateKey: "NaN",
            publicKey: "NaN",
            seed: "Mnemonic Seed will be introduced with a later software release",
            darkMode: false
        },
        txInfo: {
            txHistory: [],
            txStatus: "Complete"
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
        pageOfItems: [],
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
    mutations: {
        updateTxHistory(state, tx) {
            state.txInfo.txHistory.unshift(tx)
        },
        updateFullTxHistory(state, txHistoryUpdated) {
            
            state.txInfo.txHistory = txHistoryUpdated
            
        },
    },
    getters: {
        getTxCounter (state) {
            return state.counters.nodesOnlineCounter
        },
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

