<template>
    <div id="app">
        <!-- <router-view></router-view> -->
    </div>
</template>

<script>
import { mapActions } from 'vuex'

export default {
    name: 'App',
}
</script>

<style lang="less" scoped>
@import 'less/variables';
@import 'less/common';

#app {
    font-family: @regular-font;
}

.header {
    height: 50px;
    background-color: @extremely-dark-gray;
    padding: 0 105px;

    &__info {
        display: flex;
        line-height: 50px;
    }

    &__title {
        color: white;
        font-size: 18px;
    }

    &__flexible { flex: 1 1 0; }

    &__settings {
        color: @bluish-gray;
        min-width: 50px;
        display: flex;
        padding: 0 15px;

        svg { margin: auto; }
    }

    &__navigation {
        border-top: 2px solid @bg;
        display: flex;
    }

    &__tabs {
        display: flex;
        height: 100%;
    }

    &__tab {
        display: flex;
        height: 100%;
        margin-right: 15px;
        box-sizing: border-box;
        cursor: pointer;

        svg { margin: auto; }

        &--active {
            border-bottom: 3px solid white;

            svg path { fill: white; }
        }
    }

    &__active-tab {
        border-left: 2px solid @bg;
        margin-left: 7px;
        padding-left: 20px;
        line-height: 50px;
        color: white;
    }
}
</style>
