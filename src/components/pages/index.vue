<template>
  <section>
    <div class="wrapper">
      <card v-for="item in items" :key="item.id" :item="item" :confirm="confirm"/>
    </div>
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
      isLoading: true,
      myId: 0,
      count: 0,
      canVote: true
    }
  },
  created() {
    axios.get("/api/products").then((res) => {
      this.isLoading = false

      let items = []
      for(let item of res.data) {
        if(item.author == "konojunya") {
          this.myId = item.id
        } else {
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
    confirm() {
      this.$dialog.confirm({
        title: 'HEW投票システム',
        message: 'HEW投票システムが気に入りましたか？よければ投票してください！',
        cancelText: 'やめておく',
        confirmText: 'いいと思う！',
        onConfirm: () => this.vote()
      })
    },
    vote() {
      axios.post(`/api/products/${this.myId}/vote`).then((res) => {
        this.$toast.open({
          message: "ありがとうございます！",
          type: 'is-danger'
        })
      })
    }
  }
}
</script>

<style lang="css">
body, html {
  background-color: #fafafa;
}
.wrapper {
  height: 95vh;
  margin: 0 20px;
  display: flex;
  justify-content: flex-start;
  align-items: center;
  overflow-x: scroll;
}
@media (max-width: 640px) {
  .wrapper {
    height: auto;
    margin: 5px;
    justify-content: space-around;
    align-items: center;
    flex-wrap: wrap;
    overflow-x: inherit;
  }
}
</style>
