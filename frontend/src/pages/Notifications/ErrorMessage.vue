<template>
  <span><b>An error has occurred. </b> <br>But don't worry, funds are safe. <br>Please reach create a GitHub issue or reach out directly to <b>@armedvito</b> on Twitter or Telegram and provide the wallet.log file located in your wallet directory.<br><br> <b>Error: </b> <i>{{ this.$store.state.errorMessage }}</i></span>
</template>

<script>
  export default {
    name: 'error-message'
  }
</script>
