<template>
    <div>
    
        <!--Stats cards-->
    
    
    
        <div class="row">
            <div class="col-md-6 col-xl-4">
                <stats-card>
                    <div class="icon-big text-center" :class="`icon-success`" slot="header">
                        <i class="ti-wallet"></i>
                    </div>
                    <div class="numbers text-center" slot="content">
                        <p>$DAG Tokens</p>
                        {{wallet2.tokenAmount}}
    
                    </div>
                    <div class="stats" slot="footer">
                        <i class="ti-reload"></i> Updated {{count}} seconds ago
                    </div>
                </stats-card>
            </div>
    
    
    
            <div class="col-md-6 col-xl-4">
                <stats-card>
                    <div class="icon-big text-center" :class="`icon-danger`" slot="header">
                        <i class="ti-pulse"></i>
                    </div>
                    <div class="numbers text-center" slot="content">
                        <p>USD value</p>
                        {{wallet2.usdValue}}
    
    
                    </div>
                    <div class="stats" slot="footer">
                        <i class="ti-timer"></i> Updated {{count}} seconds ago
                    </div>
                </stats-card>
            </div>
    
    
    
            <div class="col-md-6 col-xl-4">
                <stats-card>
                    <div class="icon-big text-center" :class="`icon-info`" slot="header">
                        <i class="ti-package"></i>
                    </div>
                    <div class="numbers text-center" slot="content">
                        <p>Blocks</p>
                        {{wallet2.blocks}}
    
    
                    </div>
                    <div class="stats" slot="footer">
                        <i class="ti-reload"></i> Updated {{blockCount}} seconds ago
                    </div>
                </stats-card>
            </div>
        </div>
    
        <div class="row">
            <div class="col-md-6 col-xl-12">
                <wide-card>
                    <div class="numbers text-center col-17" slot="content">
                        <p>{{walletCard.title}}</p>
                        <hr>
                        <p style="color: #c4c4c4; padding-top: 15px; background-color: #f7f7f7; font-size: 25px; font-weight: 100; font-family: 'Inconsolata';">
                            {{wallet2.address}}
                            <p-button type="info" style="margin-bottom: 15px" icon @click.native="notifyVue('top', 'right')"><i class="fa fa-copy"></i>
                            </p-button>
                        </p>
                    </div>
                </wide-card>
            </div>
        </div>
    
        <!--Charts-->
        <div class="row">
    
            <div class="col-md-6 col-12">
                <chart-card title="Nodes Online" sub-title="Since last 24 hours" :chart-data="preferencesChart.data" chart-type="Pie">
                    <span slot="footer">
                            <i class="ti-timer"></i> Updated {{pieChartCount}} hours ago</span>
                    <div slot="legend">
                        <i class="fa fa-circle text-info"></i> Foundation Nodes
                        <i class="fa fa-circle text-success"></i> Medium Nodes
                        <i class="fa fa-circle text-danger"></i> Light Nodes
                    </div>
                </chart-card>
            </div>
    
            <div class="col-md-6 col-12">
                <chart-card title="Transactions" sub-title="The amount of transactions sent vs. received over the last year" :chart-data="activityChart.data" :chart-options="activityChart.options">
                    <span slot="footer">
                            <i class="ti-check"></i> Data information certified
                          </span>
                    <div slot="legend">Days
                        <i class="fa fa-circle text-info"></i> TX
                        <i class="fa fa-circle text-success"></i> RX
                    </div>
                </chart-card>
            </div>
    
    
    
            <div class="col-12">
                <chart-card title="Network Throughput (tps)" sub-title="24 Hours performance" :chart-data="usersChart.data" :chart-options="usersChart.options">
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
import WalletCopiedNotification from './Notifications/WalletCopied';


export default {
    components: {
        StatsCard,
        WideCard,
        ChartCard
    },
    methods: {
        getTokens: function() {
            var self = this
            window.backend.retrieveTokenAmount().then(result => {
                self.tokenAmount = result;
            });
        },
        notifyVue(verticalAlign, horizontalAlign) {
            const color = 2;
            this.$notify({
                component: WalletCopiedNotification,
                icon: "ti-check",
                horizontalAlign: horizontalAlign,
                verticalAlign: verticalAlign,
                type: this.type[color]
            })
        }

    },
    computed: {
        wallet2() {
            return this.$store.state.walletInfo;
        },
        chartData() {
            return this.$store.state.chartData;
        },
        localWallet() {
            return this.$store.getters.localWallet;
        }
    },
    mounted() {
        window.wails.Events.On("token", (amount) => {
            this.$store.state.walletInfo.tokenAmount = amount;
        });
        window.wails.Events.On("blocks", (number) => {
            this.$store.state.walletInfo.blocks = number;
        });
        window.wails.Events.On("price", (currency, dagRate) => {
            let result = dagRate * this.tokenAmount;
            this.$store.state.walletInfo.usdValue = `${currency} ${(result).toFixed(2)}`;
        });
        window.wails.Events.On("token_counter", (count) => {
            this.count = count;
        });
        window.wails.Events.On("block_counter", (blockCount) => {
            this.blockCount = blockCount;
        });
        window.wails.Events.On("chart_counter", (pieChartCount) => {
            this.pieChartCount = pieChartCount;
        });
        window.wails.Events.On("node_stats", (series, labels) => {
            this.$store.state.chartData.nodesOnline.series = series
            this.$store.state.chartData.nodesOnline.labels = labels
            // this.stats = stats
        })
    },

    /**
     * Chart data used to render stats, charts. Should be replaced with server data
     */

    data() {
        return {
            count: "0",
            // stats: [0,0,0],
            blockCount: "0",
            pieChartCount: 24,
            type: ["", "info", "success", "warning", "danger"],
            notifications: {
                topCenter: false
            },
            walletCard: {

                type: "info",
                title: "Wallet Address",
                //address: Wallet.Address

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
                    labels: this.$store.state.chartData.transactions.labels,
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
                    height: "245px"
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

</style>
