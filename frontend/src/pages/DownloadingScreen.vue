<template>
  <div class="bg">
    <center>
      <img
      v-if="this.$store.state.app.isDownloadingDependencies"
      style="margin-top: 240px;"
      src="~@/assets/img/Constellation-Logo-1.png"
      />
      <p
        v-if="this.$store.state.app.isDownloadingDependencies"
        style="color: #c4c4c4; margin-top: 5px;"
      >Downloading $DAG wallet dependencies...</p>
      <p v-if="this.$store.state.downloading.filename !== ''">
        {{this.$store.state.downloading.filename}}: {{this.$store.state.downloading.size}}
      </p>
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
  name: "downloading-screen",
  created: function() {
    window.backend.WalletApplication.CheckAndFetchWalletCLI();
  }
  // props: ["isDownloading"]
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
  font-size: 32px;
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
  --size: 32px;
  --duration: 800ms;
  text-align: center;
  height: calc(var(--size) * 2);
  width: calc(var(--size) * 3);
  position: relative;
  transform-style: preserve-3d;
  transform-origin: 50% 50%;
  margin-top: 50px;
  transform: rotateX(60deg) rotateZ(45deg) rotateY(0deg) translateZ(0px);
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

/* // Center & dribbble */

// body {
//     min-height: 100vh;
//     font-family: Roboto, Arial;
//     color: #ADAFB6;
//     display: flex;
//     justify-content: center;
//     align-items: center;
//     background: #F9FBFF;
//     .dribbble {
//         position: fixed;
//         display: block;
//         right: 20px;
//         bottom: 20px;
//         img {
//             display: block;
//             height: 28px;
//         }
//     }
// }
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
