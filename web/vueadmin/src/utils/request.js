/*
axios 二次封装
*/
import axios from "axios"
import config from "../config"
import router from './../router'
import {ElMessage} from 'element-plus'


const TOKEN_INVALID = 'Token认证失败'
const NETWORK_ERROR = '网络异常请稍后重试'

// 创建axios实例对象，添加全局配置
const service = axios.create({
    baseURL: config.baseApi,
    timeout:8000, // 8s
})

// 请求拦截、
service.interceptors.request.use((req)=>{
    // todo
    const headers = req.headers
    if (!headers.Authorization) headers.Authorization = 'Bear Mosson'
    return req
})


// 响应拦截、
service.interceptors.response.use((resp)=>{
    // todo
    const {code, data, msg} = resp.data
    if (code === 200){
        return data
    }else if (code === 40001){
        ElMessage.error(TOKEN_INVALID)
        setTimeout(()=>{
            router.push('/user/login')
        },15000)

        // 抛出异常
        return Promise.reject(TOKEN_INVALID)
    }else{
        ElMessage.error(msg || NETWORK_ERROR)
        return Promise.reject(TOKEN_INVALID)
    }    
})


/*
请求的核心函数
options 请求的配置信息
*/
function request(options){
    options.method = options.method || 'get'
    if (options.method.toLower === 'get'){
        options.params = options.data
    }
    if(typeof options.mock != 'undefined'){
        config.mock = options.mock
    }

    // 配置生产
    if(config.env === 'prod'){
        service.defaults.baseURL = config.baseApi
    }else{
        // config.mock ? config.mockApi:config.baseApi
        // config.mock 为true的时候使用 config.mockApi 否则就使用 config.baseApi
        service.defaults.baseURL = config.mock ? config.mockApi:config.baseApi
    }
    return service(options)
}
['get', 'post', 'delete', 'put', 'patch'].forEach((item)=>{
    request[item] = (url, data, options)=>{
        return request({
            url, 
            data,
            method:item,
            ...options  // 解构 是什么意思
        })
    }
})

export default request