<template>
  <div class="container">
    <div class="row">
      <div class="col-md-6 mx-auto text-center">
        <loader />
      </div>
    </div>
    <div class="row" style="height: 21rem;" />
  </div>
</template>


<script>
import Loader from "../components/Loader";
export default {
  name: "loading-screen",
  components: {
    Loader
  },
  mounted() {
    setTimeout(() => {
      this.$router.push({name: 'home'})
    }, 4000);
  }
};
</script>

<style scoped lang="scss">
</style>