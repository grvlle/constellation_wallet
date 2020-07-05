<template>
  <card>
    <div class="container">
      <div class="row">
        <div class="col-md-3"></div>
        <div class="col-md-6">
          ADDRESS: {{contact.address}}
          <br />
          NAME: {{contact.name}}
          <br />
          TAG: {{contact.tag}}
          <br />
          DESCRIPTION: {{contact.description}}
        </div>
        <div class="col-md-3">
          <button type="button" class="btn btn-primary mr-2" @click="editContact(contact)">
            <i class="fa fa-edit"></i>
              <button type="button" class="btn btn-danger mr-2" @click="deleteContact(contact)">
                <i class="fa fa-trash"></i>
                Delete                
              </button>
          <button type="submit" class="btn btn-secondary mr-2" @click="cancel">Cancel</button>
        </div>
      </div>
    </div>
  </card>
</template>

<script>
import Swal from "sweetalert2/dist/sweetalert2";

export default {
  name: "contact-details",
  components: {},
  computed: {
    contact: function() {
      return this.$store.getters["addressBook/byId"](this.$route.params.id);
    }
  },
  methods: {
    cancel: function() {
      this.$router.go(-1);
    },
    editContact: function(contact) {
      this.$router.push({
        name: "new-edit contact",
        params: { id: contact.id }
      });
    },
    deleteContact: function(contact) {
      const swalPopup = Swal.mixin({
        customClass: {
          container: this.darkMode ? "theme--dark" : "theme--light"
        }
      });
      swalPopup
        .fire({
          title: "Are you sure?",
          html:
            "You are about delete the contact with name <b>" +
            this.contact.name +
            "</b> from the address book.",
          showCancelButton: true,
          confirmButtonText: "Yes, delete contact"
        })
        .then(result => {
          if (result.value) {
            window.backend.WalletApplication.DeleteContact(contact.id).then(
              deleted => {
                if (deleted) {
                  this.$store.commit({
                    type: "addressBook/deleteContact",
                    contact
                  });
                  this.$router.push({
                    name: "address book"
                  });
                }
              }
            );
          }
        });
    }
  }
};
</script>