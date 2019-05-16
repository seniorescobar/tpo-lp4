import { dialogPaths, goals, sizes, outputTypes, experiences, instances, campaigns, feeds } from './store_constants'

// retrun { top, left, width, height }
const calculatePosition = (globalAttributes, globalImageHotspot, localImageHotspot, instance, unit) => {
    const parameters = globalImageHotspot[unit.id]
    const localParameters = localImageHotspot[instance.sizeId] && localImageHotspot[instance.sizeId][unit.id]

    if (localParameters) {
        const transform = `scale(${1 + localParameters.scale / 5 || 0}) rotate(${localParameters.angle || 0}deg)`
        return { ...unit.position, ...{ transform } } 
    } else if (parameters) {
        const isWidthOverflowing = instance.size.width / globalAttributes[unit.id].width < instance.size.height / globalAttributes[unit.id].height
        const factor = Math.max(instance.size.width / globalAttributes[unit.id].width, instance.size.height / globalAttributes[unit.id].height)

        if (isWidthOverflowing) {
            const width = globalAttributes[unit.id].width * factor
            const offsetToCenter = (width - instance.size.width) / 2
            let left = 0 - offsetToCenter
            const span = { from: 0 + offsetToCenter, to: width - offsetToCenter }
            const hotspotReach = Math.min(instance.size.height, width) * parameters.r / 2 / 10
            const maxHotspotCoordinate = Math.min(width * parameters.x + hotspotReach, width)
            const minHotspotCoordinate = Math.max(width * parameters.x - hotspotReach, 0)
            
            const minDelta = span.from - minHotspotCoordinate
            if (minDelta > 0)
                left += minDelta

            const maxDelta = span.to - maxHotspotCoordinate
            if (maxDelta < 0)
                left += maxDelta

            return { top: '0', height: '100%', left: `${left}px`, width: `${width}px` }
        } else {
            const height = globalAttributes[unit.id].height * factor
            const offsetToCenter = (height - instance.size.height) / 2
            let top = 0 - offsetToCenter
            const span = { from: 0 + offsetToCenter, to: height - offsetToCenter }
            const hotspotReach = Math.min(instance.size.width, height) * parameters.r / 2 / 10
            const maxHotspotCoordinate = Math.min(height * parameters.y + hotspotReach, height)
            const minHotspotCoordinate = Math.min(height * parameters.y - hotspotReach, 0)

            const minDelta = span.from - minHotspotCoordinate
            if (minDelta > 0)
                top += minDelta

            const maxDelta = span.to - maxHotspotCoordinate
            if (maxDelta < 0)
                top += maxDelta

            return { left: '0', width: '100%', top: `${top}px`, height: `${height}px` }
        }
    }

    return unit.position
}

export default {
    getAvailableStepIds (state) {
        return dialogPaths.find(path =>
            path.condition.isEditingSizes === state.isEditingSizes
        ).steps
    },
    getAvailableSteps (state, getters) {
        const labelsByStep = {
            'select_sizes': {
                passive: 'SELECT SIZES',
                active: 'SIZES',
            },
            'edit_sizes': {
                passive: 'EDIT SIZES',
                active: 'SIZES',
            },
            'select_experience': {
                passive: 'SELECT EXPERIENCE',
                active: 'EXPERIENCE',
            },
            content: {
                passive: 'CONTENT',
                active: 'CONTENT',
            },
            output: {
                passive: 'OUTPUT',
                active: 'OUTPUT',
            },
            finish: {
                passive: 'FINISH',
                active: 'FINISH',
            },
        }

        return getters.getAvailableStepIds.map(stepId => ({
            id: stepId,
            passiveLabel: labelsByStep[stepId].passive,
            activeLabel: labelsByStep[stepId].active,
        }))
    },
    getCurrentStepId (state, getters) {
        return getters.getAvailableStepIds[state.stepIndex]
    },
    getNextStepId (state, getters) {
        return getters.getAvailableStepIds[state.stepIndex + 1]
    },
    getAvailableSizes (state) {
        // return [...state.customSizes, ...sizes.filter(size => instances.find(instance => instance.sizeId === size.id))] // for dev
        return [...state.customSizes, ...sizes]
    },
    getSelectedSizes (state, getters) {
        return state.selectedSizesIds.map(selectedSize => getters.getAvailableSizes.find(size => selectedSize === size.id))
    },
    getSelectedInstances (state) {
        return instances.filter(instance => state.selectedSizesIds.includes(instance.sizeId))
    },
    getPopulatedInstances (state, getters) {
        return getters.getSelectedInstances.map(instance => {
            const units = instance.units.map(unit => ({ ...unit, ...{ value: state.globalAttributes[unit.id] || unit.value }, ...{ position: calculatePosition(state.globalAttributes, state.globalImageHotspot, state.localImageHotspot, instance, unit) } }))

            return { ...instance, ...{ units } }
        })
    },
    getCampaigns() {
        return campaigns
    },
    getFeeds () {
        return feeds
    },
    getGoals () {
        return goals
    },
    getOutputTypes () {
        return outputTypes
    },
    getExperiences () {
        return experiences
    },
}
