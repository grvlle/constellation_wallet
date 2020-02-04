<template>

  <div class="bg">
    
    <center>
      <div :style="'margin-top: ' + this.$store.state.app.margin + 'px'">
        <img v-if="this.$store.state.app.login" src="~@/assets/img/Constellation-Logo-1.png" />
        <p
          v-if="this.$store.state.app.login"
          style="margin-bottom: 20px; margin-top: 5px;"
        >Please enter your credentials below to access your Molly Wallet.</p>        
        <div style="height:30px;" v-if="!this.$store.state.displayLoginError"></div>
        <div style="height:30px;" v-if="this.$store.state.displayLoginError"><p style="color: firebrick; font-size: 12px;">{{this.$store.state.loginErrorMsg}}</p></div>
        <div>
          <form @submit.prevent="form">
            <div class="row">
              <div class="col-2"></div>
              <div v-if="!this.$store.state.app.register && !this.$store.state.app.import" class="col-2"></div>
              <div title="Create a New Wallet" v-if="this.$store.state.app.register && !this.$store.state.app.import" class="col-4 info-box">
                <p>
                  <br />
                  <b>Create a new wallet</b>
                  <br />
                  This section will let you create a Molly Wallet to store your <b>$DAG</b> tokens in. You simply browse to a path where you wish to save your KeyStore file, give it a name and select 'save'. <br><br /> 
                  Once the path is selected, you get to set up a password to protect the key store.<br /><br />

                  The Key Store will contain your private key. The only way to access that is with the KeyStore together with the KeyStore Password. The private key that is stored in the aforementioned file is also encrypted using a seperate Key Password. <br />
                  All three of these variables are required in order to access a $DAG Wallet.

                  <br />
                  <br />
                  <ul>
  <li><b>Key File</b><i> - Select where to save your private key. <b>You need to back this up</b> as it'll help you restore your wallet at any time. If you lose this, you will be locked out of the wallet.</i></li>
  <li><b>Store Password</b><i> - This password unlocks the keystore file. </i></li>
  <li><b>Key Password</b><i> - Extra layer of security. Both passwords will be needed when accessing/restoring a wallet.</i></li>
  <li><b>Token Label</b><i> - This will set the label of your wallet. This is <b>optional</b> and strictly for cosmetic purposes.</i></li>
</ul>
Please backup your passwords and Key Store file (key.p12) as these will allow you to restore your wallet at any time. 
                
                  </p>
              </div>


               <div title="Import Wallet" v-if="!this.$store.state.app.register && this.$store.state.app.import" class="col-4 info-box">
                <p>
                  <br />
                  <b>Import an existing wallet.</b>
                  <br />
                  This section will let you import an existing KeyStore (key.p12). Simply browse to the location of the KeyStore file, enter the Store Password as well as the Key Password to access it.<br />

                  <br />
                  <ul>
  <li><b>Key File</b><i> - Select where your <b>existing</b> private key is stored and unlock using the passwords previously set up.</i></li>
  <li><b>Store Password</b><i> - This password unlocks the keystore file. </i></li>
  <li><b>Key Password</b><i> - Extra layer of security. Both passwords will be needed when accessing/restoring a wallet.</i></li>
