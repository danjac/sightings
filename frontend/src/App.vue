<template lang="pug">
  v-app(dark)
    v-toolbar(app clipped-left fixed)
    v-content
      v-container
        router-view
    v-footer
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
</style>
