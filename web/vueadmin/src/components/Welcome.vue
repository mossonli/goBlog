<template>
  <h1>欢迎界面={{msg}}</h1>
  <el-button>去登录 button </el-button>
  <router-link to="/login">routerlink去登录</router-link>
</template>

<script>
export default {
  name:'welcome',
  props:['msg']
}
</script>



<style >

</style>
