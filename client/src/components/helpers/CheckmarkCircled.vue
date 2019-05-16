<template>
    <div class="checkmark-circled">
        <div class="checkmark-circled__wrapper">
            <div class="checkmark-circled__check"></div>
        </div>
        <svg class="checkmark-circled__circle" style="transform: rotate(-90deg)">
            <circle cx="61.5" cy="61.5" r="60"></circle>
        </svg>
    </div>
</template>

<style lang="less" scoped>
@import (reference) '../../less/variables';

.checkmark-circled {
    margin: auto;
    width: 123px;
    height: 123px;

    &__wrapper {
        position: absolute;
        margin: 0 auto;
        width: inherit;
        height: inherit;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    &__check {
        width: 16.66%;
        height: 41.66%;
        display: block;
        border: solid @light-green;
        opacity: 0;
        border-width: 0 8px 8px 0;
        transform: translateY(-6.5px) rotate(45deg);
        animation: appear 0s linear forwards;
        animation-delay: @finish-step-circle-fill-time;
    }

    &__circle {
        width: inherit;
        height: inherit;
        fill: none;
        stroke: @light-green;
        stroke-width: 3;
        stroke-dasharray: 380;
        stroke-dashoffset: 380;
        animation: spinner @finish-step-circle-fill-time linear forwards;
    }
}

@keyframes appear {
    to {
        opacity: 1;
    }
}

@keyframes spinner {
    to {
        stroke-dashoffset: 0;
    }
}
</style>
