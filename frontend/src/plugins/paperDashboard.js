import Notify from "vue-notifyjs";
import SideBar from "@/components/SidebarPlugin";
import GlobalComponents from "./globalComponents";
import MaterialIcons from "./materialIcons";
import GlobalDirectives from "./globalDirectives";
// import "es6-promise/auto";

//css assets
import "bootstrap/dist/css/bootstrap.css";
import "@/assets/sass/paper-dashboard.scss";
import "@/assets/css/themify-icons.css";

export default {
  install(Vue) {
    Vue.use(GlobalComponents);
    Vue.use(GlobalDirectives);
    Vue.use(MaterialIcons);
    Vue.use(SideBar);
    Vue.use(Notify);
  },
};
