import LoginLayout from "@/layout/login/LoginLayout.vue";
import Download from "@/pages/DownloadingScreen";
import NewWallet from "@/pages/NewWallet";
import ImportWallet from "@/pages/ImportWallet";
import CreateWallet from "@/pages/CreateWallet";
import Login from "@/pages/Login";
import Migrate from "@/pages/Migrate";
import PasswordMigration from "@/pages/PasswordMigration";
import RecoveryPhraseInfo from "@/pages/RecoveryPhraseInfo";
import RecoveryPhrase from "@/pages/RecoveryPhrase";
import PasswordMigrationComplete from "@/pages/PasswordMigrationComplete";
import Loading from "@/pages/LoadingScreen";

import DashboardLayout from "@/layout/dashboard/DashboardLayout.vue";
import Dashboard from "@/pages/Dashboard.vue";
import WalletInformation from "@/pages/WalletInformation.vue";
import About from "@/pages/About.vue";
import Settings from "@/pages/Settings.vue";
import Transactions from "@/pages/Transactions.vue";
import AdddressBook from "@/pages/AddressBook.vue";
import NewEditContact from "@/pages/NewEditContact";
import ContactDetails from "@/pages/ContactDetails";
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
        path: "create-wallet",
        name: "create wallet",
        component: CreateWallet
      },
      {
        path: "import-wallet",
        name: "import wallet",
        component: ImportWallet
      },
      {
        path: "recovery-phrase-info",
        name: "recovery phrase info",
        component: RecoveryPhraseInfo
      },
      {
        path: "recovery-phrase",
        name: "recovery phrase",
        component: RecoveryPhrase
      },
      {
        path: "login",
        name: "login",
        component: Login
      },
      {
        path: "migrate-screen",
        name: "migrate screen",
        component: Migrate
      },
      {
        path: "password-migration-screen",
        name: "password migration",
        component: PasswordMigration
      },
      {
        path: "password-migration-complete",
        name: "password migration complete",
        component: PasswordMigrationComplete
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
        path: "submit-transaction/:txAddressProvided?",
        name: "submit transaction",
        component: Transactions,
        props: true,
      },
      {
        path: "settings",
        name: "settings",
        component: Settings
      },
      {
        path: "address-book",
        name: "address book",
        component: AdddressBook
      },
      {
        path: "new-edit-contact/:id",
        name: "new-edit contact",
        component: NewEditContact
      },
      {
        path: "contact-details/:id",
        name: "contact details",
        component: ContactDetails
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
