<template>
  <div class="input-group">
    <input
      type="text"
      class="form-control"
      :disabled="true"
      aria-describedby="basic-addon2"
      :value="value"
    />
    <span class="input-group-append">
      <div v-if="action == 'SelectFile'">
        <input
          type="file"
          name="file"
          id="file"
          class="inputfile"
          @change="selectFile"
          accept=".json"
        />
        <label for="file">BROWSE</label>
      </div>

      <p-button
        v-if="action == 'SelectFile2'"
        tabIndex="-1"
        @click.native="importKey"
        class="btn"
        type="default"
      >
        <span style="display: block;">
          BROWSE
        </span>
      </p-button>
      <p-button
        v-if="action == 'SelectSaveFile'"
        tabIndex="-1"
        @click.native="SelectDirToStoreKey"
        class="btn"
        type="default"
      >
        <span style="display: block;">
          BROWSE
        </span>
      </p-button>
    </span>
  </div>
</template>

<script>
export default {
  props: {
    value: String,
    placeholder: String,
    action: String,
  },

  methods: {
    selectFile: function(e) {
      this.$emit("file", e.target.files[0]);
    },
    importKey: function() {
      window.backend.WalletApplication.ImportKey().then((result) => {
        if (result) {
          this.$store.commit("wallet/setKeystorePath", result);
        }
      });
    },
    SelectDirToStoreKey: function() {
      window.backend.WalletApplication.SelectDirToStoreKey().then((result) => {
        this.$store.commit("wallet/setKeystorePath", result);
      });
    },
  },
};
</script>

<style scoped lang="scss">
.form-control {
  background: #f9f7f7 !important;
  height: 2.25rem;
}
.inputfile {
  width: 0.1px;
  height: 0.1px;
  opacity: 0;
  overflow: hidden;
  position: absolute;
  z-index: -1;
}
.inputfile + label {
  cursor: pointer;
  padding: 0.5625rem 0.75rem;
  height: 2.25rem;
  border-radius: 0 0.25rem 0.25rem 0;
  font-size: 0.75em;
  border: 1px solid #666;
  color: white;
  background-color: #666;
  display: inline-block;
  font-family: Poppins;
  font-style: normal;
  font-weight: 500;
  font-size: 0.625rem;
  line-height: 0.9375rem;
  letter-spacing: 0.1em;
}
.inputfile + label:hover {
  background-color: #403d39;
}
</style>
