<template>
  <center>
    <div :style="'margin-top: ' + this.$store.state.app.margin + 'px'">
        <div>
      <form @submit.prevent>
        <div class="row">
          <div class="col-md-4"></div>
          <div class="col-md-4">
            <fg-input v-model="username" type="text" label="Username" placeholder="Username"></fg-input>
          </div>
        </div>
        <div class="col-md-4"></div>
        <div class="row">
          <div class="col-md-4"></div>
          <div class="col-md-4">
            <fg-input
              v-model="password"
              type="password"
              label="Password"
              placeholder="Enter Wallet Password"
            ></fg-input>
          </div>
        </div>
        <div class="col-md-4"></div>

        <div v-if="this.$store.state.app.register" class="row">
          <div class="col-md-4"></div>
          <div class="col-md-4">
            <fg-input
              v-model="alias"
              type="text"
              label="Wallet Alias"
              placeholder="Enter alias..."
            ></fg-input>
          </div>
        </div>
        <div class="col-md-4"></div>

        <div v-if="this.$store.state.app.register" class="row">
          <div class="col-md-4"></div>
          <div class="col-md-4">
            <fg-input
              v-model="storepass"
              type="password"
              label="Store Password"
              placeholder="Enter storepass..."
            ></fg-input>
          </div>
        </div>
<div class="col-md-4"></div>

        <div v-if="this.$store.state.app.register" class="row">
          <div class="col-md-4"></div>
          <div class="col-md-4">
            <fg-input
              v-model="keypass"
              type="password"
              label="Key Password"
              placeholder="Enter keypass..."
            ></fg-input>
          </div>
        </div>
        <div class="col-md-4"></div>

        <div class="row">
          <div class="col-md-4"></div>
          <div v-if="!this.$store.state.app.register" class="col-md-2">
            <p-button
              v-if="!this.$store.state.app.isLoggedIn"
              type="success"
              block
              @click.native="login"
              style="margin-top: 28px; overflow: visible;"
            >
              <span style="display: block;">
                <i class="fa fa-unlock"></i> LOGIN 
              </span>
            </p-button>
          </div>



          <div v-if="!this.$store.state.app.register" class="col-md-2">
            <p-button
              v-if="!this.$store.state.app.isLoggedIn"
              type="danger"
              block
              @click.native="newLogin"
              style="margin-top: 28px; overflow: visible;"
            >
              <span style="display: block;">
                <i class="fa fa-key"></i> CREATE A NEW WALLET
              </span>
            </p-button>
          </div>

                              <div v-if="this.$store.state.app.register" class="col-md-2">
            <p-button
              v-if="!this.$store.state.app.isLoggedIn"
              type="default"
              block
              @click.native="cancelEvent"
              style="margin-top: 28px; overflow: visible;"
            >
              <span style="display: block;">
                <i class="fa fa-close"></i> CANCEL
              </span>
            </p-button>
          </div>
        

        <div v-if="this.$store.state.app.register" class="col-md-2">
          <p-button
            v-if="!this.$store.state.app.isLoggedIn"
            type="warning"
            block
            @click.native="createLogin"
            style="margin-top: 28px; overflow: visible;"
          >
            <span style="display: block;">
              <i class="fa fa-lock"></i> CREATE!
            </span>
          </p-button>
        </div>

        <!-- <div class="clearfix"></div> -->
      </form>
      </div>
    </div>
  </center>
</template>

<script>
import TXSent from "./Notifications/TxSent";
export default {
  name: "login-screen",
  username: "",
  password: "",
  storepass: "",
  keypass: "",
  alias: "",
  access: false,
  methods: {
    newLogin: function() {
      this.$store.state.app.register = !this.$store.state.app.register;
      this.$store.state.app.margin = 160
    },
        cancelEvent: function() {
      this.$store.state.app.register = !this.$store.state.app.register;
      this.$store.state.app.margin = 250
    },
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
    createLogin: function() {
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
.body {
background-color: black; 
}
</style>
