import Vue from 'vue';
import Vuex from 'vuex';
import Wallet from './modules/wallet';
import App from './modules/app';
import Transactions from './modules/transations';
import Dashboard from './modules/dashboard';

Vue.use(Vuex);

export const store = new Vuex.Store({
    modules: {
        app: App,
        walletInfo: Wallet,
        txInfo: Transactions,
        dashboard: Dashboard
    },
    getters: {
        runningOnWindows (state) {
            return state.OS.windows
        },
        runningOnLinux (state) {
            return state.OS.linux
        },
        runningOnMacOS (state) {
            return state.OS.macOS
        }
    }
})

export default store;

