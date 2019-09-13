<template lang="pug">
  v-footer(:fixed="onFixed")#c-footer.u-row-center-middle
    span &copy; YYYY
</template>

<!-- ============================================================================ -->

<script>
    export default {
      data () {
        return {
          onFixed: true
        }
      },

      mounted() {
        this.onFixed = true;
      }
    }
</script>

<!-- ============================================================================ -->

<style lang="stylus">
  #c-footer
    background-color var(--footer-back)
    color            var(--footer-fore)
    border-top       1px solid var(--footer-border)
    bottom 0 /*下に固定*/
</style>
