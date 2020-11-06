<template>
  <div class="card">
    <div class="card-body">
      <h4 v-if="this.$route.params.id != ''" class="card-title">Edit contact</h4>
      <h4 v-else class="card-title">Create a contact</h4>
      <p class="card-category mb-3">Provide your contacts details and press save.</p>
      <form>
        <div class="form-group">
          <label for="nameInput">Name</label>
          <input
            type="text"
            class="form-control"
            id="nameInput"
            v-model.trim="name"
            placeholder="Enter name..."
          />
          <div class="validate text-danger" v-if="!$v.name.required">
            <p>Field is required</p>
          </div>
          <div class="validate text-danger" v-else />
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
        <div class="form-group">
          <label for="addressInput">DAG Address</label>
          <input
            type="text"
            class="form-control"
            id="addressInput"
            v-model.trim="address"
            placeholder="Enter address..."
          />
          <div class="validate text-danger" v-if="!$v.address.required">
            <p>Field is required</p>
          </div>
          <div
            class="validate text-danger"
            v-else-if="!$v.address.minLength || !$v.address.verifyPrefix || !$v.address.maxLength"
          >
            <p>Invalid wallet address.</p>
          </div>
          <div class="validate" v-else />
        </div>
      </form>
    </div>
    <div class="card-footer">
      <div style="float: right;">
        <button type="button" class="btn btn-secondary mr-2" @click="cancel">
          <i class="fa fa-times"></i>
          Cancel
        </button>
        <button
          type="button"
          class="btn btn-primary"
          @click="submitContact"
          :disabled="this.$v.$invalid"
        >
          <i class="fa fa-save"></i>Save
        </button>
      </div>
    </div>
  </div>
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
    cancel: function() {
      this.$router.go(-1);
    },
    submitContact: function() {
      this.$v.$touch();
      if (!this.$v.$invalid) {
        if (this.$route.params.id == "") {
          window.backend.WalletApplication.CreateContact(
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
              this.$router.push({
                name: "address book"
              });
            }
          });
        } else {
          window.backend.WalletApplication.UpdateContact(
            this.$route.params.id,
            this.address,
            this.name,
            this.tag,
            this.description
          ).then(stored => {
            if (stored) {
              let contact = {
                id: this.$route.params.id,
                address: this.address,
                name: this.name,
                tag: this.tag,
                description: this.description
              };
              this.$store.commit({ type: "addressBook/setContact", contact });
              this.$router.push({
                name: "address book"
              });
            }
          });
        }
      }
    }
  }
};
</script>

<style scoped lang="scss">
.validate {
  height: 1.25em;
}
</style>