import {
  FormGroupInput,
  Card,
  DropDown,
  Button,
  Password,
  FileSelector,
  Overlay,
  AirdropCard,
} from "../components/index";

import ViewDashboardIcon from "vue-material-design-icons/ViewDashboard.vue";
import WalletIcon from "vue-material-design-icons/Wallet.vue";
import CompareHorizontalIcon from "vue-material-design-icons/CompareHorizontal.vue";
import HelpCircleIcon from "vue-material-design-icons/HelpCircle.vue";
import CogIcon from "vue-material-design-icons/Cog.vue";
import LockIcon from "vue-material-design-icons/Lock.vue";

/**
 * You can register global components here and use them as a plugin in your main Vue instance
 */

const GlobalComponents = {
  install(Vue) {
    Vue.component("fg-input", FormGroupInput);
    Vue.component("drop-down", DropDown);
    Vue.component("card", Card);
    Vue.component("airdrop-card", AirdropCard);
    Vue.component("p-button", Button);
    Vue.component("password-input", Password);
    Vue.component("file-selector", FileSelector);
    Vue.component("page-overlay", Overlay);

    Vue.component("dashboard-icon", ViewDashboardIcon);
    Vue.component("wallet-icon", WalletIcon);
    Vue.component("tx-icon", CompareHorizontalIcon);
    Vue.component("about-icon", HelpCircleIcon);
    Vue.component("settings-icon", CogIcon);
    Vue.component("logout-icon", LockIcon);
  },
};

export default GlobalComponents;
