<template>
  <div class="card">
    <div v-if="this.$store.state.addressBook.addressBook.length">
      <div class="card-header">
        Address Book:
        <div class="input-group">
          <input
            type="text"
            class="form-control"
            placeholder="Search a contact..."
            v-model="searchFilter"
          />
        </div>
      </div>
      <div class="card-body">
        <div class="table-full-width table-responsive">
          <table class="table">
            <tbody>
              <tr
                v-for="contact in addressBookPage"
                :key="contact.id"
                @click="selectAddress(contact)"
                style="cursor: pointer"
              >
                <td>
                  <p class="card-text" style="margin-bottom: 0">
                    {{contact.name}}
                    <span class="badge badge-success float-right">{{contact.tag}}</span>
                  </p>
                  <small class="text-muted">{{contact.address}}</small>
                  <i v-if="contact.address == value" class="fa fa-check float-right"></i>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        <div class="d-flex flex-row-reverse">
          <pagination :dataset="filteredAddressBook" :pageSize="3" v-model="addressBookPage" />
        </div>
      </div>
    </div>
    <div class="text-center" v-else>
      <p
        class="card-text text-muted font-italic mt-3 mb-3"
      >Empty address book. Please add new contacts first.</p>
    </div>
  </div>
</template>

<script>
import Pagination from "../components/Pagination";

export default {
  name: "address-book-search",
  components: {
    Pagination
  },
  props: {
    value: String
  },
  data() {
    return {
      searchFilter: "",
      addressBookPage: []
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
    selectAddress(contact) {
      this.$emit("input", contact.address);
    }
  }
};
</script>

<style scoped lang="scss">
tr:hover {
  opacity: 0.8;
  font-weight: bold;
}
</style>