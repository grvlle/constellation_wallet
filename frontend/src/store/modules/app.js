const getDefaultState = () => {
  return {
    version: "v2.13.9",
    uiVersion: "v1.4.1",
    isLoggedIn: false,
    errorMessage: "",
    warningMessage: "",
    successMessage: "",
    loginErrorMsg: "",
    newRelease: "",
    network: "Main Constellation Network",
    displayLoginError: false,
    downloading: {
      filename: "",
      size: "",
    },
  };
};

const state = getDefaultState();

const actions = {
  reset({ commit }) {
    commit("resetState");
  },
};
const mutations = {
  resetState(state) {
    Object.assign(state, getDefaultState());
  },
  setNetwork(state, network) {
    state.network = network;
  },
  setIsLoggedIn(state, isLoggedIn) {
    state.isLoggedIn = isLoggedIn;
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
  },
};

export default {
  namespaced: true,
  state,
  getters: {},
  actions,
  mutations,
};
