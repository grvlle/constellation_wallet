const getDefaultState = () => {
  return {
    version: "v2.6.0",
    uiVersion: "v1.2.0",
    walletLabel: "",
    imgPath: 'faces/face-0.jpg',
    transactions: 0,
    tokenAmount: 0,
    totalBalance: 0,
    availableBalance: 0,
    nonce: 0,
    currency: "USD",
    totalValue: 0.0,
    blocks: "NaN",
    address: "N/A",
    keystorePath: "C:\\Users\\alexa\\Documents\\temp1.p12",
    saveKeystorePath: "",
    alias: "",
    privateKey: "NaN",
    publicKey: "NaN",
    seed: "Mnemonic Seed will be introduced with a later software release",
    darkMode: false,
    termsOfService: false
  }
}

const state = getDefaultState()

const actions = {
  resetWalletState({ commit }) {
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
  setBlocks(state, blocks) {
    state.blocks = blocks;
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
  setSaveKeystorePath(state, path) {
    state.saveKeystorePath = path;
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
},
}

export default {
  namespaced: true,
  state,
  getters: {},
  actions,
  mutations
}