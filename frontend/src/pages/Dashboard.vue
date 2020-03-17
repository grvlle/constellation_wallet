<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-md-4 d-flex">
        <stats-card class="stats-card">
          <div class="icon-big text-center" :class="`icon-success`" slot="header">
            <i class="fas fa-wallet"></i>
          </div>
          <div class="numbers text-center text-overflow" slot="content">
            <p>DAG</p>
            {{this.$store.state.walletInfo.tokenAmount | asDAGs}}
          </div>
          <div class="stats" slot="footer">
            <i class="ti-timer"></i>
            Updates in {{this.$store.state.counters.tokenCounter}} seconds
          </div>
        </stats-card>
      </div>
      <div class="col-md-4 d-flex">
        <stats-card class="stats-card">
          <div class="icon-big text-center" :class="`icon-danger`" slot="header">
            <i class="fas fa-search-dollar"></i>
          </div>
          <div class="numbers text-center text-overflow" slot="content">
            <p>{{this.$store.state.walletInfo.currency}}</p>
            {{this.$store.state.walletInfo.totalValue | asCurrency(this.$store.state.walletInfo.currency)}}
          </div>
          <div class="stats" slot="footer">
            <i class="ti-timer"></i>
            Updates in {{this.$store.state.counters.valueCounter}} seconds
          </div>
        </stats-card>
      </div>
      <div class="col-md-4 d-flex">
        <stats-card class="stats-card">
          <div class="icon-big text-center" :class="`icon-info`" slot="header">
            <i class="fas fa-cube"></i>
          </div>
          <div class="numbers text-center text-overflow" slot="content">
            <p>Blocks</p>
            {{wallet2.blocks}}
          </div>
          <div class="stats" slot="footer">
            <i class="ti-reload"></i>
            Updates in {{this.$store.state.counters.blockCounter}} seconds
          </div>
        </stats-card>
      </div>
    </div>
    <div class="row">
      <div class="col">
        <card title="Wallet Address" sub-title="">
          <div class="wallet-address">
            <table class="table-noheader">
              <tr>
                <td style="padding-top: 15px; padding-left: 15px; width: 90%;">
                  <span class="text-overflow">{{wallet2.address}}</span>
                  <input type="hidden" id="testing-code" :value="wallet2.address" />
                </td>
                <td style="padding-top: 10px;">
                  <p-button type="info" style="margin-bottom: 12px;" icon @click.native="copyTestingCode">
                    <i class="fa fa-copy"></i>
                  </p-button>
                </td>
              </tr>
            </table>
          </div>
        </card>
      </div>
    </div>
    <div class="row">
      <div v-if="this.$store.state.toggleDashboard.showNodesOnline" class="col-md-6 col-12 d-flex">
        <chart-card
          title="Nodes Online"
          sub-title="Since last 24 hours"
          :chart-data="this.$store.state.chartData.nodesOnline"
          chart-type="Pie"
        >
          <div slot="legend">
            <i class="fa fa-circle text-info"></i> Foundation
            <i class="fa fa-circle text-success"></i> Medium
            <i class="fa fa-circle text-danger"></i> Light
          </div>
          <span slot="footer">
            <i class="ti-timer"></i>
            Updates in {{this.$store.state.counters.nodesOnlineCounter}} seconds
          </span>
        </chart-card>
      </div>

      <div v-if="this.$store.state.toggleDashboard.showTransactions" class="col-md-6 col-12 d-flex">
        <chart-card
          title="Transactions"
          sub-title="Sent vs. received over the last year"
          :chart-data="this.$store.state.chartData.transactions"
          :chart-options="activityChart.options"
        >
          <div slot="legend">
            <i class="fa fa-circle text-info"></i> TX
            <i class="fa fa-circle text-success"></i> RX
          </div>
          <span style="padding-top: 0.625em;" slot="footer">
            <i class="ti-timer"></i>
            Updates in {{this.$store.state.counters.nodesOnlineCounter}} seconds
          </span>
        </chart-card>
      </div>

      <div v-if="this.$store.state.toggleDashboard.showThroughput" class="col-md-6 col-12 d-flex">
        <chart-card
          title="Network Throughput (tps)"
          sub-title="24 Hours performance"
          :chart-data="this.$store.state.chartData.throughput"
          :chart-options="usersChart.options"
        >
          <span slot="footer">
            <i class="ti-timer"></i>
            Updates in {{this.$store.state.counters.nodesOnlineCounter}} seconds
          </span>
          <div slot="legend">
            <i class="fa fa-circle text-info"></i> $DAG Tokens
            <i class="fa fa-circle text-success"></i> Data
          </div>
          <!-- <div slot="legend">
            <i class="fa fa-circle text-info"></i> Open
            <i class="fa fa-circle text-danger"></i> Click
            <i class="fa fa-circle text-warning"></i> Click Second Time
          </div>-->
        </chart-card>
      </div>
    </div>
  </div>
