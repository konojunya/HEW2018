<template>
  <div class="wrapper">
    <div class="content">
      <p class="name">西川 かずき</p>
      <p class="point">2point</p>
    </div>
    <div class="bg"></div>
  </div>
</template>

<style lang="scss" scoped>
.wrapper {
  width: 100%;
  height: 10vh;
  position: relative;
  margin: 5px 0;
  display: flex;
  align-items: center;
  .bg {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 10vh;
    background-color: red;
  }
  .content {
    z-index: 10;
    display: flex;
    align-items: center;
  }
}
</style>
