<template>
    <div class="template__container" :class="theme | prefix('template__container--')" @click="$emit('select', template.id)">
        <div class="template__wrapper">
            <div class="template__data">
                <div class="template__banner">
                    <svg width="32" height="32" xmlns="http://www.w3.org/2000/svg"><g fill="none" fill-rule="evenodd"><path d="M32 16c0 8.836-7.164 16-16 16S0 24.836 0 16 7.164 0 16 0s16 7.164 16 16" fill="#6A3DFF"/><path d="M20.196 12.787H13.55l.163 1.765h6.32l-.488 5.295L16 20.853l-3.546-1.006-.25-2.698h1.751l.116 1.349 1.93.557 1.898-.548.194-2.231h-5.955l-.01.033-.488-5.273h8.717l-.16 1.75zM9.072 8.16l1.298 14.085L16 23.84l5.627-1.595 1.3-14.085H9.072z" fill="#FFF"/></g></svg>
                </div>
                <div class="template__name">{{ template.templateName }}</div>
            </div>

            <div class="template__preview">
                <template-preview
                    :template="template"
                    :focus="hovered"
                ></template-preview>
            </div>

            <div v-if="isDetailsShown" class="template__details">
                <div>Format:</div>
                <div class="template__value">{{ template.format }}</div>
                <div>Size:</div>
                <div class="template__value">{{ template.size.width.split('px')[0] }}x{{ template.size.height }}</div>
                <div>Used:</div>
                <div class="template__value">{{ template.count }} instances</div>
                <div>Modified:</div>
                <div class="template__value">{{ template.modified }} </div>
                <div>Attributes:</div>
                <div class="template__value template__value--attributes">
                    <div v-for="unit in template.units" :key="unit.id" :class="getClass(unit.id) | prefix('template__value--')" @mouseover="hovered = unit.id" @mouseout="hovered = ''">{{ unit.label }}</div>
                </div>
            </div>

            <div v-else class="template__details template__details--shortened">
                <div>{{ template.templateName }}</div>
                <div>{{ template.format }}, {{ template.size.width.split('px')[0] }}x{{ template.size.height }}</div>
            </div>
        </div>
    </div>
</template>

<script>
import TemplatePreview from './TemplatePreview.vue'

export default {
    components: {
        TemplatePreview
    },
    props: {
        template: { type: Object, required: true },
        theme: { type: String, default: 'dark' },
        isDetailsShown: { type: Boolean, default: false },
        // isButtonShownOnHover, isSelected
    },
    data () {
        return {
            hovered: ''
        }
    },
    methods: {
        getClass (id) {
            return {
                highlighted: id === this.hovered
            }
        }
    }
}
</script>

<style lang="less" scoped>
@import (reference) '../../less/common';

.template {
    &__container {
        height: 100%;
        width: 100%;
        padding: 8px 32px;
        background-color: @extremely-dark-gray;
        box-sizing: border-box;
        color: white;

        &:hover { background-color: lighten(@extremely-dark-gray, 10%); }
    }

    &__wrapper {
        max-width: 970px;
        margin: 0 auto;
    }

    &__data {
        display: flex;
        height: 56px;
        line-height: 56px;
    }

    &__banner {
        display: flex;

        svg { margin: auto; }
    }

    &__name {
        flex: 1 1 0;
        margin-left: 8px;
    }

    &__count {
        color: @bluish-gray;
    }

    &__preview {
        margin: auto;
        width: fit-content;
    }

    &__details {
        margin-top: 12px;
        font-size: 12px;
        line-height: 20px;
        display: grid;
        grid-template-columns: 65px 1fr;
        color: @bluish-gray;

        &--shortened {
            display: block;
        }
    }

    &__value {
        margin-left: 16px;

        &--attributes {
            color: white;
            cursor: default;
        }

        &--highlighted {
            color: @royal-blue;
        }
    }
}

.template__container--light {
    background-color: white;
    color: black;

    &:hover { background-color: darken(white, 10%); }

    .template__value--attributes { color: black; }
}
</style>
