const getDefaultState = () => {
  return {
    version: "v2.6.0",
    uiVersion: "v1.2.0",
    email: "Molly Wallet",
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
    keystorePath: "",
    saveKeystorePath: "",
    alias: "",
    privateKey: "NaN",
    publicKey: "NaN",
    seed: "Mnemonic Seed will be introduced with a later software release",
    darkMode: false
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
    state.setAvailableBalance = available;
  },
  setTotalBalance(state, total) {
    state.setTotalBalance = total;
  },
  setBlocks(state, blocks) {
    state.blocks = blocks;
  },
  setTotalValue(state, value) {
    state.setTotalValue = value;
  },
  setAddress(state, address) {
    state.address = address;
  },
  setKeystorePath(state, path) {
    state.setKeystorePath = path;
  },
  setSaveKeystorePath(state, path) {
    state.setSaveKeystorePath = path;
  },
  setEmail(state, email) {
    state.email = email;
  },
  setDarkMode(state, darkMode) {
    state.darkMode = darkMode;
  },
  setImgPath(state, path) {
    state.imgPath = path;
  }
}

export default {
  state,
  getters: {},
  actions,
  mutations
}