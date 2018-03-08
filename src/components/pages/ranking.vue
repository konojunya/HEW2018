<template>
  <section>
    <div class="ranking">
      <router-link class="ranking-text" to="/">一覧に戻る</router-link>
    </div>
    <div class="wrapper" v-if="items.length >= 1">
      <ranking-item :item="item" :key="item.id" :index="index" :max="max" v-for="(item, index) in items"></ranking-item>
      <b-loading :active.sync="isLoading"></b-loading>
    </div>
    <div class="noresult" v-else>
      <h1>まだランキングが集計できません。</h1>
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
    axios.get("/api/ranking").then((res) => {
      this.isLoading = false
      if(res.data) {
        for(let item of res.data) {
          if(item.votes > this.max) {
            this.max = item.votes
          }
        }
        this.items = res.data
      }
    })
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
.noresult {
  h1 {
    color: white;
    text-align: center;
    font-weight: 600;
  }
}
@media (max-width: 640px) {
  .wrapper {
    width: 90%;
    height: auto;
  } 
}
</style>
