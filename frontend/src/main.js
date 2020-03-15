import Vue from "vue";
import App from "./App.vue";
import router from "./router/index";

Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from '@wailsapp/runtime';
import PaperDashboard from "./plugins/paperDashboard";
import "vue-notifyjs/themes/default.css";
import VueNotify from 'vue-notifyjs'
import VueSweetalert2 from 'vue-sweetalert2';
import {store} from './store/store'
import JwPagination from 'jw-vue-pagination';
import Vuelidate from 'vuelidate'
import ToggleButton from 'vue-js-toggle-button'
import VueProgressBar from 'vue-progressbar'

Vue.component('jw-pagination', JwPagination);

Vue.use(VueProgressBar, {
  color: '#6DECBB',
  failedColor: 'red',
  height: '0.4rem',
  thickness: '0.4rem'
})
Vue.use(ToggleButton)
Vue.use(Vuelidate)
Vue.use(VueNotify);
Vue.use(PaperDashboard);
Vue.use(VueSweetalert2);

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
