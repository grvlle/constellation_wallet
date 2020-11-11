import ViewDashboardIcon from "vue-material-design-icons/ViewDashboard.vue";
import WalletIcon from "vue-material-design-icons/Wallet.vue";
import CompareHorizontalIcon from "vue-material-design-icons/CompareHorizontal.vue";
import HelpCircleIcon from "vue-material-design-icons/HelpCircle.vue";
import CogIcon from "vue-material-design-icons/Cog.vue";
import LockIcon from "vue-material-design-icons/Lock.vue";
import ClockOutlineIcon from "vue-material-design-icons/ClockOutline";
import Cash100Icon from "vue-material-design-icons/Cash100";
import ArrowRightIcon from "vue-material-design-icons/ArrowRight";
import ArrowLeftIcon from "vue-material-design-icons/ArrowLeft";

const MaterialIcons = {
  install(Vue) {
    Vue.component("dashboard-icon", ViewDashboardIcon);
    Vue.component("wallet-icon", WalletIcon);
    Vue.component("tx-icon", CompareHorizontalIcon);
    Vue.component("about-icon", HelpCircleIcon);
    Vue.component("settings-icon", CogIcon);
    Vue.component("logout-icon", LockIcon);
    Vue.component("clock-icon", ClockOutlineIcon);
    Vue.component("fiat-icon", Cash100Icon);
    Vue.component("arrow-r-icon", ArrowRightIcon);
    Vue.component("arrow-l-icon", ArrowLeftIcon);
  },
};

export default MaterialIcons;
