<template>
  <div class="card">
    <div class="card-image">
      <figure class="img">
        <img :src="item.thumbnail" alt="thumbnail">
      </figure>
    </div>
    <button class="button" @click="vote" :disabled="!$parent.canVote">いいかも！</button>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  props: ["item","confirm"],
  data() {
    return {
      count: 0
    }
  },
  methods: {
    vote() {
      // 連打できないようにする
      if(!this.$parent.canVote) return

      this.$parent.canVote = false
      axios.post(`/api/products/${this.item.id}/vote`).then((res) => {
        this.$parent.count++
        this.$toast.open({
          message: `『${this.item.title}』に投票しました！`,
          type: 'is-success'
        })
        setTimeout(() => {
          this.$parent.canVote = true
          if(this.$parent.count !== 0 && this.$parent.count % 5 === 0) {
            this.confirm()
          }
        }, 1500)
      })
    }
  }
}
</script>


<style lang="scss" scoped>
.card {
  max-width: 30vw;
  min-width: 30vw;
  margin: 10px;
  display: flex;
  flex-direction: column;
}
.img {
  width: 100%;
  img {
    display: block;
    width: 100%;
  }
}
.button {
  margin: 5px 10px 5px auto;
  background-color: #EF5350;
  color: white;
  font-weight: 600;
  border-radius: 5px;
  border: 0;
}
</style>
