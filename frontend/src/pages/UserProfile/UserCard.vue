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
          <h5>{{normalizedAvailableBalance}}
            <br>
            <small>DAG</small>
          </h5>
        </div>
        <div class="col-4">
          <h5>{{valueInCurrency}}
            <br>
            <small>{{currency}}</small>
          </h5>
        </div>
      </div>
    </div>
  </card>
</template>

<script>
import {mapState, mapGetters} from 'vuex'
export default {
  computed: {
    ...mapState('wallet', ['imgPath', 'walletLabel', 'address', 'transactions', 'currency']),
    ...mapGetters('wallet', ['valueInCurrency', 'normalizedAvailableBalance'])
  }
};
</script>
<style scoped lang="scss">
</style>