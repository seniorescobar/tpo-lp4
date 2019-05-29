import api from 'api-client'
import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)

const state = {
    todos: [{ _id: '12312312', description: 'asdasdasd' }],
    studentId: null,
    isStudentIdSet: false,
}

const getState = () => JSON.parse(JSON.stringify(state))

export default new Vuex.Store({
    state: getState(),
    getters: {},
    actions: {
        async postTodo ({ dispatch }, payload) {
            await api.post('todo', payload)
            return dispatch('fetchAndSetTodos')
        },
        async deleteTodo (_, id) {
            return await api.delete(`todo/${id}`)
        },
        async fetchAndSetTodos ({ dispatch }) {
            const todos = await dispatch('fetchTodos');
            return await dispatch('setTodos', todos);
        },
        async fetchTodos () {
            const { data } = await api.get('todo')
            return data
        },
        setTodos ({ commit }, todos) {
            return commit('SET_TODOS', todos)
        },
        setStudentId ({ commit }, studentId) {
            return commit('SET_STUDENT_ID', studentId)
        },
        setStudentIdSet ({ commit}, isStudentIdSet) {
            return commit('SET_STUDENT_ID_SET', isStudentIdSet)
        }
    },
    mutations: {
        SET_TODOS (state, todos ) {
            state.todos = todos
        },
        SET_STUDENT_ID (state, studentId) {
            state.studentId = studentId
        },
        SET_STUDENT_ID_SET (state, isStudentIdSet) {
            state.isStudentIdSet = isStudentIdSet
        },
    }
})
