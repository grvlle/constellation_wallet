<template>
  <!-- <card class="card"> -->
    <div>
      <form @submit.prevent>
      <br>
        <div class="row">
          <div class="col-md-1"></div>
          <div class="col-md-4">
            <fg-input type="text"
                      label="Submit the amount of $DAG you wish to send"
                      :disabled="false"
                      placeholder="0">
            </fg-input>
          </div>
          <div class="col-md-1">
            <i class="fa fa-chevron-circle-right" style="color: #6DECBB; font-size: 40px; padding: 28px;"></i>
          </div>
          <div class="col-md-4">
            <fg-input type="text"
                      label="Wallet Address of Recipient"
                      :disabled="false"
                      placeholder="Enter Recipients Wallet Address">
            </fg-input>
          </div>
          <div class="col-md-1">
            <p-button type="info" block @click.native="showAlert()" style="margin-top: 28px; overflow: visible;">
            <span style="width: 80px; margin-left: -20px; overflow: hidden; white-space: nowrap; display: block; text-overflow: ellipsis;"><i class="fa fa-paper-plane"></i> SEND</span>
            </p-button>
          </div>
        </div>
        <div class="clearfix"></div>
      </form>
    </div>
  <!-- </card> -->
</template>
<script>
export default {
  data() {
    return {};
  },
  methods: {
    updateProfile() {
      alert("Your data: " + JSON.stringify(this.user));
    },
    showAlert(){
      // Use sweetalret2
      this.$swal({
        title: 'Are you sure you wish to send this transaction?',
        text: "You won't be able to revert this!",
        type: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#3085D6',
        cancelButtonColor: '#EA896E',
        confirmButtonText: 'Yes, send it!'
        }).then((result) => {
        if (result.value) {
          this.$swal(
            'Sent!',
            'Your $DAG transaction has been sent.',
            'success'
          )
        }
      });
    }
  }
};
</script>
<style>
</style>
