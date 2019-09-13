<template lang="pug">
  .c-breadcrumbs
    v-breadcrumbs(:items="items")
      template(slot="item" slot-scope="props")
        li
          nuxt-link(
            :to="`${props.item.to}`"
            :class="props.item.disabled && 'disabled'"
            exact
          ) {{props.item.label.toUpperCase()}}
</template>

<!-- ============================================================================ -->

<script>
  export default {
    props: {
      items: Array,
    },
  }
</script>

<!-- ============================================================================ -->

<style lang="stylus">
  .c-breadcrumbs
    .v-breadcrumbs
      padding 0 !important

    .v-breadcrumbs__divider
      color var(--breadcrumbs-delim-fore) !important

    a
      color var(--breadcrumbs-fore) !important

    a:hover
      color var(--breadcrumbs-fore-hover) !important

    a.nuxt-link-exact-active
      color var(--breadcrumbs-fore-active) !important

    a.nuxt-link-exact-active:hover
      color var(--breadcrumbs-fore-active-hover) !important
</style>
