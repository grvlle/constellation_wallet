<template>
  <div class="container">
    <div class="total-value text-center">{{availableBalance | asCurrency('DAG-short')}} $DAG</div>
    <ul class="timeline">
      <li class="timeline-inverted" v-for="tx in value" v-bind:key="tx.ID">
        <div class="timeline-value" :class="[address == tx.receiver ? 'receive' : 'send']">
          <div
            v-if="address == tx.receiver"
          >+ {{tx.amount | normalizeDAG | asCurrency('DAG-short')}}</div>
          <div v-else>- {{tx.amount | normalizeDAG | asCurrency('DAG-short')}}</div>
        </div>
        <div class="timeline-badge" :class="[address == tx.receiver ? 'receive' : 'send']">
          <i v-if="tx.status == 'PENDING'" class="fa fa-spinner fa-pulse"></i>
          <i v-else-if="address == tx.receiver" class="fa fa-hand-holding-usd"></i>
          <i v-else class="fa fa-hand-holding-usd fa-flip-horizontal"></i>
        </div>
        <div class="timeline-panel" :class="[address == tx.receiver ? 'receive' : 'send']">
          <div class="container" style="padding: 0;">
            <div class="row">
              <div class="col-md-3" v-if="address == tx.receiver">
                Received: {{tx.amount | normalizeDAG | asCurrency('DAG')}}
              </div>
              <div class="col-md-3" v-else>
                Send: {{tx.amount | normalizeDAG | asCurrency('DAG')}}
              </div>
              <div class="col-md-7" v-if="address == tx.receiver">From: {{tx.receiver}}</div>
              <div class="col-md-7" v-else>To: {{tx.receiver}}</div>
              <div class="col-md-2 text-right">
                <small class="text-muted">{{tx.date}}</small>
              </div>
            </div>
            <div class="row">
              <div class="col">
                <a style="color:dimgray;" :href="'https://www.dagexplorer.io/search?term=' + tx.hash">{{tx.hash}}</a>
              </div>
            </div>
          </div>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
const suffixRanges = [
  { divider: 1e18, suffix: "E" },
  { divider: 1e15, suffix: "P" },
  { divider: 1e12, suffix: "T" },
  { divider: 1e9, suffix: "G" },
  { divider: 1e6, suffix: "M" },
  { divider: 1e3, suffix: "k" },
];

import { mapState } from "vuex";

export default {
  name: "timeline",
  props: {
    value: [],
  },
  computed: {
    ...mapState("wallet", ["address", "availableBalance"]),
  },
  filters: {
    asCurrency: function (value, currency) {
      if (currency == "") {
        return "";
      } else if (currency == "DAG-short") {
        for (var i = 0; i < suffixRanges.length; i++) {
          if (value >= suffixRanges[i].divider) {
            return (
              (value / suffixRanges[i].divider).toFixed(1) +
              suffixRanges[i].suffix
            );
          }
        }
        return value.toString();
      }

      var formatter;
      if (currency == "DAG") {
        formatter = new Intl.NumberFormat(navigator.language);
      } else if (currency == "BTC") {
        formatter = new Intl.NumberFormat(navigator.language, {
          style: "currency",
          currency: "XBT",
          minimumFractionDigits: 2,
          maximumFractionDigits: 8,
        });
      } else {
        formatter = new Intl.NumberFormat(navigator.language, {
          style: "currency",
          currency: currency,
          minimumFractionDigits: 2,
          maximumFractionDigits: 2,
        });
      }
      return formatter.format(value).replace(/XBT/, "₿");
    },
    normalizeDAG: function (value) {
      return (value / 1e8).toFixed(8).replace(/\.?0+$/, "");
    },
  },
};
</script>

<style scoped lang="scss">
.total-value {
  width: 9.6rem;
}
.timeline {
  list-style: none;
  padding: 1.25rem 0 0.25rem;
  margin-bottom: 0;
  position: relative;
}

.timeline:before {
  top: 0;
  bottom: 0;
  position: absolute;
  content: " ";
  width: 0.1875rem;
  left: 4.8rem;
  margin-left: -0.09375rem;
}

.timeline > li {
  margin-bottom: 1.25rem;
  position: relative;
}

.timeline > li:before,
.timeline > li:after {
  content: " ";
  display: table;
}

.timeline > li:after {
  clear: both;
}

.timeline > li:before,
.timeline > li:after {
  content: " ";
  display: table;
}

.timeline > li:after {
  clear: both;
}

.timeline > li > .timeline-panel {
  width: 88%;
  float: left;
  border-width: 0.0625rem;
  border-radius: 0.125rem;
  padding: 0.3125rem 1.25rem 0.625rem 1.25rem;
  position: relative;
  -webkit-box-shadow: 0 0.0625rem 0.375rem rgba(0, 0, 0, 0.175);
  box-shadow: 0 0.0625rem 0.375rem rgba(0, 0, 0, 0.175);
}
.timeline > li > .timeline-panel.send {
  border-left: solid #5bc0de;
}
.timeline > li > .timeline-panel.receive {
  border-left: solid #f0ad4e;
}

.timeline > li > .timeline-panel:before {
  position: absolute;
  top: 0.875rem;
  right: -0.9375rem;
  display: inline-block;
  border-top: 0.9375rem solid transparent;
  border-bottom: 0.9375rem solid transparent;
  border-left-style: solid;
  border-left-width: 0.9375rem;
  border-right-style: solid;  
  border-right-width: 0;  
  content: " ";
}

.timeline > li > .timeline-panel:after {
  position: absolute;
  top: 1rem;
  right: -0.875rem;
  display: inline-block;
  border-top: 0.875rem solid transparent;
  border-bottom: 0.875rem solid transparent;
  border-left-style: solid;
  border-left-width: 0.875rem;
  border-right-style: solid;  
  border-right-width: 0;  
  content: " ";
}

.timeline > li > .timeline-value {
  font-size: 0.875rem;
  text-align: center;
  position: absolute;
  top: 1.25rem;
  left: 0%;
  z-index: 100;
}

.timeline > li > .timeline-badge {
  color: #fff;
  width: 1.9rem;
  height: 1.9rem;
  line-height: 1.9rem;
  font-size: 1.4em;
  text-align: center;
  position: absolute;
  top: 0.9375rem;
  left: 4.8rem;
  margin-left: -0.9375rem;
  z-index: 100;
  border-top-right-radius: 50%;
  border-top-left-radius: 50%;
  border-bottom-right-radius: 50%;
  border-bottom-left-radius: 50%;
}

.timeline > li.timeline-inverted > .timeline-panel {
  float: right;
}

.timeline > li.timeline-inverted > .timeline-panel:before {
  border-left-width: 0;
  border-right-width: 0.9375rem;
  left: -0.9375rem;
  right: auto;
}

.timeline > li.timeline-inverted > .timeline-panel:after {
  border-left-width: 0;
  border-right-width: 0.875rem;
  left: -0.875rem;
  right: auto;
}
</style>