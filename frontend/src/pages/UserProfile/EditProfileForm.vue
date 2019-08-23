<template>
  <card class="card">
    <div>
      <form @submit.prevent>
        <div class="row">
          <div class="col-md-12">
            <fg-input type="text"
                      label="Wallet Address"
                      :disabled="true"
                      placeholder="0x161D1B0bca85e29dF546AFba1360eEc6Ab4aA7Ee"
                      v-model="user.company">
            </fg-input>
          </div>
          <!-- <div class="row">
          <div class="col-md-3">

            <fg-input type="text"
                      label="Username"
                      placeholder="Username"
                      v-model="user.username">
            </fg-input>
          </div>
          <div class="col-md-4">
            <fg-input type="email"
                      label="Username"
                      placeholder="Email"
                      v-model="user.email">
            </fg-input>
          </div>
        </div> -->
        </div>

          <div class="row">
          <div class="col-md-12">
            <fg-input type="text"
                      label="Nonce"
                      :disabled="true"
                      placeholder="0"
                      v-model="user.firstName">
            </fg-input>
          </div>
          </div>
          <div class="row">
          <div class="col-md-12">
            <fg-input type="text"
                      label="Total Balance"
                      :disabled="true"
                      placeholder="0"
                      v-model="user.lastName">
            </fg-input>
          </div>
        </div>

        <div class="row">
          <div class="col-md-12">
            <fg-input type="text"
                      label="Delegated"
                      :disabled="true"
                      placeholder="0"
                      v-model="user.address">
            </fg-input>
          </div>
        </div>

        <div class="row">
          <div class="col-md-12">
            <fg-input type="text"
                      label="Deposit"
                      :disabled="true"
                      placeholder="0"
                      v-model="user.city">
            </fg-input>
          </div>
        </div>
          <div class="row">
          <div class="col-md-12">
            <fg-input type="text"
                      label="Available Balance"
                      :disabled="true"
                      placeholder="0"
                      v-model="user.country">
            </fg-input>
          </div>
          </div>
          <!-- <div class="col-md-4">
            <fg-input type="number"
                      label="Postal Code"
                      placeholder="ZIP Code"
                      v-model="user.postalCode">
            </fg-input>
          </div>
        </div>

        <div class="row">
          <div class="col-md-12">
            <div class="form-group">
              <label>About Me</label>
              <textarea rows="5" class="form-control border-input"
                        placeholder="Here can be your description"
                        v-model="user.aboutMe">

              </textarea>
            </div>
          </div>
        </div>
        <div class="text-center">
          <p-button type="info"
                    round
                    @click.native.prevent="updateProfile">
            Update Profile
          </p-button>
        </div> -->
        <div class="clearfix"></div>
      </form>
    </div>
  </card>
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
