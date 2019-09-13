<template lang="pug">
  v-dialog(v-model="model" max-width="40%" persistent)
    c-card#add-url-name
      v-form(
        ref="form"
        @submit.prevent="submit"
        v-model="valid"
      )
        c-card-title {{ $t('pages.note.new.add-url-name.dialog-title') }}

        c-card-text(ref="card-text")
          .u-col-stretch-top.u-full.u-pb16
            c-text-field.u-mb16(
              :disabled="loading"
              :label="$t('pages.note.new.add-url-name.url-name-id')"
              hint="xxxxxxx"
              @input="urlNameError = []"
              :rules="urlNameRules"
              v-model="urlName"
            )

        c-card-divider
        c-card-actions
          v-spacer
          c-btn(
            :disabled="loading || !valid"
            flat
            icon-left="fas fa-check-circle"
            :loading="loading"
            type="submit"
          ) {{ $t('pages.note.new.add-url-name.entry-btn') }}
</template>

<!-- ============================================================================ -->

<script>
    export default {
      data () {
        return {
          valid:           true,
          model:           false,
          loading:         false,

          urlNameError:    [],
          urlNameRules:    [],
          urlName:         '',
          source:          null,
        }
      },

      mounted() {
        this.model = this.value;

        this.urlNameRules  =  [
          v => !!v || this.$t('pages.note.new.add-url-name.empty-url-name'),
          v => (v && v.length >= 6 && v.length <= 20) ||this.$t('pages.note.new.add-url-name.validate-size-url-name-message'),
          v => (v && /^[a-z0-9\-]+$/.test(v)) || this.$t('pages.note.new.add-url-name.validate-url-name-message')
        ];

        this.init()
      },

      props: {
        value: Boolean,
        current: Object,
      },

      methods: {
        init () {
          this.$refs.form.reset();

          setTimeout(() => {
            this.$refs['card-text'].$el.scrollTop = 0
          }, 0)
        },

        async submit (e) {
          if (!this.$refs.form.validate()) return;
          
          this.loading = true;

          try {
            if (this.source && this.source.cancel) this.source.cancel('canceled');
            this.source = this.$source();

            let { data } = await this.$axios.post(`/www/urlname/update/${id}`)

          } catch (error) {
            
          }
        }
      }
    }
</script>

<!-- ============================================================================ -->

<style lang="stylus">
  #add-url-name
    .dropzone
      padding t(12) 0 0 0
</style>
