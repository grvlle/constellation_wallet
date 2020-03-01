<template>
  <div class="moving-arrow" :style="arrowStyle">
  </div>
</template>
<script>
export default {
  props: {
    moveY: {
      type: Number,
      default: 0
    }
  },
  computed: {
    /**
     * Styles to animate the arrow
     * @returns {{transform: string}}
     */
    arrowStyle() {
      return {
        transform: `translate3d(0em, ${this.moveY}em, 0em)`
      };
    }
  }
};
</script>
<style lang="scss">
$bg-nude: #f4f3ef !default;
.moving-arrow {
  border-right: 1.0625em solid $bg-nude;
  border-top: 1.0625em solid transparent;
  border-bottom: 1.0625em solid transparent;
  display: inline-block;
  position: absolute;
  left: 15.1875em;
  top: 16.5em;
  transition: all 0.5s cubic-bezier(0.19, 1.42, 0.79, 1);
}
@media (max-width: 992px) {
  .moving-arrow {
    display: none;
  }
}
</style>
