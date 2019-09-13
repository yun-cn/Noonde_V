<template lang="pug">
  v-app#error
</template>

<!-- ============================================================================ -->

<script>
  export default {
    props: {
      error: Object,
    }
  }
</script>

<!-- ============================================================================ -->

<style lang="stylus">
  #error
    .p-container
      min-height 80vh

</style>
