## 通过vite创建vue项目
```bash
npm init @vitejs/app projectnamw
cnpm init @vitejs/app projectnamw
yarn create @vitejs/app projectnamw
```
## 选择项目类型
```text
* vue
  vue-ts
```

## 项目结构

```text
 ~/Documents/MossonProject/goBlog/web/vueadmin  webdev ±✚  tree -L 1
.
├── index.html
├── package.json
├── public
├── src
├── vite.config.js
└── yarn.lock
```

### 初始化包

```bash
 ~/Documents/MossonProject/goBlog/web/vueadmin  webdev  yarn     // 或者使用 yarn install 
 .
├── index.html
├── node_modules
├── package.json
├── public
├── src
├── vite.config.js
└── yarn.lock

生成node包

```

## 启动

```bash
 ~/Documents/MossonProject/goBlog/web/vueadmin   webdev ±✚  yarn dev
yarn run v1.22.11
warning package.json: No license field
$ vite
Pre-bundling dependencies:
  vue
(this will be run only when your dependencies or config have changed)

  vite v2.6.13 dev server running at:

  > Local: http://localhost:3000/
  > Network: use `--host` to expose

```

## 安装插件

```bash
# 安装项目生产依赖
yarn add vue-router@next vuex@next element-plus axios -S
# 安装项目开发依赖
yarn add sass -D
```

## 目录结构

```text
├── index.html
├── dist				项目打包之后的地方
├── node_modules
├── package.json
├── public
├── src
│   ├── App.vue
│   ├── api				接口
│   ├── assets		静态文件（图片、资源）
│   ├── components	组件
│   ├── config		配置文件
│   ├── main.js
│   ├── router		路由
│   ├── utils
│   └── views
├── vite.config.js
└── yarn.lock
```

## 选择UI

```text
https://www.antdv.com/docs/vue/introduce-cn/
> 安装
yarn add ant-design-vue --dev

 yarn add babel-plugin-import --dev
```

## 修改配置

```js
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  server:{
    port:8181
  },
  plugins: [vue()]
})
```

