<template>
  <div class="form-group">
    <label v-if="label" class="control-label">
      {{ label }}
    </label>
    <div class="input-group" style="margin-bottom: 0.125em">
      <input
        :type="type"
        class="form-control"
        :disabled="false"
        :placeholder="placeholder"
        aria-describedby="basic-addon2"
        :value="value"
        @input="checkPassword($event.target.value)"
      />
      <span class="input-group-append">
        <p-button
          tabIndex="-1"
          class="btn"
          @click.native="showPassword()"
          type="default"
        >
          <i v-bind:class="btnText" />
        </p-button>
      </span>
    </div>
    <div class="validate text-danger error-message" v-if="validate">
      <p v-bind:class="{ resolved: validate && contains_eight_characters }">
        8 characters,
      </p>
      <p v-bind:class="{ resolved: validate && contains_number }">
        &nbsp;&nbsp;number,
      </p>
      <p v-bind:class="{ resolved: validate && contains_uppercase }">
        &nbsp;&nbsp;uppercase,
      </p>
      <p v-bind:class="{ resolved: validate && contains_special_character }">
        &nbsp;special character
      </p>
    </div>
    <div class="validate" v-else />
  </div>
</template>

<script>
export default {
  name: "password-input",
  props: {
    value: String,
    label: String,
    placeholder: String,
    validate: Boolean,
  },
  data: () => ({
    valid_password: false,
    password_length: 0,
    contains_eight_characters: false,
    contains_number: false,
    contains_uppercase: false,
    contains_special_character: false,
    type: "password",
    btnText: "fa fa-eye",
  }),
  methods: {
    checkPassword: function(value) {
      this.$emit("input", value);
      const format = /[ !@#$%^&*()_+\-=[\]{};':"\\|,.<>/?]/;

      this.password_length = value.length;
      if (this.password_length > 7) {
        this.contains_eight_characters = true;
      } else {
        this.contains_eight_characters = false;
      }

      this.contains_number = /\d/.test(value);
      this.contains_uppercase = /[A-Z]/.test(value);
      this.contains_special_character = format.test(value);

      if (
        this.contains_eight_characters === true &&
        this.contains_special_character === true &&
        this.contains_uppercase === true &&
        this.contains_number === true
      ) {
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
    },
  },
};
</script>

<style lang="scss" scoped>
.form-control {
  background: #f9f7f7 !important;
  color: #666666;
  height: 36px;
}
.validate {
  height: 0.625em;
  display: flex;
}
.error-message {
  p {
    font-family: Poppins;
    font-style: normal;
    font-weight: normal;
    font-size: 10px;
    line-height: 15px;
    letter-spacing: 0.05em;
    color: #eb5757;
  }

  .resolved {
    color: #219653;
  }
}
.input-group-append .btn {
  border: 1px solid #666 !important;
  color: white;
  background-color: #666;
  display: inline-block;
  font-family: Poppins;
  width: 36px;
  font-style: normal;
  font-weight: 500;
  font-size: 10px;
  line-height: 15px;
  letter-spacing: 0.1em;
}
</style>
