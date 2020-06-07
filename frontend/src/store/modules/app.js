const getDefaultState = () => {
    return {
        termsOfService: false,
        txFinished: true,
        isLoggedIn: false,
        errorMessage: "",
        warningMessage: "",
        successMessage: "",
        loginErrorMsg: "",
        newRelease: "",
        network: "MAINNET",
        displayLoginError: false,
        downloading: {
          filename: "",
          size: ""
        }
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
        state.termsOfService = termsOfService;
    },
    setTxFinished(state, setTxFinished) {
        state.termsOfService = setTxFinished;
    },
    setIsLoggedIn(state, isLoggedIn) {
        state.termsOfService = isLoggedIn;
    },
    setErrorMessage(state, message) {
        state.errorMessage = message;
    },
    setWarningMessage(state, message) {
        state.warningMessage = message;
    },
    setSuccessMessage(state, message) {
        state.successMessage = message;
    },
    setLoginErrorMessage(state, message) {
        state.loginErrorMsg = message;
    },
    setDisplayLoginError(state, val) {
        state.displayLoginError = val;
    },
    setNewRelease(state, val) {
        state.newRelease = val;
    },
    setDownloadFileName(state, name) {
        state.downloading.filename = name;
    },
    setDownloadFileSize(state, size) {
        state.downloading.size = size;
    }
}

export default {
    state,
    getters: {},
    actions,
    mutations
}