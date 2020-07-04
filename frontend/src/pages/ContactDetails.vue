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
          </button>
          <button type="submit" class="btn btn-secondary mr-2" @click="cancel">Cancel</button>
        </div>
      </div>
    </div>
  </card>
</template>

<script>
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
    }
  }
};
</script>