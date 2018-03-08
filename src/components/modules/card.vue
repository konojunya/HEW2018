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
  props: ["item", "voteme"],
  data() {
    return {
      isImageModalActive: false
    }
  },
  methods: {
    vote() {
      if(this.canVote()) {
        axios.post(`/api/products/${this.item.id}/vote`).then((res) => {
          this.$toast.open({
            message: `『${this.item.title}』に投票しました！`,
            type: 'is-success'
          })
          this.setLocalStorage()
          this.isImageModalActive = false
        })
      } else {
        this.$dialog.confirm({
          message: "1人、1票までしか投票できません。<br/>この作品に票をいれるともう1度、投票することができます。",
          cancelText: 'やめておく',
          confirmText: '投票する',
          type: 'is-danger',
          onConfirm: () => {
            this.removeLocalStorage()
            this.voteme()
          }
        })
      }
    },
    removeLocalStorage() {
      localStorage.removeItem("hew2018-vote")
    },
    setLocalStorage() {
      localStorage.setItem("hew2018-vote", this.item.id)
    },
    canVote() {
      return !localStorage.getItem("hew2018-vote")
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
