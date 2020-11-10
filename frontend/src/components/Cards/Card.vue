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
  padding: 12px !important;
}
.card-header {
  margin: 0px 24px;
  margin-top: 18px;
  padding: 0px !important;

  .card-title {
    font-style: normal;
    font-weight: normal;
    font-size: 24px;
    line-height: 20px;
    color: #666666 !important;
  }

  .card-category {
    font-style: normal;
    font-weight: normal;
    font-size: 12px;
    line-height: 20px;
    margin-top: 5px;
    color: #979797 !important;
  }
}
</style>
