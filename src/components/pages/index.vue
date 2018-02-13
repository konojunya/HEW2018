<template>
  <section>
    <card v-for="item in items" :key="item.id" :thumbnail="item.thumbnail" :title="item.title" />
  </section>
</template>

<script>
import card from '../modules/card.vue'

export default {
  name: "index",
  components: {
    card
  },
  created() {
    this.items = [
      { id: 1, title: "Dinner Match", thumbnail: "/image/presen-board.png" },
      { id: 2, title: "Dinner Match", thumbnail: "/image/presen-board.png" },
      { id: 3, title: "Dinner Match", thumbnail: "/image/presen-board.png" }
    ]
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
