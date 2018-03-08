<template>
  <div class="card">
    <b-modal :active.sync="isImageModalActive" :width="400" scroll="keep">
      <p class="img">
        <img :src="item.thumbnail">
      </p>
      <div class="button-wrapper">
        <button class="button" @click="vote">いいかも！</button>
      </div>
    </b-modal>
    <div class="card-image" @click="isImageModalActive = true">
      <figure class="img">
        <img :src="item.thumbnail" alt="thumbnail">
      </figure>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  props: ["item"],
  data() {
    return {
      isImageModalActive: false
    }
  },
  methods: {
    vote() {
      axios.post(`/api/products/${this.item.id}/vote`).then((res) => {
        this.$toast.open({
          message: `『${this.item.title}』に投票しました！`,
          type: 'is-success'
        })
        isImageModalActive = false
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
.button-wrapper {
  background-color: #fff;
  .button {
    display: block;
    width: 100%;
    background-color: #EF5350;
    color: white;
    font-weight: 600;
    border-radius: 0px;
    border: 0;
  }
}
@media (max-width: 640px) {
  .card {
    max-width: 45vw;
    min-width: 45vw;
    margin: 5px;
    display: flex;
    flex-direction: column;
  }
  .button {
    font-size: 12px;
  }
}
</style>
