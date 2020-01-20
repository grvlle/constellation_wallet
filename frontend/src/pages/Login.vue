<template>
  <div class="bg">
    <div></div>
    <center>
      <div :style="'margin-top: ' + this.$store.state.app.margin + 'px'">
        <img v-if="!this.$store.state.app.register" src="~@/assets/img/Constellation-Logo-1.png" />
        <p
          v-if="!this.$store.state.app.register"
          style="margin-bottom: 80px; margin-top: 5px;"
        >Please enter your credentials below to access your Molly Wallet.</p>
        <div>
          <form @submit.prevent>
            <div class="row">
              <div class="col-2"></div>
              <div v-if="!this.$store.state.app.register" class="col-2"></div>
              <div title="Create a New Wallet" v-if="this.$store.state.app.register" class="col-4 info-box">
                <p>
                  <br />
                  <b>Create a new wallet</b>
                  <br />
                  This section will let you create a Molly Wallet to store your $DAG tokens in. Before doing so, you need to set up the right hand side credentials. <br />

                  <br />
                  <ul>
  <li><b>Username</b><i> - This is where you select your username</i></li>
  <li><b>Password</b><i> - This is the password you will use when accessing or restoring your wallet</i></li>
  <li><b>Store Pass</b><i> - This is the storepass </i></li>
</ul>
                
                  </p>
              </div>


              <div class="col-4">  
        <div>
        <table style="width:100%;">
          <tr>
            <td style="padding: 0px; width: 81%;">
    
              <fg-input
                type="text"
                :disabled="true"
                :placeholder="this.$store.state.walletInfo.keystorePath"
                v-model="this.$store.state.walletInfo.keystorePath"
              ></fg-input>

      </td>

            <td style="padding-left: 0px;">
              <p-button
                @click.native="importKey"
                type="default"
                style="margin-top: -17px; width: 95%; float: right;"
              ><span style="display: block;">
               BROWSE
              </span></p-button>
         
            </td>
          </tr>
        </table>
              </div>
              
              <div>
                <fg-input v-model="keystorePassword" type="password" label="Keystore Password" placeholder="Enter Keystore Password ..."></fg-input>
                  </div>

                <div>
                  <fg-input
                    v-model.trim="keyPassword"
                    @change="setKeyPassword($event.target.value)"
                    type="password"
                    label="Key Password"
                    placeholder="Enter Key Password..."
                  ></fg-input>

    
                </div>
                                  <div style="height: 20px;" class="error" v-if="$v.keyPasswordValidate.minLength && $v.keyPasswordValidate.maxLength">
    
                                </div>
    
                                <div class="error" v-if="!$v.keyPasswordValidate.minLength || !$v.keyPasswordValidate.maxLength">
                                    <p class="validate">Invalid wallet address. Please verify.</p>
                                </div>

                <div v-if="this.$store.state.app.register">
                  <fg-input
                    v-model="this.$store.state.walletInfo.email"
                    type="text"
                    label="Wallet Label (optional)"
                  ></fg-input>
                </div>

                <div
                  v-if="!this.$store.state.app.register"
                  style="float: left; width: 48%; margin-top: 20px;"
                >
                  <p-button
                    v-if="!this.$store.state.app.isLoggedIn"
                    type="success"
                    block
                    @click.native="login"
                    :disabled="submitStatus === 'PENDING'"
                    style="overflow: visible;"
                  >
                    <span style="display: block;">
                      <i class="fa fa-unlock"></i> LOGIN
                    </span>
                  </p-button>
                </div>

                <div
                  v-if="!this.$store.state.app.register"
                  style="float: right; width: 48%; margin-top: 20px;"
                >
                  <p-button
                    v-if="!this.$store.state.app.isLoggedIn"
                    type="danger"
                    block
                    @click.native="newLogin"
                    style="overflow: visible;"
                  >
                    <span style="display: block;">
                      <i class="fa fa-key"></i> CREATE A NEW WALLET
                    </span>
                  </p-button>
                </div>

                <div
                  v-if="this.$store.state.app.register"
                  style="float: left; width: 48%; margin-top: 20px;"
                >
                  <p-button
                    v-if="!this.$store.state.app.isLoggedIn"
                    type="default"
                    block
                    @click.native="cancelEvent"
                    style="overflow: visible;"
                  >
                    <span style="display: block;">
                      <i class="fa fa-close"></i> CANCEL
                    </span>
                  </p-button>
                </div>

                <div
                  v-if="this.$store.state.app.register"
                  style="float: right; width: 48%; margin-top: 20px;"
                >
                  <p-button
                    v-if="!this.$store.state.app.isLoggedIn"
                    type="warning"
                    block
                    @click.native="createLogin"
                    style="overflow: visible;"
                  >
                    <span style="display: block;">
                      <i class="fa fa-lock"></i> CREATE!
                    </span>
                  </p-button>
                </div>
              </div>
              <div v-if="!this.$store.state.app.register" class="col-2"></div>
              <div class="col-2"></div>
            </div>

            <!-- <div class="clearfix"></div> -->
          </form>
        </div>
      </div>
    </center>
  </div>
