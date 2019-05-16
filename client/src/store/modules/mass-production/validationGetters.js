export default {
    getIsCurrentStepValid (state, getters) {
        return getters.getStepValidatorMap[getters.getCurrentStepId]
    },
    getStepValidatorMap (state, getters) {
        return {
            'select_sizes': true,
            'edit_sizes': true,
            'select_experience': true,
            'select_campaign': getters.getIsCampaignStepValid,
            'select_template': getters.getIsTemplateStepValid,
            content: getters.getIsContentStepValid,
            output: true
        }
    },
    getIsCampaignStepValid (state) {
        return state.campaignId !== null
    },
    getIsContentStepValid () {
        return false // TODO:
    },
    getIsSetupValid (state, getters) {
        return Object.keys(getters.getAvailableStepIds).reduce((isValid, stepId) => {
            return isValid && Boolean(getters.getStepValidatorMap[stepId])
        }, true)
    } 
}
