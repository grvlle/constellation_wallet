<template>
  <card class="card-user">
    <div slot="image">
      <img src="@/assets/img/dotz.png" style="margin-top: -9.6875em" alt="...">
    </div>
    <div>
      <div class="author">
        <img class="avatar border-white" :src="require('@/assets/img/' + this.$store.state.walletInfo.imgPath)" alt="...">
        <h4 class="title">{{this.$store.state.walletInfo.email}}
          <br>
            <small>{{this.$store.state.walletInfo.address}}</small>
        </h4>
      </div>
      <p class="description text-center">
        <br> <br>
      </p>
    </div>
    <hr>
    <div class="text-center">
      <br>
      <div class="row">
        <div class="col-4">
          <h5>{{this.$store.state.walletInfo.transactions}}
            <br>
            <small>Transactions</small>
          </h5>
        </div>
        <div class="col-4">
          <h5>{{this.$store.state.walletInfo.tokenAmount | asDAGs}}
            <br>
            <small>DAG</small>
          </h5>
        </div>
        <div class="col-4">
          <h5>{{this.$store.state.walletInfo.totalValue | asCurrency(this.$store.state.walletInfo.currency)}}
            <br>
            <small>{{this.$store.state.walletInfo.currency}}</small>
          </h5>
        </div>
      </div>
    </div>
  </card>
</template>

<script>
export default {
  filters: {
    asCurrency: function (value, currency) {
      var formatter = new Intl.NumberFormat(navigator.language, {
        style: 'currency',
        currency: currency,
        minimumFractionDigits: 0
      });
      return formatter.format(value);
    },
    asDAGs: function (value) {
      var formatter = new Intl.NumberFormat(navigator.language);
      return formatter.format(value);
    }
  },
};
</script>

<style>
</style>
