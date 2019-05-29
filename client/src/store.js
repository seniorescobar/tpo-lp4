import api from 'api-client'

const state = {
    todos: [],
    studentId: null,
    isStudentIdSet: false,
}

const getState = () => JSON.parse(JSON.stringify(state))

export default {
    state: getState(),
    getters: {
    },
    actions: {
        fetchAndSetTodos ({ commit, dispatch }) {
            return dispatch('fetchTodos')
                .then(todos => dispatch('setTodos', todos))
        },
        fetchTodos () {
            return api
                .get('todos')
                .then(res => res.data)
        },
        postTodo ({ dispatch }, payload) {
            return api
                .post('todo', payload)
                .then(() => dispatch('fetchAndSetTodos'))
        },
        setTodos ({ commit }, todos) {
            return commit('SET_TODOS', todos)
        },
        setStudentId ({ commit }, studentId) {
            return commit('SET_STUDENT_ID', studentId)
        },
        setStudentIdSet ({ commit}, isStudentIdSet) {
            return commit('SET_STUDENT_ID_SET')
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
}
