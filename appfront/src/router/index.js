import Vue from 'vue'
import Router from 'vue-router'
import login from '@/components/user/Login'
import sider from '@/components/Sider'
import testcase from '@/components/testcase/Testcase'
import graph from '@/components/data/Graph'
import upload from '@/components/tool/Upload'
import format from '@/components/tool/Format'
import App from '@/App'

Vue.use(Router)

export default new Router({
  // mode : 'history',
  routes: [
    {
      path: '/login',
      name: 'login',
      component: login,
    },
    {
      path: '/',
      name: 'index',
      redirect: {name: 'login'},
      component: App,
    },
    {
      path: '/main',
      component: sider,
      children: [
        {
          path: 'testcase',
          name: 'testcase',
          meta: {
            requireAuth: true, 
          },
          component: testcase,
        },
        {
          path: 'data/graph',
          name: 'graph',
          meta: {
            requireAuth: true,  
          },
          component: graph,
        },
        {
          path: 'tool/upload',
          name: 'upload',
          meta: {
            requireAuth: true,
          },
          component: upload,
        },
        {
          path: 'tool/format',
          name: 'format',
          meta: {
            requireAuth: true, 
          },
          component: format,
        },
      ]
    }, 
  ]
})
