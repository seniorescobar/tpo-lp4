<template>
    <div class="records-list-empty">
        <slot></slot>
        <div class="records-list-empty__title">
            <slot name="text"></slot>
        </div>
    </div>
</template>

<script>
export default {
    name: 'records-list-empty',
}
</script>

<style lang="less" scoped>
@import (reference) '~design-system/less/variables';

.records-list-empty {
    width: 100%;
    height: 540px;
    display: flex;
    background-color: white;
    flex-direction: column;
    align-items: center;
    justify-content: center;

    &__title {
        margin-top: 15px;
        font-size: 20px;
        .regular-font();
        color: @very-light-gray;
        text-align: center;
    }
}
</style>
