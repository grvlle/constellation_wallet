<template>
    <div class="row">
        <div class="col-12">
            <card :title="table1.title" :subTitle="table1.subTitle">
                <!-- <transaction-form> -->
                <div>
                    <form @submit.prevent>
                        <br>
                        <div class="row">
                            <div class="col-md-1"></div>
                            <div class="col-md-4">
                                <fg-input v-model.number="amountSubmitted" type="text" label="Submit the amount of $DAG you wish to send" placeholder="0">
                                </fg-input>
                            </div>
                            <div class="col-md-1">
                                <i class="fa fa-chevron-circle-right" style="color: #6DECBB; font-size: 40px; padding: 28px;"></i>
                            </div>
                            <div class="col-md-4">
                                <fg-input v-model="txAddress" type="text" label="Wallet Address of Recipient" placeholder="Enter Recipients Wallet Address">
                                </fg-input>
                            </div>
                            <div class="col-md-1">
                                <p-button type="info" block @click.native="tx" style="margin-top: 28px; overflow: visible;">
                                    <span style="width: 80px; margin-left: -20px; overflow: hidden; white-space: nowrap; display: block; text-overflow: ellipsis;"><i class="fa fa-paper-plane"></i> SEND</span>
                                </p-button>
                            </div>
                        </div>
                        <div class="clearfix"></div>
                    </form>
                </div>
                <br><br>
            </card>
        </div>
    
        <div class="col-12">
            <card class="card" :title="table2.title" :subTitle="table2.subTitle">
                <div class="table-full-width table-responsive">
                    <table class="table" :class="tableClass">
                        <thead>
                            <slot name="columns">
                                <th v-for="column in table2.columns" :key="column">{{column}}</th>
                            </slot>
                        </thead>
                        <tbody>
                            <tr v-for="(item, index) in pageOfItems" :key="index">
    
                                <slot :row="item">
                                    <td v-for="(column, index) in table2.columns" :key="index">

                                        <!-- {{ item[column.toLowerCase()] }} -->
                                        <p class="description" style="font-size: 15px;" v-if="index === 0">{{ item.amount | truncate}}</p>
                                        <p class="description" style="font-size: 15px;" v-if="index === 1">{{ item.address.toUpperCase() | truncate}}</p>
                                        <p class="description" style="font-size: 15px;" v-if="index === 2">{{ item.fee | truncate}}</p>
                                        <p class="description" style="font-size: 15px;" v-if="index === 3">{{ item.txhash.toUpperCase() | truncate}}</p>
                                        <p class="description" style="font-size: 15px;" v-if="index === 4">{{ item.date | truncate}}</p>
    
                                    </td>
                                </slot>
    
                            </tr>
                        </tbody>
                    </table>
                    <!-- <paper-table type="hover" :title="table2.title" :sub-title="table2.subTitle" :data="this.$store.state.txInfo.txHistory" :columns="table2.columns">
                            
                        </paper-table> -->
                    <center><jw-pagination :items="table2.data" @changePage="onChangePage"></jw-pagination></center>
                </div>
            </card>
    
        </div>
    
    </div>
</template>

<script>
// import { PaperTable } from "@/components";

const tableColumns = ["Amount", "Address", "Fee", "TxHash", "Date"];
let tableData = [];

import TxSentNotification from './Notifications/TxSent';
import Swal from 'sweetalert2'
// import TransactionHistory from '../../../JSONdata/tx.json';

export default {
    components: {
        // PaperTable
    },
    methods: {
        onChangePage(pageOfItems) {
            // update page of items
            this.pageOfItems = pageOfItems;
        },
        tx: function() {
            var self = this
            Swal.fire({
                title: 'You are about to send ' + self.amountSubmitted + ' $DAG tokens to ' + self.txAddress,
                text: "Are you sure you wish to send this transaction?",
                type: 'warning',
                showCancelButton: true,
                confirmButtonColor: '#3085D6',
                cancelButtonColor: '#EA896E',
                confirmButtonText: 'Yes, send it!'
            }).then((result) => {
                if (result.value) {
                    let amount = self.amountSubmitted
                    let address = self.txAddress
                    window.backend.WalletApplication.PrepareTransaction(amount, address)
                    Swal.fire({
                            title: 'Success!',
                            text: 'You have sent ' + self.amountSubmitted + ' $DAG tokens to address ' + self.txAddress + '.',
                            type: 'success'
                        }),
                        self.$notify({
                            component: TxSentNotification,
                            icon: "ti-check",
                            horizontalAlign: "right",
                            verticalAlign: "top",
                            type: "success"
                        })
                }
            });
        }
    },

    data() {
        return {
            amountSubmitted: 0,
            pageOfItems: [],
            txAmount: '',
            txAddress: '',
            notifications: {
                topCenter: false
            },

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
                data: this.$store.state.txInfo.txHistory
            }
        }
    },
    filters: {
        truncate: function(value) {
            if (value.length > 30) {
                value = value.substring(0, 27) + '...';
            }
            return value

        }
    },
    props: {
        columns: Array,
        data: Array,
        type: {
            type: String, // striped | hover
            default: "striped"
        },
        title: {
            type: String,
            default: ""
        },
        subTitle: {
            type: String,
            default: ""
        }
    }
}
</script>

<style>

</style>
