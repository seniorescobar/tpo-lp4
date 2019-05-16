<template>
    <div class="template-selection__wrapper">
        <div class="template-selection__goal">
            <div class="template-selection__label">I'm looking for a template for</div>
            <selectbox class="template-selection__goals" placeholder="select goal" :options="goals" size="phat" v-model="goal"></selectbox>
        </div>

        <div class="template-selection__grid">
            <template-grid :templates="templates" @select="selectTemplate"></template-grid>
        </div>
    </div>
</template>

<script>
import TemplateGrid from '../template/TemplateGrid'
import { Selectbox } from 'design-system'
import { mapState, mapGetters, mapActions } from 'vuex'

export default {
    components: {
        Selectbox,
        TemplateGrid
    },
    data () {
        return {
            goal: ''
        }
    },
    computed: {
        ...mapState('massProduction', [
            'templates'
        ]),
        ...mapGetters('massProduction', {
            goals: 'getGoals'
        })
    },
    mounted () {
        this.setTemplateToInstantiate(null)
    },
    methods: {
        ...mapActions('massProduction', [
            'nextStep',
            'setTemplateToInstantiate'
        ]),
        selectTemplate (id) {
            this.setTemplateToInstantiate(id)
            this.nextStep()
        },
    }
}
</script>

<style lang="less" scoped>
@import (reference) '../../less/common';

.template-selection {
    &__wrapper {
        position: relative;
        margin: 0 -60px;
        height: 100%;
    }

    &__goal {
        position: absolute;
        top: 80px;
        left: 0;
        right: 0;
        display: flex;
        width: 605px;
        margin: auto;
    }

    &__label {
        line-height: 71px;
        color: @bluish-gray;
        font-size: 22px;
        margin-right: 16px;
    }

    &__goals {
        max-width: 320px;
    }

    &__grid {
        position: absolute;
        top: 256px;
        left: 0;
        width: 100%;
    }
}
</style>
