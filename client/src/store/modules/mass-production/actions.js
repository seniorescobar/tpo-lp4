// import api from 'celtra-api'

export default {
    initialize ({ dispatch }) {
        dispatch('resetState')
    },
    resetState ({ commit }) {
        commit('RESET_STATE')
    },
    setDialogStartingPoint({ commit }, dialogStartingPoint) {
        commit('SET_DIALOG_STARTING_POINT', dialogStartingPoint)
    },
    nextStep ({ state, commit, getters }) {
        if (getters.getIsCurrentStepValid && getters.getNextStepId) {
            commit('SET_STEP_INDEX', state.stepIndex + 1)
        }
    },
    previousStep ({ state, commit }) {
        if (state.stepIndex - 1 >= 0) {
            commit('SET_STEP_INDEX', state.stepIndex - 1)
        }
    },
    setCampaign ({ commit }, id) {
        commit('SET_CAMPAIGN', id)
    },
    setSelectedSizesIds ({ commit }, selectedSizesIds) {
        commit('SET_SELECTED_SIZES_IDS', selectedSizesIds)
    },
    setGlobalAttribute ({ commit }, payload) {
        commit('SET_GLOBAL_ATTRIBUTE', payload)
    },
    setGlobalImageHotspot ({ commit }, payload) {
        commit('SET_GLOBAL_IMAGE_HOTSPOT', payload)
    },
    setLocalImageHotspotScale ({ commit }, payload) {
        commit('SET_LOCAL_IMAGE_HOTSPOT_SCALE', payload)
    },
    setLocalImageHotspotAngle ({ commit }, payload) {
        commit('SET_LOCAL_IMAGE_HOTSPOT_ANGLE', payload)
    },
    addCustomSize ({ commit }, customSize) {
        commit('ADD_CUSTOM_SIZE', customSize)
    },
    setExperience ({ commit }, experienceId) {
        commit('SET_EXPERIENCE', experienceId)
    },
    setExperienceName ({ commit }, experienceName) {
        commit('SET_EXPERIENCE_NAME', experienceName)
    },
}
