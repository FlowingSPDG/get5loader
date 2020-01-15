import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementUI from 'element-ui'
import api from './utils/api.vue'

// ElementUI...
import 'element-ui/lib/theme-chalk/index.css'
const locale = require('element-ui/lib/locale/lang/ja')
Vue.use(ElementUI, { locale })
Vue.config.productionTip = false

// Mixin ./utils/api.vue
Vue.mixin(api)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
