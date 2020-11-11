<template>
  <div class="form-group" :class="{ 'input-group': hasIcon }">
    <slot name="label">
      <label v-if="label" class="control-label">
        {{ label }}
      </label>
    </slot>
    <slot name="addonLeft">
      <span v-if="addonLeftIcon" class="input-group-prepend">
        <i :class="addonLeftIcon" class="input-group-text"></i>
      </span>
    </slot>
    <input
      :value="value"
      @input="$emit('input', $event.target.value)"
      v-bind="$attrs"
      class="form-control"
      aria-describedby="addon-right addon-left"
    />
    <slot></slot>
    <slot name="addonRight">
      <span v-if="addonRightIcon" class="input-group-append">
        <i :class="addonRightIcon" class="input-group-text"></i>
      </span>
    </slot>
  </div>
</template>
<script>
export default {
  inheritAttrs: false,
  name: "fg-input",
  props: {
    label: String,
    value: [String, Number, Boolean],
    addonRightIcon: String,
    addonLeftIcon: String,
  },
  computed: {
    hasIcon() {
      const { addonRight, addonLeft } = this.$slots;
      return (
        addonRight !== undefined ||
        addonLeft !== undefined ||
        this.addonRightIcon !== undefined ||
        this.addonLeftIcon !== undefined
      );
    },
  },
};
</script>

<style scoped lang="scss">
.form-control {
  background: #f2f2f2 !important;
  font-family: Poppins;
  color: #666666;
  box-sizing: border-box;
  border-radius: 0.25rem;
  border: 0.0625rem solid #c4c4c4 !important;
}

.control-label {
  font-weight: 500;
  font-size: 0.875rem;
  line-height: 1.3125rem;
  color: #979797;
}

.theme--dark {
  .form-control {
    border-color: #979797 !important;
    background: #666666 !important;
    color: #fff;
  }

  .control-label {
    color: #fff;
  }
}
</style>
