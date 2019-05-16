<template>
    <div id="app">
        <div class="title">StraightAs</div>
        <div class="navigation">
            <div v-for="tab in tabs" :key="tab.id" class="navigation__item" :class="getClass(tab.id) | prefix('navigation__item--')" @click="$router.push(tab.id)">{{ tab.label }}</div>
        </div>

        <router-view/>
    </div>
</template>

<script>

export default {
    name: 'App',
    beforeCreate () {
        this.tabs = [
            { id: 'home', label: 'Home' },
            { id: 'food', label: 'Food' },
            { id: 'events', label: 'Events' },
            { id: 'bus', label: 'Bus' },
        ]
    },
    methods: {
        getClass (tabId) {
            return {
                active: tabId === this.$route.name
            }
        }
    }
}
</script>

<style lang="less" scoped>
@import 'less/variables';
@import 'less/common';

#app {
    font-family: @regular-font;
    color: white;
}

.title {
    text-align: center;
    line-height: 32px;
    font-size: 22px;
    margin: 8px 0;
}

.navigation {
    display: flex;
    justify-content: center;
    border-bottom: 2px solid @dolphin;

    &__item {
        padding: 12px 24px;
        cursor: pointer;

        &--active {
            color: @royal-blue;
        }
    }
}
</style>
