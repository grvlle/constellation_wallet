<template>
  <div>

    <!--Stats cards-->



    <div class="row">
      <div class="col-md-6 col-xl-4">
        <stats-card title="asf">
          <div class="icon-big text-center" :class="`icon-success`" slot="header">
            <i class="ti-wallet"></i>
          </div>
          <div class="numbers text-center" slot="content">
            <p>$DAG Tokens</p>
            {{tokenAmount}}
            
            
          </div>
          <div class="stats" slot="footer">
            <i class="ti-reload"></i> Updated 42 seconds ago
          </div>
        </stats-card>
      </div>
    


      <div class="col-md-6 col-xl-4">
        <stats-card title="asf">
          <div class="icon-big text-center" :class="`icon-danger`" slot="header">
            <i class="ti-pulse"></i>
          </div>
          <div class="numbers text-center" slot="content">
            <p>USD value</p>
            {{tokenAmount}}
            
            
          </div>
          <div class="stats" slot="footer">
            <i class="ti-timer"></i> In the last hour
          </div>
        </stats-card>
      </div>



      <div class="col-md-6 col-xl-4">
        <stats-card title="asf">
          <div class="icon-big text-center" :class="`icon-info`" slot="header">
            <i class="ti-package"></i>
          </div>
          <div class="numbers text-center" slot="content">
            <p>Blocks</p>
            {{blocks}}
            
            
          </div>
          <div class="stats" slot="footer">
            <i class="ti-reload"></i> Updated now
          </div>
        </stats-card>
      </div>
      </div>

        <div class="row">
      <div class="col-md-6 col-xl-12" v-for="wallet in walletAddress" :key="wallet.address">
        <wide-card>
          <div class="numbers text-center col-17" slot="content">
            <p>{{wallet.title}}</p>  
            <hr>
            <p style="color: #c4c4c4; padding-top: 15px; background-color: #f7f7f7; font-size: 25px; font-weight: 100; font-family: 'Inconsolata';">
              {{wallet.address}} 
            <p-button type="info" style="margin-bottom: 25px;" icon @click.native="notifyVue('top', 'right')"><i class="fa fa-copy"></i>
            </p-button>
            </p>
          </div>
        </wide-card>
      </div>
    </div>

    <!--Charts-->
    <div class="row">

      <div class="col-md-6 col-12">
        <chart-card title="Nodes Online"
                    sub-title="Since last 24 hours"
                    :chart-data="preferencesChart.data"
                    chart-type="Pie">
          <span slot="footer">
            <i class="ti-timer"></i> Updated 2 days ago</span>
          <div slot="legend">
            <i class="fa fa-circle text-info"></i> Foundation Nodes
            <i class="fa fa-circle text-success"></i> Medium Nodes
            <i class="fa fa-circle text-danger"></i> Light Nodes
          </div>
        </chart-card>
      </div>

      <div class="col-md-6 col-12">
        <chart-card title="Transactions"
                    sub-title="The amount of transactions sent vs. recieved over the last year"
                    :chart-data="activityChart.data"
                    :chart-options="activityChart.options">
          <span slot="footer">
            <i class="ti-check"></i> Data information certified
          </span>
          <div slot="legend">
            <i class="fa fa-circle text-info"></i> TX
            <i class="fa fa-circle text-success"></i> RX
          </div>
        </chart-card>
      </div>

    

      <div class="col-12">
        <chart-card title="Network Throughput (tps)"
                    sub-title="24 Hours performance"
                    :chart-data="usersChart.data"
                    :chart-options="usersChart.options">
          <span slot="footer">
            <i class="ti-reload"></i> Updated 3 minutes ago
          </span>
          <!-- <div slot="legend">
            <i class="fa fa-circle text-info"></i> Open
            <i class="fa fa-circle text-danger"></i> Click
            <i class="fa fa-circle text-warning"></i> Click Second Time
          </div> -->
        </chart-card>
      </div>

    </div>

  </div>
</template>
<script>
import { StatsCard, ChartCard, WideCard } from "@/components/index";
import Chartist from 'chartist';
import NotificationTemplate from './Notifications/NotificationTemplate';

export default {
  components: {
    StatsCard,
    WideCard,
    ChartCard
  },
  methods: {
    notifyVue(verticalAlign, horizontalAlign) {
      // const color = Math.floor(Math.random() * 4 + 1);
      const color = 2;
      this.$notify({
        component: NotificationTemplate,
        icon: "ti-check",
        horizontalAlign: horizontalAlign,
        verticalAlign: verticalAlign,
        type: this.type[color]
      })
    }
    
  },
  mounted() {
    window.wails.events.on("error", (currency, number) => {
      let result = number * 2;
      this.tokenAmount = `${currency} ${number}`;
    });
    window.wails.events.on("blocks", (number) => {
      this.blocks = number;
    });
    this.tokenAmount();
    this.blockAmount();
  },

  /**
   * Chart data used to render stats, charts. Should be replaced with server data
   */

  data() {
    return {
        tokenAmount: "NaN",
        blocks: "NaN",
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
          ],
          series: [
            [287, 385, 490, 562, 594, 626, 698, 895, 952],
            [67, 152, 193, 240, 387, 435, 535, 642, 744]
          ]
        },
        options: {
          low: 0,
          high: 1000,
          showArea: true,
          height: "245px",
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
            "Jan",
            "Feb",
            "Mar",
            "Apr",
            "Mai",
            "Jun",
            "Jul",
            "Aug",
            "Sep",
            "Oct",
            "Nov",
            "Dec"
          ],
          series: [
            [542, 543, 520, 680, 653, 753, 326, 434, 568, 610, 756, 895],
            [230, 293, 380, 480, 503, 553, 600, 664, 698, 710, 736, 795]
          ]
        },
        options: {
          seriesBarDistance: 10,
          axisX: {
            showGrid: false
          },
          height: "245px"
        }
      },
      preferencesChart: {
        data: {
          labels: ["62%", "32%", "6%"],
          series: [62, 32, 6]
        },
        options: {}
      }
    };
  }

};

</script>
<style>
</style>
