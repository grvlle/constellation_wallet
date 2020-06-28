<template>
  <card title="Create a contact" sub-title="Provide your contacts details and press save.">
    <form @submit.prevent="submitContact()">
      <div class="form-group">
        <label for="addressInput">DAG Address</label>
        <input
          type="text"
          class="form-control"
          id="addressInput"
          v-model.trim="address"
          placeholder="Enter address..."
        />
        <div class="validate" v-if="!$v.address.required">
          <p>Field is required</p>
        </div>
        <div
          class="validate"
          v-else-if="!$v.address.minLength || !$v.address.verifyPrefix || !$v.address.maxLength"
        >
          <p>Invalid wallet address.</p>
        </div>
        <div class="validate" v-else />
      </div>
      <div class="form-group">
        <label for="nameInput">Name</label>
        <input
          type="text"
          class="form-control"
          id="nameInput"
          v-model.trim="name"
          placeholder="Enter name..."
        />
        <div class="validate" v-if="!$v.name.required">
          <p>Field is required</p>
        </div>
        <div class="validate" v-else />
      </div>
      <div class="form-group">
        <label for="tagInput">Tag</label>
        <input
          type="text"
          class="form-control"
          id="tagInput"
          v-model.trim="tag"
          placeholder="Enter tag..."
        />
      </div>
      <div class="form-group">
        <label for="descriptionInput">Description</label>
        <textarea
          class="form-control rounded-0"
          id="descriptionInput"
          v-model.trim="description"
          placeholder="Enter description..."
          rows="2"
        ></textarea>
      </div>
      <button type="submit" class="btn btn-primary" :disabled="this.$v.$invalid">Save</button>
    </form>
  </card>
</template>

<script>
const verifyPrefix = value =>
  value.substring(0, 3) === "DAG" || value.substring(0, 3) === "";

import { required, minLength, maxLength } from "vuelidate/lib/validators";
export default {
  name: "new-edit-contact",
  components: {},
  created: function() {
    if (this.$route.params.id != "") {
      let contact = this.$store.state.addressBook.addressBook.find(
        contact => contact.id == this.$route.params.id
      );
      this.address = contact.address;
      this.name = contact.name;
      this.tag = contact.tag;
      this.description = contact.description;
    }
  },
  data: () => ({
    address: "",
    name: "",
    tag: "",
    description: ""
  }),
  validations: {
    address: {
      required,
      minLength: minLength(40),
      maxLength: maxLength(40),
      verifyPrefix
    },
    name: {
      required
    }
  },
  methods: {
    submitContact: function() {
      this.$v.$touch();
      if (!this.$v.$invalid) {
        window.backend.WalletApplication.StoreContact(
          this.address,
          this.name,
          this.tag,
          this.description
        ).then(stored => {
          if (stored) {
            let contact = {
              address: this.address,
              name: this.name,
              tag: this.tag,
              description: this.description
            };
            this.$store.commit({ type: "addressBook/setContact", contact });
          }
        });
      }
    }
  }
};
</script>

<style scoped>
.validate {
  height: 1.25em;
}
</style>