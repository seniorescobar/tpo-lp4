<template>
    <div class="container">
        <iframe class="iframe" src="https://www.trola.si/" />
        <div class="loading">Loading</div>
    </div>
</template>

<script>
export default {
    
}
</script>

<style lang="less" scoped>
@import '../../less/common';

.container {
    position: relative;
}
.iframe {
    display: block;
    margin: 64px auto 0;
    width: 600px;
    height: 500px;
}
.loading {
    position: absolute;
    margin: 64px auto;
    top: 64px;
    left: 0;
    right: 0;
    text-align: center;
}
</style>
