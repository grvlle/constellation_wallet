<template>
  <span><b>An error has occurred. </b> <br>Please visit and report the error to our <a href="https://t.me/StardustSupport" target="_blank"><b>Molly Wallet Support</b></a> channel on Telegram.<br><br> <b>Error: </b> <i>{{ this.$store.state.app.errorMessage }}</i></span>
</template>

<script>
  export default {
    name: 'error-message'
  }
</script>
