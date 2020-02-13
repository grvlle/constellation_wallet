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
            login: true,
            margin: 70,
        },
        downloading: {
          filename: "",
          size: ""
        },
        walletInfo: {
            version: "v1.13",
            uiVersion: "v0.1.7 Beta",
            email: "user@email.com",
            imgPath: 'faces/face-0.jpg',
            transactions: 0,
            tokenAmount: 0, 
            totalBalance: 0,
            availableBalance: 0,
            nonce: 0,
            usdValue: "NaN",
            blocks: "NaN",
            address: "N/A",
            keystorePath: "",
            saveKeystorePath: "",
            alias: "",
            privateKey: "NaN",
            publicKey: "NaN",
            seed: "Mnemonic Seed will be introduced with a later software release"
        },
        validators: {
            target: "",
            alias: {
                alias_length: 0,
                contains_five_characters: false,
                valid_alias: false,
            },
            keypass: {
                password_length: 0,
                contains_eight_characters: false,
                contains_number: false,
                contains_uppercase: false,
                contains_special_character: false,
                valid_password: false
            },
            storepass: {
                password_length: 0,
                contains_eight_characters: false,
                contains_number: false,
                contains_uppercase: false,
                contains_special_character: false,
                valid_password: false
            },
            duplicate: false
        },
        txInfo: {
            txHistory: []
        },
        counters: {
            blockCounter: 5,
            tokenCounter: 60,
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
    }

})
