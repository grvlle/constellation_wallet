<template>
  <card class="card-user">
    <div slot="image">
      <img src="@/assets/img/dotz.png" style="margin-top: -9.6875em" alt="...">
    </div>
    <div>
      <div class="author">
        <img class="avatar border-white" :src="require('@/assets/img/' + imgPath)" alt="...">
        <h4 class="title">{{walletLabel}}
          <br>
            <small>{{address}}</small>
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
          <h5>{{transactions}}
            <br>
            <small>Transactions</small>
          </h5>
        </div>
        <div class="col-4">
          <h5>{{tokenAmount | asCurrency('DAG')}}
            <br>
            <small>DAG</small>
          </h5>
        </div>
        <div class="col-4">
          <h5>{{totalValue | asCurrency(currency)}}
            <br>
            <small>{{currency}}</small>
          </h5>
        </div>
      </div>
    </div>
  </card>
</template>

<script>
import {mapState} from 'vuex'
export default {
  computed: {
    ...mapState('wallet', 
      ['imgPath', 'walletLabel', 'address', 'transactions', 'tokenAmount', 'totalValue', 'currency'])
  },
  filters: {
    asCurrency: function(value, currency) {

      if (currency == "") return "";
      
      var formatter
      if (currency == "DAG") {
        formatter = new Intl.NumberFormat(navigator.language);
      } else if (currency == "BTC") {
        formatter = new Intl.NumberFormat(navigator.language, {
          style: "currency",
          currency: "XBT",
          minimumFractionDigits: 2,
          maximumFractionDigits: 8
        });
      } else {
        formatter = new Intl.NumberFormat(navigator.language, {
          style: "currency",
          currency: currency,
          minimumFractionDigits: 2,
          maximumFractionDigits: 2
        });
      }
      return formatter.format(value).replace(/XBT/,'₿');
    }
  },
};
</script>

<style>
</style>
