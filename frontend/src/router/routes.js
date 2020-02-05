import DashboardLayout from "@/layout/dashboard/DashboardLayout.vue";
// GeneralViews
import NotFound from "@/pages/NotFoundPage.vue";
import Dashboard from "@/pages/Dashboard.vue";
import WalletInformation from "@/pages/WalletInformation.vue";
import About from "@/pages/About.vue";
import Settings from "@/pages/Settings.vue";
import Transactions from "@/pages/Transactions.vue";

const routes = [

  {
    path: "/",
    component: DashboardLayout,
    redirect: "/dashboard",
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
