<template>
  <a-layout id="components-layout-demo-fixed-sider">
    <a-layout-sider :style="{ overflow: 'auto', height: '100vh', position: 'fixed', left: 0 }">
      <div>
        <svg width="3em" height="3em" viewBox="0 0 33 28" style="margin: 15 0 0 15;"><path d="M19.101 22.612l1.031-1.802h-.005L27.92 7.194h-.008l1.527-2.678h3.077L20.118 26.181 19.12 28l-9.03-16.446-.088-.166A7.881 7.881 0 0 1 9.032 7.6C9.032 3.41 12.274 0 16.26 0c3.923 0 7.127 3.305 7.225 7.406h-2.552c-.097-2.622-2.156-4.724-4.673-4.724-2.578 0-4.676 2.206-4.676 4.917 0 .858.217 1.704.63 2.452l.054.105 3.496 6.367.154.29 3.184 5.8zm-4.65 2.717L12.924 28 0 5.42h3.074l1.527 2.675.157.289 7.62 13.315.16.289 1.914 3.341zm2.687-10.006l2.031-3.554h-4.061l-1.56-2.737h9.936a7.774 7.774 0 0 1-.794 2.033l-.098.174.005.004-3.897 6.822-1.562-2.742z" fill="#FFF" fill-rule="evenodd"></path></svg>
        <h2 style="color:white;position:absolute;left:70px;top:23px">测试平台</h2>
      </div>
      <a-menu theme="dark" mode="inline" :defaultSelectedKeys="['4']">
        <a-menu-item key="testcase" @click="gotoTestcase">
          <a-icon type="appstore-o" />
          <span class="nav-text">测试用例</span>
        </a-menu-item>
        <a-sub-menu key="data">
          <span slot="title"><a-icon type="appstore-o" />数据构造</span>
          <a-menu-item key="dataGraph"  @click="gotoDataGraph">网络图</a-menu-item>
        </a-sub-menu>
        <a-sub-menu key="tool">
          <span slot="title"><a-icon type="appstore-o" />其他工具</span>
          <a-menu-item key="toolUpload"  @click="gotoToolUpload">上传数据</a-menu-item>
          <a-menu-item key="toolFormat"  @click="gotoToolFormat">格式转换</a-menu-item>
        </a-sub-menu>
      </a-menu>
    </a-layout-sider>
    <a-layout :style="{ marginLeft: '200px' }">
      <a-layout-header :style="{ background: '#fff', padding: 0 }" />
      <a-layout-content :style="{ margin: '24px 16px 0', overflow: 'initial' }">
        <router-view></router-view>
      </a-layout-content>
      <a-layout-footer :style="{ textAlign: 'center' }">
        ©2019 Created by Xiebo Zhou
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>
<script>
export default {
  name: 'App',
  methods: {
    gotoTestcase () {
      this.$router.push('/testcase');
    },
    gotoDataGraph () {
      this.$router.push('/data/graph');
    },
    gotoToolUpload() {
      this.$router.push('/tool/upload');
    },
    gotoToolFormat() {
      this.$router.push('/tool/format');
    }
  },
}
</script>
<style>
#components-layout-demo-fixed-sider .logo {
  height: 32px;
  background: rgba(255,255,255,.2);
  margin: 16px;
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
