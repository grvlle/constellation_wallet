<template>
<div class="overlay" v-bind:class="{ active: isActive }">
    <div class="overlay-text">{{text}}</div>
</div>
</template>

<script>
  export default {
    name: "page-overlay",
    props: {
      text: String,
      isActive: Boolean
    }
  }
</script>

<style>
.overlay {
  position: fixed; /* Sit on top of the page content */
  display: none; /* Hidden by default */
  width: 100%; /* Full width (cover the whole page) */
  height: 100%; /* Full height (cover the whole page) */
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0,0,0,0.5); /* Black background with opacity */
  z-index: 2; /* Specify a stack order in case you're using a different order for other elements */
  cursor: pointer; /* Add a pointer on hover */
}

.active {
    display: block;
}

.overlay-text {
  position: absolute;
  top: 50%;
  left: 50%;
  font-size: 2rem;
  color: white;
  transform: translate(-50%,-50%);
  -ms-transform: translate(-50%,-50%);
}


</style>