<template>
  <card class="card">
    <div>
      <form @submit.prevent>
        <div class="row">
          <div class="col-md-12">
            <input
              type="text"
              label="Wallet Address"
              class="form-control"
              :disabled="true"
              placeholder="No DAG Address found."
              v-model="address"
            />
          </div>
        </div>
        <div class="row">
          <div class="col-md-12">
            <input
              type="text"
              label="Available Balance"
              class="form-control"
              :disabled="true"
              placeholder="0"
              v-model="normalizedAvailableBalance"
            />
          </div>
        </div>
        <!--        <div class="row">-->
        <!--          <div class="col-md-12">-->
        <!--            <fg-input-->
        <!--              type="text"-->
        <!--              label="Wallet Version"-->
        <!--              :disabled="true"-->
        <!--              placeholder="0"-->
        <!--              v-model="uiVersion"-->
        <!--            ></fg-input>-->
        <!--          </div>-->
        <!--        </div>-->
        <!--        <div class="row">-->
        <!--          <div class="col-md-12">-->
        <!--            <fg-input-->
        <!--              type="text"-->
        <!--              label="Alias"-->
        <!--              :disabled="true"-->
        <!--              placeholder="0"-->
        <!--              v-model="alias"-->
        <!--            ></fg-input>-->
        <!--          </div>-->
        <!--        </div>-->
        <div class="row">
          <div class="col-md-12">
            <input
              type="text"
              label="Private Key File"
              :disabled="true"
              placeholder="0"
              class="form-control"
              v-model="keystorePath"
            />
          </div>
        </div>
        <div class="clearfix"></div>
      </form>
    </div>
  </card>
</template>

<script>
import { mapState, mapGetters } from "vuex";
export default {
  computed: {
    ...mapState("app", ["version", "uiVersion"]),
    ...mapState("wallet", ["address", "alias", "keystorePath"]),
    ...mapGetters("wallet", ["normalizedAvailableBalance"]),
  },
};
</script>
<style scoped lang="scss"></style>
