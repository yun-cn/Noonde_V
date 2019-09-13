import Vue from 'vue';
import Cookies from 'js-cookie';
import VueI18n from 'vue-i18n';

Vue.use(VueI18n);

export default ({ app, store }) => {
  app.i18n = new VueI18n({
    locale: Cookies.get('language') || 'ja',
    fallbackLocale: 'ja',
    messages: {
      'ja': require('json-loader!yaml-loader!~/assets/locales/ja.yaml'),
      'en-US': require('json-loader!yaml-loader!~/assets/locales/en-US.yaml')
    }
  })
}
