<template>
  <section>
    <div class="ranking">
      <router-link class="ranking-text" to="/ranking">ランキングをみてみる</router-link>
    </div>
    <div class="wrapper">
      <card v-for="item in items" :key="item.id" :item="item" :voteme="voteme" />
    </div>
    <b-loading :active.sync="isLoading"></b-loading>
  </section>
</template>

<script>
import card from '../modules/card.vue'
import axios from 'axios'

const MY_ID = 1

export default {
  components: {
    card
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

      let items = []
      for(let item of res.data) {
        if(item.id != MY_ID) {
          items.push(item)
        }
      }
      this.items = this.shuffle(items)
    })
  },
  methods: {
    shuffle(array) {
      for(let i = array.length - 1; i > 0; i--) {
        const r = Math.floor(Math.random() * (i + 1))
        const tmp = array[i]
        array[i] = array[r]
        array[r] = tmp
      }
      return array
    },
    voteme() {
      axios.post(`/api/products/${MY_ID}/vote`).then((res) => {
        this.$toast.open({
          message: `ありがとうございます！`,
          type: 'is-success'
        })
      })
    }
  }
}
</script>

<style lang="scss">
section {
  padding-top: 2rem;
}
.wrapper {
  height: 80vh;
  margin: 0 20px;
  display: flex;
  justify-content: flex-start;
  align-items: center;
  overflow-x: scroll;
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
    height: auto;
    padding-bottom: 20px;
    margin: 0 5px;
    justify-content: space-around;
    align-items: center;
    flex-wrap: wrap;
    overflow-x: inherit;
  }
}
</style>
