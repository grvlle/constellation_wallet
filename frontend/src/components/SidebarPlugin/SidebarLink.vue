<template>
  <component :is="tag" @click.native="hideSidebar" class="nav-item" v-bind="$attrs" tag="li">
    <slot>
      <p class="nav-item">
        <a class="nav-link">
          <i v-if="icon" :class="icon"></i>
          {{name}}
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
      default: true
    },
    addLink: {
      default: () => {}
    },
    removeLink: {
      default: () => {}
    }
  },
  props: {
    name: String,
    icon: String,
    tag: {
      type: String,
      default: "router-link"
    }
  },
  methods: {
    hideSidebar() {
      if (this.autoClose) {
        this.$sidebar.displaySidebar(false);
      }
    },
    isActive() {
      return this.$el.classList.contains("active");
    }
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
  }
};
</script>

<style>
@import url("https://fonts.googleapis.com/css?family=Poppins&display=swap");
</style>
