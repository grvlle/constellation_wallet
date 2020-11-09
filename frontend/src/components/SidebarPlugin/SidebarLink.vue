<template>
  <component
    :is="tag"
    @click.native="hideSidebar"
    class="nav-item"
    v-bind="$attrs"
    tag="li"
  >
    <slot>
      <p class="nav-item">
        <a class="nav-link">
          <dashboard-icon class="s-icon" v-if="icon === 'dashboard'" />
          <wallet-icon class="s-icon" v-if="icon === 'wallet'" />
          <tx-icon class="s-icon" v-if="icon === 'tx'" />
          <about-icon class="s-icon" v-if="icon === 'about'" />
          <i v-else-if="icon" :class="icon"></i>
          {{ name }}
        </a>
      </p>
    </slot>
  </component>
</template>

<script>
export default {
  name: "sidebar-link",
  inheritAttrs: false,
  inject: {
    autoClose: {
      default: true,
    },
    addLink: {
      default: () => {},
    },
    removeLink: {
      default: () => {},
    },
  },
  props: {
    name: String,
    icon: String,
    tag: {
      type: String,
      default: "router-link",
    },
  },
  methods: {
    hideSidebar() {
      if (this.autoClose) {
        this.$sidebar.displaySidebar(false);
      }
    },
    isActive() {
      return this.$el.classList.contains("active");
    },
  },
  mounted() {
    if (this.addLink) {
      this.addLink(this);
    }
  },
  beforeDestroy() {
    if (this.$el && this.$el.parentNode) {
      this.$el.parentNode.removeChild(this.$el);
    }
    if (this.removeLink) {
      this.removeLink(this);
    }
  },
};
</script>

<style scoped lang="scss">
.theme--light a:focus {
  color: #57b0e2 !important;
}

.nav-link {
  font-weight: normal;
  font-size: 14px;
  line-height: 24px;
  text-transform: none;
  color: $gray-input-bg !important;
}

.nav-item:first-child {
  margin-top: 10px;
}

.s-icon {
  margin-left: 12px;
  margin-right: 8px;

  svg {
    width: 18px;
    height: 18px;
  }
}
</style>
