import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

import Wallet from "../../../JSONdata/wallet.json"
import TXHistory from "../../../../../../../.dag/txhistory.json"
//import TransactionHistory from '../../../JSONdata/tx.json';
import ChartData from '../../../JSONdata/chart_data.json';

export const store = new Vuex.Store({
    state: {
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
            txHistory: TXHistory
        },
        counters: {
            blockCounter: 5,
            tokenCounter: 60,
            nodesOnlineCounter: 24
        },
        chartData: {
            nodesOnline: {
                labels: ChartData.nodes_online.labels,
                series: ChartData.nodes_online.series
            },
            transactions: {
                labels:    [ChartData.transactions.labels],
                seriesOne: ChartData.transactions.series_one,
                seriesTwo: ChartData.transactions.series_two
            },
            throughput: {
                labels:    [ChartData.throughput.labels],
                seriesOne: ChartData.throughput.series_one,
                seriesTwo: ChartData.throughput.series_two
            }
        }
    },
    mutations: {
        updateTxHistory(state, tx) {
            state.txInfo.txHistory.unshift(tx)
        },
        clearTxHistory(state, clear) {
            if (clear == true) {
            state.txInfo.txHistory = []
            }
        }
    }

})
