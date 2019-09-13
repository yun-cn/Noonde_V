<template lang="pug">
  .c-chip
    v-chip(
      small
      label
      disabled
      :color="backColor"
      :text-color="foreColor"
    )
      template(v-if="to")
        nuxt-link.c-chip-link(:to="to" exact)
          slot

      template(v-else)
        span.c-chip-span
          slot
</template>

<style lang="stylus">
  .c-chip
    .v-chip
      margin    0     !important
      height    auto  !important
      font-size t(13) !important

      .v-chip__content
        padding     0      !important
        height      auto   !important
        white-space normal !important
        line-height 1.1    !important

      .c-chip-span
        padding t(4) t(6) !important

      .c-chip-link
        color   var(--chip-link-fore) !important
        padding t(4) t(6)             !important
</style>

<script>
  export default {
    computed: {
      backColor () {
        if (this.to) return this.$color.chipLinkBack
        return this.$color.chipBack
      },

      foreColor () {
        if (this.to) return this.$color.chipLinkFore
        return this.$color.chipFore
      },
    },

    props: {
      to: String,
    },
  }
</script>
