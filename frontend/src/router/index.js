import Vue from 'vue'
import Router from 'vue-router'
import ReportList from '@/components/ReportList'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'ReportList',
      component: ReportList
    }
  ]
})
