<template>
  <div class="card bg-light">
    <div class="card-header">
      Address Book:
      <div class="input-group">
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
      </div>
    </div>
    <div class="card-body">
      <div class="list-group list-group-flush">
        <div
          class="list-group-item list-group-item-action"
          style="cursor: pointer;"
          v-for="contact in filteredAddressBook"
          :key="contact.id"
          @click="selectAddress(contact)"
        >
          <p class="card-text" style="margin-bottom: 0">
            {{contact.name}}
            <span class="badge badge-success float-right">{{contact.tag}}</span>
          </p>
          <small class="text-muted">{{contact.address}}</small>
          <i v-if="contact.address == value" class="fa fa-check float-right"></i>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "address-book-search",
  props: {
    value: String
  },
  data() {
    return {
      searchFilter: ""
    };
  },
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
    selectAddress(contact) {
        this.$emit('input', contact.address)
    }
  }
};
</script>

<style>
</style>