import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);
export const store = new Vuex.Store({
    state: {
        errorMessage: "",
        warningMessage: "",
        successMessage: "",
        loginErrorMsg: "",
        displayLoginError: false,
        app: {
            toc: false,
            txFinished: true,
            isLoading: false,
            isLoggedIn: false,
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
            version: "2.1.0-rc",
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
            seed: "Mnemonic Seed will be introduced with a later software release"
        },
        txInfo: {
            txHistory: [],
            txStatus: "Complete"
        },
        notificationInfo: {
            txNotifications: [],
            systemNotifications: []
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
        addTxHistory(state, tx) {
            state.txInfo.txHistory.unshift(tx)
        },
        updateFullTxHistory(state, txHistoryUpdated) {
            state.txInfo.txHistory = txHistoryUpdated
        },
        addTxNotification(state, notification) {
            state.notificationInfo.txNotifications.unshift(notification)
        },
        updateTxNotification (state, notification) {
            let foundIndex = state.notificationInfo.txNotifications.findIndex(x => x.id == notification.id);
            state.notificationInfo.txNotifications[foundIndex] = notification;
        },
        updateTxNotificationsAsRead (state) {
            state.notificationInfo.txNotifications.filter(n => n.read == false)
              .forEach(f => f.read = true);
        },
        deleteTxNotifications (state) {
            state.notificationInfo.txNotifications = [];
        },
        addSystemNotification(state, notification) {
            state.notificationInfo.systemNotifications.unshift(notification)
        },
        updateSystemNotification (state, notification) {
            let foundIndex = state.notificationInfo.systemNotifications.findIndex(x => x.id == notification.id);
            state.notificationInfo.systemNotifications[foundIndex] = notification;
        },
        updateSystemNotificationsAsRead (state) {
              state.notificationInfo.systemNotifications.filter(n => n.read == false)
              .forEach(f => f.read = true);
        },
        deleteSystemNotifications (state) {
            state.notificationInfo.systemNotifications = [];
        }
    }

})
