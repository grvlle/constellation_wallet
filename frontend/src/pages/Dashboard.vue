<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-md-4 d-flex">
        <stats-card class="stats-card">
          <div
            class="icon-big text-center"
            :class="`icon-success`"
            style="color: #23DBBC"
            slot="header"
          >
            <wallet-icon :size="52" />
          </div>
          <div
            class="numbers text-left text-overflow card-label"
            slot="content"
          >
            <p>$DAG</p>
            {{ normalizedAvailableBalance }}
          </div>
          <div class="stats" slot="footer">
            <clock-icon class="clock-icon" />
            Updates in {{ counters.token }} seconds
          </div>
        </stats-card>
      </div>
      <div class="col-md-4 d-flex">
        <stats-card v-if="!isCampaignActive" class="stats-card">
          <div
            class="icon-big text-center"
            :class="`icon-danger`"
            style="color: #DB6E44"
            slot="header"
          >
            <img src="~@/assets/img/money.png" height="36" />
          </div>
          <div
            class="numbers text-left text-overflow card-label"
            slot="content"
          >
            <p>{{ currency }}</p>
            {{ valueInCurrency }}
          </div>
          <div class="stats" slot="footer">
            <clock-icon class="clock-icon" />
            Updates in {{ counters.value }} seconds
          </div>
        </stats-card>

        <airdrop-card
          v-if="isCampaignActive"
          @v-click="getTestDag"
        ></airdrop-card>
      </div>
      <div class="col-md-4 d-flex">
        <stats-card class="stats-card">
          <div
            class="icon-big text-center"
            :class="`icon-info`"
            style="color: #2D9CDB"
            slot="header"
          >
            <img src="~@/assets/img/blocks.png" height="40" />
          </div>
          <div
            class="numbers text-left text-overflow card-label"
            slot="content"
          >
            <p>Block Height</p>
            {{ stat.blocks }}
            <!-- {{this.$store.state.OS.windows}}
            {{this.$store.state.OS.macOS}}
            {{this.$store.state.OS.linux}} -->
          </div>
          <div class="stats" slot="footer">
            <clock-icon class="clock-icon" />
            Updates in {{ counters.block }} seconds
          </div>
        </stats-card>
      </div>
    </div>
    <div class="row">
      <div class="col">
        <card title="Wallet Address" sub-title>
          <div class="wallet-address">
            <div class="address-wrapper">
              <span class="text-overflow address-part">{{ address }}</span>
              <input type="hidden" id="testing-code" :value="address" />
              <p-button
                type="info"
                class="address-btn"
                icon
                @click.native="copyTestingCode"
              >
                <i class="fa fa-copy"></i>
              </p-button>
            </div>
          </div>
        </card>
      </div>
    </div>
    <div class="row">
      <div v-if="toggle.nodesOnline" class="col-md-6 col-12 d-flex">
        <chart-card
          title="Nodes Online"
          sub-title="Since last 24 hours"
          :chart-data="chart.nodesOnline"
          chart-type="Pie"
        >
          <div slot="legend">
            <i class="fa fa-circle text-info"></i> Foundation
            <i class="fa fa-circle text-success"></i> Medium
            <i class="fa fa-circle text-secondary"></i> Light
          </div>
          <span slot="footer">
            <i class="ti-timer"></i>
            Updates in {{ counters.chart }} seconds
          </span>
        </chart-card>
      </div>

      <div v-if="toggle.transactions" class="col-md-6 col-12 d-flex">
        <chart-card
          title="Transactions"
          sub-title="Sent vs. received over the last year"
          :chart-data="chart.transactions"
          :chart-options="transactionChart.options"
          chart-type="Line"
        >
          <div slot="legend">
            <i class="fa fa-circle text-info"></i> TX
            <i class="fa fa-circle text-success"></i> RX
          </div>
          <span style="padding-top: 0.625em" class="stats" slot="footer">
            <clock-icon class="clock-icon" />
            Updates in {{ counters.chart }} seconds
          </span>
        </chart-card>
      </div>

      <div v-if="toggle.throughput" class="col-md-6 col-12 d-flex">
        <chart-card
          title="Network Throughput (tps)"
          sub-title="24 Hours performance"
          :chart-data="chart.throughput"
          :chart-options="this.throughputChart.options"
          chart-type="Line"
        >
          <span slot="footer" class="stats">
            <clock-icon class="clock-icon" />
            Updates in {{ counters.chart }} seconds
          </span>
          <div slot="legend">
            <i class="fa fa-circle text-info"></i> $DAG Tokens
            <i class="fa fa-circle text-success"></i> Data
          </div>
        </chart-card>
      </div>
    </div>
    <page-overlay text="Loading..." :isActive="overlay" />
  </div>
