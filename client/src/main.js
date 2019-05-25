import Vue from 'vue'
import App from './App.vue'
import store from './store/store'
import router from './router'
import Helpers from './helpers'
import VueMaterial from 'vue-material'
import 'vue-material/dist/vue-material.min.css'
import 'vue-material/dist/theme/default-dark.css'

Helpers.install(Vue)
Vue.config.productionTip = false
Vue.use(VueMaterial)

new Vue({
  store,
  router,
  render: h => h(App),
}).$mount('#app')