</template>

<script>
import ErrorNotification from "./Notifications/ErrorMessage";
import { required, minLength, maxLength, between } from 'vuelidate/lib/validators'

export default {
  name: "login-screen",
  keystorePassword: "",
  keyPasswordValidate: "",
  storepass: "",
  keypass: "",
  alias: "",
  access: false,
  submitStatus: null,
  validations: {
      keyPasswordValidate: {
          required,
          minLength: minLength(10),
          maxLength: maxLength(40),
      },
      storekeyPasswordValidate: {
          required,
          inBetween: between(0.00000001, 3711998690),
      }
  },
  methods: {
        setKeyPassword(value) {
            this.keyPasswordValidate = value
            this.$v.keyPasswordValidate.$touch()
        },
    importKey: function() {
      window.backend.WalletApplication.ImportKey().then(
        result => {
          if (result) {
          this.$store.state.walletInfo.keystorePath = result;
          }
          // handle err
        }
      );
    },
    newLogin: function() {
      this.$store.state.app.register = !this.$store.state.app.register;
      this.$store.state.app.margin = 180;
    },
    cancelEvent: function() {
      this.$store.state.app.register = !this.$store.state.app.register;
      this.$store.state.app.margin = 100;
    },
    login: function() {
      var self = this;

        window.backend.WalletApplication.Login(self.$store.state.walletInfo.keystorePath, self.keystorePassword, self.keyPassword).then(
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
      window.backend.WalletApplication.CreateWallet(self.$store.state.walletInfo.keystorePath, self.keystorePassword, self.keyPassword
      ).then(result => {
        if (result) {
          window.backend.WalletApplication.Login(self.$store.state.walletInfo.keystorePath, self.keystorePassword, self.keyPassword
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

<style scoped lang="scss">

.vue-notifyjs.notifications {
  .alert {
    z-index: 10000;
  }
  .list-move {
    transition: transform 0.3s, opacity 0.4s;
  }
  .list-item {
    display: inline-block;
    margin-right: 10px;
  }
  .list-enter-active {
    transition: transform 0.2s ease-in, opacity 0.4s ease-in;
  }
  .list-leave-active {
    transition: transform 1s ease-out, opacity 0.4s ease-out;
  }
  .list-enter {
    opacity: 0;
    transform: scale(1.1);
  }
  .list-leave-to {
    opacity: 0;
    transform: scale(1.2, 0.7);
  }
}

.fadeout {
  animation: fadeout 2s backwards;
}

@keyframes fadeout {
  to {
    opacity: 0;
    visibility: hidden;
  }
}

p.validate {
    font-size: 10px;
    color: firebrick;
    margin-top: -5px;
}

body,
html {
  height: 100%;
}

.info-box {
text-align: left;
}

.bg {
  /* The image used */
  background-image: linear-gradient(
      rgba(255, 255, 255, 0.2),
      rgba(255, 255, 255, 0.2)
    ),
    url("~@/assets/img/nodes2.jpg");

  /* Full height */
  height: 100%;
  position: absolute;
  width: 100%;

  /* Center and scale the image nicely */
  background-position: center;
  background-repeat: no-repeat;
  background-size: cover;
}

</style>
