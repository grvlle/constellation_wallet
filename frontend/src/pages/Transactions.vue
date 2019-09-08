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
                                <fg-input type="text" label="Submit the amount of $DAG you wish to send" :disabled="false" placeholder="0">
                                </fg-input>
                            </div>
                            <div class="col-md-1">
                                <i class="fa fa-chevron-circle-right" style="color: #6DECBB; font-size: 40px; padding: 28px;"></i>
                            </div>
                            <div class="col-md-4">
                                <fg-input type="text" label="Wallet Address of Recipient" :disabled="false" placeholder="Enter Recipients Wallet Address">
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
                    <paper-table type="hover" :title="table2.title" :sub-title="table2.subTitle" :data="table2.data" :columns="table2.columns">
    
                    </paper-table>
                </div>
            </card>
        </div>
    
    </div>
</template>

<script>
import { PaperTable } from "@/components";

const tableColumns = ["Id", "Amount", "Address", "Date"];
let tableData = [];

// Temp dummy numbers for transaction
let amount = 123;
let address = "0x161D1B0bca85e29dF546AFba1360eEc6Ab4aA7Ee";

import TxSentNotification from './Notifications/TxSent';
import Transactions from '../../../tx.json'

//this.table2.data = []

export default {
    components: {
        PaperTable
    },
    methods: {
        tx: function() {
            var self = this
            this.tokenSendInit(
                this.$swal({
                    title: 'You are about to send ' + self.txAmount + ' $DAG tokens to ' + self.txAddress,
                    text: "Are you sure you wish to send this transaction?",
                    type: 'warning',
                    showCancelButton: true,
                    confirmButtonColor: '#3085D6',
                    cancelButtonColor: '#EA896E',
                    confirmButtonText: 'Yes, send it!'
                }).then((result) => {
                    if (result.value) {
                        window.backend.sendTransaction(amount, address).then(transaction => {
                            self.txFull = transaction;
                            self.txAmount = transaction.Amount;
                            self.txAddress = transaction.Address;
                        })
                        this.$swal({
                                title: 'Success!',
                                text: 'You have sent ' + self.txAmount + ' $DAG tokens to address ' + self.txAddress + '.',
                                type: 'success'
                            }),
                            this.$notify({
                                component: TxSentNotification,
                                icon: "ti-check",
                                horizontalAlign: "right",
                                verticalAlign: "top",
                                type: "success"
                            })
                    }
                })
            );
        }
    },
    data() {
        return {
            txFull: {},
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
                data: [Transactions]
            }
        }
    }
}
</script>

<style>

</style>
