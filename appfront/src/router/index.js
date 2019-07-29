import Vue from 'vue'
import Router from 'vue-router'
import login from '@/components/user/Login'
import testcase from '@/components/Testcase'
import graph from '@/components/data/Graph'
import upload from '@/components/tool/Upload'
import format from '@/components/tool/Format'
import App from '@/App'

Vue.use(Router)

export default new Router({
  // mode : 'history',
  routes: [
    {
      path: '/',
      name: 'index',
      redirect: '/testcase',
      component: App,
    },
    {
      path: '/login',
      name: 'login',
      component: login,
    },
    {
      path: '/testcase',
      name: 'testcase',
      meta: {
        requireAuth: true, 
      },
      component: testcase,
    },
    {
      path: '/data/graph',
      name: 'graph',
      meta: {
        requireAuth: true,  
      },
      component: graph,
    },
    {
      path: '/tool/upload',
      name: 'uploadData',
      meta: {
        requireAuth: true,
      },
      component: upload,
    },
    {
      path: '/tool/format',
      name: 'changeFormat',
      meta: {
        requireAuth: true, 
      },
      component: format,
    },
  ]
})
