// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import router from './router'
import store from '@/store'
import App from './App'
import Echarts from 'echarts'
import echartsTheme from './lib/echartTheme'


Echarts.registerTheme('macarons', echartsTheme)
Vue.config.productionTip = false
Vue.prototype.$echarts = Echarts
/* eslint-disable no-new */
window.$Vue = new Vue({
  el: '#app',
  router,
  store,
  components: {App},
  template: '<App/>'
})
