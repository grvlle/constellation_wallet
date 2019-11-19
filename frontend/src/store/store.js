import Vue from 'vue';
import Vuex from 'vuex';



import Wallet from "../../../JSONdata/wallet.json"
//import ChartData from '../../../../../../../.dag/chart_data.json';

Vue.use(Vuex);
export const store = new Vuex.Store({
    state: {
        errorMessage: "None",
        walletInfo: {
        tokenAmount: Wallet.balance, 
        usdValue: "NaN",
        blocks: "NaN",
        address: Wallet.address,
        privateKey: "NaN",
        publicKey: "NaN",
        seed: "witch collapse practice feed shame open despair creek road again ice least"
        },
        txInfo: {
            txHistory: []
        },
        counters: {
            blockCounter: 5,
            tokenCounter: 60,
            nodesOnlineCounter: 24
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
