<template>
  <div>
    <label v-if="label" class="control-label">
      {{label}}
    </label>
    <div class="input-group">
      <input
        :type="type"
        class="form-control"
        :disabled="false"
        :placeholder="placeholder"
        aria-describedby="basic-addon2"
        :id="id"
        :value="value"
        @input="checkPassword($event.target.value)"
      >
      <div class="input-group-append">
        <p-button class="btn" @click.native="showPassword()" type="danger">
          <i v-bind:class="btnText" />
        </p-button>
      </div>
    </div>
    
    <div style="height: 30px; margin-top: -15px;" v-if="storepass_validations">             
      <p v-if="!this.$store.state.validators.storepass.contains_eight_characters" class="validate"> 8 Characters Long, </p> 
      <p v-if="!this.$store.state.validators.storepass.contains_number" class="validate"> Number,</p> 
      <p v-if="!this.$store.state.validators.storepass.contains_uppercase" class="validate"> Uppercase, </p> 
      <p v-if="!this.$store.state.validators.storepass.contains_special_character" class="validate"> Special Character </p>     
    </div>
    <div style="height: 30px; margin-top: -15px;" v-if="!storepass_validations"/> 

    <div style="height: 30px; margin-top: -15px;" v-if="keypass_validations">             
      <p v-if="!this.$store.state.validators.keypass.contains_eight_characters" class="validate"> 8 Characters Long, </p> 
      <p v-if="!this.$store.state.validators.keypass.contains_number" class="validate"> Number,</p> 
      <p v-if="!this.$store.state.validators.keypass.contains_uppercase" class="validate"> Uppercase, </p> 
      <p v-if="!this.$store.state.validators.keypass.contains_special_character" class="validate"> Special Character </p>     
    </div>
    <div style="height: 30px; margin-top: -15px;" v-if="!keypass_validations"/>

  </div>
</template>
 
<script>
  export default {
    name: "password-input",
    props: {
      password_type: String,
      label: String,
      placeholder: String,
      id: String,
      value: String
    },
    data: () => ({
      password_length: 0,
      contains_eight_characters: false,
      contains_number: false,
      contains_uppercase: false,
      contains_special_character: false,
      type: "password",
      btnText: "fa fa-eye"
    }),
    computed: {
      storepass_validations: function () {
        if (this.password_type == 'storepass' && 
          !this.$store.state.validators.duplicate && 
          !this.$store.state.app.login &&
          !this.$store.state.validators.storepass.valid_password
        ) {
          return true;
        } else {
          return false;
        }
      },
      keypass_validations: function () {
        if (this.password_type == 'keypass' && 
          !this.$store.state.validators.duplicate && 
          !this.$store.state.app.login &&
          !this.$store.state.validators.keypass.valid_password) {
          return true;
        } else {
          return false;
        }
      }
    },
    methods: {
      checkPassword: function(value) {
        this.$emit("input", value);

        this.$store.state.validators.target = value
        const format = /[ !@#$%^&*()_+\-=[\]{};':"\\|,.<>/?]/;

        var validator = null;
        if (this.$props.password_type == 'storepass') {
          validator = this.$store.state.validators.storepass;
        } else if (this.$props.password_type == 'keypass') {
          validator = this.$store.state.validators.keypass;
        }

        validator.password_length = value.length;
        if (validator.password_length >= 8) {
          validator.contains_eight_characters = true;
        } else {
          validator.contains_eight_characters = false;
        }

        validator.contains_number = /\d/.test(value);
        validator.contains_uppercase = /[A-Z]/.test(value);
        validator.contains_special_character = format.test(value);

        if (validator.contains_eight_characters === true &&
          validator.contains_special_character === true &&
          validator.contains_uppercase === true &&
          validator.contains_number === true) {
          validator.valid_password = true;			
        } else {
          validator.valid_password = false;
        }
      },
      showPassword: function() {
        if (this.type === "password") {
          this.type = "text";
          this.btnText = "fa fa-eye-slash";
        } else {
          this.type = "password";
          this.btnText = "fa fa-eye";
        }
      }
    }
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