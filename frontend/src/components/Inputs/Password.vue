<template>
  <div class="form-group">
    <label v-if="label" class="control-label">
      {{label}}
    </label>
    <div class="input-group" style="margin-bottom: 2px">
      <input
        :type="type"
        class="form-control"
        :disabled="false"
        :placeholder="placeholder"
        aria-describedby="basic-addon2"
        :value="value"
        @input="checkPassword($event.target.value)" />
      <span class="input-group-btn">
        <p-button class="btn" @click.native="showPassword()" type="default">
          <i v-bind:class="btnText" />
        </p-button>
      </span>
    </div>
    <div class="validate" v-if="!valid_password && !this.$store.state.app.login">             
      <p v-if="!this.contains_eight_characters"> 8 Characters Long, </p> 
      <p v-if="!this.contains_number"> Number,</p> 
      <p v-if="!this.contains_uppercase"> Uppercase, </p> 
      <p v-if="!this.contains_special_character"> Special Character </p>     
    </div>
    <div class="validate" v-else/> 

  </div>
</template>
 
<script>
  export default {
    name: "password-input",
    props: {
      value: String,
      label: String,
      placeholder: String
    },
    data: () => ({
      valid_password: false,
      password_length: 0,
      contains_eight_characters: false,
      contains_number: false,
      contains_uppercase: false,
      contains_special_character: false,
      type: "password",
      btnText: "fa fa-eye"
    }),
    methods: {
      checkPassword: function(value) {
        this.$emit("input", value);
        const format = /[ !@#$%^&*()_+\-=[\]{};':"\\|,.<>/?]/;

        this.password_length = value.length;
        if (this.password_length >= 8) {
          this.contains_eight_characters = true;
        } else {
          this.contains_eight_characters = false;
        }

        this.contains_number = /\d/.test(value);
        this.contains_uppercase = /[A-Z]/.test(value);
        this.contains_special_character = format.test(value);

        if (this.contains_eight_characters === true &&
          this.contains_special_character === true &&
          this.contains_uppercase === true &&
          this.contains_number === true) {
          this.valid_password = true;
          this.$emit("valid");		
        } else {
          this.valid_password = false;
          this.$emit("invalid");	
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
.validate {
  height: 10px;
  display: flex;
}
.validate > p {
  /*flex: 1;*/
  font-size: 10px;
  color: firebrick;
  margin-top: 0px;
  margin-right: 2px;
}
</style>