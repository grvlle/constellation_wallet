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
  padding: 0.75rem !important;
}
.card-header {
  margin: 0 1.5rem;
  margin-top: 1.125rem;
  padding: 0 !important;

  .card-title {
    font-style: normal;
    font-weight: normal;
    font-size: 1.5rem;
    line-height: 1.25rem;
    color: #666666 !important;
  }

  .card-category {
    font-style: normal;
    font-weight: normal;
    font-size: 0.75rem;
    line-height: 1.25rem;
    margin-top: 0.3125rem;
    color: #979797 !important;
  }
}

.theme--dark {
  .card {
    background: #181818;

    .card-title,
    .card-category {
      color: #f2f2f2 !important;
    }
  }
}
</style>
