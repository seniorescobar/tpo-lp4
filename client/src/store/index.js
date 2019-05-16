import Vue from 'vue'
import Vuex from 'vuex'
import mockData from './constants/mockData'
// import api from 'client-api'

Vue.use(Vuex)

const state = {}

export const getState = () => {
    return JSON.parse(JSON.stringify(state))
}

export default new Vuex.Store({
    namespaced: true,
    state: getState(),
    getters: {
        constants: () => ({}),
        mockData: () => mockData,
    },
    actions: {},
    mutations: {},
})