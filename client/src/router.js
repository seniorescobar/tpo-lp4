import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from './components/views/Home'
import Food from './components/views/Food'
import Events from './components/views/Events'
import Bus from './components/views/Bus'

Vue.use(VueRouter)

export default new VueRouter({
    mode: 'history',
    base: __dirname,
    routes: [
        { path: '/', name: 'home', component: Home },
        { path: '/home', redirect: '/' },
        { path: '/food', name: 'food', component: Food },
        { path: '/events', name: 'events', component: Events },
        { path: '/bus', name: 'bus', component: Bus },
    ]
})