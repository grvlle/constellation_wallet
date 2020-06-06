const getDefaultState = () => {
    return {
        txHistory: [],
        txStatus: "Complete"
    }
}

const state = getDefaultState()

const getters = {
    getTxCounter (state) {
        return state.counters.nodesOnlineCounter
    }
}

const actions = {
    resetTxInfoState({ commit }) {
        commit('resetState')
    }
}
const mutations = {
    resetState(state) {
        Object.assign(state, getDefaultState())
    },
    updateTxHistory(state, tx) {
        state.txInfo.txHistory.unshift(tx)
    },
    updateFullTxHistory(state, txHistoryUpdated) {
        state.txInfo.txHistory = txHistoryUpdated
    },
}

export default {
    state,
    getters,
    actions,
    mutations
}