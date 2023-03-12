<template>
  <a-layout class="layout">
    <a-layout-header>
      <div class="title">课程表查询系统</div>
      <a-menu
        v-model:selectedKeys="selectedKeys"
        theme="dark"
        mode="horizontal"
        :style="{ lineHeight: '64px' }"
      >
        <a-menu-item key="1">主页</a-menu-item>
        <a-menu-item key="2">课程表</a-menu-item>
      </a-menu>
      <div class="right">
          <p>登录</p>
        </div>
    </a-layout-header>
    <a-layout-content style="padding: 0 50px">
      <a-breadcrumb style="margin: 16px 0">
        <a-breadcrumb-item>Home</a-breadcrumb-item>

        
      </a-breadcrumb>
      <div :style="{ background: '#fff', padding: '24px', minHeight: '350px' }" >{{containerContext}}</div>
    </a-layout-content>
    <a-layout-footer style="text-align: center">
      
Footer
    </a-layout-footer>
  </a-layout>
</template>
<script lang="ts">
import { defineComponent, ref } from 'vue';
export default defineComponent({
  setup() {
    const selectedKeys = ref<string[]>(['1']);
    return {
      selectedKeys,
      containerContext: ref<string>("No Data")
    };
  },
});


</script>
<style>
.site-layout-content {
  min-height: 280px;
  padding: 24px;
  background: #fff;
}
#components-layout-demo-top .logo {
  float: left;
  width: 120px;
  height: 31px;
  margin: 16px 24px 16px 0;
  background: rgba(255, 255, 255, 0.3);
}
.ant-row-rtl #components-layout-demo-top .logo {
  float: right;
  margin: 16px 0 16px 24px;
}

[data-theme='dark'] .site-layout-content {
  background: #141414;
}
.ant-layout-header{
  display:flex
}
.title{
  color:white;
  font-size: 25px;
  margin-right: 20px;
}
.right{
  width:50px;
  align-items: right;
}
</style>
