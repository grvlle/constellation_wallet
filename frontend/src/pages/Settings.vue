<template>
  <div class="row">
    <div class="col-md-6">
      <card>
        <h4 class="card-title">Settings</h4>
        <p class="card-category">Configure the Constellation Wallet to your personal preferences</p>

        <br />
        <p>This section let's you personalize and tailor your wallet to your needs. The right hand side lets you manage your wallet keys.</p>
        <br />

        <label class="control-label">
          <p>Toggle Dashboard Widgets</p>
        </label>
        <p
          class="small"
          style="margin-bottom: -5px; margin-top: -20px;"
        >This option let's you customize your dashboard by toggling/untoggling widgets from the Dashboard.</p>

        <div style="padding: 5px;">
          <br />
          <table style="width:100%; ">
            <tr style="background-color: #f9f9f9; border-bottom: 3px solid white;">
              <td style="padding-top: 15px; padding-left: 15px;">
                <p>
                  Toggle
                  <b>Nodes Online</b> widget
                </p>
              </td>
              <td style="padding-top: 10px;">
                <toggle-button
                  @change="toggleNodesOnline"
                  :value="this.$store.state.toggleDashboard.showNodesOnline"
                  color="#82C7EB"
                  :sync="true"
                  :labels="true"
                />
              </td>
            </tr>
            <tr style="background-color: #f9f9f9; border-bottom: 3px solid white;">
              <td style="padding-top: 15px; padding-left: 15px;">
                <p>
                  Toggle
                  <b>Transactions</b> widget
                </p>
              </td>
              <td>
                <toggle-button
                  @change="toggleTransactions"
                  :value="this.$store.state.toggleDashboard.showTransactions"
                  color="#82C7EB"
                  :sync="true"
                  :labels="true"
                />
              </td>
            </tr>
            <tr style="background-color: #f9f9f9; border-bottom: 3px solid white;">
              <td style="padding-top: 15px; padding-left: 15px;">
                <p>
                  Toggle
                  <b>Throughput</b> widget
                </p>
              </td>
              <td>
                <toggle-button
                  @change="toggleThroughput"
                  :value="this.$store.state.toggleDashboard.showThroughput"
                  color="#82C7EB"
                  :sync="true"
                  :labels="true"
                />
              </td>
            </tr>
          </table>
        </div>
        <br />
        <label class="control-label">
          <p>Edit Wallet Information</p>
        </label>
        <p
          class="small"
          style="margin-bottom: -5px; margin-top: -20px;"
        >This option let's you customize your Wallet Card under Wallet Information with your own label and profile picture.</p>
        <br />
        <form>
          <table style="width:100%;">
            <tr>
              <td style="padding: 0px; width: 81%;">
                <fg-input type="text" placeholder="Enter a new Wallet Label" v-model="newLabel"></fg-input>
              </td>

              <td style="padding: 0px;">
                <p-button
                  style="margin-top: -17px; width: 95%; float: right;"
                  @click.native="submitLabel()"
                  type="info"
                >Apply</p-button>
              </td>
            </tr>
          </table>
        </form>
        <table style="width:100%;">
          <tr>
            <td style="padding: 0px; width: 81%;">
              <fg-input
                type="text"
                :disabled="true"
                :placeholder="this.$store.state.walletInfo.imgPath"
                v-model="this.$store.state.walletInfo.imgPath"
              ></fg-input>
            </td>

            <td style="padding: 0px;">
              <p-button
                style="margin-top: -17px; width: 95%; float: right;"
                @click.native="uploadImage()"
                type="info"
              >Browse</p-button>
            </td>
          </tr>
        </table>
      </card>
    </div>
    <div class="col-md-6">
      <card>
        <h4 class="card-title">Backup / Restore</h4>
        <p class="card-category">Backup your Wallet</p>
        <br />

        <form @submit.prevent>
          <!-- Below can be revived when Mnemonic Seed is supported -->
          <!-- <label class="control-label">
                                                                                              <p>Mnemonic Seed</p>
                                                                                            </label>
                                                                                            <p class="small" style="margin-bottom: -5px; margin-top: -20px;">This is the seed that will be used to backup and restore your wallet.</p>
                                                                                          <div class="row">
                                                                                            
                                                                                            <div class="input-group col-md-12">
                                                                                              
                                                                                              <input :type=seed class="form-control" :disabled="true" placeholder="Mnemonic Seed" v-model="this.$store.state.walletInfo.seed">
                                                                                              <div class="input-group-append">
                                                                                                <p-button class="btn" @click.native="showPassword2()" type="danger"><i v-bind:class="btnText"/></p-button>
                                                                                              </div>
                                                                                              
                                                                                          </div>
          </div>-->

          <label class="control-label">
            <p>Public Key</p>
          </label>
          <p class="small" style="margin-bottom: -5px; margin-top: -20px;">Base64 Encoded public key</p>
          <div class="row">
            <div class="col-md-12">
              <fg-input
                type="text"
                :disabled="true"
                placeholder="Public Key"
                v-model="this.$store.state.walletInfo.publicKey"
              ></fg-input>
            </div>
          </div>
          <label class="control-label">
            <p>Private Key</p>
          </label>
          <p
            class="small"
            style="margin-bottom: -5px; margin-top: -20px;"
          >Base64 Encoded private key</p>
          <div class="row">
            <div class="input-group col-md-12">
              <input
                :type="type"
                label="Private Key"
                class="form-control"
                :disabled="true"
                placeholder="Private Key"
                v-model="this.$store.state.walletInfo.privateKey"
                aria-describedby="basic-addon2"
              />
              <div class="input-group-append">
                <p-button class="btn" @click.native="showPassword()" type="danger">
                  <i v-bind:class="btnText" />
                </p-button>
              </div>
            </div>
          </div>
        </form>
        <label class="control-label">
          <p>Import / Export Keys</p>
        </label>
        <p class="small" style="margin-top: -20px;">
          Select Import if you wish to restore your wallet from a previously exported file.
          <br />Select Export to export your keys into an encrypted .pem file on your filesystem. Store this file in cold storage for optimal security.
        </p>
        <p-button @click.native="importKeys()" style="margin-right: 10px;" type="success">
          Import
          <i class="fas fa-file-import" />
        </p-button>

        <p-button @click.native="exportKeys()" type="danger">
          Export
          <i class="fas fa-file-export" />
        </p-button>
      </card>
    </div>
  </div>
