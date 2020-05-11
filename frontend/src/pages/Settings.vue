<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-12">
        <card
          title="Settings"
          style="padding-bottom: 1em;"
          sub-title="Configure the Molly Wallet to your personal preferences"
        ></card>
      </div>
    </div>
    <div class="row">
      <div class="col-md-6 d-flex">
        <card
          title="Dashboard Widgets"
          sub-title="Customize your dashboard by activating / deactivating widgets"
        >
          <div class="container">
            <table class="table-noheader">
              <tr>
                <td>
                  <p>Nodes Online (dummy data)</p>
                </td>
                <td align="right">
                  <toggle-button
                    @change="toggleNodesOnline"
                    :value="this.$store.state.toggleDashboard.showNodesOnline"
                    color="#5fd1fa"
                    :sync="true"
                    :labels="true"
                  />
                </td>
              </tr>
              <tr>
                <td>
                  <p>Transactions (dummy data)</p>
                </td>
                <td align="right">
                  <toggle-button
                    @change="toggleTransactions"
                    :value="this.$store.state.toggleDashboard.showTransactions"
                    color="#5fd1fa"
                    :sync="true"
                    :labels="true"
                  />
                </td>
              </tr>
              <tr>
                <td>
                  <p>Throughput (dummy data)</p>
                </td>
                <td align="right">
                  <toggle-button
                    @change="toggleThroughput"
                    :value="this.$store.state.toggleDashboard.showThroughput"
                    color="#5fd1fa"
                    :sync="true"
                    :labels="true"
                  />
                </td>
              </tr>
            </table>
          </div>
        </card>
      </div>
      <div class="col-md-6 d-flex">
        <card
          title="Wallet Information"
          sub-title="Customize your Wallet with your own label and profile picture"
        >
          <div class="container">
            <form>
              <div class="input-group">
                <input
                  type="text"
                  class="form-control"
                  placeholder="Enter a new Wallet Label..."
                  v-model="newLabel"
                />
                <span class="input-group-append">
                  <p-button @click.native="submitLabel()" type="success" style="width:6rem;">Apply</p-button>
                </span>
              </div>
              <div class="input-group">
                <input
                  type="text"
                  :disabled="true"
                  class="form-control"
                  :placeholder="this.$store.state.walletInfo.imgPath"
                  v-model="this.$store.state.walletInfo.imgPath"
                />
                <span class="input-group-append">
                  <p-button @click.native="uploadImage()" type="default" style="width:6rem;">Browse</p-button>
                </span>
              </div>
            </form>
          </div>
        </card>
      </div>
      <div class="col-md-6 d-flex">
        <card title="Backup / Restore" sub-title="Backup your Wallet">
          <div class="container">
            <form @submit.prevent>
              <!-- Below can be revived when Mnemonic Seed is supported -->
              <!-- 
              <label class="control-label">
                <p>Mnemonic Seed</p>
              </label>
              <p class="small" style="margin-bottom: -5px; margin-top: -20px;">
                This is the seed that will be used to backup and restore your wallet.
              </p>
              <div class="row">
                <div class="input-group col-md-12">
                  <input :type=seed class="form-control" :disabled="true" placeholder="Mnemonic Seed" v-model="this.$store.state.walletInfo.seed">
                  <div class="input-group-append">
                    <p-button class="btn" @click.native="showPassword2()" type="danger"><i v-bind:class="btnText"/></p-button>
                  </div>
                </div>
              </div>-->
              <div class="row">
                <div class="col-12">
                  <label class="control-label" style="margin-bottom: 0;">
                    <p style="margin-bottom: 0;">Path to private key (key.p12)</p>
                  </label>
                  <fg-input
                    type="text"
                    :disabled="true"
                    placeholder="Path to private key (key.p12)"
                    v-model="this.$store.state.walletInfo.keystorePath"
                  ></fg-input>
                </div>
              </div>
              <div class="row">
                <div class="col-12">
                  <label class="control-label" style="margin-bottom: 0;">
                    <p style="margin-bottom: 0;">Mnemonic Seed (coming soon)</p>
                  </label>
                  <p
                    class="small"
                    style="margin-bottom: 2px;"
                  >This will be used to restore your wallet</p>
                  <div class="input-group">
                    <input
                      :type="type"
                      :disabled="true"
                      class="form-control"
                      label="Private Key"
                      placeholder="Mnemonic Seed (coming soon)"
                      v-model="this.$store.state.walletInfo.seed"
                      aria-describedby="basic-addon2"
                    />
                    <div class="input-group-append">
                      <p-button class="btn" @click.native="showPassword()" type="danger">
                        <i v-bind:class="btnText" />
                      </p-button>
                    </div>
                  </div>
                </div>
              </div>
            </form>
          </div>
        </card>
      </div>
      <!-- <div class="col-md-6 d-flex">
        <card
          title="Import / Export Keys"
          sub-title="Restore your Wallet from a previously exported file or create a new export"
        >
          <div class="container" style="margin-top: auto;">
            <div class="row">
              <div class="col">
                <p>This feature is currently disabled since the team wallet does not support it for now.</p>
                <p class="small">
                  Select Import if you wish to restore your wallet from a previously exported file.
                  Select Export to export your keys into an encrypted .pem file on your filesystem.
                  Store this file in cold storage for optimal security.
                </p>
              </div>
            </div>
            <div class="row">
              <div class="col-md-6 pr-md-2 mb-3">
                <p-button @click.native="importKeys()" type="info" block :disabled="true">
                  <span style="display: block;">
                    <i class="fa fa-file-import"></i> IMPORT
                  </span>
                </p-button>
              </div>
              <div class="col-md-6 pl-md-2 mb-3">
                <p-button @click.native="exportKeys()" type="info" block :disabled="true">
                  <span style="display: block;">
                    <i class="fa fa-file-export"></i> EXPORT
                  </span>
                </p-button>
              </div>
            </div>
          </div>
        </card>
      </div> -->
      <div class="col-md-6 d-flex">
        <card
          title="Display settings"
          sub-title="Customize your Molly Wallet display settings"
        >
          <div class="container">
            <div class="row settings">
              <div class="col-6">
                <p>Dark Mode [BETA]</p>
              </div>
              <div class="col-6" align="right">
                <toggle-button
                  @change="toggleDarkMode"
                  :value="this.$store.state.walletInfo.darkMode"
                  color="#5fd1fa"
                  :sync="true"
                  :labels="true"
                />
              </div>
            </div>
            <div class="row settings">
              <div class="col-6">
                <p>Currency</p>
              </div>
              <div class="col-6" align="right">
                <vue-select 
                  @input="setCurrency"
                  :value="this.$store.state.walletInfo.currency"
                  :options="['BTC', 'EUR', 'USD']">
                </vue-select>
              </div>
            </div>
          </div>
        </card>
      </div>
    </div>
  </div>
