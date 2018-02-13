<template>
  <section>
    <card v-for="item in items" :key="item.id" :thumbnail="item.thumbnail" :title="item.title" />
  </section>
</template>

<script>
import card from '../modules/card.vue'
import axios from 'axios'

export default {
  name: "index",
  components: {
    card
  },
  data() {
    return {
      items: []
    }
  },
  created() {
    axios.get("/api/products").then((res) => {
      this.items = res.data
    })
  }
}
</script>

<style lang="css">
body, html {
  width: 100%;
  height: 100%;
  background-color: #fafafa;
}
* {
  padding: 0;
  margin: 0;
  box-sizing: border-box;
}
section {
  display: flex;
  padding: 5px 20px;
  justify-content: space-between;
  flex-wrap: wrap;
}
</style>
