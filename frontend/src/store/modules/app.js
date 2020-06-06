const getDefaultState = () => {
    return {
        termsOfService: false,
        txFinished: true,
        isLoggedIn: false
    }
}

const state = getDefaultState()

const actions = {
    resetAppState({ commit }) {
        commit('resetState')
    }
}
const mutations = {
    resetState(state) {
        Object.assign(state, getDefaultState())
    }
}

export default {
    state,
    getters: {},
    actions,
    mutations
}