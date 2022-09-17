import './shim';

Vue.config.productionTip = false;
Vue.config.devtools = true;

import Vue from "vue";
import App from "./App.vue";
import router from "./router/index";
import * as Wails from '@wailsapp/runtime';
import PaperDashboard from "./plugins/paperDashboard";
import "vue-notifyjs/themes/default.css";
import VueNotify from 'vue-notifyjs';
import {store} from './store/store';
import Vuelidate from 'vuelidate';
import VueProgressBar from 'vue-progressbar';
import VueSelect from 'vue-select';
import 'vue-select/dist/vue-select.css';
import IdleVue from 'idle-vue'
import './plugins/globalMethods';

const ONE_MINUTE = 60 * 1000;
const FIFTEEN_MINUTES = 15 * ONE_MINUTE;

const eventsHub = new Vue();
Vue.use(IdleVue, {
  eventEmitter: eventsHub,
  store,
  idleTime: FIFTEEN_MINUTES,
  startAtIdle: false
})
Vue.use(VueProgressBar, {
  color: '#6DECBB',
  failedColor: 'red',
  height: '0.4rem',
  thickness: '0.4rem'
})
Vue.use(Vuelidate)
Vue.use(VueNotify);
Vue.use(PaperDashboard);
Vue.use(VueSelect);

Wails.Init(() => {
  new Vue({
    router,
    store: store,
    render: h => h(App),
    mounted() {
      this.$router.replace('/')
    }
  }).$mount("#app");
});
