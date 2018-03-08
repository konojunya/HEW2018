<template>
  <section>
    <div class="ranking">
      <router-link class="ranking-text" to="/">一覧に戻る</router-link>
    </div>
    <div class="wrapper">
      <ranking-item :item="item" :key="item.id" :index="index" :max="max" v-for="(item, index) in items"></ranking-item>
      <b-loading :active.sync="isLoading"></b-loading>
    </div>
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

<style lang="scss" scoped>
section {
  padding-top: 2rem;
}
.wrapper {
  width: 80%;
  margin: 0 auto;
  height: 100vh;
  display: flex;
  align-items: center;
  flex-direction: column;
  justify-content: center;
}
.ranking {
  width: 250px;
  margin-bottom: 2rem;
  padding: 1rem 1.5rem;
  background-color: rgb(210, 12, 84);
  border-top-right-radius: 2.5rem;
  border-bottom-right-radius: 2.5rem;
  border: 1px solid #FFF;
  border-left-width: 0;
  .ranking-text {
    color: #FFF;
    font-weight: 600;
    text-decoration: none;
  }
}
@media (max-width: 640px) {
  .wrapper {
    width: 90%;
    height: auto;
  } 
}
</style>