</ul>
If you're able to authenticate against the Key Store and Private Key, your Key Store will be unlocked and you'll be able to access your wallet.                
                  </p>
              </div>


              <div class="col-4">  
        <div v-if="this.$store.state.app.login && !this.$store.state.app.register && !this.$store.state.app.import">
          <p>Select your private key (key.p12)</p>
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
         <div v-if="this.$store.state.app.import && !this.$store.state.app.login && !this.$store.state.app.register">
        <p>Select the private key you wish to import.</p>
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

        <div v-if="!this.$store.state.app.import && !this.$store.state.app.login && this.$store.state.app.register">
        <p>Select a directory to store your private key (key.p12) in</p>
        <table style="width:100%;">
          <tr>
            <td style="padding: 0px; width: 81%;">
    
              <fg-input
                type="text"
                :disabled="true"
                :placeholder="this.$store.state.walletInfo.saveKeystorePath"
                v-model="this.$store.state.walletInfo.saveKeystorePath"
              ></fg-input>

      </td>

            <td style="padding-left: 0px;">
              <p-button
                @click.native="SelectDirToStoreKey"
                type="default"
                style="margin-top: -17px; width: 95%; float: right;"
              ><span style="display: block;">
               BROWSE
              </span></p-button>
         
            </td>
          </tr>
        </table>
              </div>
                <div class="fg-style">
                  <fg-input
                    type="text"
                    v-model="alias"
                    @input="checkAlias(alias)"
                    :placeholder="this.$store.state.walletInfo.alias"
                    label="Key Alias"
                  ></fg-input>
                </div>
                <div style="height: 30px; margin-top: -30px;" v-if="!this.$store.state.validators.duplicate && !this.$store.state.app.login && !this.$store.state.validators.alias.valid_alias">
      
                  <p v-if="!this.$store.state.validators.alias.contains_five_characters" class="validate"> Alias has to be atleast 5 characters long. </p>  
                
                </div> 
                
                <div class="fg-style">
                <fg-input @input="checkPassword(keystorePassword)" type="password" v-model="keystorePassword" label="Keystore Password" placeholder="Enter Keystore Password ..." />
                </div>
                 <div style="height: 30px; margin-top: -30px;" v-if="!this.$store.state.validators.duplicate && !this.$store.state.app.login && !this.$store.state.validators.storepass.valid_password">             
                            <p v-if="!this.$store.state.validators.storepass.contains_eight_characters" class="validate"> 8 Characters Long, </p> 
                            <p v-if="!this.$store.state.validators.storepass.contains_number" class="validate"> Number,</p> 
                            <p v-if="!this.$store.state.validators.storepass.contains_uppercase" class="validate"> Uppercase, </p> 
                            <p v-if="!this.$store.state.validators.storepass.contains_special_character" class="validate"> Special Character </p>     
                </div>
                <div class="fg-style">
                  <fg-input
                    @input="checkKeyPassword(keyPasswordValidate)"
                    v-model="keyPasswordValidate"
                    type="password"
                    label="Key Password"
                    placeholder="Enter Key Password..."
                  ></fg-input>
                </div>
                <!-- <div style="height: 30px; margin-top: -30px;" v-if="this.$store.state.app.register && this.$store.state.validators.duplicate && this.keyPasswordValidate !== ''">
                <p class="validate"> Keystore Password cannot be the same as the Key Password</p>
                </div> -->
                  
                 <div style="height: 30px; margin-top: -30px;" v-if="!this.$store.state.validators.duplicate && !this.$store.state.app.login && !this.$store.state.validators.valid_password">
               
                            <p v-if="!this.$store.state.validators.contains_eight_characters" class="validate"> 8 Characters Long, </p> 
                            <p v-if="!this.$store.state.validators.contains_number" class="validate"> Number,</p> 
                            <p v-if="!this.$store.state.validators.contains_uppercase" class="validate"> Uppercase, </p> 
                            <p v-if="!this.$store.state.validators.contains_special_character" class="validate"> Special Character </p> 
                          
                </div> 



                <div v-if="!this.$store.state.app.import && !this.$store.state.app.login && this.$store.state.app.register">
                  <fg-input
                    type="text"
                    v-model="newWalletLabel"
                    :placeholder="this.$store.state.walletInfo.email"
                    label="Wallet Label (optional)"
                  ></fg-input>
                </div>

                <div
                  v-if="!this.$store.state.app.import && this.$store.state.app.login && !this.$store.state.app.register"
                  style="margin-top: 20px;"
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
                  v-if="!this.$store.state.app.import && this.$store.state.app.login && !this.$store.state.app.register"
                  style="float: left; width: 48%; margin-top: 20px;"
                >
                            <p-button
                    v-if="!this.$store.state.app.isLoggedIn"
                    type="info"
                    block
                    @click.native="showImportView"
                    style="overflow: visible;"
                  >
                    <span style="display: block;">
                      <i class="fas fa-file-import"> </i> IMPORT WALLET
                    </span>
                  </p-button>

                </div>

                <div
                  v-if="!this.$store.state.app.import && this.$store.state.app.login && !this.$store.state.app.register"
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
                  v-if="!this.$store.state.app.import && !this.$store.state.app.login && this.$store.state.app.register"
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
                  v-if="!this.$store.state.app.import && !this.$store.state.app.login && this.$store.state.app.register"
                  style="float: right; width: 48%; margin-top: 20px;"
                >
                <!-- :disabled="!this.$store.state.validators.valid_password && this.keystorePassword !== '' && this.keyPasswordValidate !== ''" -->
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
                <div
                  v-if="this.$store.state.app.import && !this.$store.state.app.login && !this.$store.state.app.register"
                  style="float: left; width: 48%; margin-top: 20px;"
                >
                  <p-button
                    v-if="!this.$store.state.app.isLoggedIn"
                    type="default"
                    block
                    @click.native="cancelImportView"
                    style="overflow: visible;"
                  >
                    <span style="display: block;">
                      <i class="fa fa-close"></i> CANCEL
                    </span>
                  </p-button>
                </div>
                                <div
                  v-if="this.$store.state.app.import && !this.$store.state.app.login && !this.$store.state.app.register"
                  style="float: right; width: 48%; margin-top: 20px;"
                >
                <!-- :disabled="!this.$store.state.validators.valid_password && this.keystorePassword !== '' && this.keyPasswordValidate !== ''" -->
                  <p-button
                    v-if="!this.$store.state.app.isLoggedIn"
                    type="info"
                    block
                    
                    @click.native="importWallet()"
                    style="overflow: visible;"
                  >
                    <span style="display: block;">
                      <i class="fas fa-file-import"></i> IMPORT KEY!
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
import { required, minLength } from "vuelidate/lib/validators";

