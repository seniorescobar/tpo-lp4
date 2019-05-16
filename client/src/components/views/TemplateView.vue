<template>
    <div class="template">
        <template-grid v-if="templates.length" :templates="templates" theme="light" @select="openDialog"></template-grid>
        <mass-production v-if="isDialogOpen" @close="closeDialog"></mass-production>
    </div>
</template>

<script>
import TemplateGrid from '../template/TemplateGrid.vue'
import MassProduction from '../mass-production/MassProduction.vue'
import { mapState, mapActions } from 'vuex'

export default {
    name: 'TemplateView',
    components: {
        TemplateGrid,
        MassProduction
    },
    data () {
        return {
            isDialogOpen: false
        }
    },
    computed: {
        ...mapState('massProduction', [
            'templates'
        ])
    },
    methods: {
        ...mapActions('massProduction', [
            'resetState',
            'setDialogStartingPoint',
            'setTemplateToInstantiate',
            'fetchCampaigns',
        ]),
        closeDialog () {
            this.isDialogOpen = false
            this.resetState()
        },
        openDialog (templateId) {
            this.resetState()
            this.setDialogStartingPoint('template')
            this.fetchCampaigns()
            this.setTemplateToInstantiate(templateId)
            this.isDialogOpen = true
        }
    }
}
</script>

<style lang="less" scoped>
@import (reference) '../../less/common';

</style>
