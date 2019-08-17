// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'
import Antd from 'ant-design-vue'
import 'ant-design-vue/dist/antd.css'
import axios from 'axios'

Vue.config.productionTip = false
Vue.use(Antd)

//异步请求前在header里加入token
axios.interceptors.request.use(
  config => {
    if (config.url === '/auth' || config.url === '/register'){  // 如果是登录和注册操作，则不需要携带header里面的token
    } else {
      if (localStorage.getItem('Authorization')) {
        config.headers.Authorization = localStorage.getItem('Authorization');
      }
    }
    return config;
  },
  error => {
    return Promise.reject(error);
  }
)

// 异步请求后，判断token是否过期
axios.interceptors.response.use(
  response => {
    if (response) {
      switch (response.data.code) {
        case 400:    // 请求头中没有token
          localStorage.removeItem('Authorization');
          localStorage.removeItem('username');
          router.replace({
            path: '/login',
            // query: {redirect: router.currentRoute.fullPath}
          })
        case 401:    // token验证失败
          localStorage.removeItem('Authorization');
          localStorage.removeItem('username');
          router.replace({
            path: '/login',
            // query: {redirect: router.currentRoute.fullPath}
          })
        case 403:
          localStorage.removeItem('Authorization');
          localStorage.removeItem('username');
          router.replace({
            path: '/login',
            // query: {redirect: router.currentRoute.fullPath}
          })
      }
    }
    return response;
  },
  error => {
    return Promise.reject(error.response.data); // 接口返回的错误信息
  }
)

// 异步请求前判断请求的连接是否需要token
router.beforeEach((to, from, next) => {
  if (to.matched.some(res => res.meta.requireAuth)) {
    if (localStorage.getItem('Authorization')) {
      next();
    } else {
      next({
        path: '/login',
        query: {redirect: to.fullPath},
      })
    }
  } else {
    next();
  }
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  render: h => h(App),
  template: '<App/>'
})