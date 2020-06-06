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
  }
}

export default {
  state,
  getters: {},
  actions,
  mutations
}