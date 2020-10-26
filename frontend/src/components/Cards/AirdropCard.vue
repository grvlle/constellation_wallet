<template>
  <button
    class="airdrop-card"
    @mouseenter="handleMouseEnter"
    @mouseleave="handleMouseLeave"
    @click="onClick"
  >
    <img :src="image" style="width: 100%;" />
  </button>
</template>
<script>
import airdrop1 from "../../assets/img/airdrop1.png";
import airdrop2 from "../../assets/img/airdrop2.png";

export default {
  name: "airdrop-card",
  data() {
    return {
      image: airdrop1,
    };
  },
  methods: {
    handleMouseEnter() {
      this.image = airdrop2;
    },
    handleMouseLeave() {
      this.image = airdrop1;
    },
    onClick() {
      this.$emit("v-click");
    },
  },
};
</script>
<style scoped lang="scss">
.airdrop-card {
  background: transparent;
  padding: 0px;
  border: 0px;
  padding-bottom: 17.5px;
}
</style>
