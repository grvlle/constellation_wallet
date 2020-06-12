<template>
  <card>
    <div slot="header">
      <h4 v-if="$slots.title || title" class="card-title">
        <slot name="title">{{title}}</slot>
      </h4>
      <p class="card-category">
        <slot name="subTitle">{{subTitle}}</slot>
      </p>
  </div>
  

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
      count: 0,
      onWindows: false,
      onMacOS: false,
      onLinux: false,
    };
  },
  computed: {
    counter() {
      return this.$store.state.dashboard.counters.nodesOnlineCounter;
    },
    UserRunningOnWindows() {
      return this.$store.state.OS.windows;
    },
    UserRunningOnMacOS() {
      return this.$store.state.OS.macOS;
    },
    UserRunningOnLinux() {
      return this.$store.state.OS.linux;
    }
  },
  watch: {
    counter(val) {
      this.count = val;
      if (this.count == 1) {
        this.chartSelected.update();
      }
    },
    UserRunningOnWindows(val) {
        this.onWindows = val;
    },
    UserRunningOnMacOS(val) {
        this.onMacOS = val;
    },
    UserRunningOnLinux(val) {
        this.onLinux = val;
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
      );

      /* eslint-disable no-console */
      // console.log("asfg", Chartist.Svg.isSupported("http://www.w3.org/TR/SVG11/feature#SVG-animation"))

      // if (this.onWindows == false && this.onLinux == true || this.onMacOS == true) {
      // console.log("asfg", chart.Svg.isSupported("http://www.w3.org/TR/SVG11/feature#SVG-animation"))
      // }



      return chart;
    },

    includeAnimations(chart) {
      var delays = 0;
      var durations = 3000;

      chart.on("draw", function(data) {
        if (data.type === "line" || data.type === "area") {
          data.element.animate({
            d: {
              begin: 1000 * data.index,
              dur: 3000,
              from: data.path
                .clone()
                .scale(1, 0)
                .translate(0, data.chartRect.height())
                .stringify(),
              to: data.path.clone().stringify(),
              easing: "easeOutQuart"
            }
          });
          data.element.animate({
            opacity: {
              // The delay when we like to start the animation
              begin: delays,
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
              begin: 4000,
              dur: durations,
              from: 0,
              to: 1,
              easing: "easeOutQuart"
            }
          });
        }
      });
      
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
              // Animations are not yet supported on Windows
      if (Chartist.Svg.isSupported("AnimationEventsAttribute")) {
        this.includeAnimations(this.chartSelected)
      }
      });
    });
  }
};
</script>

<style>
</style>
