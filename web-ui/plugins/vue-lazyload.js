import Vue from 'vue';
import VueLazyload from 'vue-lazyload';

Vue.use(VueLazyload, {
  preLoad: 1.0,
  attempt: 3,
  lazyComponent: true,
});
