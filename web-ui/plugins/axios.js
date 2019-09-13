export default function ({ app, $axios, store }) {
  $axios.defaults.baseURL = app.$env.BASE_URL || 'http://127.0.1:3000';

  $axios.onRequest(config => {
    config.headers.common['Content-Type'] = 'application/json';

    if (store.state.auth.token) {
      config.headers.common['Authorization'] = `Bearer ${store.state.auth.token}`
    }
  })
}