</template>

<script>
import { StatsCard, ChartCard } from "@/components/index";
import Chartist from "chartist";
import WalletCopiedNotification from "./Notifications/WalletCopied";
import WalletCopiedFailedNotification from "./Notifications/WalletCopiedFailed";

export default {
  components: {
    StatsCard,
    ChartCard
  },

  methods: {
    copyTestingCode() {
      let testingCodeToCopy = document.querySelector("#testing-code");
      testingCodeToCopy.setAttribute("type", "text");
      testingCodeToCopy.select();

      try {
        var successful = document.execCommand("copy");
        successful ? "successful" : "unsuccessful";
        this.addNotification("top", "right", 2, WalletCopiedNotification);
      } catch (err) {
        this.addNotification("top", "right", 4, WalletCopiedFailedNotification);
        alert("Oops, unable to copy");
      }

      /* unselect the range */
      testingCodeToCopy.setAttribute("type", "hidden");
      window.getSelection().removeAllRanges();
    },
    getTokens: function() {
      var self = this;
      window.backend.retrieveTokenAmount().then(result => {
        self.tokenAmount = result;
      });
    },
    addNotification(verticalAlign, horizontalAlign, color, copied) {
      setTimeout(() => {
        this.$notifications.clear();
      }, 6000);
      this.$notify({
        component: copied,
        icon: "ti-check",
        horizontalAlign: horizontalAlign,
        verticalAlign: verticalAlign,
        type: this.type[color],
        onClick: () => {
          this.$notifications.clear();
        }
      });
    }
  },
  computed: {
    wallet2() {
      return this.$store.state.walletInfo;
    },
    chartData() {
      return this.$store.state.chartData;
    }
  },
  filters: {
    asCurrency: function (value, currency) {
      var formatter = new Intl.NumberFormat(navigator.language, {
        style: 'currency',
        currency: currency,
        minimumFractionDigits: 0
      });
      return formatter.format(value);
    },
    asDAGs: function (value) {
      var formatter = new Intl.NumberFormat(navigator.language);
      return formatter.format(value);
    }
  },

  /**
   * Chart data used to render stats, charts. Should be replaced with server data
   */

  data() {
    return {
      type: ["", "info", "success", "warning", "danger"],
      notifications: {
        topCenter: false
      },
      usersChart: {
        data: {
          labels: [
            "9:00AM",
            "12:00AM",
            "3:00PM",
            "6:00PM",
            "9:00PM",
            "12:00PM",
            "3:00AM",
            "6:00AM"
          ], //this.$store.state.chartData.throughput.labels,
          series: [
            this.$store.state.chartData.throughput.seriesOne,
            this.$store.state.chartData.throughput.seriesTwo
          ]
        },
        options: {
          low: 0,
          high: 1000,
          showArea: true,
          height: "15.3125em",
          axisX: {
            showGrid: false
          },
          lineSmooth: Chartist.Interpolation.simple({
            divisor: 3
          }),
          showLine: true,
          showPoint: false
        }
      },
      activityChart: {
        data: {
          labels: [
            "Jan  ",
            "Feb  ",
            "Mar  ",
            "Apr  ",
            "Mai  ",
            "Jun  ",
            "Jul  ",
            "Aug  ",
            "Sep  ",
            "Oct  ",
            "Nov  ",
            "Dec  "
          ], //this.$store.state.chartData.transactions.labels,
          series: [
            this.$store.state.chartData.transactions.seriesOne,
            this.$store.state.chartData.transactions.seriesTwo
          ]
        },
        options: {
          seriesBarDistance: 10,
          axisX: {
            showGrid: false
          },
          height: "15.3125em"
        }
      },
      preferencesChart: {
        data: {
          labels: this.$store.state.chartData.nodesOnline.labels,
          series: this.$store.state.chartData.nodesOnline.series
        },
        options: {}
      }
    };
  }
};
</script>

<style>

.text-overflow {
  display: block;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}

.stats-card .card {
  width: 100%;
  display: flex;
  flex-direction: column;
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
  padding-top: 0em;
  padding-bottom: 0.3em;
  font-size: 1.5625rem;
  font-weight: 100;
}

.wallet-address > p-button {
  margin-bottom: 10em;
}
</style>
