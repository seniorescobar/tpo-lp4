import Vue from 'vue'
import VueRouter from 'vue-router'
import { Login, Register, Main, Home, Food, Events, Bus } from './components/views'

Vue.use(VueRouter)

const router =  new VueRouter({
    mode: 'history',
    base: __dirname,
    routes: [
        { path: '/app', component: Main,
            children: [
                { path: 'home', name: 'home', component: Home },
                { path: 'food', name: 'food', component: Food },
                { path: 'events', name: 'events', component: Events },
                { path: 'bus', name: 'bus', component: Bus },
                { path: '*', redirect: '/app/home' },
            ]
        },
        { path: '/register', name: 'register', component: Register },
        { path: '/login', name: 'login', component: Login },
        { path: '*', redirect: '/app' },
    ]
})

const publicPages = ['login', 'register']

router.beforeEach((to, from, next) => {
    const authRequired = !publicPages.includes(to.name)
    const loggedIn = localStorage.getItem('user')
  
    if (authRequired && !loggedIn) {
        return next({ path: '/login' })
    }

    next()
})

export default router