</template>

<script>
import ImgUploaded from "./Notifications/ImageUploaded";
import Swal from "sweetalert2";
export default {
  methods: {
    submitLabel: function() {
      this.$store.state.walletInfo.email = this.newLabel;
    },
    toggleNodesOnline: function() {
      this.$store.state.toggleDashboard.showNodesOnline = !this.$store.state
        .toggleDashboard.showNodesOnline;
    },
    toggleTransactions: function() {
      this.$store.state.toggleDashboard.showTransactions = !this.$store.state
        .toggleDashboard.showTransactions;
    },
    toggleThroughput: function() {
      this.$store.state.toggleDashboard.showThroughput = !this.$store.state
        .toggleDashboard.showThroughput;
    },
    importKeys: function() {
      window.backend.WalletApplication.ImportKeys();
    },
    exportKeys: function() {
      window.backend.WalletApplication.ExportKeys();
    },
    uploadImage: function() {
      window.backend.WalletApplication.UploadImage().then(path => {
        if (path === "None") {
          Swal.fire({
            title: "Failed!",
            text:
              "Unable to change wallet image. Make sure that the image resolution is not larger than 200x200",
            type: "error"
          });
        } else {
          Swal.fire({
            title: "Are you sure?",
            html:
              "You are about to upload " +
              path +
              ". This will replace your existing image.",
            type: "warning",
            showCancelButton: true,
            confirmButtonColor: "#6DECBB",
            confirmButtonText: "Yes, upload image!"
          }).then(result => {
            if (result.value) {
              this.$store.state.walletInfo.imgPath = path;
              Swal.fire({
                title: "Success!",
                text: "You have uploaded a new wallet picture",
                type: "success"
              }),
                setTimeout(() => {
                  this.$notifications.clear();
                }, 6000);
              this.$notify({
                component: ImgUploaded,
                icon: "ti-check",
                horizontalAlign: "right",
                verticalAlign: "top",
                type: "success",
                onClick: () => {
                  this.$notifications.clear();
                }
              });
            }
          });
        }
      });
    },

    showPassword: function() {
      if (this.type === "password") {
        this.type = "text";
        this.btnText = "fa fa-eye-slash";
      } else {
        this.type = "password";
        this.btnText = "fa fa-eye";
      }
    },
    showPassword2: function() {
      if (this.seed === "password") {
        this.seed = "text";
        this.btnText = "fa fa-eye-slash";
      } else {
        this.seed = "password";
        this.btnText = "fa fa-eye";
      }
    }
  },

  data() {
    return {
      newLabel: "",
      type: "password",
      seed: "password",
      btnText: "fa fa-eye"
    };
  }
};
</script>

<style>
</style>
