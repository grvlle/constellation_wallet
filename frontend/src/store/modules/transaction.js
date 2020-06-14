const getDefaultState = () => {
    return {
        txHistory: [],
        txStatus: "Complete"
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
    namespaced: true,
    state,
    getters: {},
    actions,
    mutations
}