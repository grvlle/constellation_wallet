<template>
  <div>
    <div class="form-group" >
      <slot name="label">
        <label v-if="label" class="control-label">
          {{label}}
        </label>
      </slot>
      <password 
        v-model="password"
        :toggle="true"
        :badge=true
        :secureLength=8
        @input="checkPassword(password)" 
      />
    </div>
    <div style="height: 30px; margin-top: -20px;" v-if="!this.$store.state.validators.duplicate && !this.$store.state.app.login && !this.$store.state.validators.storepass.valid_password">             
              <p v-if="!this.$store.state.validators.storepass.contains_eight_characters" class="validate"> 8 Characters Long, </p> 
              <p v-if="!this.$store.state.validators.storepass.contains_number" class="validate"> Number,</p> 
              <p v-if="!this.$store.state.validators.storepass.contains_uppercase" class="validate"> Uppercase, </p> 
              <p v-if="!this.$store.state.validators.storepass.contains_special_character" class="validate"> Special Character </p>     
    </div>
  </div>
</template>
 
<script>
  import Password from 'vue-password-strength-meter'
  export default {
    name: "password-input",
    components: { Password },
    props: {
      label: String
    },
    methods: {
      checkPassword: function() {
        this.$store.state.validators.target = this.password
        this.$store.state.validators.storepass.password_length = this.password.length;
        const format = /[ !@#$%^&*()_+\-=[\]{};':"\\|,.<>/?]/;

        if (this.$store.state.validators.storepass.password_length > 8) {
          this.$store.state.validators.storepass.contains_eight_characters = true;
        } else {
          this.$store.state.validators.storepass.contains_eight_characters = false;
        }

        this.$store.state.validators.storepass.contains_number = /\d/.test(this.password);
        this.$store.state.validators.storepass.contains_uppercase = /[A-Z]/.test(this.password);
        this.$store.state.validators.storepass.contains_special_character = format.test(this.password);

        if (this.$store.state.validators.storepass.contains_eight_characters === true &&
          this.$store.state.validators.storepass.contains_special_character === true &&
          this.$store.state.validators.storepass.contains_uppercase === true &&
          this.$store.state.validators.storepass.contains_number === true) {
          this.$store.state.validators.storepass.valid_password = true;			
        } else {
          this.$store.state.validators.storepass.valid_password = false;
        }
      }
    },
    data: () => ({
      password: null
    })
  }
</script>

<style lang="scss" scoped>
  p.validate {
      font-size: 10px;
      color: firebrick;
      margin-top: 0px;
      margin-right: 2px;
      float: left;
  }
</style>