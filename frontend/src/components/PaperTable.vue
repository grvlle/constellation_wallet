<template>
  <table class="table" :class="tableClass">
    <thead>
    <slot name="columns">
      <th v-for="column in columns" :key="column">{{column}}</th>
    </slot>
    </thead>
    <tbody>
    <tr v-for="(item, index) in data" :key="index">
      
      <slot :row="item">
        <td v-for="(column, index) in columns" :key="index">
          <!-- {{ item[column.toLowerCase()] }} -->
          <p class="description" style="font-size: 1rem;" v-if="index === 0">{{ item.amount | truncate}}</p>
          <p class="description" style="font-size: 1rem;" v-if="index === 1">{{ item.address | truncate}}</p>
          <p class="description" style="font-size: 1rem;" v-if="index === 2">{{ item.fee | truncate}}</p>
          <p class="description" style="font-size: 1rem;" v-if="index === 3">{{ item.txhash | truncate}}</p>
          <p class="description" style="font-size: 1rem;" v-if="index === 4">{{ item.date | truncate}}</p>


        
          
          
        </td>
      </slot>
      
    </tr>
    </tbody>
  </table>
  
  
</template>
<script>
export default {
  name: 'paper-table',
  props: {
    columns: Array,
    data: Array,
    type: {
      type: String, // striped | hover
      default: "striped"
    },
    title: {
      type: String,
      default: ""
    },
    subTitle: {
      type: String,
      default: ""
    }
  },
  computed: {
    tableClass() {
      return `table-${this.type}`;
    }
  //   truncString: function(){
  //     if(this.username.length > 5) {
  //        return this.username.slice(0,4);
  //     }
  //     return this.username;
  //  }
  },
  filters: {
    truncate: function(value) {
      if (value.length > 30) {
        value = value.substring(0, 27) + '...';
      }
      return value

    }
  },
  methods: {
    hasValue(item, column) {
      return item[column.toLowerCase()] !== "undefined";
    },
    itemValue(item, column) {
      return item[column.toLowerCase()];
    }
  }
};
</script>

<style scoped lang="scss">
</style>