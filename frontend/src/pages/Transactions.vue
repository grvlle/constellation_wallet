<template>
    <div class="row">
      <div class="col-12">
        <card :title="table1.title" :subTitle="table1.subTitle">
          <!-- <div slot="raw-content" class="table-responsive">
            <paper-table :data="table1.data" :columns="table1.columns">

            </paper-table>
          </div> -->
          <br><br>
        </card>
      </div>

      <div class="col-12">
        <card class="card-plain">
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
    id: 1,
    date: "2019-02-13 10:12 AM",
    amount: "$36.738",
    hash: "Niger",
    block: "Oud-Turnhout"
  },
  {
    id: 2,
    date: "2019-02-13 10:12 AM",
    amount: "$23,789",
    hash: "Curaçao",
    block: "Sinaai-Waas"
  },
  {
    id: 3,
    date: "2019-02-13 10:12 AM",
    amount: "$56,142",
    hash: "Netherlands",
    block: "Baileux"
  },
  {
    id: 4,
    date: "2019-02-13 10:12 AM",
    amount: "$38,735",
    hash: "Korea, South",
    block: "Overland Park"
  },
  {
    id: 5,
    date: "2019-02-13 10:12 AM",
    amount: "$63,542",
    hash: "Malawi",
    block: "Feldkirchen in Kärnten"
  }
];

export default {
  components: {
    PaperTable
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
        title: "Table on Plain Background",
        subTitle: "Here is a subtitle for this table",
        columns: [...tableColumns],
        data: [...tableData]
      }
    };
  }
};
</script>
<style>
</style>
