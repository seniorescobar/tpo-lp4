<template>
    <div id="app">
        <router-view/>
    </div>
</template>

<script>

export default {
    name: 'App',
}
</script>

<style lang="less">
@import 'less/variables';
@import 'less/common';

#app {
    font-family: @regular-font;
    color: white;
}

.md-theme-default.md-theme-default {
    background-color: @purple;
}
</style>