</template>

<script>
import ImgUploaded from "./Notifications/ImageUploaded";
import Swal from "sweetalert2";
import VueSelect from 'vue-select';

export default {
  components: {
    VueSelect
  },
  methods: {
    submitLabel: function() {
      const swalPopup = Swal.mixin({
        customClass: {
          container: this.$store.state.walletInfo.darkMode
            ? "theme--dark"
            : "theme--light"
        }
      });
      if (this.newLabel === "") {
        swalPopup.fire({
          title: "Failed!",
          text: "Unable to change wallet label. No new label entered.",
          type: "error"
        });
      } else {
        swalPopup
          .fire({
            title: "Are you sure?",
            html:
              "You are about change wallet label to " +
              this.newLabel +
              ". This will replace your wallet label.",
            type: "warning",
            showCancelButton: true,
            confirmButtonColor: "#6DECBB",
            confirmButtonText: "Yes, change label!"
          })
          .then(result => {
            if (result.value) {
              this.$store.state.walletInfo.email = this.newLabel;
              window.backend.WalletApplication.StoreWalletLabelInDB(
                this.newLabel
              );
              swalPopup.fire({
                title: "Success!",
                text: "You have set a new wallet label",
                type: "success"
              });
            }
          });
      }
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
    toggleDarkMode: function() {
      this.$store.state.walletInfo.darkMode = !this.$store.state.walletInfo.darkMode;
      window.backend.WalletApplication.StoreDarkModeStateDB(this.$store.state.walletInfo.darkMode);
    },
    setCurrency: function(value) {
      this.$store.state.walletInfo.currency = value;
      window.backend.WalletApplication.StoreCurrencyStateDB(this.$store.state.walletInfo.currency);
    },
    importKeys: function() {
      window.backend.WalletApplication.ImportKeys();
    },
    exportKeys: function() {
      window.backend.WalletApplication.ExportKeys();
    },
    uploadImage: function() {
      window.backend.WalletApplication.UploadImage().then(path => {
        const swalPopup = Swal.mixin({
          customClass: {
            container: this.$store.state.walletInfo.darkMode
              ? "theme--dark"
              : "theme--light"
          }
        });
        if (path === "None") {
          swalPopup.fire({
            title: "Failed!",
            text:
              "Unable to change wallet image. Make sure that the image resolution is not larger than 200x200",
            type: "error"
          });
        } else {
          swalPopup
            .fire({
              title: "Are you sure?",
              html:
                "You are about to upload " +
                path +
                ". This will replace your existing image.",
              type: "warning",
              showCancelButton: true,
              confirmButtonColor: "#6DECBB",
              confirmButtonText: "Yes, upload image!"
            })
            .then(result => {
              if (result.value) {
                this.$store.state.walletInfo.imgPath = path;
                swalPopup.fire({
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

<style lang="scss">
.vs__clear {
  display: none;
}
</style>
