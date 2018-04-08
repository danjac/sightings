import Vue from 'vue'
import Router from 'vue-router'
import ReportList from '@/components/ReportList'
import ReportDetail from '@/components/ReportDetail'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'ReportList',
      component: ReportList
    },
    {
      path: '/:id',
      name: 'ReportDetail',
      component: ReportDetail
    }
  ]
})
