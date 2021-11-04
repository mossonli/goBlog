/*
api管理
*/

import request from './../utils/request'
export default{
    login(params){
        return request({
            url:'/user/login',
            data:params,
            method:'post',
            mock:true
        })
    }
}