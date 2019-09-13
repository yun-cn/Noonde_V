<template lang="pug">
  .c-bars.u-col-left-middle
    .c-bar.c-bar-1
    .c-bar.c-bar-2.u-my2
    .c-bar.c-bar-3
</template>

<!-- ============================================================================ -->

<script>
  export default {
    props: {
      widths: Array,
    },
  }
</script>

<!-- ============================================================================ -->

<style lang="stylus">
  .c-bars
    width 100%

    .c-bar
      border-radius   p(2, 2)
      height          p(20, 20)
      background      linear-gradient(90deg, var(--bar-start), var(--bar-end))
      background-size 400%
      animation       gradient 2s ease infinite

    .c-bar-1
      width 80%

    .c-bar-2
      width 40%

    .c-bar-3
      width 60%

  @keyframes gradient
    0% {
      background-position 100%
    }
    50% {
      background-position 0%
    }
    100% {
      background-position 100%
    }
</style>
