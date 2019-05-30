import Vue from 'vue'
import App from './App.vue'
import store from './store'
import router from './router'
import Helpers from './helpers'
import VueMaterial from 'vue-material'
import 'vue-material/dist/vue-material.min.css'
import 'vue-material/dist/theme/default-dark.css'

Vue.config.productionTip = false
Helpers.install(Vue)
Vue.use(VueMaterial)

import VCalendar from 'v-calendar'
import 'v-calendar/lib/v-calendar.min.css'

Vue.use(VCalendar);

new Vue({
  store,
  router,
  render: h => h(App),
}).$mount('#app')
