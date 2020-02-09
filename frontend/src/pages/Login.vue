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
          <form ref='textareaform' @submit.prevent="form">
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

        <div style="margin-top: 20px;" v-if="!this.$store.state.app.import && !this.$store.state.app.login && this.$store.state.app.register">
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
                    @input.native="checkAlias(alias)"
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
                    v-model="keyPasswordValidate"
                    @input="checkKeyPassword(keyPasswordValidate)"
                    type="password"
                    label="Key Password"
                    placeholder="Enter Key Password..."
                  />
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
                    @click.native="login()"
                  :disabled="!this.$store.state.validators.valid_password || !this.$store.state.validators.storepass.valid_password || this.loginInProgress && this.alias !== '' && this.keystorePassword !== '' && this.keyPasswordValidate !== ''"

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
                    @click.native="showImportView()"
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
                    @click.native="newLogin()"
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
                    @click.native="cancelEvent()"
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
                
                  <p-button
                    v-if="!this.$store.state.app.isLoggedIn"
                    type="warning"
                    block
                    :disabled="!this.$store.state.validators.valid_password || !this.$store.state.validators.storepass.valid_password || this.loginInProgress && this.alias !== '' && this.keystorePassword !== '' && this.keyPasswordValidate !== ''"
                    @click.native="createLogin()"
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
                    @click.native="cancelImportView()"
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
                
                  <p-button
                    v-if="!this.$store.state.app.isLoggedIn"
                    type="info"
                    block
                    :disabled="!this.$store.state.validators.valid_password || !this.$store.state.validators.storepass.valid_password || this.loginInProgress && this.alias !== '' && this.keystorePassword !== '' && this.keyPasswordValidate !== ''"
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
import Swal from "sweetalert2";

