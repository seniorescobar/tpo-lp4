<template>
    <mass-production-dialog
        :steps="steps"
        :step-id="stepId"
        :is-valid="isSetupValid"
        @next-step="nextStep"
        @previous-step="previousStep"
        @closed="$emit('close')"
    >
        <transition :name="stepAnimation" mode="out-in">
            <mass-production-campaign-selection-step   v-if="stepId === 'select_campaign'"/> <!-- ignore for now -->
            <mass-production-sizes-selection-step      v-if="stepId === 'select_sizes'" />
            <mass-production-experience-selection-step v-if="stepId === 'select_experience'"/>
            <mass-production-template-selection-step   v-if="stepId === 'select_template'"/> <!-- ignore for now -->
            <mass-production-content-step              v-if="stepId === 'content'"/>
            <mass-production-output-selection-step     v-if="stepId === 'output'"/>
            <mass-production-finish-step               v-if="stepId === 'finish'" @close="$emit('close')"/>
        </transition>
    </mass-production-dialog>
</template>

<script>
import { Dialog } from 'design-system'
import MassProductionSizesSelectionStep from './MassProductionSizesSelectionStep.vue'
import MassProductionOutputSelectionStep from './MassProductionOutputSelectionStep.vue'
import MassProductionExperienceSelectionStep from './MassProductionExperienceSelectionStep.vue'
import MassProductionCampaignSelectionStep from './MassProductionCampaignSelectionStep.vue'
import MassProductionTemplateSelectionStep from './MassProductionTemplateSelectionStep.vue'
import MassProductionContentStep from './MassProductionContentStep.vue'
import MassProductionFinishStep from './MassProductionFinishStep.vue'
import { mapGetters, mapActions, mapState } from 'vuex'

export default {
    name: 'MassProduction',
    components: {
        MassProductionSizesSelectionStep,
        MassProductionOutputSelectionStep,
        MassProductionExperienceSelectionStep,
        MassProductionCampaignSelectionStep,
        MassProductionTemplateSelectionStep,
        MassProductionContentStep,
        MassProductionFinishStep,
        MassProductionDialog: Dialog
    },
    data () {
        return {
            stepAnimation: '',
        }
    },
    computed: {
        ...mapState('massProduction', [
            'stepIndex',
        ]),
        ...mapGetters('massProduction', {
            steps: 'getAvailableSteps',
            stepId: 'getCurrentStepId',
            isSetupValid: 'getIsSetupValid',
        })
    },
    watch: {
        stepIndex (current, previous) {
            this.stepAnimation = current > previous ? 'step-next' : 'step-previous'
        },
    },
    methods: {
        ...mapActions('massProduction', [
            'nextStep',
            'previousStep',
        ]),
    },
}
</script>

<style lang="less" scoped>
@import (reference) '../../less/common';

</style>
