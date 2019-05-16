<template>
    <div class="new-dialog__step-wrapper">
        <div class="campaign-selection__wrapper">
            <div class="campaign-selection__container">
                <input-element v-model="name" placeholder="Experience name"/>
            </div>
        </div>

        <div class="new-dialog__button-container">
            <dialog-button :disabled="campaignId === null" @click="nextStep">Next step</dialog-button>
        </div>
    </div>    
</template>

<script>
import { mapState, mapActions } from 'vuex'
import { DialogButton, Input } from 'design-system'

export default {
    components: {
        DialogButton,
        inputElement: Input,
    },
    computed: {
        ...mapState('massProduction', [
            'experienceName'
        ]),
        name: {
            get: () => this.experienceName,
            set: (v) => this.setExperienceName(v)
        }
    },
    methods: {
        ...mapActions('massProduction', [
            'nextStep',
            'setExperienceName',
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