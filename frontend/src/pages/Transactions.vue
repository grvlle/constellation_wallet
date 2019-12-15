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
                                <fg-input v-model.number="txAmountValidation" @change="sendAmount($event.target.value)" type="text" label="Submit the amount of $DAG you wish to send" placeholder="0">
                                </fg-input>
                                <div style="height: 25px;" class="error" v-if="$v.txAmountValidation.inBetween">
    
                                </div>
    
                                <div class="error" v-if="!$v.txAmountValidation.inBetween">
                                    <p class="validate">Invalid amount. Please specify a number between 0.00000001 - 3,711,998,690.</p>
                                </div>
                            </div>
                            <div class="col-md-1">
                                <i class="fa fa-chevron-circle-right" style="color: #6DECBB; font-size: 40px; padding: 28px;"></i>
                            </div>
                            <div class="col-md-4">
                                <fg-input v-model.trim="txAddressValidation" @change="setName($event.target.value)" type="text" label="Wallet Address of Recipient" placeholder="Enter Recipients Wallet Address">
                                </fg-input>
                                <div style="height: 25px;" class="error" v-if="$v.txAddressValidation.minLength && $v.txAddressValidation.maxLength && $v.txAddressValidation.verifyPrefix">
    
                                </div>
    
                                <div class="error" v-if="!$v.txAddressValidation.minLength || !$v.txAddressValidation.verifyPrefix || !$v.txAddressValidation.maxLength">
                                    <p class="validate">Invalid wallet address. Please verify.</p>
                                </div>
    
                            </div>
                            <div class="col-md-1">
                                <p-button type="info" block @click.native="tx" :disabled="submitStatus === 'PENDING'" style="margin-top: 28px; overflow: visible;">
                                    <span style="width: 80px; margin-left: -20px; overflow: hidden; white-space: nowrap; display: block; text-overflow: ellipsis;"><i class="fa fa-paper-plane"></i> SEND</span>
                                </p-button>
    
    
                            </div>
                        </div>
                        <!-- <div class="clearfix"></div> -->
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
                            <slot txAddressValidation="columns">
                                <th v-for="column in table2.columns" :key="column">{{column}}</th>
                            </slot>
                        </thead>
                        <tbody>
                            <tr v-for="(item, index) in this.$store.state.pageOfItems" :key="index">
    
                                <slot :row="item">
                                    <td v-for="(column, index) in table2.columns" :key="index">
    
                                        <!-- {{ item[column.toLowerCase()] }} -->
                                        <p class="description" style="font-size: 15px;" v-if="index === 0"><b>{{ item.amount | dropzero}}</b> $DAG</p>
                                        <p class="description" style="font-size: 15px;" v-if="index === 1">{{ item.address.toUpperCase() | truncate}}</p>
                                        <p class="description" style="font-size: 15px;" v-if="index === 2">{{ item.fee | truncate}}</p>
                                        <a id="txhash" href="#" rel="noopener noreferrer" target="_blank">
                                            <p style="font-size: 15px;" v-if="index === 3">{{ item.txhash.toUpperCase() | truncate}}</p>
                                        </a>
                                        <p class="description" style="font-size: 15px;" v-if="index === 4">{{ item.date | truncate}}</p>
    
                                    </td>
                                </slot>
    
                            </tr>
                        </tbody>
                    </table>
                    <!-- <paper-table type="hover" :title="table2.title" :sub-title="table2.subTitle" :data="this.$store.state.txInfo.txHistory" :columns="table2.columns">
                                                                                            
                                                                                        </paper-table> -->
                    <center>
                        <jw-pagination :items="table2.data" @changePage="onChangePage"></jw-pagination>
                    </center>
                </div>
            </card>
    
        </div>
    
    </div>
</template>

<script>
const tableColumns = ["Amount", "Address", "Fee", "TxHash", "Date"];
let tableData = [];
const verifyPrefix = (value) => value.substring(0, 3) === "DAG" || value.substring(0, 3) === "";

