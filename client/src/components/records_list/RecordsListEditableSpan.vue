<template>
    <span
        v-click-outside="stopEditing"
        v-if="editing"
        ref="span"
        :class="[theme] | prefix('records-list__editable-span--')"
        contenteditable="true"
        class="records-list__editable-span records-list__editable-span--editing"
        spellcheck="false"
        @paste="handlePaste"
        @click="handleClick"
        @keydown="checkLength"
        @keydown.space.stop
        @keydown.enter.stop.prevent="stopEditing"
        @keydown.esc.stop="cancelEditing">
        {{ text }}
    </span>
    <span
        v-else
        :class="[theme] | prefix('records-list__editable-span--')"
        class="records-list__editable-span">
        {{ text }}
    </span>
</template>

<script>
export default {
    props: {
        value: { type: String, required: true },
        editing: { type: Boolean, required: true },
        theme: { type: String, required: false, default: 'light' },
        maxCharacters: { type: Number, required: false, default: null },
        maxVisibleLength: { type: Number, required: false, default: null },
        trackName: { type: String, default: 'editableSpan' },
    },
    computed: {
        text () {
            return this.maxVisibleLength ? this.$options.filters.middleEllipsis(this.value, this.maxVisibleLength) : this.value
        },
    },
    watch: {
        editing (val) {
            if (val) {
                this.$nextTick(() => {
                    const range = document.createRange()
                    range.selectNodeContents(this.$refs.span)
                    const sel = window.getSelection()
                    sel.removeAllRanges()
                    sel.addRange(range)
                    this.$refs.span.focus()
                    this.$emit('focus')
                    this.$root.$emit('tracking-event', { type: 'input', label: this.trackName, trigger: 'focus' })
                })
            }
        },
    },
    methods: {
        cancelEditing (ev) {
            if (this.editing) {
                window.getSelection().removeAllRanges()
                this.$emit('cancel')
            }
        },
        stopEditing (ev) {
            if (this.editing) {
                window.getSelection().removeAllRanges()
                this.$emit('input', this.$refs.span.textContent)
            }
        },
        handleClick (ev) {
            if (this.editing) {
                ev.stopPropagation()
                ev.preventDefault()
            }
        },
        checkLength (ev) {
            const allowedKeys = [8, 37, 39]
            if (window.getSelection().isCollapsed && !allowedKeys.includes(ev.keyCode) && this.maxCharacters && this.$refs.span.textContent.length >= this.maxCharacters) {
                ev.preventDefault()
            }
        },
        handlePaste () {
            this.$nextTick(() => {
                let withoutLinebreaks = this.$refs.span.innerText.replace(/\r?\n|\r/g, '')
                if (this.maxCharacters) {
                    withoutLinebreaks = withoutLinebreaks.slice(0, this.maxCharacters)
                }
                if (this.$refs.span.innerText !== withoutLinebreaks) {
                    this.$refs.span.innerText = withoutLinebreaks
                }
            })
        },
    },
}
</script>

<style lang="less" scoped>
@import (reference) '~design-system/less/variables';

.records-list__editable-span {
    display: block;
    max-width: 100%;
    overflow: hidden;
    white-space: nowrap;
    outline: none;

    &--editing {
        user-select: text;
        cursor: text;
    }
}
</style>
