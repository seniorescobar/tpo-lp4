<template>
    <div class="component">
        <div class="component__dropdown">
            <div @click="isDropdownOpen = !isDropdownOpen" class="component__dropdown-trigger">
                {{ label }}
                <svg width="7" height="3" xmlns="http://www.w3.org/2000/svg"><path fill="#D4D4D9" d="M3.5 3L7 0H0z" fill-rule="evenodd" opacity=".8"/></svg>
            </div>
            <div v-if="isDropdownOpen" class="component__dropdown-container">
                <default-list
                    :items="options"
                    :value="current"
                    theme="light"
                    size="condensed"
                    @select="switchComponent"
                ></default-list>
            </div>
        </div>
        <div class="component__slot">
            <slot :name="current"/>
        </div>
    </div>
</template>

<script>
import { DefaultList } from 'design-system'

export default {
    components: {
        DefaultList,
    },
    props: {
        options: { type: Array, required: true },
        default: { type: String }
    },
    data () {
        return {
            current: this.default || this.options[0].id,
            isDropdownOpen: false,
        }
    },
    computed: {
        label () {
            return this.options.find(option => option.id === this.current).label
        }
    },
    methods: {
        switchComponent (item) {
            this.current = item.id
            this.isDropdownOpen = false
        }
    }
}
</script>


<style lang="less" scoped>
@import (reference) '../../less/common';

.component {
    position: relative;

    &__dropdown {
        position: absolute;
        right: 0;
        z-index: 1;
        cursor: pointer;
    }

    &__dropdown-trigger {
        display: flex;
        color: @light-smoke;
        font-size: 11px;
        line-height: 16px;

        svg { margin: auto 5px; }
    }

    &__dropdown-container {
        width: fit-content;
        position: absolute;
        left: 0;
        top: 20px;
        width: fit-content;
        background-color: #fff;
    }

    &__slot {
        margin-top: 20px;
    }
}
</style>