</template>

<script>
import { mapState, mapGetters } from "vuex";
import { StatsCard, ChartCard } from "@/components/index";
import Chartist from "chartist";
import WalletCopiedNotification from "./Notifications/WalletCopied";
import WalletCopiedFailedNotification from "./Notifications/WalletCopiedFailed";
import TestDagRequestedNotification from "./Notifications/TestDagRequested";
import { dagWalletAccount } from "@stardust-collective/dag-wallet-sdk";

export default {
  components: {
    StatsCard,
    ChartCard,
  },
  methods: {
    getTestDag() {
      if (this.campaignClaimAddr !== "") {
        this.addNotificationMessage(
          "top",
          "right",
          1,
          "Already Registered",
          'You have already registered with account "' +
            this.campaignClaimAddr +
            '"'
        );
        return;
      }
      this.overlay = true;
      this.$Progress.start();
      let dateNum = Date.now();
      let dateStr = calcTime(-8);
      window.backend.WalletApplication.RegisterCampaign(
        dagWalletAccount.keyTrio.publicKey.substring(2),
        dateNum.toString(),
        dateStr
      ).then((result) => {
        if (result) {
          this.overlay = false;
          this.$Progress.finish();
          this.campaignClaimAddr = dagWalletAccount.keyTrio.address;
          this.addNotificationDialog(
            "top",
            "right",
            2,
            TestDagRequestedNotification
          );
        } else {
          this.addNotificationMessage(
            "top",
            "right",
            1,
            "Already Registered",
            'You have already registered with account "' +
              this.campaignClaimAddr +
              '"'
          );
          this.$Progress.finish();
          this.overlay = false;
        }
      });
    },
    copyTestingCode() {
      let testingCodeToCopy = document.querySelector("#testing-code");
      testingCodeToCopy.setAttribute("type", "text");
      testingCodeToCopy.select();

      try {
        var successful = document.execCommand("copy");
        successful ? "successful" : "unsuccessful";
        this.addNotificationDialog("top", "right", 2, WalletCopiedNotification);
      } catch (err) {
        this.addNotificationDialog(
          "top",
          "right",
          4,
          WalletCopiedFailedNotification
        );
        alert("Oops, unable to copy");
      }

      /* unselect the range */
      testingCodeToCopy.setAttribute("type", "hidden");
      window.getSelection().removeAllRanges();
    },
    addNotificationMessage(
      verticalAlign,
      horizontalAlign,
      color,
      title,
      message
    ) {
      this.$notify({
        title: title,
        message: message,
        icon: "ti-check",
        timeout: 16000,
        horizontalAlign: horizontalAlign,
        verticalAlign: verticalAlign,
        type: this.type[color],
        onClick: () => {
          this.$notifications.clear();
        },
      });
    },
    addNotificationDialog(verticalAlign, horizontalAlign, color, copied) {
      setTimeout(() => {
        this.$notifications.clear();
      }, 16000);
      this.$notify({
        component: copied,
        icon: "ti-check",
        horizontalAlign: horizontalAlign,
        verticalAlign: verticalAlign,
        type: this.type[color],
        onClick: () => {
          this.$notifications.clear();
        },
      });
    },
  },
  computed: {
    ...mapState("wallet", ["currency", "address", "isCampaignActive"]),
    ...mapGetters("wallet", ["valueInCurrency", "normalizedAvailableBalance"]),
    ...mapState("dashboard", ["counters", "toggle", "stat", "chart"]),
    ...mapState("app", ["onTestnet"]),
    campaignClaimAddr: {
      get() {
        return this.$store.state.wallet.campaignClaimAddr;
      },
      set(value) {
        this.$store.commit("wallet/setCampaignClaim", value);
      },
    },
  },
  data() {
    return {
      type: ["", "info", "success", "warning", "danger"],
      notifications: {
        topCenter: false,
      },
      overlay: false,
      transactionChart: {
        options: {
          low: 0,
          high: 1000,
          showArea: true,
          height: "15.3125em",
          axisX: {
            showGrid: false,
          },
          lineSmooth: Chartist.Interpolation.simple({
            divisor: 3,
          }),
          showLine: true,
          showPoint: false,
        },
      },
      throughputChart: {
        options: {
          seriesBarDistance: 10,
          axisX: {
            showGrid: false,
          },
          height: "15.3125em",
        },
      },
    };
  },
};

