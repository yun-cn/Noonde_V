<template lang="pug">
  v-card-actions.c-card-actions.u-wrap
    slot
</template>

<style lang="stylus">
  .c-card-actions
    background-color var(--card-actions-back)     !important
    border-bottom    1px solid var(--card-border) !important
    border-left      1px solid var(--card-border) !important
    border-right     1px solid var(--card-border) !important

  .v-dialog
    .c-card-actions
      border-bottom none !important
      border-left   none !important
      border-right  none !important
</style>
