<template lang="pug">
  v-card-title.c-card-title
    slot
</template>

<style lang="stylus">
  .c-card-title
    background-color var(--card-title-back) !important
    color            var(--card-title-fore) !important
    font-weight      500                    !important
    padding          t(12) t(18)            !important
    font-size        t(15)                  !important

    i
      color var(--card-title-fore) !important
</style>
