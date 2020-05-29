<template>
  <span>There's a new version of Molly Wallet available. Click this notification to upgrade to v.{{ this.$store.state.newRelease }}.</span>
</template>

<script>
  export default {
    name: 'newRelease'
  }
</script>
