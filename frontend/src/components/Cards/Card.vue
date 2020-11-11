<template>
  <div class="card">
    <div class="card-image" v-if="$slots.image">
      <slot name="image"></slot>
    </div>
    <div class="card-header" v-if="$slots.header || title">
      <slot name="header">
        <h4 class="card-title">{{ title }}</h4>
        <p class="card-category" v-if="subTitle">{{ subTitle }}</p>
      </slot>
    </div>
    <div
      class="card-body"
      v-if="$slots.default"
      v-bind:class="{ 'tight-card-body': tightPadding }"
    >
      <slot></slot>
    </div>
    <slot name="raw-content"></slot>
    <div class="card-footer" v-if="$slots.footer">
      <hr />
      <slot name="footer"></slot>
    </div>
  </div>
</template>
<script>
export default {
  name: "card",
  props: {
    title: String,
    subTitle: String,
    tightPadding: Boolean,
  },
};
</script>

<style scoped lang="scss">
.tight-card-body {
  padding: 0.75em !important;
}
.card-header {
  margin: 0em 1.5em;
  margin-top: 1.125em;
  padding: 0em !important;

  .card-title {
    font-style: normal;
    font-weight: normal;
    font-size: 1.5em;
    line-height: 1.25em;
    color: #666666 !important;
  }

  .card-category {
    font-style: normal;
    font-weight: normal;
    font-size: 0.75em;
    line-height: 1.25em;
    margin-top: 0.3125em;
    color: #979797 !important;
  }
}
</style>