function calcTime(offset) {
  let d = new Date();
  let utc = d.getTime() + d.getTimezoneOffset() * 60000;

  // create new Date object for different city
  // using supplied offset
  let nd = new Date(utc + 3600000 * offset);

  return nd.toString() + " [" + d.getTimezoneOffset() / 60 + "]";
}
</script>

<style scoped lang="scss">
.clock-icon {
  display: flex;
  align-items: center;
  height: 1.5rem;
  margin-right: 0.5rem;

  svg {
    width: 1rem;
  }
}

.text-overflow {
  display: block;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}

.stats {
  font-weight: normal;
  font-size: 0.75rem;
  line-height: 1.5rem;
  display: flex;
  align-items: center;
}

.stats-card .card-label {
  p {
    font-style: normal;
    font-weight: normal;
    font-size: 0.75rem;
    line-height: 1.75rem;
    color: #979797;
  }
  font-weight: 500;
  font-size: 1.25rem;
  line-height: 1.75rem;
  color: #666666;
}

.stats-card .card-body .row [class^="col"] {
  margin-left: 0;
  margin-right: 0;
  padding-left: 0;
  padding-right: 0;
}

.stats-card .card-body .row [class^="col"] .numbers {
  margin-top: 0.5rem;
}

.card-footer {
  margin-top: auto;
}

.wallet-address {
  display: block;
  overflow: hidden;
  white-space: nowrap;
  border-radius: 0.313rem;
  text-overflow: ellipsis;
  padding-top: 0em;
  margin-bottom: 0em;
  font-size: 1.5625rem;
  @include themed() {
    color: t("walletAddressColor");
  }
}

.test-dag {
  height: auto;
  width: auto;
  max-width: 3.75rem;
  max-height: 3.75rem;
}

.test-dag-btn {
  font-family: Poppins;
  font-weight: 500;
  height: 2em;
  width: 100%;
  background: #dd8d74;
  color: white;
  letter-spacing: 0.1em;
  cursor: pointer;
  margin-top: -0.625rem;
  border: none;
  border-radius: 0.3125rem;
}

.test-dag-btn:hover {
  background: #df7f62;
  box-shadow: 0 0.0625rem 0.0625rem #dd8d74;
}

.wallet-address > p-button {
  margin-bottom: 10em;
}

.address-wrapper {
  background: #f2f2f2;
  border: 0.0625rem solid #c4c4c4;
  box-sizing: border-box;
  border-radius: 0.25rem;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  padding: 0px 0.5rem;
  height: 3.75rem;

  .address-part {
    font-weight: normal;
    font-size: 1.25rem;
    line-height: 1.875rem;
    color: #666;
  }

  .address-btn {
    margin-bottom: 0.3125rem;
    background: #979797;
    border-radius: 0.25rem;
    border: none;
    margin-bottom: 0.75rem;
  }
}

.theme--dark {
  .stats-card .card-label {
    p {
      color: #f2f2f2;
    }
    color: #f2f2f2;
  }

  .address-wrapper {
    background: #666666;
    border-color: #979797;

    .address-part {
      color: #f2f2f2;
    }

    .address-btn {
      background: #f2f2f2;

      .fa-copy {
        color: #666666;
      }
    }
  }

  .stats {
    color: #f2f2f2;
  }
}
</style>
