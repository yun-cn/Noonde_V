<template lang="pug">
.c-search.u-col-stretch-middle
  v-combobox(
    append-icon=""
    @blur="$emit('blur')"
    chips
    clearable
    deletable-chips
    @focus="$emit('focus')"
    :hide-details="!hint"
    :hint="hint"
    item-disabled="true"
    @keydown="$emit('keydown')"
    multiple
    outline
    :persistent-hint="!!hint"
    prepend-inner-icon="fas fa-search"
    single-line
    small-chips
    v-model="model"
  )

  .u-row-left-middle.u-wrap(v-if="(items || []).length > 0")
    template(v-for="item, i) in items")
      template(v-if="item.title")
        .c-search-title {{item.title}}

      template(v-else-if="item.delim")
        .c-search-delim
]
      template(v-else)
        v-chip.c-search-chip(
          @click="toggle($event, item)"
          :color="chipBackColor(item)"
          disabled
          small
          :text-color="chipForeColor(item)"
        ) {{item.label}}

</template>

<!-- ============================================================================ -->

<script>
  export default {
    data () {
      return {
        init:  true,
        model: null,
      }
    },

    methods: {
      active (item) {
        return this.$include(this.tags, item.value)
      },

      chipBackColor (item) {
        if (this.active(item)) return this.$color.searchChipBackActive
        return this.$color.searchChipBack
      },

      chipForeColor (item) {
        if (this.active(item)) return this.$color.searchChipForeActive
        return this.$color.searchChipFore
      },

      toggle (e, item) {
        let temps = this.tags
        if (this.$include(temps, item.value)) {
          temps = this.$remove(temps, item.value)
        } else {
          for (let item2 of this.items) {
            if (item.key && item2.key == item.key && item.value != item2.value) {
              if (this.$include(temps, item2.value)) {
                temps = this.$remove(temps, item2.value)
              }
            }
          }
          temps.push(item.value)
        }

        this.$emit('update:tags', temps)
        this.$emit('input', this.model)
      },
    },

    mounted () {
      this.init  = true
      this.model = this.value
    },

    props: {
      value: Array,
      tags:  Array,
      hint:  String,
      items: Array,
    },

    updated () {
      this.init = false
    },

    watch: {
      value (val) {
        this.model = val
      },

      model (val) {
        if (!this.init) this.$emit('input', val)
      },
    },
  }
</script>

<!-- ============================================================================ -->

<style lang="stylus">
  .c-search
    .c-search-title
      color         var(--search-chip-title) !important
      margin-top    t(6)                     !important
      margin-left   0                        !important
      margin-right  t(6)                     !important
      margin-bottom 0                        !important
      font-size     t(12)                    !important

    .c-search-delim
      background-color var(--search-chip-delim) !important
      height           t(24)                    !important
      width            t(4)                     !important
      margin-top       t(6)                     !important
      margin-left      0                        !important
      margin-right     t(6)                     !important
      margin-bottom    0                        !important

    .c-search-chip
      margin-top    t(6) !important
      margin-left   0    !important
      margin-right  t(6) !important
      margin-bottom 0    !important

      .v-chip__content
        cursor pointer !important

    // normal
    .v-input
      .v-input__slot
        border         1px solid var(--search-border) !important
        min-height     auto                           !important
        padding-top    t(6)                           !important
        padding-bottom t(6)                           !important

      .v-input__icon--prepend-inner > i
        color var(--search-icon) !important

      .v-input__icon--clear > i
        color var(--search-icon) !important

      .v-select__selections .v-chip
        background-color var(--search-chip-back-active) !important
        color            var(--search-chip-fore-active) !important

      .v-input__prepend-inner
        margin-right t(6) !important
        margin-top   t(4) !important

      .v-input__append-inner
        margin-top t(4) !important

      input
        margin-top 0     !important
        font-size  t(14) !important

      .v-text-field__details
        padding-left  t(4) !important
        padding-right t(4) !important

      .v-messages__message
        line-height 1.2 !important

    // focused
    .v-input.v-input--is-focused
      .v-input__slot
        border 1px solid var(--search-border-focused) !important

      .v-input__icon--prepend-inner > i
        color var(--search-icon-focused) !important

      .v-input__icon--clear > i
        color var(--search-icon-focused) !important

      .v-select__selections .v-chip
        background-color var(--search-chip-back-focused) !important
        color            var(--search-chip-fore-focused) !important

      input
        color       var(--search-fore-focused) !important
        caret-color var(--search-fore-focused) !important
</style>
