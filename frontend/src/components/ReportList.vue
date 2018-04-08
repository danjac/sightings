<template lang="pug">
v-data-table(
  :total-items="pagination.totalItems"
  :pagination.sync="pagination"
  :items="pagination.items"
  :headers="headers"
  :rows-per-page-items="[10, 20, 40]"
)
  template(slot="items" slot-scope="props")
    td
      a(href='#') {{ formatDate(props.item.occurred_at) }}
    td {{ props.item.shape }}
    td {{ props.item.location }}
</template>

<script>
import moment from 'moment'
import axios from 'axios'

export default {
  data() {
    return {
      headers: [
        {
          text: 'Date',
          value: 'occurred_at',
          sortable: false
        },
        {
          text: 'Shape',
          value: 'shape',
          sortable: false
        },
        {
          text: 'Place',
          value: 'location',
          sortable: false
        }
      ],
      pagination: {
        loading: false,
        page: parseInt(this.$route.page) || 1,
        rowsPerPage: 20,
        totalItems: 0,
        items: [],
        sortBy: 'occurred_at'
      }
    }
  },
  methods: {
    formatDate(value) {
      return moment(value).format('MMMM Do YYYY')
    },
    async fetchData(params) {
      if (this.pagination.loading) {
        return
      }
      this.pagination.loading = true

      const response = await axios.get('/reports/', { params })

      this.pagination.items = response.data.items
      this.pagination.totalItems = response.data.total
      this.pagination.pages = response.data.pages
      this.pagination.loading = false
    }
  },
  watch: {
    pagination: {
      handler() {
        this.fetchData({
          page: this.pagination.page,
          per_page: this.pagination.rowsPerPage
        })
      }
    }
  }
}
</script>
