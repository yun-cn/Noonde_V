<template lang="pug">
  #c-header-user
    v-toolbar(app flat color="primary" dense clipped-left)
      v-toolbar-title.u-fs30.u-fw9
        v-toolbar-side-icon.hidden-lg-and-up(dark @click="drawer = !drawer")
        nuxt-link.white--text.hidden-sm-and-down.u-fs24.u-fw9(to="/") YYYY

      v-spacer

      v-menu(offset-y)
        v-btn(icon dark slot="activator")
          v-icon more_vert

    v-navigation-drawer(app clipped v-model="drawer" width="210")
      v-list
        v-list-tile(
          active-class="c-active-class"
          :key="i"
          nuxt
          ripple
          :to="item.to"
          v-for="(item, i) in items"
        )
          v-list-tile-action.u-row-center-middle
            v-icon.c-icon(v-html="item.icon")

          v-list-tile-content.c-list-tile-content
            v-list-tile-title(v-text="item.title")

</template>

<!-- ============================================================================ -->

<script>
  export default {
    data () {
      let drawer = true;
      if (process.client && window.innerWidth < 1264) {
        drawer = false
      }

      let items = [
        { icon: 'loyalty', title: this.$t('components.c-header-user.navbar-recommended'), to: '/'},
        { icon: 'videogame_asset', title: this.$t('components.c-header-user.navbar-space'), to: '/space'},
        { icon: 'local_mall', title: this.$t('components.c-header-user.navbar-listing'), to: '/listing'},
        { icon: 'create', title: this.$t('components.c-header-user.navbar-contribution'), to: '/'},
      ];

      return {
        drawer: drawer,
        items: items,
      }
    }
  }
</script>

<!-- ============================================================================ -->

<style lang="stylus">
  #c-header-user
    .v-navigation-drawer
      background-color var(--nav-back) !important
      a
        color var(--nav-title-fore) !important
    .c-list-tile-content
      color var(--nav-title-fore) !important
    .c-icon
      color var(--nav-icon-fore) !important
    .c-active-class
      background-color var(--nav-back-active) !important
    .c-list-tile-content
      color var(--nav-title-fore-active) !important
    .c-icon
      color var(--nav-icon-fore-active) !important
    .v-navigation-drawer__border
      background-color var(--nav-border) !important
</style>
