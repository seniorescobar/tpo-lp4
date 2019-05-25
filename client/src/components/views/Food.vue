<template>
    <div class="container">
        <iframe class="iframe" src="https://maps.google.com/maps?q=ljubljana%20restaurants&t=&z=15&ie=UTF8&iwloc=&output=embed" frameborder="0" scrolling="no" marginheight="0" marginwidth="0"></iframe>
        <div class="loading">Loading</div>
    </div>
</template>

<style lang="less" scoped>
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
