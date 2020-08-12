<template>
  <ul v-if="pageCount > 1" class="pagination justify-content-center">
    <li class="page-item" :class="pageNumber == 0 ? 'disabled' : ''">
      <a class="page-link" style="cursor: pointer;" @click="prevPage">
        <i class="fa fa-angle-left"></i>
      </a>
    </li>
    <li
      class="page-item"
      :class="page == pageNumber + 1 ? 'active' : ''"
      v-for="page in pageCount"
      :key="page"
    >
      <a class="page-link" style="cursor: pointer;" @click="gotoPage(page)">{{page}}</a>
    </li>
    <li class="page-item" :class="pageNumber >= pageCount - 1 ? 'disabled' : ''">
      <a class="page-link" style="cursor: pointer;" @click="nextPage">
        <i class="fa fa-angle-right"></i>
      </a>
    </li>
  </ul>
</template>

<script>
export default {
  name: "pagination",
  props: {
    dataset: Array,
    pageSize: Number,
    value: Array
  },
  data() {
    return {
      pageNumber: 0
    };
  },
  computed: {
    pageCount() {
      let l = this.dataset.length,
        s = this.pageSize;
      return Math.ceil(l / s);
    },
    paginatedData() {
      const start = this.pageNumber * this.pageSize,
        end = start + this.pageSize,
        pageData = this.dataset.slice(start, end);
      return pageData;
    }
  },
  watch: {
    paginatedData: {
      handler() {
        this.$emit("input", this.paginatedData);
      },
      immediate: true
    }
  },
  methods: {
    nextPage() {
      this.pageNumber++;
    },
    prevPage() {
      this.pageNumber--;
    },
    gotoPage(page) {
      this.pageNumber = page - 1;
    }
  }
};
</script>

<style scoped lang="scss">
.disabled .page-link {
  opacity: 0.7;
}
</style>