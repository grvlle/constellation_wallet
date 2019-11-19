<template>
    <card>
        <template slot="header">
          <h4 v-if="$slots.title || title" class="card-title">
            <slot name="title">
              {{title}}
            </slot>
          </h4>
          <p class="card-category">
            <slot name="subTitle">
              {{subTitle}}
            </slot>
          </p>
</template>
    <div>
      <div :id="chartId" class="ct-chart"></div>
      <div class="footer">
        <div class="chart-legend">
          <slot name="legend"></slot>
        </div>
        <hr>
        <div class="stats">
          <slot name="footer"></slot>
        </div>
        <div class="pull-right">
        </div>
      </div>
    </div>

  </card>
</template>

<script>
import Card from "./Card.vue";
export default {
    name: "chart-card",
    components: {
        Card
    },
    props: {
        footerText: {
            type: String,
            default: ""
        },
        title: {
            type: String,
            default: ""
        },
        subTitle: {
            type: String,
            default: ""
        },
        chartType: {
            type: String,
            default: "Line" // Line | Pie | Bar
        },
        chartOptions: {
            type: Object,
            default: () => {
                return {};
            }
        },
        chartData: {
            type: Object,
            default: () => {
                return {
                    labels: [],
                    series: []
                };
            }
        }
    },
    data() {
        return {
            chartId: "no-id",
            chartSelected: "",
            chartUpdated: false
        };
    },

    methods: {


        updateChart() {
          this.chartSelected.update();
          var self = this;
          setTimeout(self.updateChart, 1000)
        },
        /***
         * Initializes the chart by merging the chart options sent via props and the default chart options
         */

        initChart(Chartist) {
            const chartIdQuery = `#${this.chartId}`;
            var chart
            chart = Chartist[this.chartType](
                chartIdQuery,
                this.chartData,
                this.chartOptions
            );
            return chart
        },
        /***
         * Assigns a random id to the chart
         */
        updateChartId() {
            const currentTime = new Date().getTime().toString();
            const randomInt = this.getRandomInt(0, currentTime);
            this.chartId = `div_${randomInt}`;
        },
        getRandomInt(min, max) {
            return Math.floor(Math.random() * (max - min + 1)) + min;
        }
    },
    mounted() {
        this.updateChartId();


        import ('chartist').then((Chartist) => {
            let ChartistLib = Chartist.default || Chartist;
            this.$nextTick(() => {
                this.chartSelected = this.initChart(ChartistLib);
                // this.updateChart();
            });
        });

        window.setTimeout(this.updateChart, 0);
    }
};
</script>

<style>

</style>
