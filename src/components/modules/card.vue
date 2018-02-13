<template>
  <div class="card">
    <div class="card-image">
      <figure class="img">
        <img :src="thumbnail" alt="thumbnail">
      </figure>
    </div>
    <button class="button is-primary" @click="vote">投票する</button>
  </div>
</template>

<script>
export default {
  props: ["thumbnail", "title"],
  methods: {
    vote() {
      this.$toast.open({
        message: `『${this.title}』に投票しました！`,
        type: 'is-success'
      })
    }
  }
}
</script>


<style lang="scss" scoped>
:scope {
  display: block;
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
  }
}
</style>
