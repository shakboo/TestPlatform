import Vue from 'vue';
import Vuex from 'vuex';
Vue.use(Vuex);
 
export default new Vuex.Store({
 
  state: {
    // 存储token
    Authorization: localStorage.getItem('Authorization') ? localStorage.getItem('Authorization') : '',
    username: localStorage.getItem('username') ? localStorage.getItem('username') : '',
    role: localStorage.getItem('role') ? localStorage.getItem('role') : ''
  },

  mutations: {
    // 修改token，并将token存入localStorage
    changeLogin (state,user) {
        state.Authorization = user.Authorization;
        state.username = user.username;
        localStorage.setItem('Authorization', user.Authorization);
        localStorage.setItem('username', user.username);
        localStorage.setItem('role', user.role)
    }
  }
});