/*
环境配置
*/

// 开发环境
const env = import.meta.env.MODE || 'prod';
console.log("env", env)

const EnvConfig = {
    dev:{
        baseApi:'/api/v1/',
        mockApi:'https://www.fastmock.site/mock/c1c302e8baed9894c48c17e4738c092e/api'
    },
    test:{
        baseApi:'//test.futurefe.com/api',
        mockApi:'https://www.fastmock.site/mock/c1c302e8baed9894c48c17e4738c092e/api'
    },
    prod:{
        baseApi:'//futurefe.com/api',
        mockApi:'https://www.fastmock.site/mock/c1c302e8baed9894c48c17e4738c092e/api'
    }
}
export default {
    env,
    mock:true,
    ...EnvConfig[env]
}