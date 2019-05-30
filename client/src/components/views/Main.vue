<template>
    <div id="main">
        <div class="header">
            <div class="header__action" @click="logout">Log out</div>
            <div class="header__title">StraightAs</div>
            <div></div>
        </div>

        <div class="navigation">
            <div v-for="tab in tabs" :key="tab.id" class="navigation__item" :class="getClass(tab.id) | prefix('navigation__item--')" @click="$router.push(tab.id)">{{ tab.label }}</div>
        </div>

        <router-view/>
    </div>
</template>

<script>

export default {
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
        },
        logout () {
            localStorage.removeItem("user")
            this.$router.push('/login')
        }
    }
}
</script>

<style lang="less" scoped>
@import '../../less/common';

.header {
    display: grid;
    grid-template-columns: 100px 1fr 100px;
    padding: 8px 0;

    &__action {
        text-decoration: underline;
        text-align: center;
        cursor: pointer;
    }
    &__title {
        text-align: center;
        line-height: 32px;
        font-size: 22px;
    }
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