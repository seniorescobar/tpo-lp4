<template>
    <div :class="{'records-list-row__header--stuck': isStuck, 'records-list-row__header--empty': isEmpty}" class="records-list-row records-list-row__header">
        <div ref="stickySentinel" class="records-list-header__sticky-sentinel"></div>
        <div class="records-list-row__cell records-list-row__cell--checkbox" v-if="showCheckbox">
            <checkbox :value="allSelected" theme="light" @input="toggle"></checkbox>
        </div>
        <slot></slot>
    </div>
</template>

<script>
import { Checkbox } from 'design-system'

export default {
    components: {
        checkbox: Checkbox,
    },
    props: {
        allSelected: { type: Boolean, default: false },
        isEmpty: { type: Boolean, default: false },
        trackName: { type: String, default: 'header' },
        showCheckbox: { type: Boolean, default: true },
    },
    data () {
        return {
            isStuck: false,
        }
    },
    watch: {
        isStuck (val) {
            this.$emit('stick', val)
        },
    },
    mounted () {
        if (window.IntersectionObserver) {
            this.observer = new IntersectionObserver((entries) => {
                this.isStuck = entries[0].boundingClientRect.bottom < 0
            }, { threshold: 0 })

            this.observer.observe(this.$refs.stickySentinel)
        }
    },
    destroyed () {
        if (this.observer) {
            this.observer.disconnect()
        }
    },
    methods: {
        toggle (val, event) {
            // Prevent checkbox spacebar toggle since we handle this in global listener on reportsList
            if (!(event instanceof KeyboardEvent)) {
                this.$emit('selectAll', val)
                this.$root.$emit('tracking-event', { type: this.trackName, label: 'toggleAll', trigger: 'click', data: { value: val } })
            }
        },
    },
}
</script>

<style lang="less" scoped>
@import (reference) '~design-system/less/variables';
@import './records_list_common';

.records-list-header__sticky-sentinel {
    position: absolute;
    left: 0;
    right: 0;
    visibility: hidden;
    height: 1px;
    top: -2px;
}

.records-list-row__sort-arrow {
    fill: @gunpowder;
    margin-left: 5px;
    opacity: 0;
    transition: opacity 200ms ease-out;

    &--active {
        opacity: 1;
    }

    &--ascending {
        transform: rotate(180deg);
    }
}

.records-list-row__header {
    background-color: @white;
    border-bottom: 2px solid @extremely-light-gray;
    position: sticky;
    top: 0;
    height: 60px;
    width: 100%;
    z-index: 15;
    padding-top: 15px;
    box-sizing: border-box;

    .records-list-row__cell {
        color: @gunpowder;
    }

    &--empty {
        .records-list-row__cell {
            opacity: 0.4;
            pointer-events: none;
        }
    }
}
</style>
