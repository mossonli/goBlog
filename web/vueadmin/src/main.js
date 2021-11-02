import { createApp } from 'vue'
import App from './App.vue'

// 引入element-plus https://element-plus.gitee.io/zh-CN/guide/quickstart.html
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

// 跨域
// import axios from 'axios'
import request from './utils/request' // 在request里面封装了 axios 
import config from './config'

// 引入路由
import router from './router'




console.log("环境变==>", import.meta.env)

const app = createApp(App)

app.config.globalProperties.$request = request

// axios.get(config.mockApi+'/v1/user/login').then((res)=>{
//     console.log(res)
// })

app.use(router)

app.use(ElementPlus)

app.mount('#app')
