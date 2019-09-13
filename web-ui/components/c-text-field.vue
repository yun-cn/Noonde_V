<template lang="pug">
  .c-text-field.u-row-left-bottom(:class="klass")
    v-text-field(
      :append-icon="appendIcon"
      clearable
      @click:append="$emit('clickAppend')"
      :disabled="disabled"
      :error="error"
      :error-messages="errorMessages"
      :hide-details="hideDetails"
      :hint="hint"
      :label="label"
      :maxlength="maxlength || 171"
      outline
      persistent-hint
      :prepend-inner-icon="prependInnerIcon"
      :rules="rules"
      :type="type"
      v-model="model"
    )

    .c-append.u-nowrap(v-if="append") {{append}}
</template>

<style lang="stylus">

  // normal
  .c-text-field
    padding-top t(20) !important

    .c-append
      color         var(--text-field-append) !important
      font-size     t(12) !important
      margin-left   t(6)  !important
      margin-bottom t(26) !important

    // normal
    .v-input
      .v-input__slot
        border          1px solid var(--text-field-border) !important
        margin-bottom   0                                  !important
        min-height      auto                               !important
        padding-top     t(8)                               !important
        padding-bottom  t(8)                               !important
        font-size       t(14)                              !important

      .v-input__icon--prepend-inner > i
        color var(--text-field-icon) !important

      .v-text-field__slot > .v-label
        color       var(--text-field-fore) !important
        margin-top  t(-50)                 !important
        margin-left t(-16)                 !important
        transform   none                   !important
        font-size   t(12)                  !important
        max-width   initial                !important

      .v-text-field__slot > input
        color       var(--text-field-fore-active) !important
        caret-color var(--text-field-fore-active) !important
        margin-top  t(2)                          !important
        padding     0                             !important

      .v-text-field__slot > input:-webkit-autofill
        -webkit-box-shadow: 0 0 0 1000px white inset !important

      .v-input__icon--clear > i
        color var(--text-field-icon) !important

      .v-input__icon--append > i
        color var(--text-field-icon) !important

      .v-input__prepend-inner
        margin-top 0 !important

      .v-input__append-inner
        margin-top 0 !important

      .v-text-field__details
        margin-top    t(4) !important
        margin-bottom 0    !important
        padding-left  0    !important

      .v-messages__message
        color var(--text-field-fore) !important

    // focused
    .v-input.v-input--is-focused
      .v-input__slot
        border 1px solid var(--text-field-border-focused) !important

      .v-input__icon--prepend-inner > i
        color var(--text-field-icon-focused) !important

      .v-text-field__slot > .v-label
        color var(--text-field-fore-focused) !important

      .v-text-field__slot > input
        color       var(--text-field-fore-focused) !important
        caret-color var(--text-field-fore-focused) !important

      .v-input__icon--clear > i
        color var(--text-field-icon-focused) !important

      .v-input__icon--append > i
        color var(--text-field-icon-focused) !important

      .v-messages__message
        color var(--text-field-fore-focused) !important

    // disabled
    .v-input.v-input--is-disabled
      .v-input__slot
        border 1px solid var(--disabled-border) !important

      .v-input__icon--prepend-inner > i
        color var(--disabled-icon) !important

      .v-text-field__slot > .v-label
        color var(--disabled-fore) !important

      .v-text-field__slot > input
        color var(--disabled-fore) !important

      .v-input__icon--clear > i
        display none !important

      .v-input__icon--append > i
        color var(--disabled-icon) !important

      .v-messages__message
        color var(--disabled-fore) !important

    // error
    .v-input.error--text
      .v-input__slot
        border 1px solid var(--error-border) !important

      .v-input__icon--prepend-inner > i
        color var(--error-icon) !important

      .v-text-field__slot > .v-label
        color var(--error-fore) !important

      .v-text-field__slot > input
        color       var(--error-fore) !important
        caret-color var(--error-fore) !important

      .v-input__icon--clear > i
        color var(--error-icon) !important

      .v-input__icon--append > i
        color var(--error-icon) !important

      .v-messages__message
        color var(--error-fore) !important

  // no label
  .c-text-field.c-text-field-no-label
    padding-top 0 !important

  // hide details
  .c-text-field.c-text-field-hide-details
    .c-append
      margin-bottom t(12) !important

  // prepend inner icon
  .c-text-field.c-text-field-prepend-inner-icon
    .v-text-field__slot > .v-label
      margin-left t(-44) !important

  // disabled
  .c-text-field.c-text-field-disabled
    .c-append
      color var(--disabled-fore) !important

  // error
  .c-text-field.c-text-field-error
    .c-append
      color var(--error-fore) !important
</style>

<script>
  export default {
    computed: {
      klass () {
        return {
          'c-text-field-no-label':           !this.label,
          'c-text-field-prepend-inner-icon': !!this.prependInnerIcon,
          'c-text-field-hide-details':       this.hideDetails,
          'c-text-field-error':              this.error,
          'c-text-field-disabled':           this.disabled,
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
      append:           String,
      appendIcon:       String,
      disabled:         Boolean,
      error:            Boolean,
      errorMessages:    Array,
      hideDetails:      Boolean,
      hint:             String,
      label:            String,
      maxlength:        String,
      prependInnerIcon: String,
      rules:            Array,
      type:             String,
      value:            String,
    },

    updated () {
      this.init = false
    },

    watch: {
      value (val) {
        this.model = val
      },

      model (val) {
        if (this.init) return
        this.$emit('input', val)
      },
    },
  }
</script>
