<template>
  <span><b>An error has occurred. </b> <br>But don't worry, your funds are safe. <br>Please create a <a href="https://github.com/StardustCollective/molly_wallet/issues"><b>GitHub issue</b></a> or visit our <a href="https://t.me/StardustSupport"><b>Moly Wallet Support</b></a> channel on Telegram.<br><br> <b>Error: </b> <i>{{ this.$store.state.app.errorMessage }}</i></span>
</template>

<script>
  export default {
    name: 'error-message'
  }
</script>
