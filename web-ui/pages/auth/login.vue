<template lang="pug">
  #auth-login
    c-header

    v-content
      .u-col-center-middle.p-container
        c-card
          v-form(
            ref="form"
            @submit.prevent="submit"
            v-model="valid"
          )
            c-card-title {{ $t('pages.auth.login.title') }}

            c-card-text.p-card-text
              .u-col-stretch-middle.u-pa12
                .p-error.u-mb8.u-fs14(v-if="error") {{error}}

                c-text-field.u-mb16(
                  :disabled="loading"
                  :label="$t('pages.auth.login.email')"
                  prepend-inner-icon="fas fa-envelope"
                  :rules="emailRules"
                  v-model="email"
                )

                c-text-field.u-mb16(
                  :append-icon="mask ? 'far fa-eye' : 'far fa-eye-slash'"
                  @clickAppend="() => (mask = !mask)"
                  :disabled="loading"
                  :label="$t('pages.auth.login.password')"
                  prepend-inner-icon="fas fa-key"
                  :rules="passwordRules"
                  :type="mask ? 'password' : 'text'"
                  v-model="password"
                )

                c-checkboxes(
                  hide-details
                  :items="items"
                  v-model="rememberMe"
                )

            c-card-divider

            c-card-actions
              c-btn(
                @click="email='yun313350095@gmail.com'; password='xiaozhu521'"
                flat
                simple
                v-if="$env.DEBUG"
              ) Debug

              v-spacer

              c-btn(
                flat
                :loading="loading"
                type="submit"
                :valid="valid"
              ) {{ $t('pages.auth.login.login-btn') }}

    c-footer
</template>

<style lang="stylus">
  #auth-login
    .p-container
      min-height 80vh

    .p-card-text
      max-width p(480, 340)
      min-width p(480, 340)

    .p-error
      color var(--error-fore)
</style>

<script>
  export default {
    asyncData () {
      return {
        email:      '',
        emailRules: [],
        error:      false,

        items: [
          { label: 'ログイン情報を記録する', value: true },
        ],

        loading:       false,
        password:      '',
        passwordRules: [],
        rememberMe:    [],
        mask:          true,
        valid:         true,
      }
    },

    mounted () {
      this.emailRules   =  [
        v => !!v || this.$t('pages.auth.sign.email-error'),
        v => /^\w+([.\+-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/.test(v) || this.$t('pages.auth.login.validate-email-message')
      ];
      this.passwordRules = [
        v => !!v || this.$t('pages.auth.sign.password-error'),
        v => (v && v.length >= 6 && v.length <= 16) ||this.$t('pages.auth.login.validate-password-error')
      ];
      this.init()
    },

    methods: {
      init() {
        this.$refs.form.reset();
        this.warn = false;
        this.rememberMe = []
      },

      async submit () {
        if (!this.$refs.form.validate()) return;

        this.loading = true;

        try {
          let { data } = await this.$axios.post('/www/token/create', {
            email:    this.email,
            password: this.password,
          });
          console.log(data);

          this.$auth.set(data.token, data.user, this.rememberMe[0]);

          // this.$router.replace('/')

        } catch (error) {
          console.log(error);
          this.error = this.$t('pages.auth.login.login-error')
        }

        this.loading = false
      },
    },
  }
</script>