export default {
  name: "login-screen",
  newWalletLabel: "",
  keystorePassword: '',
  keyPasswordValidate: '',
  loginInProgress: false,
  storepass: "",
  keypass: "",
  alias: "",
  doneLoading: false,
  password_length: 0,
  contains_eight_characters: false,
  contains_number: false,
  contains_uppercase: false,
  contains_special_character: false,
  valid_password: false,
  access: false,
  submitStatus: null,
  termsOfService: "This HTML scroll box has had color added. You can add color to the background of your scroll box. You can also add color to the scroll bars",
  methods: {
    importWallet: function() {
      var self = this;
      self.$Progress.start();
      self.loginInProgress = true;
      window.backend.WalletApplication.ImportWallet(self.$store.state.walletInfo.keystorePath, self.keystorePassword, self.keyPasswordValidate, self.alias
      ).then(walletImported => {
        if (walletImported) {
          window.backend.WalletApplication.Login(self.$store.state.walletInfo.keystorePath, self.keystorePassword, self.keyPasswordValidate, self.alias
          ).then(loggedIn => {
            self.access = loggedIn;
            if (self.access) {
              self.loginInProgress = false;
              self.$store.state.app.isLoading = self.access;
              self.$store.state.app.isLoggedIn = self.access;
              self.$Progress.finish();
              setTimeout(() => {
                self.$store.state.app.isLoading = false;
              }, 8000);
            } else {
              self.loginInProgress = false;
              self.$Progress.fail();
            }
          });
        } else {
          self.loginInProgress = false;
          self.$Progress.fail();
        }
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
      const format = /[ !@#$%^&*()_+\-=[\]{};':"\\|,.<>/?]/;
      
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
      const format = /[ !@#$%^&*()_+\-=[\]{};':"\\|,.<>/?]/;
      
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
      this.alias = ''
      this.keystorePath = ''
      this.keyPasswordValidate = ''
      this.keystorePassword = ''
      this.$store.state.walletInfo.keystorePath = ''
      this.$store.state.walletInfo.alias = ''
      this.$store.state.walletInfo.keystorePassword = ''
      this.$store.state.walletInfo.keyPasswordValidate = ''
      this.$store.state.app.import = !this.$store.state.app.import;
      this.$store.state.app.login = !this.$store.state.app.login;
      this.$store.state.app.margin = 150;
    },
    cancelImportView: function() {
      this.alias = ''
      this.keystorePath = ''
      this.keyPasswordValidate = ''
      this.keystorePassword = ''
      this.$store.state.walletInfo.keystorePath = ''
      this.$store.state.walletInfo.alias = ''
      this.$store.state.walletInfo.keystorePassword = ''
      this.$store.state.walletInfo.keyPasswordValidate = ''
      this.$store.state.app.import = !this.$store.state.app.import;
      this.$store.state.app.login = !this.$store.state.app.login;
      this.$store.state.app.margin = 70;
    },
    newLogin: function() {
      this.alias = ''
      this.keystorePath = ''
      this.keyPasswordValidate = ''
      this.keystorePassword = ''
      this.$store.state.walletInfo.keystorePath = ''
      this.$store.state.walletInfo.alias = ''
      this.$store.state.walletInfo.keystorePassword = ''
      this.$store.state.walletInfo.keyPasswordValidate = ''
      this.$store.state.app.register = !this.$store.state.app.register;
      this.$store.state.app.login = !this.$store.state.app.login;
      this.$store.state.app.margin = 100;
    },
    cancelEvent: function() {
      this.alias = ''
      this.keystorePath = ''
      this.keyPasswordValidate = ''
      this.keystorePassword = ''
      this.$store.state.walletInfo.keystorePath = ''
      this.$store.state.walletInfo.alias = ''
      this.$store.state.walletInfo.keystorePassword = ''
      this.$store.state.walletInfo.keyPasswordValidate = ''
      this.$store.state.app.register = !this.$store.state.app.register;
      this.$store.state.app.login = !this.$store.state.app.login;
      this.$store.state.app.margin = 70;
    },
    login: function() {
      var self = this;
        self.$Progress.start();
        self.loginInProgress = true
        window.backend.WalletApplication.Login(self.$store.state.walletInfo.keystorePath, self.keystorePassword, self.keyPasswordValidate, self.alias)
        .then(result => {
          self.access = result;
          if (self.access) {
            window.backend.WalletApplication.SetWalletTag().then(walletTag =>
            self.$store.state.walletInfo.email = walletTag
          )
            window.backend.WalletApplication.SetImagePath().then(imagePath =>
            self.$store.state.walletInfo.imgPath = imagePath
          )
            self.loginInProgress = false;
            self.$store.state.app.isLoading = self.access;
            self.$store.state.app.isLoggedIn = self.access;
            self.$Progress.finish();
            setTimeout(() => {
              self.$store.state.app.isLoading = false;
            }, 8000);
          } else {
            self.loginInProgress = false;
            self.$Progress.fail();
          }
        
        }
      );
    },
    createLogin: function() {
      // if (this.$store.state.validators.valid_password) {
      var self = this;
      self.$Progress.start();
      self.loginInProgress = true;
      if (self.newWalletLabel !== '') {
          self.$store.state.walletInfo.email = self.newWalletLabel
          window.backend.WalletApplication.StoreWalletLabelInDB(self.newWalletLabel)
      }
      window.backend.WalletApplication.CreateWallet(self.$store.state.walletInfo.keystorePath, self.keystorePassword, self.keyPasswordValidate, self.alias
      ).then(walletCreated => {
        if (walletCreated) {
          window.backend.WalletApplication.Login(self.$store.state.walletInfo.keystorePath, self.keystorePassword, self.keyPasswordValidate, self.alias
          ).then(loggedIn => {
            self.access = loggedIn;
            if (self.access) {
              self.loginInProgress = false;
              self.$store.state.app.isLoading = self.access;
              self.$store.state.app.isLoggedIn = self.access;
                  setTimeout(function () {
                    self.$store.state.app.isLoading = false
                    Swal.fire({
                      html:
                      '<div style="overflow: scroll; padding: 20px; width: 1160px; height: 600px;">'+
                        '<p style="text-align: center;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>TERMS OF SERVICE</strong></p>'+
            '<p style="text-align: center;background: transparent;margin-bottom: 0.11in;line-height: 108%;">Last updated: 02.04.2020</p>'+
            '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">Welcome to Molly, a free tool for interacting directly with the Hypergraph Blockchain. Please read these terms and conditions (the <strong>"</strong><u>Terms of Service</u><strong>"</strong>) carefully. They apply to your use of the Constellation Network, Inc.’s, and any subsidiary, parent, or affiliate thereof (collectively, the “<u>Company</u>") Molly Wallet desktop application (the&nbsp;"<u>Site</u>") and any related services offered by the Company through the Site. The Site and related services offered through the Site shall hereinafter be referred to as the “<u>Services</u>.” This Terms of Service together with any additional posted guidelines or rules applicable to related services and features, and the Privacy Policy (as hereinafter defined) shall hereinafter be referred to as the “<u>Agreement</u>.”</p>'+
            '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>1.&nbsp;</strong><u><strong>Introduction</strong></u></p>'+
            '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">This Agreement sets forth the legally binding terms for your use of the Services. By using the Services, you agree to be bound by this Agreement, any additional posted guidelines or rules applicable to specific services and features, and our&nbsp;Privacy Policy. If you are accepting this Agreement on behalf of a company or other legal entity, you represent and warrant that you have the authority to bind such entity to the terms set forth herein. If you do not have such authority or do not agree to be bound by this Agreement, you may not access or use the Services. You must agree to this Agreement when you create a wallet via the Site, perform a transaction via the Site, and/or otherwise use the Services.</p>'+
            '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">THIS AGREEMENT CONTAINS IMPORTANT INFORMATION REGARDING YOUR LEGAL RIGHTS, REMEDIES, AND OBLIGATIONS. PLEASE NOTE THAT SECTION 10 OF THIS AGREEMENT INCLUDES AN ARBITRATION AGREEMENT. EXCEPT FOR CERTAIN TYPES OF DISPUTES MENTIONED IN THAT CLAUSE, YOU AND THE COMPANY AGREE THAT DISPUTES BETWEEN US WILL BE RESOLVED BY MANDATORY BINDING ARBITRATION, AND YOU WAIVE ANY RIGHT TO PARTICIPATE IN A CLASS ACTION LAWSUIT OR CLASS ARBITRATION.</p>'+
            '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">The Company may, in its sole discretion, modify or revise this Agreement at any time, and you agree to be bound by such modifications or revisions. Although we may attempt to notify you when major changes are made to this Agreement, you should periodically review the most up-to-date version, which will always be posted at '+
            'http://www.constellationnetwork.com. Your continued use of the Services constitutes your acceptance of such changes.</p>'+
            '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>2.&nbsp;</strong><u><strong>Services Eligibility and Information</strong></u></p>'+
            '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>2.1 Services Eligibility</strong></p>'+
            '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">The Services are offered and available to users who are 18 years of age or older. Any registration by, use of or access to the Services by anyone under 18 is unauthorized and in violation of this Agreement. By using the Services, you represent and warrant that you are 18 years of age or older and that you agree to abide by this Agreement.</p>'+
            '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>2.2 Services Information and Limitations</strong></p>'+
            '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">The Site is a free interface that allows you to interact directly with the Company’s Hypergraph blockchain (the “<u>Blockchain</u>”), while remaining in full control of your keys.</p>'+
            '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">When you access certain features of the Services, you will be able to create a wallet and/or access a wallet to perform a variety of transactions. You will receive a private key and set up a password, but you do not and will not create an account with the Company. You also will not give us any of your tokens, and your tokens are not on the Site or otherwise in the custody of the Company. The only data that leaves your computer is the data that is recorded on the Hypergraph blockchain. The Company does not collect or hold your private keys or information, and the Company cannot access accounts; recover keys, passwords, or other information; reset passwords; or reverse transactions. You are solely responsible for your use of the Services, including without limitation for storing, backing-up, and maintaining the confidentiality of your private keys, passwords, and information, and for the security of any transactions you perform using the Site. You expressly relieve and release the Company from any and all liability and/or loss arising from your use of the Services.</p>'+
            '<p style="margin-bottom: 0.11in;line-height: 108%;text-align: left;background: transparent;">Prior to using the Services for any purpose, we highly recommend that you read our guides, https://constellationnetwork.io/token-swap-information/ on for some recommendations on how to be proactive about your security. In addition, we recommend that you review the additional FAQs, tips, and guidelines provided on the Site.</p>'+
            '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>3.&nbsp;</strong><u><strong>Rights and Restrictions</strong></u></p>'+
            '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>3.1 Our Proprietary Rights</strong></p>'+
            '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">Our Services are protected by copyright, trademark, and other laws of the United States and foreign countries. Except as expressly provided in this Agreement, we (or our licensors) exclusively own all right, title and interest in and to the Services, including all associated intellectual property rights. You may not remove, alter or obscure any copyright, trademark, service mark or other proprietary rights notices incorporated in or accompanying the Services, including in any content therefrom. You acknowledge and agree that any feedback, comments or suggestions you may provide regarding the Services will be the sole and exclusive property of the Company, and you hereby irrevocably assign to us all of your right, title and interest in and to the foregoing. The Company reserves the right to discontinue any aspect of the Site at any time.</p>'+
            '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>3.2 Restrictions</strong></p>'+
            '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">You agree not to (i) interfere with, damage, impair, or disable the Services\' operation, by any means (whether through automated means or otherwise), including uploading or otherwise disseminating viruses, worms, spyware, adware or other malicious code on the Site or to any users of the Services; (ii) use any robot, spider, scraper, or other automated means to access the Services for any purpose without our express consent or bypass our robot exclusion headers or similar measures; (iii) remove, circumvent, disable, damage or otherwise interfere with the Services\' security-related features, features that prevent or restrict the use or copying of any part of the Services, or features that enforce limitations of the Services; (iv) take any action that imposes, or may impose, in our sole discretion, an unreasonable or disproportionately large load on our technology infrastructure or otherwise make excessive traffic demands of the Service; (v) use the Services for any illegal or unauthorized purpose nor may you, in the use of the Service, violate any laws in your jurisdiction (including but not limited to intellectual property laws); or (vi) use, including as part of trademarks and/or as part of domain names, in connection with any product or service in any manner that is likely to cause confusion and may not be copied, imitated, or used, in whole or in part, without the prior written permission of the Company.</p>'+
            '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>4.&nbsp;</strong><u><strong>Privacy</strong></u></p>'+
            '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">The Company’s Privacy Policy applies to the collection and use of “<u>Personal Information</u>,” which is defined as information about an individual that (either by itself or when combined with information from other available sources) allows that individual to be identified. The Company’s Privacy Policy is available at&nbsp;http://www.constellationnetwork.io&nbsp;and its terms are made a part of this Agreement by this reference. You understand that by using the Services you consent to the collection, use and disclosure of your Personal Information as set forth in our Privacy Policy, and to have your Personal Information collected, used, transferred to and processed in the United States.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>5.&nbsp;</strong><u><strong>Third Party Content and Services</strong></u></p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">The Site may contain links and/or access to third-party websites, applications (together, “<u>Third-party Sites</u>”), and/or application programming interfaces (“<u>APIs</u>”) that are not owned or controlled by the Company. Access to and use of any Third-party Sites or APIs is at your own risk and we are not responsible for (i) the accuracy or reliability of information on Third-party Sites or APIs; (ii) the acts or omissions of the operators of Third-party Sites or APIs (or their partners or affiliates); (iii) any loss or damage incurred in connection with the use of any Third-party Sites or APIs, or (iv) any transaction you consummate in connection with your use or access of any Third-party Sites or APIs. We encourage you to be aware when you leave the Site and/or access third-party APIs, and to read the terms and privacy policy of each other Third-party Site that you visit or API that you access. We provide such links and access merely as a convenience, and the inclusion of such links or access does not imply an endorsement. Upon leaving the Site and/or when using third-party APIs, this Agreement shall no longer govern, provided, however, this Section 5 shall apply. By using the Site, you expressly relieve and release the Company from any and all liability arising from your use of any Third-party Sites or APIs.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>6.&nbsp;</strong><u><strong>Disclaimers, Limitations of Liability</strong></u></p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">YOUR USE OF THE SITE, ITS CONTENT AND ANY SERVICES OFFERED THROUGH THE SITE IS AT YOUR OWN RISK. THE SITE, ITS CONTENT AND ANY SERVICES OFFERED THROUGH THE SITE ARE PROVIDED ON AN "AS IS" AND "AS AVAILABLE" BASIS, WITHOUT ANY WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED. NEITHER THE COMPANY NOR ANY PERSON OR ENTITY ASSOCIATED WITH THE COMPANY MAKES ANY WARRANTY OR REPRESENTATION WITH RESPECT TO THE COMPLETENESS, SECURITY, RELIABILITY, QUALITY, ACCURACY OR AVAILABILITY OF THE SITE OR THE SERVICES. WITHOUT LIMITING THE FOREGOING, NEITHER THE COMPANY NOR ANY PERSON OR ENTITY ASSOCIATED WITH THE COMPANY REPRESENTS OR WARRANTS THAT THE SITE, ITS CONTENT OR ANY SERVICES OFFERED THROUGH THE SITE WILL BE ACCURATE, RELIABLE, ERROR-FREE OR UNINTERRUPTED, THAT DEFECTS WILL BE CORRECTED, THAT THE SITE OR THE SERVER THAT MAKES IT AVAILABLE ARE FREE OF VIRUSES OR OTHER HARMFUL COMPONENTS OR THAT THE SITE OR ANY SERVICES OFFERED THROUGH THE SITE WILL OTHERWISE MEET YOUR NEEDS OR EXPECTATIONS.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">THE COMPANY HEREBY DISCLAIMS ALL WARRANTIES OF ANY KIND, WHETHER EXPRESS OR IMPLIED STATUTORY OR OTHERWISE, INCLUDING BUT NOT LIMITED TO ANY WARRANTIES OF MERCHANTABILITY, NON-INFRINGEMENT AND FITNESS FOR PARTICULAR PURPOSE.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">IN NO EVENT WILL THE COMPANY, ITS AFFILIATES OR THEIR LICENSORS, SERVICE PROVIDERS, EMPLOYEES, AGENTS, OFFICERS OR DIRECTORS BE LIABLE FOR DAMAGES OF ANY KIND, UNDER ANY LEGAL THEORY, ARISING OUT OF OR IN CONNECTION WITH YOUR USE, OR INABILITY TO USE, THE SITE, ANY WEBSITES LINKED TO IT, ANY CONTENT ON THE SITE OR SUCH OTHER WEBSITES, OR ANY SERVICES OFFERED THROUGH THE SITE OR SUCH OTHER WEBSITES, INCLUDING ANY DIRECT, INDIRECT, SPECIAL, INCIDENTAL, CONSEQUENTIAL OR PUNITIVE DAMAGES, INCLUDING BUT NOT LIMITED TO, PERSONAL INJURY, PAIN AND SUFFERING, EMOTIONAL DISTRESS, LOSS OF REVENUE, LOSS OF PROFITS, LOSS OF BUSINESS OR ANTICIPATED SAVINGS, LOSS OF USE, LOSS OF GOODWILL, LOSS OF DATA, AND WHETHER CAUSED BY TORT (INCLUDING NEGLIGENCE), BREACH OF CONTRACT OR OTHERWISE, EVEN IF FORESEEABLE.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">The limitations and disclaimers in this section do not purport to limit liability or alter your rights beyond what is permitted by applicable law. The Company’s liability shall be limited to the extent permitted by law.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">We provide the Site for use only by persons located in the United States. We make no claims that the Site or any of its content is accessible or appropriate outside of the United States. Access to the Site may not be legal by certain persons or in certain countries. If you access the Site from outside the United States, you do so on your own initiative and are responsible for compliance with local laws. Similarly, the Site has been translated into a variety of languages. However, the Company can only verify the validity and accuracy of the information provided in English and, because of this, the English version of the Site is the official text.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>7.&nbsp;</strong><u><strong>Risks</strong></u></p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>7.1 Sophistication and Risk of Cryptographic Systems</strong></p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">By utilizing the Services or interacting with the Site, you represent that you understand the inherent risks associated with cryptographic systems; and warrant that you have an understanding of the usage and intricacies of native cryptographic tokens and smart contract based tokens.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>7.2 Risk of Regulatory Actions in One or More Jurisdictions</strong></p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">The Company, the Site, and the Services could be impacted by one or more regulatory inquiries or regulatory action, which could impede or limit the ability of the Company to continue to develop, or which could impede or limit your ability to access or use the Services.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>7.3 Risk of Weaknesses or Exploits in the Field of Cryptography</strong></p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">You acknowledge and understand that cryptography is a progressing field. Advances in code cracking or technical advances such as the development of quantum computers may present risks to the Services, which could result in the theft or loss of your cryptographic tokens or property. To the extent possible, the Company intends to update the protocol underlying the Services to account for any advances in cryptography and to incorporate additional security measures, but does not guarantee or otherwise represent full security of the system. By using the Services, you acknowledge these inherent risks.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>8.&nbsp;</strong><u><strong>Indemnity</strong></u></p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">You agree to defend, indemnify and hold harmless the Company, its officers, directors, employees and agents from and against any and all claims, damages, obligations, losses, liabilities, costs or debt, and expenses (including but not limited to attorneys’ fees) arising from your use of and access to the Services, or your violation of this Agreement. This defense and indemnification obligation will survive this Agreement and your use of the Services.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>9.&nbsp;</strong><u><strong>Assignment</strong></u></p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">This Agreement, and any rights and licenses granted hereunder, may not be transferred or assigned by you, but may be assigned by the Company without restriction.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>10.&nbsp;</strong><u><strong>Arbitration Agreement; Class Waiver; Jury Waiver</strong></u></p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">PLEASE READ THIS SECTION 10 (THE “<u>ARBITRATION AGREEMENT</u>”) CAREFULLY BECAUSE IT REQUIRES YOU AND THE COMPANY TO AGREE TO RESOLVE ALL DISPUTES THROUGH BINDING INDIVIDUAL ARBITRATION, UNLESS OTHERWISE NOTED.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>10.1 Applicability of Arbitration Agreement.</strong></p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">All claims and disputes (excluding claims for injunctive or other equitable relief as set forth below) in connection with this Agreement that cannot be resolved informally or in small claims court shall be resolved, to the extent permitted by applicable law, by binding arbitration on an individual basis under the terms of this Arbitration Agreement. This Arbitration Agreement applies to you and the Company, and to any subsidiaries, affiliates, agents, employees, predecessors in interest, successors, and assigns, as well as all authorized or unauthorized users or beneficiaries of services or goods provided under the Agreement.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>10.2 Notice Requirement and Informal Dispute Resolution.&nbsp;</strong></p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">Before either party may seek arbitration, the party must first send to the other party a written notice of dispute ("<u>Notice</u>") describing the nature and basis of the claim or dispute, and the requested relief. A Notice to the Company should be sent to: 480 5th Street, San Francisco, CA 94107. After the Notice is received, you and the Company may attempt to resolve the claim or dispute informally. If you and the Company do not resolve the claim or dispute within 30 days after the Notice is received, either party may begin an arbitration proceeding. The amount of any settlement offer made by any party may not be disclosed to the arbitrator until after the arbitrator has determined the amount of the award, if any, to which either party is entitled.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>10.3 Arbitration Rules.&nbsp;</strong></p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">Arbitration shall be initiated through the American Arbitration Association ("<u>AAA</u>"), an established alternative dispute resolution provider ("<u>ADR Provider</u>") that offers arbitration as set forth in this section. If AAA is not available to arbitrate, the parties shall agree to select an alternative ADR Provider. The rules of the ADR Provider shall govern all aspects of this arbitration, including but not limited to the method of initiating and/or demanding arbitration, except to the extent such rules conflict with the Agreement. The AAA Consumer Arbitration Rules ("<u>Arbitration Rules</u>") governing the arbitration are available online at www.adr.org or by calling the AAA at 1-800-778-7879. The arbitration shall be conducted by a single, neutral arbitrator. Any claims or disputes where the total amount of the award sought is less than ten thousand U.S. Dollars (US $10,000.00) may be resolved through binding non-appearance-based arbitration, at the option of the party seeking relief. For claims or disputes where the total amount of the award sought is ten thousand U.S. Dollars (US $10,000.00) or more, the right to a hearing will be determined by the Arbitration Rules. Any hearing will be held in the city of San Francisco, California, unless the parties agree otherwise. Any judgment on the award rendered by the arbitrator may be entered in any court of competent jurisdiction.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>10.4 Additional Rules for Non-Appearance Based Arbitration.&nbsp;</strong></p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">If non-appearance arbitration is elected, the arbitration shall be conducted by telephone, online, and/or based solely on written submissions; the specific manner shall be chosen by the party initiating the arbitration.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>10.5 Authority of Arbitrator.&nbsp;</strong></p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">If arbitration is initiated, the arbitrator will decide the rights and liabilities, if any, of you and the Company, and the dispute will not be consolidated with any other matters or joined with any other cases or parties. The arbitrator shall have the authority to grant motions dispositive of all or part of any claim. The arbitrator shall have the authority to award monetary damages and to grant any non-monetary remedy or relief available to an individual under applicable law, the AAA Rules, and the Agreement. The arbitrator shall issue a written award and statement of decision describing the essential findings and conclusions on which the award is based, including the calculation of any damages awarded. The arbitrator has the same authority to award relief on an individual basis that a judge in a court of law would have. The award of the arbitrator is final and binding upon you and the Company.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>10.6 Waiver of Jury Trial.&nbsp;</strong></p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">THE PARTIES HEREBY WAIVE THEIR CONSTITUTIONAL AND STATUTORY RIGHTS TO GO TO COURT AND HAVE A TRIAL IN FRONT OF A JUDGE OR A JURY, instead electing that all claims and disputes shall be resolved by arbitration under this Arbitration Agreement, excluding claims for injunctive or other equitable relief as set forth below. Arbitration procedures are typically more limited, more efficient, and less costly than court proceedings and are subject to very limited review by a court. If any litigation should arise between you and the Company in any state or federal court in a suit to vacate or enforce an arbitration award or otherwise, YOU AND THE COMPANY WAIVE ALL RIGHTS TO A JURY TRIAL, instead electing that the dispute be resolved by a judge.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>10.7 Waiver of Class or Consolidated Actions.&nbsp;</strong></p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">ALL CLAIMS AND DISPUTES WITHIN THE SCOPE OF THIS ARBITRATION AGREEMENT MUST BE ARBITRATED OR LITIGATED ON AN INDIVIDUAL BASIS AND NOT ON A CLASS BASIS, AND CLAIMS OF MORE THAN ONE CUSTOMER OR USER CANNOT BE ARBITRATED OR LITIGATED JOINTLY OR CONSOLIDATED WITH THOSE OF ANY OTHER CUSTOMER OR USER.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>10.8 Severability.&nbsp;</strong></p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">If any part or parts of this Arbitration Agreement are found under the law to be invalid or unenforceable by a court of competent jurisdiction, then such specific part or parts shall be of no force and effect and shall be severed, and the remainder of the Arbitration Agreement shall continue in full force and effect.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>10.9 Right to Waive.&nbsp;</strong></p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">Any or all of the rights and limitations set forth in this Arbitration Agreement may be waived by the party against whom the claim is asserted. Such waiver shall not waive or affect any other portion of this Arbitration Agreement.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>10.10 Survival of Agreement.&nbsp;</strong></p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">This Arbitration Agreement will survive the termination of your relationship with the Company.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>10.11 Small Claims Court.&nbsp;</strong></p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">Notwithstanding the foregoing, either you or the Company may bring an individual action in small claims court.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>10.12 Equitable Relief.</strong></p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">Notwithstanding the foregoing, either party may seek equitable relief before a court of competent jurisdiction for the alleged unlawful use of copyrights, trademarks, trade names, logos, trade secrets, or patents or emergency equitable relief before a court of competent jurisdiction to maintain the status quo pending arbitration. A request for interim measures shall not be deemed a waiver of any other rights or obligations under this Arbitration Agreement.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>11.&nbsp;</strong><u><strong>Governing Law, Venue, and Jurisdiction</strong></u></p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">To the extent the parties are permitted under this Agreement to initiate litigation in court, both you and the Company agree that all claims and disputes, including statutory claims and disputes, arising out of or relating to this Agreement, shall be governed in all respects by the substantive law of the state of California, without regard to its conflict of law principles. You and the Company hereby consent to submit to the jurisdiction of the federal and state courts sitting in the state of California for any actions, suits, or proceedings arising out of or relating to this Agreement, that are not subject to the Arbitration Agreement.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">YOU AND THE COMPANY AGREE THAT ANY CAUSE OF ACTION ARISING OUT OF OR RELATED TO THE SERVICES MUST COMMENCE WITHIN ONE (1) YEAR AFTER THE CAUSE OF ACTION ACCRUES. OTHERWISE, SUCH CAUSE OF ACTION IS PERMANENTLY BARRED TO THE EXTENT PERMITTED BY LAW.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;"><strong>12.&nbsp;</strong><u><strong>Entire Agreement</strong></u></p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">This Agreement any other legal notices published by the Company on the Site shall constitute the entire agreement between you and the Company concerning the Services. If any provision of this Agreement is deemed invalid by a court of competent jurisdiction, the invalidity of such provision shall not affect the validity of the remaining provisions of this Agreement, which shall remain in full force and effect. No waiver of any term of this Agreement shall be deemed a further or continuing waiver of such term or any other term, and the Company’s failure to assert any right or provision under this Agreement shall not constitute a waiver of such right or provision.</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">Contact information</p>'+
                      '<p style="text-align: justify;background: transparent;margin-bottom: 0.11in;line-height: 108%;">We welcome your comments or questions about this Agreement. You may contact us at legal@constellationnetwork.io</p>'+
                      '</div>',
                      showCloseButton: true,
                      showCancelButton: true,
                      width: "1200px",
                      focusConfirm: false,
                      confirmButtonText: '<i class="fa fa-thumbs-up"></i> I have read and I Agree to Terms!',
                      confirmButtonAriaLabel: "I have read and I Agree to Terms",
                      cancelButtonText: '<i class="fa fa-thumbs-down"></i> Close Application',
                      cancelButtonAriaLabel: "Close Application"
                  }).then(result => {
              if (result.value) {
                self.$store.state.app.toc = true;
                Swal.fire({
                  title: "Terms of Service Accepted!",
                  text: "Thank you for accepting the Terms of Service. Enjoy your wallet experience.",
                  type: "success"
                });
              } else {
                  Swal.fire({
                  title: "Terms of Service Rejected!",
                  text: "You need to accept the Terms of Service to use this product.",
                  type: "error"
                });
                
              }
              // self.$Progress.finish();
              }) }, 8000);
              
            } else {
              self.loginInProgress = false;
              self.$Progress.fail();
            }
            
          });
        } else {
          self.loginInProgress = false;
          self.$Progress.fail();
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
