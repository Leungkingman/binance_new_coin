import Vue from 'vue'
import App from './App.vue'
import Lockr from 'lockr'
import 'assets/css/common.css'
// 引入element
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
// 全局配置elementui的dialog不能通过点击遮罩层关闭
Vue.use(ElementUI);

// 引入封装的router
import router from '@/router/index'

Vue.config.productionTip = false

window.router = router
window.Lockr = Lockr

new Vue({
  render: h => h(App),
  router,
}).$mount('#app')
