<template>
    <div>
        <div v-for="(row, index) in rows" ref="row" :key="index" class="records-list-row">
            <div class="records-list-row__cell records-list-row__cell--checkbox"></div>
            <div v-for="(column, name) in columns"
                 :class="{'records-list-row__cell--circle': column.loaderType === 'circle'}"
                 :style="{ width: column.width }"
                 :key="name"
                 class="records-list-row__cell">
            </div>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        rows: { type: Number, required: true },
        columns: { type: Object, required: true },
    },
    mounted () {
        this.$nextTick(() => {
            this.$refs.row.forEach((row) => {
                Array.from(row.querySelectorAll('.records-list-row__cell:not(:first-child)')).reduce((widthAcc, cell) => {
                    // Set delay according to cumulative width of previous cells, make it look like gradient travels between cells
                    cell.style.animationDelay = `${Math.ceil(widthAcc / 6)}ms`
                    cell.style.minWidth = `${Math.ceil(cell.clientWidth * (0.4 + 0.25 * Math.random()))}px` // Random width between 40% and 65%
                    return widthAcc + cell.clientWidth
                }, 0)
            })
        })
    },
}
</script>

<style lang="less" scoped>
@import (reference) '~design-system/less/variables';
@import './records_list_common';

.records-list-row__cell {
    position: relative;

    &:after {
        content: "";
        position: absolute;
        height: 33%;
        top: 33%;
        min-width: inherit;
        animation: placeHolderShimmer 1.2s linear infinite forwards;
        animation-delay: inherit;
        background: linear-gradient(to right, @gallery-smoke 8%, #ddd 18%, @gallery-smoke 33%);
        background-size: 5000px 10px;
    }

    &--circle:after {
        height: 32px;
        width: 32px;
        border-radius: 50%;
        min-width: 0;
        top: auto;
    }
}

@keyframes placeHolderShimmer {
    0% {
        background-position: 50% 0;
    }
    100% {
        background-position: -50% 0;
    }
}
</style>