export default {
  name: "login-screen",
  newWalletLabel: "",
  keystorePassword: '',
  keyPasswordValidate: '',
  storepass: "",
  keypass: "",
  alias: "",
  password_length: 0,
  contains_eight_characters: false,
  contains_number: false,
  contains_uppercase: false,
  contains_special_character: false,
  valid_password: false,
  access: false,
  submitStatus: null,
  methods: {
    importWallet: function() {
      var self = this;
      self.$Progress.start();
      window.backend.WalletApplication.ImportWallet(self.$store.state.walletInfo.keystorePath, self.keystorePassword, self.keyPasswordValidate, self.alias
      ).then(walletImported => {
        if (walletImported) {
          window.backend.WalletApplication.Login(self.$store.state.walletInfo.keystorePath, self.keystorePassword, self.keyPasswordValidate, self.alias
          ).then(loggedIn => {
            self.access = loggedIn;
            if (self.access) {
              self.$store.state.app.isLoading = self.access;
              self.$store.state.app.isLoggedIn = self.access;
              self.$Progress.finish();
              setTimeout(() => {
                self.$store.state.app.isLoading = false;
              }, 8000);
            }
          });
        }
        self.$Progress.finish();
      });
    },
    checkAlias: function() {
      this.$store.state.validators.target = this.alias;
      this.$store.state.validators.alias.alias_length = this.alias.length;

      if (this.$store.state.validators.alias.alias_length >= 5) {
        this.$store.state.validators.alias.contains_five_characters = true;
      } else {
        this.$store.state.validators.alias.contains_five_characters = false;
      }

    },
    checkKeyPassword: function() {
      this.$store.state.validators.target = this.keyPasswordValidate
      this.$store.state.validators.password_length = this.keyPasswordValidate.length;
      const format = /[ !@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]/;
      
      if (this.$store.state.validators.password_length >= 8) {
        this.$store.state.validators.contains_eight_characters = true;
      } else {
        this.$store.state.validators.contains_eight_characters = false;
      }
      
      this.$store.state.validators.contains_number = /\d/.test(this.keyPasswordValidate);
      this.$store.state.validators.contains_uppercase = /[A-Z]/.test(this.keyPasswordValidate);
      this.$store.state.validators.contains_special_character = format.test(this.keyPasswordValidate);
      
      if (this.$store.state.validators.contains_eight_characters === true &&
          this.$store.state.validators.contains_special_character === true &&
          this.$store.state.validators.contains_uppercase === true &&
          this.$store.state.validators.contains_number === true) {
            this.$store.state.validators.valid_password = true;			
      } else {
        this.$store.state.validators.valid_password = false;
      }
    },
    checkPassword: function() {
      this.$store.state.validators.target = this.keystorePassword
      this.$store.state.validators.storepass.password_length = this.keystorePassword.length;
      const format = /[ !@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]/;
      
      if (this.$store.state.validators.storepass.password_length > 8) {
        this.$store.state.validators.storepass.contains_eight_characters = true;
      } else {
        this.$store.state.validators.storepass.contains_eight_characters = false;
      }
      
      this.$store.state.validators.storepass.contains_number = /\d/.test(this.keystorePassword);
      this.$store.state.validators.storepass.contains_uppercase = /[A-Z]/.test(this.keystorePassword);
      this.$store.state.validators.storepass.contains_special_character = format.test(this.keystorePassword);
      
      if (this.$store.state.validators.storepass.contains_eight_characters === true &&
          this.$store.state.validators.storepass.contains_special_character === true &&
          this.$store.state.validators.storepass.contains_uppercase === true &&
          this.$store.state.validators.storepass.contains_number === true) {
            this.$store.state.validators.storepass.valid_password = true;			
      } else {
        this.$store.state.validators.storepass.valid_password = false;
      }
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
    SelectDirToStoreKey: function() {
      window.backend.WalletApplication.SelectDirToStoreKey().then(
        result => {
          this.$store.state.walletInfo.saveKeystorePath = result;
          this.$store.state.walletInfo.keystorePath = result;
          
          // handle err
        }
      );
    },
    showImportView: function() {
      this.$store.state.app.import = !this.$store.state.app.import;
      this.$store.state.app.login = !this.$store.state.app.login;
      this.$store.state.app.margin = 150;
    },
    cancelImportView: function() {
      this.$store.state.app.import = !this.$store.state.app.import;
      this.$store.state.app.login = !this.$store.state.app.login;
      this.$store.state.app.margin = 70;
    },
    newLogin: function(event) {
      this.$store.state.app.register = !this.$store.state.app.register;
      this.$store.state.app.login = !this.$store.state.app.login;
      this.$store.state.app.margin = 100;
    },
    cancelEvent: function(event) {
      this.$store.state.app.register = !this.$store.state.app.register;
      this.$store.state.app.login = !this.$store.state.app.login;
      this.$store.state.app.margin = 70;
    },
    login: function() {
      var self = this;
        self.$Progress.start();
        window.backend.WalletApplication.Login(self.$store.state.walletInfo.keystorePath, self.keystorePassword, self.keyPasswordValidate, self.alias)
        .then(result => {
          self.access = result;
          window.backend.WalletApplication.SetWalletTag().then(walletTag =>
            self.$store.state.walletInfo.email = walletTag
          )
          window.backend.WalletApplication.SetImagePath().then(imagePath =>
            self.$store.state.walletInfo.imgPath = imagePath
          )
          if (self.access) {
            self.$store.state.app.isLoading = self.access;
            self.$store.state.app.isLoggedIn = self.access;
            self.$Progress.finish();
            setTimeout(() => {
              self.$store.state.app.isLoading = false;
            }, 8000);
          }
        
        }
      );
    },
    createLogin: function() {
      // if (this.$store.state.validators.valid_password) {
      var self = this;
      self.$Progress.start();
      self.$store.state.walletInfo.email = self.newWalletLabel
      
      window.backend.WalletApplication.StoreWalletLabelInDB(self.newWalletLabel)
      window.backend.WalletApplication.CreateWallet(self.$store.state.walletInfo.keystorePath, self.keystorePassword, self.keyPasswordValidate, self.alias
      ).then(walletCreated => {
        if (walletCreated) {
          window.backend.WalletApplication.Login(self.$store.state.walletInfo.keystorePath, self.keystorePassword, self.keyPasswordValidate, self.alias
          ).then(loggedIn => {
            self.access = loggedIn;
            if (self.access) {
              self.$store.state.app.isLoading = self.access;
              self.$store.state.app.isLoggedIn = self.access;
              self.$Progress.finish();
              setTimeout(() => {
                self.$store.state.app.isLoading = false;
              }, 8000);
            }
          });
        }
      });
      // }
    }
  },
  
}


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
    margin-top: 0px;
    margin-right: 2px;
    float: left;
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
  overflow:hidden;

  /* Center and scale the image nicely */
  background-position: center;
  background-repeat: no-repeat;
  background-size: cover;
}

.fg-style {
  margin-bottom: 30px;
}

</style>
