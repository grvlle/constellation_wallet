<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-md-6">
        <card title="Graph #1" sub-title="Grafana embedded test">
            <div class="embed-responsive embed-responsive-16by9">
            <iframe 
                class="embed-responsive-item" 
                :src="'http://grafana.constellationnetwork.io:3000/d-solo/4a_rfgjWz/alerting?orgId=1&refresh=10s&from=1589865856111&to=1589869456111&panelId=20' + grafanaTheme" allowfullscreen>
            </iframe>
            </div>
        </card>
      </div>
      <div class="col-md-6">
        <card title="Graph #2" sub-title="Grafana embedded test">
            <div class="embed-responsive embed-responsive-16by9">
            <iframe 
                class="embed-responsive-item" 
                :src="'http://grafana.constellationnetwork.io:3000/d-solo/4a_rfgjWz/alerting?orgId=1&refresh=10s&from=1589865856111&to=1589869456111&panelId=18' + grafanaTheme" allowfullscreen>
            </iframe>
            </div>
        </card>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  computed: {
    grafanaTheme: function() {
      if (this.$store.state.walletInfo.darkMode) {
        return "&theme=dark"
      } else {
        return "&theme=light"
      }
    }
  }
};
</script>

<style lang="scss">
</style>
