<template>
  <span><b>An error has occurred! :(</b> <br><br>Please reach out to <b>@armedvito</b> on twitter and provide the log file.<br><br> <b>Error: </b> <i>{{ this.$store.state.errorMessage }}</i></span>
</template>

<script>
  export default {
    name: 'error-message'
  }
</script>
