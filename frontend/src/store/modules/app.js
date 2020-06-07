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
    },
    setTermsOfService(state, termsOfService) {
        this.termsOfService = termsOfService;
    },
    setTxFinished(state, setTxFinished) {
        this.termsOfService = setTxFinished;
    },
    setIsLoggedIn(state, isLoggedIn) {
        this.termsOfService = isLoggedIn;
    }
}

export default {
    state,
    getters: {},
    actions,
    mutations
}