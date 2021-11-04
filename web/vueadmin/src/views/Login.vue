<template>
  <div class="login-wrapper">
    <div class="modal">
      <!-- status-icon 框右侧的勾/×  ref 操作原生dom的句柄-->
      <el-form ref="userForm" :model="user" status-icon :rules="rules">
        <div class="title">Luffy</div>
        <el-form-item prop="userName">
          <!-- prop 绑定校验字段 -->
          <el-input type="text" prefix-icon="el-icon-user" v-model="user.username"/>
        </el-form-item>
        <el-form-item prop="userPwd">
          <el-input type="password" prefix-icon="el-icon-view" v-model="user.password"/>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" class="btn-login" @click="login">登录</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
// import Welcome from '../views/Welcome.vue'
export default {
  name: 'login',
  data(){
    return {
      user:{
        username:'',
        password:''
      },
      rules:{
        username:[
          {
            required:true, message:"请输入用户名", trigger:"blur"
          }
        ],
        password:[
          {
            required:true, message:"请输入密码", trigger:"blur"
          }
        ]
      }
    }
  },
  mounted () {
    // 方式1 
    // this.$request({
    //   method: 'get',
    //   url:'/user/login',
    //   data: {
    //     name:'mosson'
    //   }
    // }).then((resp)=>{
    //   console.log(resp)
    // })

    // 方式2 需要在request 里面讲请求方法加到request里面
    // this.$request.get("/user/login", { name: 'mosson' }, { mock: true }).then((resp) => {
    //   console.log(resp)
    // })

  },
  methods: {
    login(){
      this.$refs.userForm.validate((valid)=>{
        if(valid){
          this.$api.login(this.user).then((resp)=>{
            this.$store.commit('saveUserInfo', resp)
            this.$router.push('/welcome')
          })
        }else{
          return false
        }
      })
    },
    goHome () {
      this.$router.push('/welcome')
    }
  }
}
</script>



<style lang="scss">
.login-wrapper {
  display:flex;
  justify-content: center;
  align-items: center;
  background-color: #f9fcff;
  width: 100vw;
  height: 100vh;
  .modal{
    width:500px;
    padding: 50px;
    background-color: #fff;
    border-radius: 4px;
    box-shadow: 0px 0px 10px 3px #c7c9cb4d;
    .title{
      font-size: 50px;
      line-height: 1.5;
      text-align: center;
      margin-bottom: 30px;
    }
    .btn-login{

    }
  }
}
</style>
