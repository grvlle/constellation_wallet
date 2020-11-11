<template>
  <card class="card-user">
    <div slot="image">
      <img
        src="@/assets/img/dotz.png"
        style="margin-top: -9.6875em"
        alt="..."
      />
    </div>
    <div>
      <div class="author">
        <img
          class="avatar border-white"
          :src="require('@/assets/img/' + imgPath)"
          alt="..."
        />
        <h4 class="title">
          {{ walletLabel || "hello world" }}
          <small>{{ address }}</small>
        </h4>
      </div>
      <p class="description text-center">
        <br />
        <br />
      </p>
    </div>
    <hr />
    <div class="text-center">
      <br />
      <div class="row">
        <div class="col-4">
          <h5>
            {{ transactions }}
            <br />
            <small>Transactions</small>
          </h5>
        </div>
        <div class="col-4">
          <h5>
            {{ normalizedAvailableBalance }}
            <br />
            <small>DAG</small>
          </h5>
        </div>
        <div class="col-4">
          <h5>
            {{ valueInCurrency }}
            <br />
            <small>{{ currency }}</small>
          </h5>
        </div>
      </div>
    </div>
  </card>
</template>

<script>
import { mapState, mapGetters } from "vuex";
export default {
  computed: {
    ...mapState("wallet", [
      "imgPath",
      "walletLabel",
      "address",
      "transactions",
      "currency",
    ]),
    ...mapGetters("wallet", ["valueInCurrency", "normalizedAvailableBalance"]),
  },
};
</script>
<style scoped lang="scss">
hr {
  border-top: 0.0625em solid #c4c4c4;
}

.card-user {
  h5 {
    font-weight: 500;
    font-size: 0.875em;
    line-height: 1.5em;
    color: #666666;

    small {
      font-weight: normal;
      font-size: 0.75em;
      line-height: 1.5em;
      color: #979797;
    }
  }

  .author {
    .title {
      font-weight: 500;
      font-size: 1.125em;
      line-height: 2.25em;
      color: #666666;
      display: flex;
      flex-direction: column;

      small {
        font-weight: normal;
        font-size: 0.75em;
        line-height: 2.25em;
        color: #979797;
      }
    }
  }
}
</style>
