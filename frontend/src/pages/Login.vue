<template>
  <center>
    <div style="margin-top: 50px;">
      <form @submit.prevent>
        <div class="row">
          <div class="col-md-3"></div>
          <div class="col-md-6">
            <fg-input v-model="username" type="text" label="Enter Username" placeholder="Username"></fg-input>
          </div>
        </div>
        <div class="row">
          <div class="col-md-3"></div>
          <div class="col-md-6">
            <fg-input
              v-model="password"
              type="password"
              label="Enter Password"
              placeholder="Enter Wallet Password"
            ></fg-input>
          </div>
        </div>
        <div class="row">
          <div class="col-md-4"></div>
          <div class="col-md-2">
            <p-button
              v-if="!this.$store.state.app.isLoggedIn"
              type="info"
              block
              @click.native="login"
              style="margin-top: 28px; overflow: visible;"
            >
              <span style="display: block;">
                <i class="fa fa-paper-plane"></i> LOGIN TO EXISTING WALLET
              </span>
            </p-button>
          </div>

          <div class="col-md-2">
            <p-button
              v-if="!this.$store.state.app.isLoggedIn"
              type="warning"
              block
              @click.native="newLogin"
              style="margin-top: 28px; overflow: visible;"
            >
              <span style="display: block;">
                <i class="fa fa-lock"></i> CREATE A NEW WALLET
              </span>
            </p-button>
          </div>
        </div>

        <!-- <div class="clearfix"></div> -->
      </form>
    </div>
  </center>
</template>

<script>
import TXSent from "./Notifications/TxSent";
export default {
  name: "login-screen",
  username: "",
  password: "",
  access: false,
  methods: {
    login: function() {
      var self = this;

      window.backend.WalletApplication.Login(self.username, self.password).then(
        result => {
          self.access = result;
          if (self.access) {
            self.$store.state.app.isLoading = self.access;
            self.$store.state.app.isLoggedIn = self.access;
            setTimeout(() => {
              self.$store.state.app.isLoading = false;
            }, 8000);
          }
        }
      );
    },
    newLogin: function() {
      var self = this;
      window.backend.WalletApplication.CreateUser(
        self.username,
        self.password
      ).then(result => {
        if (result) {
          window.backend.WalletApplication.Login(
            self.username,
            self.password
          ).then(result => {
            self.access = result;
            if (self.access) {
              self.$store.state.app.isLoading = self.access;
              self.$store.state.app.isLoggedIn = self.access;
              setTimeout(() => {
                self.$store.state.app.isLoading = false;
              }, 8000);
            }
          });
        }
      });
    }
  }
};
</script>

<style scoped>
body {
  background-color: black;
}
</style>
