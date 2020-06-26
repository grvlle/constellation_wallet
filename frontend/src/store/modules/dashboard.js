const getDefaultState = () => {
    return {
        counters: {
            block: 5,
            token: 30,
            value: 60,
            chart: 24
        },
        toggle: {
            nodesOnline: false,
            transactions: true,
            throughput: true,
        },
        stat: {
            blocks: "NaN",
        },
        chart: {
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
    reset({ commit }) {
        commit('resetState')
    }
}
const mutations = {
    resetState(state) {
        Object.assign(state, getDefaultState())
    },
    setTokenCounter(state, counter) {
        state.counters.token = counter
    },
    setValueCounter(state, counter) {
        state.counters.value = counter
    },
    setBlockCounter(state, counter) {
        state.counters.block = counter
    },
    setChartCounter(state, counter) {
        state.counters.chart = counter
    },
    setShowNodesOnline(state, val) {
        state.toggle.nodesOnline = val;
    },
    setShowTransactions(state, val) {
        state.toggle.transactions = val;
    },
    setShowThroughput(state, val) {
        state.toggle.throughput = val;
    },
    setBlocks(state, blocks) {
        state.stat.blocks = blocks;
      },
    setNodeOnlineChart(state, obj) {
        state.chart.nodesOnline.series = obj.series;
        state.chart.nodesOnline.labels = obj.labels;
    },
    setTransactionStatsChart(state, obj) {
        state.chart.transactions.series = [obj.seriesOne, obj.seriesTwo];
        state.chart.transactions.labels = obj.labels;
    },
    setNetworkStatsChart(state, obj) {
        state.chart.throughput.series = [obj.seriesOne, obj.seriesTwo];
        state.chart.throughput.labels = obj.labels;
    }
}

export default {
    namespaced: true,
    state,
    getters: {},
    actions,
    mutations
}