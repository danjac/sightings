<template lang="pug">
div
  h1 "Reports go here"
  ul(v-if="!loading")
    li(v-for="report in reports") {{ report.location }}
  div(v-if="loading")
    h2 Loading....
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      reports: [],
      loading: false,
      error: false
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    async fetchData() {
      console.log('fetching data')
      this.loading = true
      this.error = false
      this.reports = []
      try {
        const response = await axios.get('/reports/')
        this.loading = false
        this.reports = response.data.items
      } catch (e) {
        this.loading = false
        this.error = true
        console.log('error', e)
      }
    }
  },
  watch: {
    $route: 'fetchData'
  }

}
</script>