import TxSentNotification from './Notifications/TxSent';
import Swal from 'sweetalert2'
import { required, minLength, maxLength, between } from 'vuelidate/lib/validators'

export default {
    methods: {
        isFloat: function(n) {
            return n === +n && n !== (n | 0);
        },

        isInteger: function(n) {
            return n === +n && n === (n | 0);
        },

        onChangePage(pageOfItems) {
            // update page of items
            this.$store.state.pageOfItems = pageOfItems;
        },
        sendAmount(value) {
            this.txAmountValidation = value
            this.$v.txAmountValidation.$touch()
        },
        setName(value) {
            this.txAddressValidation = value
            this.$v.txAddressValidation.$touch()
        },
        tx: function() {
            var self = this
            self.$v.$touch()
            if (self.$v.$invalid) {
                self.submitStatus = 'ERROR'
            } else {
                // do your submit logic here
                self.submitStatus = 'PENDING'
                setTimeout(() => {
                    self.submitStatus = 'OK'
                }, 500)
            }


            if (self.submitStatus != 'ERROR') {
                Swal.mixin({
                    progressSteps: ['1', '2'],
                }).queue([{
                        title: "Are you sure?",
                        html: 'You are about to send <b>' + self.txAmountValidation + '</b> $DAG tokens to ' + self.txAddressValidation.toUpperCase(),
                        type: 'warning',
                        showCancelButton: true,
                        confirmButtonColor: '#5FD1FB',
                        confirmButtonText: 'Yes, please proceed!'
                    },
                    {
                        title: 'Set a fee to prioritize your transaction.',
                        html: 'This is <b>optional</b>, enter 0 for no fee.',
                        input: 'number',
                        inputValue: 0,
                        confirmButtonText: 'Send transaction',
                        confirmButtonColor: '#6DECBB',
                        showCancelButton: true,
                        inputValidator: (value) => {
                            return new Promise((resolve) => {
                                if (value < 0 || value >= 3711998690) {
                                    resolve('Please enter a value between 0 and 3711998690')
                                } else {
                                    resolve()
                                }
                            })
                        }
                    },
                ]).then((result) => {
                    if (result.value) {
                        let amount = self.txAmountValidation
                        let address = self.txAddressValidation
                        let fee = result.value
                        window.backend.WalletApplication.PrepareTransaction(parseInt(amount), parseInt(fee[1]), address)
                        Swal.fire({
                                title: 'Success!',
                                text: 'You have sent ' + self.txAmountValidation + ' $DAG tokens to address ' + self.txAddressValidation + '.',
                                type: 'success'
                            }),
                            setTimeout(() => {
                                this.$notifications.clear();
                            }, 6000)
                        self.$notify({
                            component: TxSentNotification,
                            icon: "ti-check",
                            horizontalAlign: "right",
                            verticalAlign: "top",
                            type: "success",
                            onClick: () => {
                                this.$notifications.clear();
                            }
                        })
                    }
                });
            }
        }

    },

    data() {
        return {
            txAddressValidation: '',
            txAmountValidation: null,
            submitStatus: null,
            amountSubmitted: null,
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
        dropzero: function(value) {
            if (!value) return ''
            let index
            value = value.toString().split("")
            index = value.length - 8
            value = value.splice(0, index)
            return value.join('')
        },
        truncate: function(value) {
            if (value.length > 30) {
                value = value.substring(0, 27) + '...';
            }
            return value

        }
    },

    validations: {
        txAddressValidation: {
            required,
            minLength: minLength(40),
            maxLength: maxLength(40),
            verifyPrefix,
        },
        txAmountValidation: {
            inBetween: between(0.00000001, 3711998690),
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
txhash a {
    color: blue;
}

txhash a:visited {
    color: blue
}

txhash p {
    font-weight: bold
}

p.validate {
    font-size: 10px;
    color: firebrick;
    margin-top: -5px;
}
</style>
