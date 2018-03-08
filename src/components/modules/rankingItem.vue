<template>
  <div class="wrapper">
    <div class="content">
      <p class="name">{{item.author}}</p>
      <p class="point">{{item.votes}} 票</p>
    </div>
    <div class="bg" :style="style"></div>
  </div>
</template>

<script>
import bgColor from '../../config/ranking-bgcolor.json'

export default {
  props: ["item", "max", "index"],
  data() {
    return {
      style: {
        width: '0%',
        backgroundColor: "transparent"
      }
    }
  },
  created() {
    const width = this.item.votes / parseInt(this.max) * 100
    this.style.width = `${width}%`
    this.style.backgroundColor = bgColor[this.index]
  }
}
</script>

<style lang="scss" scoped>
.wrapper {
  width: 100%;
  height: 15vh;
  position: relative;
  margin: 10px 0;
  display: flex;
  align-items: center;
  .bg {
    position: absolute;
    top: 0;
    left: 0;
    width: 0%;
    height: 15vh;
  }
  .content {
    width: 100%;
    z-index: 10;
    display: flex;
    align-items: center;
    justify-content: flex-start;
    margin: 0;
    color: white;
    font-weight: 600;
    .name {
      margin: 0 auto 0 30px;
      font-size: 1.5rem;
    }
    .point {
      margin-right: 30px;
      font-size: 2rem;
    }
  }
}
@media (max-width: 640px) {
  .wrapper {
    height: 10vh;
    .bg {
      height: 10vh;
    }
    .content {
      .name {
        font-size: 16px;
      }
      .point {
        font-size: 16px;
      }
    }
  }
}
</style>
