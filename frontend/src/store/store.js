import Vue from 'vue';
import Vuex from 'vuex';
import Wallet from './modules/wallet';
import App from './modules/app';
import Transaction from './modules/transaction';
import Dashboard from './modules/dashboard';

Vue.use(Vuex);

export const store = new Vuex.Store({
    modules: {
        app: App,
        wallet: Wallet,
        transaction: Transaction,
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

