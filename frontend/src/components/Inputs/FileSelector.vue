<template>
  <div class="input-group">
    <input type="text"
      class="form-control"
      :disabled="true"
      aria-describedby="basic-addon2"
      :value=value />
    <span class="input-group-append">
      <div v-if="action == 'SelectFile'">
        <input type="file" name="file" id="file" class="inputfile" @change="selectFile" accept=".p12,.json" />
        <label for="file">BROWSE</label>
      </div>

      <p-button v-if="action == 'SelectFile2'"
        tabIndex="-1" 
        @click.native="importKey" 
        class="btn"
        type="default">
        <span style="display: block;">
          BROWSE
        </span>
      </p-button>
      <p-button v-if="action == 'SelectSaveFile'"
        tabIndex="-1" 
        @click.native="SelectDirToStoreKey" 
        class="btn"
        type="default">
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
    action: String
  },

  methods: {
    selectFile: function(e) {
      this.$emit('file', e.target.files[0])
    },
    importKey: function() {
      window.backend.WalletApplication.ImportKey().then(
        result => {
          if (result) {
            this.$store.commit('wallet/setKeystorePath', result);
          }
        }
      );
    },
    SelectDirToStoreKey: function() {
      window.backend.WalletApplication.SelectDirToStoreKey().then(
        result => {
          this.$store.commit('wallet/setKeystorePath', result);
        }
      );
    }
  }
}
</script>

<style scoped lang="scss">
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
  padding: 9px 12px;
  height: 36px;
  border-radius: 0 0.25rem 0.25rem 0;
  font-size: 0.75em;
  font-weight: 600;
  border: 1px solid #DDDDDD;
  color: white;
  background-color: #66615B;
  display: inline-block;
}
.inputfile + label:hover {
  background-color: #403D39;
}
</style>
