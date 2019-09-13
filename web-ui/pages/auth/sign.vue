<template lang="pug">
  #auth-sign
    c-header

    v-content
      .u-col-center-middle.p-container
        c-card
          v-form(
            ref="form"
            @submit.prevent="submit"
            v-model="valid"
          )
            c-card-title {{ $t('pages.auth.sign.title') }}

            c-card-text.p-card-text
              .u-col-stretch-middle.u-pa12
                .p-error.u-mb8.u-fs14(v-if="error") {{error}}

                c-text-field.u-mb16(
                  :disabled="loading"
                  :error-messages="emailError"
                  hint="(例) someone@gmail.com"
                  :label= "$t('pages.auth.sign.email')"
                  @input="emailError = []"
                  :rules="emailRules"
                  v-model="email"
                  prepend-inner-icon="fas fa-envelope"
                )

                c-text-field.u-mb16(
                  :append-icon="mask ? 'far fa-eye' : 'far fa-eye-slash'"
                  @clickAppend="() => (mask = !mask)"
                  :disabled="loading"
                  hint="(例) mysecretpassword"
                  :label="$t('pages.auth.sign.password')"
                  :error-messages="passwordError"
                  @input="passwordError = []"
                  prepend-inner-icon="fas fa-key"
                  :rules="passwordRules"
                  :type="mask ? 'password' : 'text'"
                  v-model="password"
                )

                c-text-field.u-mb16(
                  :disabled ="loading"
                  :error-messages="nicknameError"
                  hint="(例) 山田 太郎"
                  @input="nicknameError = []"
                  :label="$t('pages.auth.sign.nick-name')"
                  :rules="nicknameRules"
                  prepend-inner-icon="fas fa-id-card-alt"
                  v-model="nickname"
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
              ) {{ $t('pages.auth.sign.create-btn') }}
    c-footer
</template>

<!-- ============================================================================ -->

<script>
  import Identicon from 'identicon.js'

  export default {
      data() {
        return {
          croppie:       null,
          email:         '',
          emailError:    [],
          emailRules:    [],

          password:      '',
          passwordError: [],
          passwordRules: [],

          nickname:       '',
          nicknameError:  [],
          nicknameRules:  [],

          timeout1:      null,
          timeout2:      null,
          valid:         true,
          error:         false,
          loading:       false,
          source:        null,
          src:           '',
          mask:          true,
        }
      },

      destroyed () {
        if (this.croppie)  this.croppie.destroy();
        if (this.timeout1) clearTimeout(this.timeout1)
      },

      mounted() {

        this.emailRules   =  [
          v => !!v || this.$t('pages.auth.sign.email-error'),
          v => /^\w+([.\+-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/.test(v) || this.$t('pages.auth.sign.validate-email-message')
        ];
        this.nicknameRules = [
          v => !!v || this.$t('pages.auth.sign.nick-name-error'),
          v => (v && v.length >= 2 && v.length <= 16) ||this.$t('pages.auth.sign.validate-name-error')
        ];
        this.passwordRules = [
          v => !!v || this.$t('pages.auth.sign.password-error'),
          v => (v && v.length >= 6 && v.length <= 16) ||this.$t('pages.auth.sign.validate-password-error')
        ];

        this.init()
      },

      methods: {
        init() {
          this.$refs.form.reset();
          this.warn = false;
          this.setIdenticon()
        },

        setIdenticon() {
          const chars = '0123456789ABCDEFGHIJ';
          let rand = '';
          for (let i = 1; i <= 15; i++) {
            rand += chars.charAt(Math.floor(Math.random() * chars.length))
          }
          let data = new Identicon(rand, {
            margin: 0.2,
            size: 300,
          }).toString();
          this.src = `data:image/png;base64,${data}`;
        },

        async submit() {
          if (!this.$refs.form.validate()) return;

          this.loading = true;

          try {
            let { data } = await this.$axios.post('/www/user/create', {
              email: this.email,
              password: this.password,
              nickname: this.nickname,
              avatar_data: this.src,
            })
          } catch (error) {
            if (error.message == 'canceled') return;

            if (error.response && error.response.data && error.response.data.detail) {
              let detail = error.response.data.detail;

              if (detail.email && detail.email.code == 'P00004') {
                this.emailError = [this.$t('pages.auth.sign.invalid-value-email-message')]
              }
            }
          }
          this.loading = false
        }
      },

      props: {
        value: Boolean,
      },

      watch: {
        model (val) {
          this.$emit('input', val)
        },
      }
    }
</script>

<!-- ============================================================================ -->

<style lang="stylus">
  #auth-sign
    .p-container
      min-height 80vh

    .p-card-text
      max-width p(480, 340)
      min-width p(480, 340)
</style>
