import Vue from 'vue'
import Router from 'vue-router'
import login from '@/components/user/Login'
import sider from '@/components/Sider'
import testcase from '@/components/testcase/Testcase'
import graph from '@/components/data/Graph'
import upload from '@/components/tool/Upload'
import format from '@/components/tool/Format'
import App from '../App'

Vue.use(Router)

export default new Router({
  // mode : 'history',
  routes: [
    // 登录
    {
      path: '/login',
      name: 'login',
      component: login,
    },
    // 主界面
    {
      path: '/',
      name: 'home',
      meta: {
        requireAuth: true, 
      },
      component: sider,
      children: [
        {
          path: 'testcase',
          name: 'testcase',
          component: testcase,
        },
        {
          path: 'data/graph',
          name: 'graph',
          component: graph,
        },
        {
          path: 'tool/upload',
          name: 'upload',
          component: upload,
        },
        {
          path: 'tool/format',
          name: 'format',
          component: format,
        },
      ]
    }, 
  ]
})
