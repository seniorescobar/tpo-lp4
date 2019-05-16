<template>
    <div class="campaign">
        <div class="campaign__navigation">
            <div class="campaign__tab campaign__tab--active">Creatives</div>
            <div class="campaign__tab">Placements</div>
            <div class="campaign__tab">Data</div>
            <div class="campaign__tab">Audiences</div>
            <div class="campaign__tab">Strategy</div>
            <div class="campaign__tab">Reports</div>
        </div>

        <div class="campaign__creatives">
            <div class="creatives__header">
                <div class="creatives__new-buttons">
                    <div class="creatives__new creatives__new--template" @click="openDialog"><svg width="16" height="16" xmlns="http://www.w3.org/2000/svg"><path d="M6 9H1a1 1 0 0 1-1-1V1a1 1 0 0 1 1-1h5a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1zm0 7H1a1 1 0 0 1-1-1v-3a1 1 0 0 1 1-1h5a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1zm9-10h-5a1 1 0 0 1-1-1V1a1 1 0 0 1 1-1h5a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1zm0 10h-5a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h5a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1z" fill="#000014" fill-rule="evenodd"/></svg></div>
                    <div class="creatives__new creatives__new--creative"><svg width="22" height="22" xmlns="http://www.w3.org/2000/svg"><g stroke="#000" stroke-width="2" fill="none" fill-rule="evenodd" stroke-linecap="square"><path d="M11 1v20M21 11H1"/></g></svg></div>
                </div>
            </div>

            <div class="creatives__content">
                
            </div>
        </div>

        <mass-production v-if="isDialogOpen" @close="closeDialog"></mass-production>
    </div>
</template>

<script>
// import { RecordsListHeader, RecordsListRow, RecordsListFooter } from 'records-list'
import MassProduction from '../mass-production/MassProduction.vue'
import { mapActions, mapGetters } from 'vuex'

export default {
    name: 'CampaignView',
    components: {
        MassProduction
        // RecordsListHeader, RecordsListRow, RecordsListFooter
    },
    data () {
        return {
            isDialogOpen: false
        }
    },
    computed: {
        ...mapGetters({
            creatives: 'getCreatives',
        })
    },
    methods: {
        ...mapActions('massProduction', [
            'resetState',
            'setDialogStartingPoint',
            'setCampaign'
        ]),
        closeDialog () {
            this.isDialogOpen = false
            this.resetState()
        },
        openDialog () {
            this.resetState()
            this.setDialogStartingPoint('campaign')
            this.setCampaign('campaign1')
            this.isDialogOpen = true
        }
    }
}
</script>

<style lang="less" scoped>
@import (reference) '../../less/common';

.campaign {
    &__navigation {
        background: @extremely-dark-gray;
        border-top: 2px solid @bg;
        height: 50px;
        display: flex;
        padding: 0 105px;
    }

    &__tab {
        line-height: 50px;
        color: @bluish-gray;
        box-sizing: border-box;
        margin-right: 30px;
        font-size: 14px;

        &--active {
            color: white;
            border-bottom: 3px solid white;
        }
    }

    &__creatives {
        margin: 60px 105px;
        background-color: white;
        position: relative;
    }
}

.creatives {
    &__header {
        height: 60px;
    }

    &__new {
        height: 60px;
        width: 60px;
        border-radius: 50%;
        margin: 0 5px;
        display: flex;
        cursor: pointer;

        &-buttons {
            display: flex;
            position: absolute;
            right: 0;
            height: 60px;
            top: -30px;
            padding-right: 25px;
        }

        &--creative { background-color: @light-green; }

        &--template { background-color: @light-smoke; }

        svg { margin: auto; }
    }

    // &__content {}
}
</style>
