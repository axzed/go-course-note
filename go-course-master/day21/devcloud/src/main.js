import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

// UI 组件
import Element from 'element-ui'
import "element-ui/lib/theme-chalk/index.css"

// 加载 svg icons
import './icons'

// 加载全局样式
import './styles/index.scss'

// 加载全局指令
import './directives'

// 加载全局过滤器
import './filters'

Vue.config.productionTip = false;
Vue.use(Element, {
  size: 'mini'
})

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount("#app");
