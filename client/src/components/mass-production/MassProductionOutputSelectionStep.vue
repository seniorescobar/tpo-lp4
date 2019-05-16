<template>
    <div class="new-dialog__step-wrapper">
        <div class="new-dialog__content-wrapper">
            output selection
        </div>

        <div class="new-dialog__button-container">
            <dialog-button @click="nextStep">Next step</dialog-button>
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
        ]),
        ...mapGetters('massProduction', {
        })
    },
    methods: {
        ...mapActions('massProduction', [
            'nextStep'
        ]),
    }
}
</script>

<style lang="less" scoped>
@import (reference) '../../less/common';
@import '../../less/dialogs_common';
@import '../../less/dialogs_columns_layout';

.new-dialog__content-wrapper {
    height: 100%;
}
</style>