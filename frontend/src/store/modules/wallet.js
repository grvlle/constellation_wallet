const getDefaultState = () => {
  return {
    walletLabel: "",
    imgPath: 'faces/face-0.jpg',
    transactions: "",
    tokenAmount: "",
    totalBalance: "",
    availableBalance: "",
    currency: "USD",
    currencyRates: [],
    isCampaignActive: false,
    campaignClaimAddr: "",
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
  normalizedAvailableBalance: (state) => {
    return (state.availableBalance / 1e8).toFixed(8).replace(/\.?0+$/, "");
  },
  valueInCurrency: (state) => {
    let currencyRate = state.currencyRates.find(v => v.currency === state.currency)
    if (currencyRate === undefined) {
      return "..."
    } else {
      let formatOptions, value
      value = state.availableBalance/1e8 * currencyRate.tokenprice
      if (state.currency == "BTC") {
        formatOptions = {
          style: "currency",
          currency: "XBT",
          minimumFractionDigits: 2,
          maximumFractionDigits: 8
        };
      } else {
        formatOptions = {
          style: "currency",
          currency: state.currency,
          minimumFractionDigits: 2,
          maximumFractionDigits: 2
        };
      }      
      let formatter = new Intl.NumberFormat(navigator.language, formatOptions);
      return formatter.format(value).replace(/XBT/,'₿');
    }
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
  setCurrencyRates(state, value) {
    state.currencyRates = value;
  },
  setAddress(state, address) {
    state.address = address;
  },
  setCampaignStatus(state, status) {
    state.isCampaignActive = status;
  },
  setCampaignClaim(state, address) {
    state.campaignClaimAddr = address;
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
