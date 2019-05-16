import Vue from 'vue'
import { getState } from './state'

const getResetState = (currentState, excludedStateKeys) => {
    const resetState = getState()

    for (const key of excludedStateKeys) {
        resetState[key] = currentState[key]
    }

    return resetState
}

export default {
    RESET_STATE (state, excludedStateKeys = []) {
        state = { ...state, ...getResetState(state, excludedStateKeys)}
    },
    SET_STEP_INDEX (state, index) {
        state.stepIndex = index
    },
    SET_DIALOG_STARTING_POINT (state, dialogStartingPoint) {
        state.dialogStartingPoint = dialogStartingPoint
    },
    SET_CAMPAIGN (state, id) {
        state.campaignId = id
    },
    SET_SELECTED_SIZES_IDS (state, selectedSizesIds) {
        state.selectedSizesIds = selectedSizesIds
    },
    ADD_CUSTOM_SIZE (state, customSize) {
        state.customSizes.push(customSize)
    },
    SET_GLOBAL_ATTRIBUTE (state, { id, value }) {
        Vue.set(state.globalAttributes, id, value)
    },
    SET_EXPERIENCE (state, experienceId) {
        state.experienceId = experienceId
    },
    SET_EXPERIENCE_NAME (state, experienceName) {
        state.experienceName = experienceName
    },
    SET_GLOBAL_IMAGE_HOTSPOT (state, { id, x, y, r}) {
        Vue.set(state.globalImageHotspot, id, { x, y, r })
    },
    SET_LOCAL_IMAGE_HOTSPOT_SCALE (state, { instanceId, assetId, scale }) {
        Vue.set(state.localImageHotspot, instanceId, { [assetId]: { ...(state.localImageHotspot[instanceId] && state.localImageHotspot[instanceId][assetId]), ...{ scale } } } )
    },
    SET_LOCAL_IMAGE_HOTSPOT_ANGLE (state, { instanceId, assetId, angle }) {
        Vue.set(state.localImageHotspot, instanceId, { [assetId]: { ...(state.localImageHotspot[instanceId] && state.localImageHotspot[instanceId][assetId]), ...{ angle } } } )
    },
}
