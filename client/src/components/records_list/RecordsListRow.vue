<template>
    <div :class="{ 'records-list-row--selected': selected }" class="records-list-row" tabindex="0" @focus="onFocus">
        <div class="records-list-row__cell records-list-row__cell--checkbox" v-if="showCheckbox">
            <checkbox :value="selected" theme="white" tabindex="-1" @input="toggle"></checkbox>
        </div>
        <slot></slot>
    </div>
</template>

<script>
import { Checkbox } from 'design-system'

export default {
    name: 'records-list-row',
    components: {
        checkbox: Checkbox,
    },
    props: {
        selected: { type: Boolean, default: false },
        showCheckbox: { type: Boolean, default: true },
    },
    methods: {
        toggle (val, event) {
            // Prevent checkbox spacebar toggle since we handle this in global listener on reportsList
            if (!(event instanceof KeyboardEvent)) {
                this.$emit('toggle', val, event)
            }
        },
        onFocus () {
            this.$emit('focus')
        },
    },
}
</script>

<style lang="less">
@import (reference) '~design-system/less/variables';

.records-list-row {
    .records-list__editable-span--editing::selection {
        background: fade(@royal-blue, 99%);
        color: white;
    }
}

.records-list-row--selected {
    .records-list__editable-span--editing::selection {
        color: @gunpowder;
        background: fade(white, 99%); // Needs to be rgba with 0.99 opacity, otherwise it's 50% transparent
    }
}
</style>

<style lang="less" scoped>
@import (reference) '~design-system/less/variables';
@import  './records_list_common';

@keyframes row-highlight {
    10%  { background-color: @white; }
    30%  { background-color: @cream-smoke; }
    50%  { background-color: @cream-smoke; }
    100% { background-color: @white; }
}

.records-list-row {
    position: relative;
    background-color: white;
    transition: background-color @records-list-row-animation-time ease-out;

    &:focus {
        outline: none;
    }

    &--highlight {
        animation: row-highlight @records-list-row-highlight-time ease-out;
    }

    &--focused, &:hover:not(&--highlight):not(&--selected) {
        background-color: fade(@extremely-light-gray, 50%);
    }

    .records-list__editable-span--editing {
        color: @gunpowder;
    }

    &--selected {
        background-color: @royal-blue;
        animation: 0;

        .records-list__editable-span--editing {
            color: white;
        }

        &.records-list-row--focused, &:hover {
            background-color: mix(@royal-blue, @light-gray, 80%);
        }
    }

    &.adding-items-enter-active, &.adding-items-leave-active {
        transition: transform @records-list-row-animation-time;
    }

    &.adding-items-leave-to, &.adding-items-leave-active {
        position: absolute;
        z-index: 1;
    }

    &.adding-items-enter {
        transform: translateY(-60px);
    }

    &.adding-items-move {
        transition: transform @records-list-row-animation-time;
    }

    &--removing-items {
        z-index: 5;
    }

    &.removing-items-enter-active, &.removing-items-leave-active {
        transition: transform @records-list-row-animation-time;
    }

    &.removing-items-leave-to, &.removing-items-leave-active {
        transform: translateY(-60px);
        position: absolute;
        z-index: 1;
    }

    &.removing-items-enter {
        transform: translateY(60px);
    }

    &.removing-items-move {
        transition: transform @records-list-row-animation-time;
    }
}
</style>
