<template>
    <div class="template-grid__container" :class="theme | prefix('template-grid__container--')">
        <div class="template-grid__toolbar">
            <div v-if="isFeaturedShown" class="template-grid__action" :class="getClass('featured') | prefix('template-grid__action--')" @click="filteredBy = 'featured'">
                <svg class="template-grid__action-icon" width="16" height="16" xmlns="http://www.w3.org/2000/svg"><path d="M0 0v4c0 2.065 1.604 4 4.141 4a3.983 3.983 0 0 0 2.65 2.793c-.179.938-.509 2.111-1.124 3.207H4v2h8v-2h-1.667c-.615-1.096-.945-2.27-1.123-3.207A3.983 3.983 0 0 0 11.859 8C14.388 8 16 6.072 16 4V0H0zm2 4V2h2v4c-1.103 0-2-.897-2-2zm12 0c0 1.103-.897 2-2 2V2h2v2z" fill="#858594" fill-rule="nonzero"/></svg>
                <div class="template-grid__action-label">Featured</div>
            </div>

            <div class="template-grid__action" :class="getClass('brand') | prefix('template-grid__action--')" @click="filteredBy = 'brand'">
                <svg class="template-grid__action-icon" width="16" height="16" xmlns="http://www.w3.org/2000/svg"><path d="M15.447 8.105c.339.17.553.516.553.895v3H0V9c0-.379.214-.725.553-.895L6 6V3.721l-2.317-.773A1 1 0 0 1 3 2V0h10v2a1 1 0 0 1-.683.948L10 3.721V6l5.447 2.105zM13 16H3a1 1 0 0 1-1-1v-1h12v1a1 1 0 0 1-1 1z" fill="#858594" fill-rule="nonzero"/></svg>
                <div class="template-grid__action-label">Brand</div>
            </div>

            <div class="template-grid__action" :class="getClass('product') | prefix('template-grid__action--')" @click="filteredBy = 'product'">
                <svg class="template-grid__action-icon" width="16" height="14" xmlns="http://www.w3.org/2000/svg"><path d="M13 5c.6 0 1 .4 1 1v7c0 .6-.4 1-1 1H3c-.6 0-1-.4-1-1V6c0-.6.4-1 1-1h3v4l2-1 2 1V5h3zM0 0h16v4H0V0z" fill="#858594" fill-rule="nonzero"/></svg>
                <div class="template-grid__action-label">Product</div>
            </div>

            <div class="template-grid__action">
                <svg class="template-grid__action-icon" width="16" height="16" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"><defs><path d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8 8-3.6 8-8-3.6-8-8-8zm1 7h3v2H9v3H7V9H4V7h3V4h2v3z" id="a"/></defs><use fill="#858594" fill-rule="nonzero" xlink:href="#a"/></svg>
                <div class="template-grid__action-label">FILTER</div>
            </div>

            <div class="template-grid__action--flexible"></div>

            <div class="template-grid__action">
                <checkbox class="template-grid__action-toggle" :isToggle="true" v-model="isDetailsShown">
                    <div class="template-grid__action-label">Show details</div>
                </checkbox>
            </div>
        </div>

        <div v-for="templateGroup in activeTemplates" :key="templateGroup.label" class="template-grid__group">
            <div v-if="activeTemplates.length > 1" class="template-grid__group-title">{{ templateGroup.label }}</div>

            <template-layout
                :templates="templateGroup.templates"
                :is-details-shown="isDetailsShown"
                :theme="theme"
                @select="id => $emit('select', id)"
            ></template-layout>
        </div>
    </div>
</template>

<script>
import TemplateLayout from './TemplateLayout.vue'
import { Checkbox } from 'design-system'
import { groupBy } from 'lodash'

export default {
    components: {
        TemplateLayout,
        Checkbox
    },
    props: {
        templates: { type: Array, required: true },
        isFeaturedShown: { type: Boolean, default: true },
        isCreateButtonShown: { type: Boolean, default: false },
        isViewFormatShown: { type: Boolean, default: false },
        theme: { type: String, default: 'dark' }
    },
    data () {
        return {
            filteredBy: 'featured', // featured / brands / product
            isDetailsShown: false
        }
    },
    computed: {
        groupedTemplates () {
            return {
                featured: { featured: this.templates.filter(template => template.isFeatured) },
                product: groupBy(this.templates, 'product'),
                brand: groupBy(this.templates, 'brand'),
            }
        },
        mappedTemplates () {
            return {
                featured: this.mapTemplate('featured'),
                product: this.mapTemplate('product'),
                brand: this.mapTemplate('brand')
            }
        },
        activeTemplates () {
            return this.mappedTemplates[this.filteredBy]
        }
    },
    methods: {
        mapTemplate (type) {
            return Object.keys(this.groupedTemplates[type]).map(key => ({
                label: key,
                templates: this.groupedTemplates[type][key]
            }))
        },
        getClass (filter) {
            return {
                active: this.filteredBy === filter
            }
        }
    }
}
</script>

<style lang="less" scoped>
@import (reference) '../../less/common';

.template-grid {
    &__toolbar {
        height: 54px;
        background-color: @extremely-dark-gray;
        padding-left: 32px;
        display: flex;
    }

    &__action {
        line-height: 54px;
        margin-right: 24px;
        display: flex;
        cursor: pointer;

        &-label {
            color: @bluish-gray;
            font-size: 12px;
        }
        
        &-icon {
            margin: auto 8px auto 0;
        }

        &-toggle {
            margin: auto;
            margin-top: 14px;
        }

        &--flexible {
            flex: 1 1 0;
        }

        &--active {
            .template-grid__action {
                &-label { color: @royal-blue; }
                &-icon path { fill: @royal-blue; }
            }
        }
    }

    &__group {
        margin-top: 2px;

        &-title {
            height: 54px;
            line-height: 54px;
            color: white;
            font-size: 22px;
            padding-left: 32px;
            margin: 2px 0;
            background-color: @extremely-dark-gray;
        }
    }
}

.template-grid__container--light {
    .template-grid__toolbar {
        background-color: unset;
    }

    .template-grid__group-title {
        background-color: white;
        color: black;
    }
}
</style>
