<template>
    <div class="tab__container">
        <div v-for="instance in instances" :key="instance.id" class="tab" :class="getClass(instance.id) | prefix('tab--')" @click="$emit('input', instance.id)">
            <div class="tab__content">
                <div class="tab__text tab__text--bold">{{ instance.label }}</div>
                <div class="tab__text">{{ instance.size }}</div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        instances: { type: Array, required: true },
        value: { type: String },
        disabled: { type: Boolean, default: false }
    },
    methods: {
        getClass (instanceId) {
            return {
                disabled: this.disabled,
                selected: !this.disabled && this.value === instanceId
            }
        }
    }
}
</script>

<style lang="less" scoped>
@import (reference) '../less/common';

.tab {
    margin-right: 16px;
    padding: 8px;
    box-sizing: border-box;
    background-color: @extremely-dark-gray;
    border-radius: 3px;
    height: 44px;
    display: flex;
    cursor: pointer;

    &__container { display: flex; }

    &__content { margin: auto; }

    &__text {
        color: white;
        font-size: 12px;

        &--bold { font-family: @demi-font; }
    }

    &--disabled {
        cursor: default;

        &__text { color: @bluish-gray; }
    }

    &--selected { background-color: @royal-blue; }
}
</style>
