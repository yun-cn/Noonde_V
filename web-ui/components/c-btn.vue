<template lang="pug">
  .c-btn(:class="klass")
    v-btn(
      @click="$emit('click', $event)"
      :color="backColor"
      :disabled="disabled"
      :flat="flat"
      :icon="!!icon"
      :loading="loading"
      :style="{ color: `${foreColor} !important` }"
      :to="to"
      :type="type"
      :nuxt="!!to"
      exact
    )
      .u-row-center-middle
        v-icon.c-icon-left(small v-if="iconLeft") {{iconLeft}}

        span(v-if="!icon")
          slot

        v-icon.c-icon-right(small v-if="iconRight") {{iconRight}}

        v-icon(v-if="icon") {{icon}}

      v-icon.u-loader(small slot="loader" :color="foreColor") fas fa-sync-alt
</template>

<style lang="stylus">

  // normal
  .c-btn
    background-color transparent

    .v-btn
      box-shadow    none  !important
      margin        0     !important
      font-size     t(16) !important
      border-radius t(2)  !important
      height        auto  !important
      min-width     auto  !important
      padding-right t(16) !important
      padding-left  t(16) !important

    .v-btn__content
      margin-top    t(6) !important
      margin-bottom t(6) !important

    .c-icon-left
      margin-right t(6)  !important
      margin-top   t(-2) !important

    .c-icon-right
      margin-left t(6)  !important
      margin-top  t(-2) !important

  // flat
  .c-btn.c-btn-flat
    .v-btn
      border-radius t(2)  !important
      font-size     t(14) !important

  // mini
  .c-btn.c-btn-mini
    .v-btn
      border-radius t(2)  !important
      padding-right t(10) !important
      padding-left  t(10) !important
      font-size     t(14) !important

    .v-btn__content
      margin-top    t(4) !important
      margin-bottom t(4) !important
</style>

<script>
  export default {
    computed: {
      klass () {
        return {
          'c-btn-flat': this.flat,
          'c-btn-mini': this.mini,
        }
      },

      backColor () {
        if (this.disabled) return this.$color.disabledBack

        if (!(this.inverted && this.flat) && (this.inverted || this.flat)) {
          if (this.danger)    return this.$color.btnDangerFore
          if (this.error)     return this.$color.btnErrorFore
          if (this.important) return this.$color.btnImportantFore
          if (this.simple)    return this.$color.btnSimpleFore
          return this.$color.btnFore
        }

        if (this.danger)    return this.$color.btnDangerBack
        if (this.error)     return this.$color.btnErrorBack
        if (this.important) return this.$color.btnImportantBack
        if (this.simple)    return this.$color.btnSimpleBack
        return this.$color.btnBack
      },

      foreColor () {
        if (this.disabled) return this.$color.disabledFore

        if (!(this.inverted && this.flat) && (this.inverted || this.flat)) {
          if (this.danger)    return this.$color.btnDangerBack
          if (this.error)     return this.$color.btnErrorBack
          if (this.important) return this.$color.btnImportantBack
          if (this.simple)    return this.$color.btnSimpleBack
          return this.$color.btnBack
        }

        if (this.danger)    return this.$color.btnDangerFore
        if (this.error)     return this.$color.btnErrorFore
        if (this.important) return this.$color.btnImportantFore
        if (this.simple)    return this.$color.btnSimpleFore
        return this.$color.btnFore
      },
    },

    props: {
      danger:    Boolean,
      disabled:  Boolean,
      error:     Boolean,
      flat:      Boolean,
      icon:      String,
      iconLeft:  String,
      iconRight: String,
      important: Boolean,
      inverted:  Boolean,
      loading:   Boolean,
      simple:    Boolean,
      to:        String,
      type:      String,
      mini:      Boolean,
    },
  }
</script>
