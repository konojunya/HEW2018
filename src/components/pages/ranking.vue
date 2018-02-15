<template>
  <section class="container">
    <ranking-item :item="item" :key="item.id" :index="index" :max="max" v-for="(item, index) in items"></ranking-item>
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
      max: 0,
      isLoading: true
    }
  },
  created() {
    axios.get("/api/products").then((res) => {
      this.isLoading = false
      if(res.data){
        this.items = this.rankingSort(res.data).slice(0, 5)
      }
    })
  },
  methods: {
    rankingSort(items) {
      return items.sort((a, b) => {
        let x = a["votes"]
        let y = b["votes"]
        if(x > y) {
          if(this.max < x) {
            this.max = x
          }
          return -1
        }
        if(x < y) {
          if(this.max < y) {
            this.max = y
          }
          return 1
        }
        return 0
      })
    }
  }
}
</script>

<style lang="scss">
body, html {
  background-color: #1A237E;
}
.container {
  width: 80%;
  height: 100vh;
  display: flex;
  align-items: center;
  flex-direction: column;
  justify-content: center;
}
</style>
