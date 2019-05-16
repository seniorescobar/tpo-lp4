import { getState } from './state'
import actions from './actions'
import getters from './getters'
import validationGetters from './validationGetters'
import mutations from './mutations'

const store = {
    namespaced: true,
    state: getState(),
    getters: { ...getters, ...validationGetters },
    mutations: mutations,
    actions: actions,
}

export default store
