<template>
  <a-layout id="components-layout-demo-fixed-sider">
    <a-layout-sider 
    style="min-height: 100vh"
    collapsible
    v-model="collapsed">
      <div v-if="!collapsed">
        <svg width="3em" height="3em" viewBox="0 0 33 28" style="margin: 15 0 0 15;"><path d="M19.101 22.612l1.031-1.802h-.005L27.92 7.194h-.008l1.527-2.678h3.077L20.118 26.181 19.12 28l-9.03-16.446-.088-.166A7.881 7.881 0 0 1 9.032 7.6C9.032 3.41 12.274 0 16.26 0c3.923 0 7.127 3.305 7.225 7.406h-2.552c-.097-2.622-2.156-4.724-4.673-4.724-2.578 0-4.676 2.206-4.676 4.917 0 .858.217 1.704.63 2.452l.054.105 3.496 6.367.154.29 3.184 5.8zm-4.65 2.717L12.924 28 0 5.42h3.074l1.527 2.675.157.289 7.62 13.315.16.289 1.914 3.341zm2.687-10.006l2.031-3.554h-4.061l-1.56-2.737h9.936a7.774 7.774 0 0 1-.794 2.033l-.098.174.005.004-3.897 6.822-1.562-2.742z" fill="#FFF" fill-rule="evenodd"></path></svg>
        <h2 style="color:white;position:absolute;left:70px;top:23px">测试平台</h2>
      </div>
      <a-menu theme="dark" mode="inline" :defaultSelectedKeys="['testcaseML']" :defaultOpenKeys="['testcase']">
        <a-sub-menu key="testcase">
          <span slot="title"><a-icon type="file" /><span>测试用例</span></span>
          <a-menu-item key="testcaseML" @click="gotoTestcaseML">机器学习平台</a-menu-item>
          <a-menu-item key="testcaseSJ" @click="gotoTestcaseSJ">数据治理平台</a-menu-item>
        </a-sub-menu>
        <a-sub-menu key="data">
          <span slot="title"><a-icon type="appstore-o" /><span>数据构造</span></span>
          <a-menu-item key="dataGraph"  @click="gotoDataGraph">网络图</a-menu-item>
        </a-sub-menu>
        <a-sub-menu key="tool">
          <span slot="title"><a-icon type="tool" /><span>其他工具</span></span>
          <a-menu-item key="toolUpload"  @click="gotoToolUpload">上传数据</a-menu-item>
          <a-menu-item key="toolFormat"  @click="gotoToolFormat">格式转换</a-menu-item>
        </a-sub-menu>
      </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-layout-header :style="{ background: '#fff', padding: 0 }">
        <a-tooltip placement="top" title="注销">
          <a-icon id="logout" type="logout" @click="logout" />
        </a-tooltip>
        <a-tooltip placement="top" :title="username">
          <a-icon type="user" id="user" />
        </a-tooltip>
      </a-layout-header>
      <a-layout-content>
        <div :style="{ margin: '5px 0px 0px 0px', overflow: 'initial', padding: '0px' }">
          <router-view :key="this.$route.fullPath"></router-view>
        </div>
      </a-layout-content>
      <a-layout-footer style="text-align: center">
        ©2019 Created by Xiebo Zhou
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>
<script>
import axios from 'axios';
export default {
  name: 'Sider',
  data () {
    return {
      username: localStorage.getItem('username'),
      collapsed: false,
    }
  },
  mounted () {
    if (localStorage.getItem('Authorization')) {
      this.$router.push({name: 'testcase-ml'});
    } else {
      this.$router.push({name: 'login'});
    }
  },
  methods: {
    gotoTestcaseML () {
      this.$router.push({name: 'testcase-ml'});
    },
    gotoTestcaseSJ () {
      this.$router.push({name: 'testcase-sj'});
    },
    gotoDataGraph () {
      this.$router.push({name: 'graph'});
    },
    gotoToolUpload () {
      this.$router.push({name: 'upload'});
    },
    gotoToolFormat () {
      this.$router.push({name: 'format'});
    },
    logout () {
      localStorage.removeItem('Authorization');
      this.$router.push({name: 'login'});
      this.$message.success("成功注销");   
    },
  },
}
</script>
<style>
#components-layout-demo-fixed-sider .logo {
  height: 32px;
  background: rgba(255,255,255,.2);
  margin: 16px;
}

#logout, #user {
  float: right;
  font-size: 18px;
  line-height: 64px;
  padding: 0 24px;
  cursor: pointer;
  transition: color .3s;
}
#logout:hover, #user:hover {
  color: #1890ff;
}
/*
a {
  text-decoration: none;
  color: white;
}

 .router-link-active {
   text-decoration: none;
   color: white;
 }
*/
</style>
