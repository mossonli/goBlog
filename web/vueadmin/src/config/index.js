/*
环境配置
*/

// 开发环境
const env = import.meta.env.MODE || 'prod';
console.log("env", env)

const EnvConfig = {
    dev:{
        baseApi:'http://127.0.0.1:3001/api/v1/',
        mockApi:'https://www.fastmock.site/mock/f7a1268debd84cb716b2e8af635f88f0/api'
    },
    test:{
        baseApi:'//test.futurefe.com/api',
        mockApi:'https://www.fastmock.site/mock/f7a1268debd84cb716b2e8af635f88f0/api'
    },
    prod:{
        baseApi:'//futurefe.com/api',
        mockApi:'https://www.fastmock.site/mock/f7a1268debd84cb716b2e8af635f88f0/api'
    }
}
export default {
    env,
    mock:true,
    namespace:'manager',
    ...EnvConfig[env]
}