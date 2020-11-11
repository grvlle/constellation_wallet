<template>
  <div class="container">
    <div v-if="value.length">
      <div class="balance text-center">
        {{ normalizedAvailableBalance | asCurrency("DAG-short") }} $DAG
      </div>
      <ul class="timeline">
        <li
          class="timeline-inverted d-flex"
          v-for="tx in value"
          v-bind:key="tx.ID"
        >
          <div class="timeline-label">
            <div
              class="timeline-value"
              :class="[
                tx.status.toLowerCase(),
                address == tx.receiver ? 'receive' : 'send',
              ]"
            >
              <span v-if="address == tx.receiver">+</span>
              <span v-else>-</span>
              {{ tx.amount | normalizeDAG | asCurrency("DAG-short") }}
            </div>
            <div
              class="timeline-badge"
              :class="[
                tx.status.toLowerCase(),
                address == tx.receiver ? 'receive' : 'send',
              ]"
            >
              <i
                v-if="tx.status == 'Pending'"
                class="fa fa-spinner fa-spin"
              ></i>
              <i v-else-if="tx.status == 'Error'" class="fa fa-times"></i>
              <arrow-l-icon
                v-else-if="address == tx.receiver"
                :size="16"
                fillColor="#ffffff"
              />
              <arrow-r-icon v-else :size="16" fillColor="#ffffff" />
            </div>
          </div>
          <div
            class="timeline-panel"
            :class="[
              tx.status.toLowerCase(),
              address == tx.receiver ? 'receive' : 'send',
            ]"
          >
            <div class="container" style="padding: 0;">
              <div class="row">
                <div class="col-10">
                  <div class="row">
                    <div
                      class="col-md-3 text-truncate"
                      v-if="address == tx.receiver"
                    >
                      received:&nbsp;{{
                        tx.amount | normalizeDAG | asCurrency("DAG")
                      }}
                    </div>
                    <div class="col-md-3 text-truncate" v-else>
                      send:&nbsp;{{
                        tx.amount | normalizeDAG | asCurrency("DAG")
                      }}
                    </div>
                    <div class="col-md-9 text-truncate">
                      <div v-if="address == tx.receiver">
                        <span>from:&nbsp;</span>
                        <span v-html="displayContact(tx.sender)" />
                      </div>
                      <div v-else>
                        <span>to:&nbsp;</span>
                        <span v-html="displayContact(tx.receiver)" />
                      </div>
                    </div>
                  </div>
                  <div class="row">
                    <div class="col text-truncate">
                      <button
                        class="link text-muted"
                        @click="
                          openURL(
                            'https://www.dagexplorer.io/search?term=' + tx.hash
                          )
                        "
                        rel="noopener noreferrer"
                        target="_blank"
                      >
                        {{ tx.hash }}
                      </button>
                      >
                    </div>
                  </div>
                </div>
                <div class="col-md-2 text-right">
                  <small class="text-muted">{{ tx.date }}</small>
                </div>
              </div>
            </div>
          </div>
        </li>
      </ul>
    </div>
    <div v-else class="text-center" style="height: 10rem;">
      <p class="card-text text-muted font-weight-bold mt-4">NO TRANSACTIONS</p>
    </div>
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

import { mapState, mapGetters } from "vuex";

export default {
  name: "timeline",
  props: {
    value: Array,
  },
  computed: {
    ...mapState("wallet", ["address"]),
    ...mapGetters("wallet", ["normalizedAvailableBalance"]),
  },
  methods: {
    displayContact: function(value) {
      // let contactInfo = "<span class='text-danger'>unknown</span>";
      let contact = this.$store.getters["addressBook/search"](value);
      if (contact.length) {
        return "<span>" + contact[0].name + "</span>";
      }
      return value;
    },
    openURL(url) {
      window.backend.WalletApplication.OpenBrowser(url);
    },
  },
  filters: {
    asCurrency: function(value, currency) {
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
      } else {
        var formatter = new Intl.NumberFormat(navigator.language);
        return formatter.format(value);
      }
    },
    normalizeDAG: function(value) {
      return (value / 1e8).toFixed(8).replace(/\.?0+$/, "");
    },
  },
};
</script>

<style scoped lang="scss">
.timeline-label {
  width: 9.6rem;
}

