<template lang="pug">
  v-card-text.c-card-text(:style="style" :class="klass")
    slot
</template>

<style lang="stylus">
  .c-card-text
    background-color var(--card-text-back)        !important
    color            var(--card-text-fore)        !important
    border-left      1px solid var(--card-border) !important
    border-right     1px solid var(--card-border) !important
    padding          t(12) t(18)                  !important
    overflow-y       auto                         !important
    font-size        t(14)                        !important

  .c-card-text.c-card-text-no-actions
    border-bottom 1px solid var(--card-border) !important

  .v-dialog
    .c-card-text
      border-left  none !important
      border-right none !important
      max-height   70vh !important
</style>

<script>
  export default {
    computed: {
      klass () {
        return {
          'c-card-text-no-actions': this.noActions,
        }
      },

      style () {
        return {
          maxHeight: this.height && `${this.height} !important`
        }
      },
    },

    props: {
      height:    String,
      noActions: Boolean,
    },
  }
</script>
