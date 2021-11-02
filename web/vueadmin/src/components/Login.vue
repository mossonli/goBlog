<script>
import Welcome from './Welcome.vue'
export default {
  name:'login',
  components:{Welcome},
  mounted(){
    // 方式1 
    this.$request({
      method: 'get',
      url:'/user/login',
      data: {
        name:'mosson'
      }
    }).then((resp)=>{
      console.log(resp)
    })
    // 方式2 需要在request 里面讲请求方法加到request里面
    this.$request.get("/user/login", {name:'mosson'},{mock:true}).then((resp)=>{
      console.log(resp)
    })

  },
  methods:{
    goHome(){
      this.$router.push('/welcome')
    }
  }
}
</script>

<template>
  <h1>登录界面</h1>
  <el-button @click="goHome">回到首页</el-button>
</template>

<style scoped>
</style>
