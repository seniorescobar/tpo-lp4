<template>
    <div class="new-dialog__step-wrapper">
        <div class="campaign-selection__wrapper">
            <div class="campaign-selection__container">
                <selectbox :options="campaigns" :value="campaignId" label="Use template in" placeholder="Please select a campaign" size="phat" @input="setCampaign"/>
            </div>
        </div>

        <div class="new-dialog__button-container">
            <dialog-button :disabled="campaignId === null" @click="nextStep">Next step</dialog-button>
        </div>
    </div>    
</template>

<script>
import { mapState, mapActions, mapGetters } from 'vuex'
import { DialogButton, Selectbox } from 'design-system'

export default {
    components: {
        DialogButton,
        Selectbox
    },
    computed: {
        ...mapState('massProduction', [
            'campaignId',
        ]),
        ...mapGetters('massProduction', {
            campaigns: 'getCampaigns'
        })
    },
    methods: {
        ...mapActions('massProduction', [
            'setCampaign',
            'nextStep',
        ]),
    }
}
</script>

<style lang="less" scoped>
@import (reference) '../../less/common';
@import '../../less/dialogs_common';
@import '../../less/dialogs_columns_layout';

.campaign-selection {
    &__wrapper {
        height: 100%;
        display: flex;
    }

    &__container {
        margin: auto;
        width: 600px;
        height: fit-content;
    }
}
</style>