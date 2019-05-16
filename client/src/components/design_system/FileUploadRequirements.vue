<template>
    <div class="file-upload-requirements">
        <p v-for="requirement in requirements" :key="requirement.name" class="file-upload-requirements__requirement">
            {{ requirement.name }}:
            <span class="file-upload-requirements__requirement-value">{{ requirement.value }}</span>
        </p>
    </div>
</template>

<script>
export default {
    props: {
        requirements: { type: Array, required: true },
    },
}
</script>

<style scoped lang="less">
@import (reference) './less/common';
@import './less/typography';

.file-upload-requirements {
    color: @dolphin;
    font-size: 11px;
    margin-top: 15px;
    line-height: 15px;
    letter-spacing: 0.5px;
    position: absolute;
    z-index: @z-heaven;
    .regular-font();

    .file-upload-requirements__requirement-value {
        .medium-font();
        font-weight: bold;
    }
}
</style>