.link {
  background: transparent;
  border: none;
  padding: 0;

  &:hover {
    text-decoration: underline;
  }
}

.balance {
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
  width: 0.0625rem;
  border-left: 0.0625rem dashed;
  left: 4.8rem;
  margin-left: -0.09375rem;
  @include themed() {
    border-color: #c4c4c4;
  }
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

.timeline-value {
  font-size: 0.875rem;
  text-align: center;
  position: absolute;
  top: 1.25rem;
  left: 0%;
  z-index: 100;
}

.timeline-value.receive {
  @include themed() {
    color: #dbb018;
  }
}

.timeline-value.send {
  @include themed() {
    color: #2d9cdb;
  }
}

.timeline-value.error {
  @include themed() {
    color: t("dangerColor");
  }
}

.timeline-badge {
  width: 1.9rem;
  height: 1.9rem;
  line-height: 1.9rem;
  font-size: 0.875rem;
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
  display: flex;
  justify-content: center;
  align-items: center;
}

.timeline-badge .fa {
  @include themed() {
    color: t("cardBackgroundColor");
  }
}

.timeline-badge.receive {
  @include themed() {
    background-color: #dbb018;
  }
}

.timeline-badge.send {
  @include themed() {
    background-color: #2d9cdb;
  }
}

.timeline-badge.error {
  @include themed() {
    background-color: t("dangerColor");
  }
}

.timeline-badge.pending {
  @include themed() {
    background-color: t("borderColor");
  }
}
.timeline-badge.pending .fa {
  @include themed() {
    color: t("cardTextColor");
  }
}
.timeline-badge.pending.send {
  border-style: solid;
  @include themed() {
    border-color: #2d9cdb;
  }
}
.timeline-badge.pending.receive {
  border-style: solid;
  @include themed() {
    border-color: #dbb018;
  }
}

.timeline-panel:before {
  position: absolute;
  top: 1.175rem;
  right: -0.75rem;
  display: inline-block;
  border-top: 0.75rem solid transparent;
  border-bottom: 0.75rem solid transparent;
  border-left-style: solid;
  border-left-width: 0.75rem;
  border-right-style: solid;
  border-right-width: 0;
  content: " ";
}

// .timeline-panel:after {
//   position: absolute;
//   top: 1rem;
//   right: -0.875rem;
//   display: inline-block;
//   border-top: 0.875rem solid transparent;
//   border-bottom: 0.875rem solid transparent;
//   border-left-style: solid;
//   border-left-width: 0.875rem;
//   border-right-style: solid;
//   border-right-width: 0;
//   content: " ";
//   @include themed() {
//     border-left-color: t("cardTableColor");
//     border-right-color: t("cardTableColor");
//   }
// }

.timeline-inverted .timeline-panel:before {
  border-left-width: 0;
  border-right-width: 0.9375rem;
  left: -0.9375rem;
  right: auto;
}

.timeline-panel.send:before {
  @include themed() {
    border-left-color: #2d9cdb;
    border-right-color: #2d9cdb;
  }
}

.timeline-panel.receive:before {
  @include themed() {
    border-left-color: #dbb018;
    border-right-color: #dbb018;
  }
}

.timeline-panel.error:before {
  @include themed() {
    border-left-color: t("dangerColor");
    border-right-color: t("dangerColor");
  }
}

.timeline-inverted .timeline-panel:after {
  border-left-width: 0;
  border-right-width: 0.875rem;
  left: -0.875rem;
  right: auto;
}

.timeline-panel {
  width: 100%;
  padding: 0.75rem;
  position: relative;
  background: #f2f2f2;
  border: 0.0625rem solid #c4c4c4;
  box-sizing: border-box;
  box-shadow: 0 0 0.3125rem rgba(0, 0, 0, 0.25);
  -webkit-box-shadow: 0 0 0.3125rem rgba(0, 0, 0, 0.25);
  border-radius: 0.375rem;
}

// .timeline-panel.send {
//   @include themed() {
//     border-left-color: #2d9cdb;
//   }
// }

// .timeline-panel.receive {
//   @include themed() {
//     border-left-color: #dbb018;
//   }
// }

// .timeline-panel.error {
//   @include themed() {
//     border-left-color: t("dangerColor");
//   }
// }
</style>
