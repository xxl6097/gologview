import { createApp } from 'vue'
import App from './App.vue'

import ElementPlus from "element-plus";
// import "element-plus/theme-chalk/dark/css-vars.css";
// 2 eui 样式
import "element-plus/dist/index.css";

createApp(App).use(ElementPlus).mount('#app')
