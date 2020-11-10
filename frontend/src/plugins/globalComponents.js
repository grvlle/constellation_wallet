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
  },
};

export default GlobalComponents;
