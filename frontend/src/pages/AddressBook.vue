<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <card>
          <div class="input-group mb-3">
            <input
              type="text"
              class="form-control"
              placeholder="Search a contact..."
              v-model="searchFilter"
            />
            <button
              class="btn bg-transparent"
              @click="clearSearch"
              style="margin-left: -40px; z-index: 100; border: none; color:darkgray;"
            >
              <i class="fa fa-times"></i>
            </button>
            <button type="button" class="btn btn-primary ml-2" @click="createContact">
              <i class="fa fa-plus"></i>
            </button>
          </div>
        </card>
      </div>
    </div>
    <div class="row">
      <div class="col-md-4 mb-4" v-for="contact in filteredAddressBook" :key="contact.id">
        <div class="card h-100">
          <div class="card-body">
            <h5 class="card-title mb-2">
              {{contact.name}}
              <span class="badge badge-success float-right">{{contact.tag}}</span>
            </h5>
            <p class="card-text">{{contact.description}}</p>
          </div>
          <div class="card-footer">
            <small class="text-muted">{{contact.address}}</small>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data: () => ({
    searchFilter: ""
  }),
  created: function() {
    window.backend.WalletApplication.GetAddressBook().then(ab => {
      let addressBook;
      try {
        addressBook = JSON.parse(ab);
      } catch (e) {
        addressBook = [];
      }
      this.$store.commit({ type: "addressBook/setAddressBook", addressBook });
    });
  },
  computed: {
    filteredAddressBook: function() {
      if (this.searchFilter == "") {
        return this.$store.state.addressBook.addressBook;
      } else {
        return this.$store.getters["addressBook/search"](this.searchFilter);
      }
    }
  },
  methods: {
    clearSearch: function() {
      this.searchFilter = "";
    },    
    createContact: function() {
      this.$router.push({
        name: "new-edit contact",
        params: { id: "" }
      });
    }
  }
};
</script>