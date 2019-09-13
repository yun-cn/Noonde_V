<template lang="pug">
  .c-checkboxes
    .c-div.u-col-left-top(:class="klass")
      label(v-if="label") {{label}}

      .c-checkbox-div(:class="{ 'c-checkbox-div-row': row }")
        v-checkbox(
          v-for="(item, i) in items"
          :disabled="item.disabled"
          hide-details
          :key="i"
          :label="item.label"
          v-model="model"
          :value="item.value"
        )

      .c-messages(v-if="!hideDetails")
        p(v-if="!error") {{hint}}
        div(v-if="(errorMessages || []).length > 0")
          p {{errorMessages[0]}}
</template>

<style lang="stylus">
  .c-checkboxes

    // normal
    .c-div
      & > label
        color         var(--checkboxes-fore) !important
        font-size     t(12)                  !important
        margin-bottom t(4)                   !important

      .c-messages
        color      var(--checkboxes-fore) !important
        font-size  t(12)                  !important
        margin-top t(4)                   !important

      .c-checkbox-div
        display         flex       !important
        flex-direction  column     !important
        justify-content flex-start !important
        align-items     flex-start !important
        margin-bottom   t(-2)      !important

      // row
      .c-checkbox-div.c-checkbox-div-row
        flex-direction  row        !important
        justify-content left       !important
        align-items     flex-start !important
        flex-wrap       wrap       !important

      // normal - checkbox
      .v-input
        margin-top    0     !important
        padding-top   0     !important
        margin-right  t(12) !important
        margin-bottom t(2)  !important

        .v-label
          animation none                   !important
          color     var(--checkboxes-fore) !important
          font-size t(14) !important

        .v-icon
          color var(--checkboxes-icon) !important

        .v-input--selection-controls__input
          margin-right t(4) !important

      // active - checkbox
      .v-input.v-input--is-label-active
        .v-label
          color var(--checkboxes-fore-active) !important

        .v-icon
          color var(--checkboxes-icon-active) !important

        .v-ripple__container
          color var(--checkboxes-icon-active) !important

      // disabled - checkbox
      .v-input.v-input--is-disabled
        .v-label
          color var(--disabled-fore) !important

        .v-icon
          color var(--disabled-icon) !important

    // error
    .c-div.c-checkboxes-error
      & > label
        color var(--error-fore) !important

      .c-messages
        color      var(--error-fore) !important

      // normal - checkbox
      .v-input
        .v-label
          color var(--error-fore) !important

        .v-icon
          color var(--error-fore) !important

      // active - checkbox
      .v-input.v-input--is-label-active
        .v-label
          color var(--error-fore) !important

        .v-icon
          color var(--error-fore) !important

        .v-ripple__container
          color var(--error-fore) !important

      // disabled - checkbox
      .v-input.v-input--is-disabled
        .v-label
          color var(--disabled-fore) !important

        .v-icon
          color var(--disabled-icon) !important
</style>

<script>
  export default {
    computed: {
      klass () {
        return {
          'c-checkboxes-error': this.error || (this.errorMessages || []).length > 0,
        }
      },
    },

    data () {
      return {
        init:  true,
        model: null,
      }
    },

    mounted () {
      this.init  = true
      this.model = this.value
    },

    props: {
      error:         Boolean,
      errorMessages: Array,
      hideDetails:   Boolean,
      hint:          String,
      items:         Array,
      label:         String,
      row:           Boolean,
      rules:         Array,
      value:         Array,
    },

    updated () {
      this.init = false
    },

    watch: {
      value (val) {
        this.model = val
      },

      model (val) {
        if (this.init) return;
        this.$emit('input', val)
      },
    },
  }
</script>
