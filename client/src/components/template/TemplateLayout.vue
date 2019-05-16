<template>
    <div class="template-grid__layout" :class="getClass | prefix('template-grid__layout--')">
        <div class="grid__row grid__row--1">
            <template-details :template="templates[0]" :is-details-shown="isDetailsShown" @select="select" :theme="theme"/>
            <template-details :template="templates[1]" :is-details-shown="isDetailsShown" @select="select" :theme="theme"/>
            <template-details :template="templates[2]" :is-details-shown="isDetailsShown" @select="select" :theme="theme"/>
            <div class="grid__row--1--4">
                <template-details :template="templates[3]" :is-details-shown="isDetailsShown" @select="select" :theme="theme"/>
                <template-details :template="templates[4]" :is-details-shown="isDetailsShown" @select="select" :theme="theme"/>
            </div>
        </div>
        <div class="grid__row grid__row--2">
            <template-details :template="templates[5]" :is-details-shown="isDetailsShown" @select="select" :theme="theme"/>
        </div>
        <div class="grid__row grid__row--3">
            <template-details :template="templates[6]" :is-details-shown="isDetailsShown" @select="select" :theme="theme"/>
            <template-details :template="templates[7]" :is-details-shown="isDetailsShown" @select="select" :theme="theme"/>
        </div>
    </div>
</template>

<script>
/**
 * This is a static representation of the dynamic grid layout (also found in VCE)
 */

import TemplateDetails from './TemplateDetails.vue'

export default {
    components: {
        TemplateDetails
    },
    props: {
        templates: { type: Array, required: true },
        isDetailsShown: { type: Boolean, default: false },
        theme: { type: String, default: 'dark' },
    },
    computed: {
        getClass () {
            return {
                detailed: this.isDetailsShown
            }
        }
    },
    methods: {
        select (id) {
            this.$emit('select', id)
        }
    }
}
</script>

<style lang="less" scoped>
@import (reference) '../../less/common';

.template-grid__layout {
    display: grid;
    grid-template-columns: 1fr;
    grid-template-rows: 740px 390px 440px;
    grid-gap: 2px;
}

.grid {
    &__row {
        &--1 {
            display: grid;
            grid-template-columns: 364fr 322fr 364fr 384fr;
            grid-gap: 2px;

            &--4 {
                display: grid;
                grid-template-columns: 1fr;
                grid-template-rows: 190fr 548fr;
                grid-gap: 2px;
            }
        }
        &--3 {
            display: grid;
            grid-template-columns: 792fr 646fr;
            grid-gap: 2px;
        }
    }
}

.template-grid__layout--detailed {
    grid-template-rows: 860px 485px 560px;

    .grid__row--1--4 { grid-template-rows: 330fr 548fr; }
}
</style>
