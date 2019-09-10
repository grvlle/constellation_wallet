import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

import Wallet from "../../../JSONdata/wallet.json"
import TransactionHistory from '../../../JSONdata/tx.json';
import ChartData from '../../../JSONdata/chart_data.json';

export const store = new Vuex.Store({
    state: {
        walletInfo: {
        tokenAmount: Wallet.balance, 
        usdValue: "$ " + (Wallet.balance * Wallet.token_price.DAG.USD).toFixed(2),
        blocks: "NaN",
        address: Wallet.address
        },
        txInfo: {
            txHistory: [TransactionHistory]
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
    getters:{
        localWallet: state => {
        var localWallet = state.walletInfo.map(walletItem => {
            return {
                tokenAmount: walletItem.tokenAmount,
                usdValue: walletItem.usdValue,
                blocks: walletItem.blocks
            }
        });
        return localWallet;
    }
    }
})
