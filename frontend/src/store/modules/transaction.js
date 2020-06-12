const getDefaultState = () => {
    return {
        txHistory: [],
        txStatus: "Complete"
    }
}

const state = getDefaultState()

const actions = {
    resetTransactionsState({ commit }) {
        commit('resetState')
    }
}
const mutations = {
    resetState(state) {
        Object.assign(state, getDefaultState())
    },
    updateTxHistory(state, tx) {
        state.txHistory.unshift(tx)
    },
    updateFullTxHistory(state, obj) {
      state.txHistory = obj.txHistoryFull
    },
    updateTxStatus(state, status) {
        state.txStatus = status
    }
}

export default {
    state,
    getters: {},
    actions,
    mutations
}