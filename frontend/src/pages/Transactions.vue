<template>
    <div class="row">
      <div class="col-12">
        <card :title="table1.title" :subTitle="table1.subTitle">
          <transaction-form>
          </transaction-form>
          <!-- <div slot="raw-content" class="table-responsive">
            <paper-table :data="table1.data" :columns="table1.columns">

            </paper-table>
          </div> -->
          <br><br>
        </card>
      </div>

      <div class="col-12">
        <card class="card" :title="table2.title" :subTitle="table2.subTitle">
          <div class="table-full-width table-responsive">
            <paper-table type="hover" :title="table2.title" :sub-title="table2.subTitle" :data="table2.data"
                         :columns="table2.columns">

            </paper-table>
          </div>
        </card>
      </div>

    </div>
</template>
<script>
import { PaperTable } from "@/components";
const tableColumns = ["Id", "Amount", "Hash", "Block", "Date"];
const tableData = [
  {
    id: 5,
    date: "2019-02-13 10:12 AM",
    amount: "$36.738",
    hash: "0000011c3edf76bf3e9c65626ba0e253ea78e8647158aec4fb3",
    block: "0000019437c37db98412d5615d68f9a7610f7befa59582a65cb"
  },
  {
    id: 4,
    date: "2019-02-13 10:12 AM",
    amount: "$23,789",
    hash: "000075fe262c5ed829663813b00a8d3cab6dd144daf89d28d",
    block: "000014c1e9be05fdb49fc9285ec619fa6ae8fcf46bf7e37ff7"
  },
  {
    id: 3,
    date: "2019-02-13 10:12 AM",
    amount: "$56,142",
    hash: "000019437c37db98412d5615d68f9a7610f7befa59582a65cb",
    block: "000014c1e9be05fdb49fc9285ec619fa6ae8fcf46bf7e37ff7"
  },
  {
    id: 2,
    date: "2019-02-13 10:12 AM",
    amount: "$38,735",
    hash: "000011c3edf76bf3e9c65626ba0e253ea78e8647158aec4fb3",
    block: "000019437c37db98412d5615d68f9a7610f7befa59582a65cb"
  },
  {
    id: 1,
    date: "2019-02-13 10:12 AM",
    amount: "$63,542",
    hash: "000075fe262c5ed829663813b00a8d3cab6dd144daf89d28d",
    block: "000019437c37db98412d5615d68f9a7610f7befa59582a65cb"
  }
];

import TransactionForm from "./UserProfile/TransactionForm.vue";

export default {
  components: {
    PaperTable,
    TransactionForm
  },
  data() {
    return {
      table1: {
        title: "Transactions",
        subTitle: "Submit a $DAG Transaction",
        columns: [...tableColumns],
        data: [...tableData]
      },
      table2: {
        title: "Transaction History",
        subTitle: "Table containing all previous transactions",
        columns: [...tableColumns],
        data: [...tableData]
      }
    };
  }
};
</script>
<style>
</style>
