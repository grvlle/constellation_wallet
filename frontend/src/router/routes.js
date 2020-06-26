import LoginLayout from "@/layout/login/LoginLayout.vue";
import Download from "@/pages/DownloadingScreen";
import NewWallet from "@/pages/NewWallet";
import Login from "@/pages/Login";
import Loading from "@/pages/LoadingScreen";

import DashboardLayout from "@/layout/dashboard/DashboardLayout.vue";
import Dashboard from "@/pages/Dashboard.vue";
import WalletInformation from "@/pages/WalletInformation.vue";
import About from "@/pages/About.vue";
import Settings from "@/pages/Settings.vue";
import Transactions from "@/pages/Transactions.vue";

import TermsOfService from "@/pages/TermsOfService.vue";
import NotFound from "@/pages/NotFoundPage.vue";

const routes = [
  {
    path: "/",
    component: LoginLayout,
    redirect: "/download",
    children: [
      {
        path: "download",
        name: "download",
        component: Download
      },
      {
        path: "new-wallet",
        name: "new wallet",
        component: NewWallet
      },
      {
        path: "login",
        name: "login",
        component: Login
      },
      {
        path: "accept-terms-of-service",
        name: "accept terms of service",
        component: TermsOfService
      },
      {
        path: "loading",
        name: "loading",
        component: Loading
      }
    ]
  },
  {
    path: "/home",
    name: "home",
    component: DashboardLayout,
    redirect: "/home/dashboard",
    children: [
      {
        path: "dashboard",
        name: "dashboard",
        component: Dashboard
      },
      {
        path: "wallet-info",
        name: "wallet information",
        component: WalletInformation
      },
      {
        path: "about",
        name: "about",
        component: About
      },
      {
        path: "submit-transaction",
        name: "submit transaction",
        component: Transactions
      },
      {
        path: "settings",
        name: "settings",
        component: Settings
      },
      {
        path: "terms-of-service",
        name: "terms of service",
        component: TermsOfService
      },
    ]
  },
  { path: "*", component: NotFound }
];

/**
 * Asynchronously load view (Webpack Lazy loading compatible)
 * The specified component must be inside the Views folder
 * @param  {string} name  the filename (basename) of the view to load.
function view(name) {
   var res= require('../components/Dashboard/Views/' + name + '.vue');
   return res;
};**/

export default routes;
