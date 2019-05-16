<template>
    <div class="template-preview" :style="getStyle">
        <div v-for="(unit, index) in units" :key="index" class="unit" :style="unit.position">
            <img v-if="unit.type === 'asset'" :ref="unit.id" :src="imageSource(unit)" class="unit__asset" :style="unit.style">
            <div v-if="unit.type === 'text'" class="unit__text" :style="unit.style">{{ unit.value }}</div>
            <button v-if="unit.type === 'button'" class="unit__button" :style="unit.style">{{ unit.value }}</button>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        template: { type: Object, required: true },
        isOverflowHidden: { type: Boolean, default: true },
        exposedId: { type: String, default: '' },
    },
    computed: {
        getStyle () {
            return {
                width: `${this.template.size.width}px`,
                height: `${this.template.size.height}px`,
                overflow: this.isOverflowHidden ? 'hidden' : 'visible'
            }
        },
        units () {
            return this.template.units.filter(unit => !this.exposedId || this.exposedId === unit.id)
        }
    },
    methods: {
        imageSource (unit) {
            if (typeof unit.value === 'string')
                return `images/${unit.value}`
            else
                this.readFile(unit)
        },
        readFile (unit) {
            const fileReader = new FileReader()
            fileReader.onloadend =  () => this.$refs[unit.id][0].src = fileReader.result
            fileReader.readAsDataURL(unit.value.file)
        }
    }
}
</script>

<style lang="less" scoped>
@import (reference) '../../less/common';

.template-preview {
    position: relative;
}

.unit {
    position: absolute;
    box-sizing: border-box;
    border: 2px solid transparent;

    &__asset {
        width: 100%;
        height: 100%;
        object-fit: cover;
    }

    &__text {
        text-align: center;
        color: black;
        cursor: default;
    }

    &__button {
        width: 100%;
        height: 100%;
        background-color: black;
        color: white;
        text-transform: uppercase;
        border: none;
        outline: none;
    }
}
</style>
