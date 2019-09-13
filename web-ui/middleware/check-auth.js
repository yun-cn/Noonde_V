export default function ({ app, store, req }) {
  if (process.server && !req) return

  const auth = process.server ? app.$auth.getFromCookie(req) : app.$auth.getFromLocalStorage();
  store.commit('merge', ['auth', auth])
}
