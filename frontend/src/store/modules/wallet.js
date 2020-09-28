const getDefaultState = () => {
  return {
    walletLabel: "",
    imgPath: 'faces/face-0.jpg',
    transactions: "",
    tokenAmount: "",
    totalBalance: "",
    availableBalance: "",
    currency: "USD",
    totalValue: "",
    address: "",
    keystorePath: "",
    alias: "",
    publicKey: "",
    darkMode: false,
    termsOfService: false
  }
}

const state = getDefaultState()

const getters = {
  getNormalizedAvailableBalance: (state) => {
    return (state.availableBalance / 1e8).toFixed(8).replace(/\.?0+$/, "");
  }
}

const actions = {
  reset({ commit }) {
    commit('resetState')
  }
}
const mutations = {
  resetState(state) {
    Object.assign(state, getDefaultState())
  },
  setCurrency(state, currency) {
    state.currency = currency;
  },
  setTokenAmount(state, amount) {
    state.tokenAmount = amount;
  },
  setAvailableBalance(state, available) {
    state.availableBalance = available;
  },
  setTotalBalance(state, total) {
    state.totalBalance = total;
  },
  setTotalValue(state, value) {
    state.totalValue = value;
  },
  setAddress(state, address) {
    state.address = address;
  },
  setAlias(state, alias) {
    state.alias = alias;
  },
  setKeystorePath(state, path) {
    state.keystorePath = path;
  },
  setLabel(state, label) {
    state.walletLabel = label;
  },
  setDarkMode(state, darkMode) {
    state.darkMode = darkMode;
  },
  setImgPath(state, path) {
    state.imgPath = path;
  },
  setTermsOfService(state, termsOfService) {
    state.termsOfService = termsOfService;
  }
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}