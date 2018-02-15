<template>
  <section class="container">
    <div v-for="item in items">
      <ranking-item></ranking-item>
    </div>
    <b-loading :active.sync="isLoading"></b-loading>
  </section>
</template>

<script>
import axios from 'axios'
import rankingItem from '../modules/rankingItem.vue'

export default {
  components: {
    rankingItem
  },
  data() {
    return {
      items: [],
      isLoading: true
    }
  },
  created() {
    axios.get("/api/products").then((res) => {
      this.isLoading = false
      this.items = this.rankingSort(res.data)
    })
  },
  methods: {
    rankingSort(items) {
      return items.sort((a, b) => {
        let x = a["votes"]
        let y = b["votes"]
        if(x > y) return -1
        if(x < y) return 1
        return 0
      })
    }
  }
}
</script>

<style lang="scss">
.container {
  width: 80%;
  margin: 0 auto;
}
</style>
