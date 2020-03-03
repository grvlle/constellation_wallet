<template>
  <div class="bg">
    <center>
      <img
      v-if="this.$store.state.app.isLoading && this.$store.state.app.isLoggedIn"
      style="margin-top: 12.5em;"
      src="~@/assets/img/Constellation-Logo-1.png"
      />
      <p
        v-if="this.$store.state.app.isLoading && this.$store.state.app.isLoggedIn"
        style="color: #c4c4c4; margin-top: 0.3125em;"
      >Getting your $DAG Wallet ready...</p>
    </center>
    <center>
      <div class="boxes">
        <div class="box">
          <div></div>
          <div></div>
          <div></div>
          <div></div>
        </div>
        <div class="box">
          <div></div>
          <div></div>
          <div></div>
          <div></div>
        </div>
        <div class="box">
          <div></div>
          <div></div>
          <div></div>
          <div></div>
        </div>
        <div class="box">
          <div></div>
          <div></div>
          <div></div>
          <div></div>
        </div>
      </div>
    </center>
  </div>
</template>

<script>
export default {
  name: "loading-screen"
  // props: ["isLoading"]
};
</script>

<style lang="scss" scoped>
.bg {
  /* The image used */
  background-image: linear-gradient(
      rgba(255, 255, 255, 0.2),
      rgba(255, 255, 255, 0.2)
    ),
    url("~@/assets/img/nodes2.jpg");

  /* Full height */
  height: 100%;
  position: absolute;
  width: 100%;

  /* Center and scale the image nicely */
  background-position: center;
  background-repeat: no-repeat;
  background-size: cover;
}

.loader {
  background: #f2f2f2;
  background-repeat: no-repeat;
  background-size: cover;
  color: white;
  display: block;
  font-size: 2rem;
  overflow: hidden;
  padding-top: 35vh;
  position: fixed;
  text-align: center;
}

$color: #5c8df6;
$colorRight: darken(#5c8df6, 15%);
$colorLeft: darken(#5c8df6, 5%);
$shadow: #dbe3f4;
.boxes {
  --size: 2em;
  --duration: 800ms;
  text-align: center;
  height: calc(var(--size) * 2);
  width: calc(var(--size) * 3);
  position: relative;
  transform-style: preserve-3d;
  transform-origin: 50% 50%;
  margin-top: 3.125em;
  transform: rotateX(60deg) rotateZ(45deg) rotateY(0deg) translateZ(0em);
  .box {
    width: var(--size);
    height: var(--size);
    top: 0;
    left: 0;
    position: absolute;
    transform-style: preserve-3d;
    &:nth-child(1) {
      transform: translate(100%, 0);
      animation: box1 var(--duration) linear infinite;
    }
    &:nth-child(2) {
      transform: translate(0, 100%);
      animation: box2 var(--duration) linear infinite;
    }
    &:nth-child(3) {
      transform: translate(100%, 100%);
      animation: box3 var(--duration) linear infinite;
    }
    &:nth-child(4) {
      transform: translate(200%, 0);
      animation: box4 var(--duration) linear infinite;
    }
    & > div {
      --background: #{$color};
      --top: auto;
      --right: auto;
      --bottom: auto;
      --left: auto;
      --translateZ: calc(var(--size) / 2);
      --rotateY: 0deg;
      --rotateX: 0deg;
      position: absolute;
      width: 100%;
      height: 100%;
      background: var(--background);
      top: var(--top);
      right: var(--right);
      bottom: var(--bottom);
      left: var(--left);
      transform: rotateY(var(--rotateY)) rotateX(var(--rotateX))
        translateZ(var(--translateZ));
      &:nth-child(1) {
        --top: 0;
        --left: 0;
      }
      &:nth-child(2) {
        --background: #{$colorRight};
        --right: 0;
        --rotateY: 90deg;
      }
      &:nth-child(3) {
        --background: #{$colorLeft};
        --rotateX: -90deg;
      }
      &:nth-child(4) {
        --background: #{$shadow};
        --top: 0;
        --left: 0;
        --translateZ: calc(var(--size) * 3 * -1);
      }
    }
  }
}

@keyframes box1 {
  0%,
  50% {
    transform: translate(100%, 0);
  }
  100% {
    transform: translate(200%, 0);
  }
}

@keyframes box2 {
  0% {
    transform: translate(0, 100%);
  }
  50% {
    transform: translate(0, 0);
  }
  100% {
    transform: translate(100%, 0);
  }
}

@keyframes box3 {
  0%,
  50% {
    transform: translate(100%, 100%);
  }
  100% {
    transform: translate(0, 100%);
  }
}

@keyframes box4 {
  0% {
    transform: translate(200%, 0);
  }
  50% {
    transform: translate(200%, 100%);
  }
  100% {
    transform: translate(100%, 100%);
  }
}

html {
  -webkit-font-smoothing: antialiased;
}

* {
  box-sizing: border-box;
  &:before,
  &:after {
    box-sizing: border-box;
  }
}

.fadeout {
  animation: fadeout 2s forwards;
}

@keyframes fadeout {
  to {
    opacity: 0;
    visibility: hidden;
  }
}
</style>
