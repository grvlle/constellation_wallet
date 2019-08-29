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
            <p-button type="info" block style="margin-top: 28px; overflow: visible;">
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
    return {
      user: {
        company: "0x161D1B0bca85e29dF546AFba1360eEc6Ab4aA7Ee",
        username: "12",
        city: "2523523",
        firstName: "325",
        lastName: "21",
        address: "53",
        country: "12512351235"
      }
    };
  },
  methods: {
    updateProfile() {
      alert("Your data: " + JSON.stringify(this.user));
    }
  }
};
</script>
<style>
</style>
