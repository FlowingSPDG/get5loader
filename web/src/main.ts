import Vue from 'vue'
import VueI18n from 'vue-i18n'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementUI from 'element-ui'
import api from './utils/api.vue'

// ElementUI...
import 'element-ui/lib/theme-chalk/index.css'
const locale = require('element-ui/lib/locale/lang/ja')
Vue.use(ElementUI, { locale })

// Mixin ./utils/api.vue
Vue.mixin(api)

// i18n
Vue.use(VueI18n)
const translations = require('./translations/translations.json')

const i18n = new VueI18n({
  locale: 'en', // Use English by default
  messages: translations
})

new Vue({
  router,
  store,
  i18n: i18n,
  render: h => h(App)
}).$mount('#app')
