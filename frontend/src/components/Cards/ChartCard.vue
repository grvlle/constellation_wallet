<template>
  <card>
    <template slot="header">
      <h4 v-if="$slots.title || title" class="card-title">
        <slot name="title">{{title}}</slot>
      </h4>
      <p class="card-category">
        <slot name="subTitle">{{subTitle}}</slot>
      </p>
  </card>
  </template>
  
<template>
    <card>
      <div>
        <div :id="chartId" class="ct-chart"></div>
        <div class="footer">
          <div class="chart-legend">
            <slot name="legend"></slot>
          </div>
          <hr />
          <div class="stats">
            <slot name="footer"></slot>
          </div>
          <div class="pull-right"></div>
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
      chartUpdated: false,
      count: 0
    };
  },
  computed: {
    counter() {
      return this.$store.state.counters.nodesOnlineCounter;
    }
  },
  watch: {
    counter(val) {
      this.count = val;
      if (this.count == 6) {
        this.chartSelected.update();
      }
    }
  },

  methods: {
    /***
     * Initializes the chart by merging the chart options sent via props and the default chart options
     */

    initChart(Chartist) {
      const chartIdQuery = `#${this.chartId}`;
      var chart;

      chart = Chartist[this.chartType](
        chartIdQuery,
        this.chartData,
        this.chartOptions,
        this.chartTask
      );

      var delays = 80;
      var durations = 500;

      chart.on("draw", function(data) {
        if (data.type === "line" || data.type === "area") {
          data.element.animate({
            d: {
              begin: 2000 * data.index,
              dur: 3000,
              from: data.path
                .clone()
                .scale(1, 0)
                .translate(0, data.chartRect.height())
                .stringify(),
              to: data.path.clone().stringify(),
              easing: Chartist.Svg.Easing.easeOutQuart
            }
          });
          data.element.animate({
            opacity: {
              // The delay when we like to start the animation
              begin: delays + 1000,
              // Duration of the animation
              dur: durations,
              // The value where the animation should start
              from: 0,
              // The value where it should end
              to: 1
            }
          });
        }
        if (data.type === "point") {
          data.element.animate({
            x1: {
              begin: delays,
              dur: durations,
              from: data.x - 10,
              to: data.x,
              easing: "easeOutQuart"
            },
            x2: {
              begin: delays,
              dur: durations,
              from: data.x - 10,
              to: data.x,
              easing: "easeOutQuart"
            },
            opacity: {
              begin: delays,
              dur: durations,
              from: 0,
              to: 1,
              easing: "easeOutQuart"
            }
          });
        }
      });

      return chart;
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

    import("chartist").then(Chartist => {
      let ChartistLib = Chartist.default || Chartist;
      this.$nextTick(() => {
        this.chartSelected = this.initChart(ChartistLib);
      });
    });
  }
};
</script>

<style>
</style>
