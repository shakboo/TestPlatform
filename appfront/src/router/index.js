import Vue from 'vue'
import Router from 'vue-router'
import testcase from '@/components/Testcase'
import graph from '@/components/data/Graph'
import upload from '@/components/tool/Upload'
import format from '@/components/tool/Format'
import App from '@/App'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'index',
      redirect: '/testcase',
      component: App,
    },
    {
      path: '/testcase',
      name: 'testcase',
      component: testcase,
    },
    {
      path: '/data/graph',
      name: 'graph',
      component: graph
    },
    {
      path: '/tool/upload',
      name: 'uploadData',
      component: upload
    },
    {
      path: '/tool/format',
      name: 'changeFormat',
      component: format
    },
  ]
})
