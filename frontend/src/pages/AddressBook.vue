<template>
  <div class="container">
    <div class="row">
      <div class="col-md-3">
        <card>
          <div class="input-group mb-3">
            <input
              type="text"
              class="form-control"
              placeholder="Search a contact..."
              v-model="searchFilter"
            />
          </div>
          <div class="list-group list-group-flush text-center">
            <div
              class="list-group-item list-group-item-action"
              v-for="contact in filteredAddressBook"
              :key="contact.id"
            >{{contact.name}}</div>
          </div>
        </card>
      </div>
      <div class="col-md-9">
        <router-view></router-view>
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
    this.$router.push({
      name: "new-edit contact",
      params: { id: "" }
    });

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
  }
};
</script>