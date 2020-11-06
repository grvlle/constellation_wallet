<template>
  <card>
    <div>
      <div class="row">
        <div class="col-5" v-if="$slots.header">
          <slot name="header"></slot>
        </div>
        <div class="col-xl-12" v-if="$slots.content">
          <slot name="content"></slot>
        </div>
      </div>
      <div v-if="$slots.footer">
        <slot name="footer"></slot>
      </div>
    </div>

  </card>
</template>
<script>
import Card from "./Card.vue";

export default {
  name: "wide-card",
  components: {
    Card
  }
};
</script>

<style scoped lang="scss">
</style>