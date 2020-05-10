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
</style>
