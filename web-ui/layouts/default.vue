<template lang="pug">
  v-app#app(v-scroll="onScroll"  :style="$style()")
    nuxt
</template>

<!-- ============================================================================ -->

<script>
    export default {
      methods: {
        onScroll (e) {
          this.offsetTop = window.pageYOffset || document.documentElement.scrollTop
        }
      }
    }
</script>

<!-- ============================================================================ -->

<style lang="stylus">
  #app
    .u-scroll-up
      right     t(8)
      bottom    t(36)
      opacity   0
      transform scale(0)
      z-index   5
      &.visible
        opacity   1
        transform scale(1)
</style>
