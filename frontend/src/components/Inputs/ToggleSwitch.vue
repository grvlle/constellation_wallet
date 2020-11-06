<template>
  <label class="switch">
    <input type="checkbox" :checked="value" @change="changeHandler" />
    <span class="slider round"></span>
  </label>
</template>

<script>
export default {
  name: "toggle-switch",
  props: {
    value: Boolean,
  },
  methods: {
    changeHandler() {
      this.$emit("input", !this.value);
    },
  },
};
</script>

<style scoped lang="scss">
.switch {
  position: relative;
  display: inline-block;
  width: 2.5rem;
  height: 1.375rem;
  margin-bottom: 0;
}

.switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  @include themed() {
    background-color: t("inputPlaceholderColor");
  }
  -webkit-transition: 0.4s;
  transition: 0.4s;
}

.slider:before {
  position: absolute;
  content: "";
  height: 1.125rem;
  width: 1.125rem;
  left: 0.125rem;
  bottom: 0.125rem;
  @include themed() {
    background-color: t("cardTableColor");
  }
  -webkit-transition: 0.4s;
  transition: 0.4s;
}


input:checked + .slider:before {
  @include themed() {
    background-color: t("inputBackgroundColor");
  }
}
input:checked + .slider {
  @include themed() {
    background-color: t("infoColor");
  }
}


.slider {
  box-shadow: 0 0 0.125rem;
  @include themed() {
    color: t("inputBorderColor");
  }
}

input:checked + .slider:before {
  -webkit-transform: translateX(1.0625rem);
  -ms-transform: translateX(1.0625rem);
  transform: translateX(1.0625rem);
}

.slider.round {
  border-radius: 2.125rem;
}

.slider.round:before {
  border-radius: 50%;
}
</style>

