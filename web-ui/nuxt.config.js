const pkg = require('./package');
const webpack = require('webpack');

module.exports = {
  mode: 'universal',

  head: {
    title: pkg.name,
    meta: [
      {charset: 'utf-8'},
      {name: 'viewport', content: 'width=device-width, initial-scale=1'},
      {hid: 'description', name: 'description', content: pkg.description}
    ],
    link: [
      {rel: 'icon', type: 'image/x-icon', href: '/favicon.ico'},
      {rel: 'stylesheet', href: 'https://fonts.googleapis.com/css?family=Roboto:300,400,500,700|Material+Icons'},
      {rel: 'stylesheet', href: 'https://fonts.googleapis.com/earlyaccess/notosansjapanese.css'},
      {rel: 'stylesheet', href: 'https://use.fontawesome.com/releases/v5.5.0/css/all.css'},
    ]
  },

  loading: {color: '#fff'},

  css: [
    '~/assets/css/require.styl',
    '~/assets/css/common.styl',
    '~/assets/css/app.styl',
    'quill/dist/quill.snow.css',
    'quill/dist/quill.bubble.css',
    'quill/dist/quill.core.css'
  ],

  plugins: [
    '@/plugins/i18n',
    '@/plugins/axios',
    '@/plugins/vuetify',
    '@/plugins/vue-lazyload',
    '@/plugins/components',
    '@/plugins/injects',
    { src: '@/plugins/vue-quill-editor.js', ssr: false },
  ],

  styleResources: {
    stylus: './assets/css/funcs.styl'
  },

  manifest: {
    name: 'Sample',
    display: 'standalone',
  },

  modules: [
    '@nuxtjs/axios',
    '@nuxtjs/style-resources',
    '@nuxtjs/pwa',
    ['nuxt-env', {keys: ['BASE_URL', 'DEBUG']}],
  ],

  router: {
    middleware: [
      'check-auth'
    ]
  },

  build: {
    cssSourceMap: false,

    extractCSS: true,
    extend(config, ctx) {
      if (ctx.isDev && process.client) {
        config.module.rules.push({
          enforce: 'pre',
          test: /\.(js|vue)$/,
          loader: 'eslint-loader',
          exclude: /(node_modules)/
        })
      }
    },
    plugins: [
      new webpack.ProvidePlugin({
        'window.Quill': 'quill/dist/quill.js',
        'Quill': 'quill/dist/quill.js'
      })
    ]
  }
};
