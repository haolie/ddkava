import Vue from 'vue'
import Router from 'vue-router'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
// import CetcUI from 'cetc-ui'
import MessageBox from 'element-ui/packages/message-box'
import Message from 'element-ui/packages/message'
import api from '@/api'
import http from '@/http'
import lib from '@/lib'
import bus from '@/lib/bus'
import drag from '@/lib/drag'
import directives from '@/directives'

// Object.defineProperty(Vue.prototype, '$bus', {value: bus})
Vue.use(Router)
Vue.use(ElementUI)
// Vue.use(CetcUI)
Vue.prototype.$alert = MessageBox.alert
Vue.prototype.$message = Message
Vue.prototype.$confirm = MessageBox.confirm
Vue.use(api)
Vue.use(http)
Vue.use(lib)
Vue.use(bus)
Vue.use(drag)

directives.init(Vue)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'WebClient',
      component: () => import('@/views/mainView/mainView')
    },
    // {
    //   path: '/netProfit',
    //   name: '净利润',
    //   component: () => import('@/views/netProfit/netProfit')
    // }
  ]
})
