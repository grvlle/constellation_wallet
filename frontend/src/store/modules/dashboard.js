const getDefaultState = () => {
    return {
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
    }
}

const state = getDefaultState()

const actions = {
    resetDashboardState({ commit }) {
        commit('resetState')
    }
}
const mutations = {
    resetState(state) {
        Object.assign(state, getDefaultState())
    },
    setTokenCounter(state, counter) {
        state.counters.tokenCounter = counter
    },
    setValueCounter(state, counter) {
        state.counters.valueCounter = counter
    },
    setBlockCounter(state, counter) {
        state.counters.blockCounter = counter
    },
    setNodesOnlineCounter(state, counter) {
        state.counters.nodesOnlineCounter = counter
    },
    setShowNodesOnline(state, val) {
        state.toggleDashboard.showNodesOnline = val;
    },
    setShowTransactions(state, val) {
        state.toggleDashboard.showTransactions = val;
    },
    setShowThroughput(state, val) {
        state.toggleDashboard.showThroughput = val;
    },
    setNodeOnlineChart(state, obj) {
        state.chartData.nodesOnline.series = obj.series;
        state.chartData.nodesOnline.labels = obj.labels;
    },
    setTransactionStatsChart(state, obj) {
        state.chartData.transactions.series = [obj.seriesOne, obj.seriesTwo];
        state.chartData.transactions.labels = obj.labels;
    },
    setNetworkStatsChart(state, obj) {
        state.chartData.throughput.series = [obj.seriesOne, obj.seriesTwo];
        state.chartData.throughput.labels = obj.labels;
    }
}

export default {
    namespaced: true,
    state,
    getters: {},
    actions,
    mutations
}