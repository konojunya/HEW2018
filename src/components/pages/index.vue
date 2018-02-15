<template>
  <section>
    <card v-for="item in items" :key="item.id" :item="item"/>
    <b-loading :active.sync="isLoading"></b-loading>
  </section>
</template>

<script>
import card from '../modules/card.vue'
import axios from 'axios'

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
      res.data.map((item) => {
        console.log(item)
      })
      this.items = res.data
    })
  }
}
</script>

<style lang="css">
section {
  display: flex;
  padding: 5px 20px;
  justify-content: space-between;
  flex-wrap: wrap;
}
</style>
