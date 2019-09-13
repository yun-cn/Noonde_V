<template lang="pug">
  #space-index
    c-header-user

    v-content
      .u-row-between-center.u-wrap.u-row-reverse.u-mb16.search-postion
        c-search.p-search.u-full(
          @input="page = 1; reload($event)"
          :items="opts"
          :tags.sync="tags"
          v-model="searches"
        )

</template>

<!-- ============================================================================ -->
<script>
    export default {

    }
</script>

<!-- ============================================================================ -->

<style lang="stylus">
  #space-index
    .search-postion
      margin-top 30px
      margin-left 30px
      margin-right 30px
</style>
