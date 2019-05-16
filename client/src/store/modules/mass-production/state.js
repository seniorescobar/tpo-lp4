const state = {
    stepIndex: 0,
    dialogStartingPoint: '',
    isEditingSizes: false,
    campaignId: null,
    selectedSizesIds: [],
    customSizes: [],
    experienceId: null,
    globalAttributes: {}, 
    globalImageHotspot: {},
    localImageHotspot: {}, // scale, rotate, left, top
    experienceName: null,
}

// for dev
state.selectedSizesIds = ['mediumRectangle', 'leaderboard', 'mobile']
// state.experienceId = 'brandPackshotWithModel'
// state.stepIndex = 2

export const getState = () => {
    return JSON.parse(JSON.stringify(state))
}
