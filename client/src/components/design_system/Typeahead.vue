<template>
    <div v-click-outside="close" class="typeahead">
        <input-element ref="input" v-bind="inputData" :error="inputError" :label="label" :track-name="trackName" class="typeahead__input" @keyup.enter="selectFirstItem()" @keyup.down="onDown" @focus="onInputFocus" @input="onInput" @blur="onBlur"></input-element>

        <template v-if="showSuggestions">
            <scrollable-list v-if="suggestions.length > 0" ref="list" :items="suggestions" :num-items="10" :highlight-query="value" class="typeahead__suggestions" theme="light" @select="onSelect" @blur="onBlur"/>
            <div v-else-if="noItemsText" class="typeahead__no-items-text">{{ noItemsText }}</div>
        </template>
    </div>
</template>

<script>
import Input from './Input.vue'
import ScrollableList from './ScrollableList.vue'
import debounce from 'lodash.debounce'

export default {
    components: {
        inputElement: Input,
        ScrollableList,
    },
    props: {
        label: { type: String },
        value: { type: [String, Number], default: '' },
        getSuggestions: { type: Function, required: true },
        noItemsText: { type: String, default: 'No items' },
        isValid: { type: Function, required: false },
        trackName: { type: String, default: 'typeahead' },
    },
    data () {
        return {
            isOpen: false,
        }
    },
    computed: {
        inputData () {
            return {
                ...this.$attrs,
                value: this.value,
                error: null,
            }
        },
        showSuggestions () {
            return this.isOpen && (this.isValueValid || this.suggestions.length > 0) && !(this.value === null || typeof this.value === 'string' && this.value.length <= 2)
        },
        suggestions () {
            return this.getSuggestions(this.value)
        },
        inputError () {
            if (this.showSuggestions || this.suggestions.filter(item => !item.disabled).length > 0 || !this.isValid || ( this.value === null || typeof this.value === 'string' && this.value.length <= 2)) {
                return null
            }
            return this.isValid(this.value)
        },
        isValueValid () {
            return this.isValid ? this.isValid(this.value) === null : true
        },
    },
    watch: {
        value (text) {
            this.$nextTick(() => {
                this.highlightFirstItem()
                // Reset input state to normal
                if (text && text.length <= 2) {
                    this.$refs.input.errorMessage = null
                }
            })
        },
        isOpen () {
            this.$nextTick(() => {
                this.highlightFirstItem()
                // Reset input state to normal
                if (this.isOpen) {
                    this.$refs.input.errorMessage = null
                }
            })
        },
    },
    methods: {
        focus () {
            this.$refs.input.focus()
        },
        onInputFocus () {
            this.isOpen = true
            this.$emit('focus')
        },
        close () {
            this.isOpen = false
            this.$emit('blur')
        },
        onBlur (ev) {
            if (!this.$el.contains(ev.relatedTarget)) {
                this.close()
            }
        },
        onInput (v) {
            this.isOpen = true
            this.$emit('input', v)
            if (!this.onInputTrackingDebounced) {
                this.onInputTrackingDebounced = debounce(() => {
                    this.$root.$emit('tracking-event', { type: 'input', label: this.trackName, trigger: 'search' })
                }, 1000)
            }
            this.onInputTrackingDebounced()
        },
        onSelect (suggestion) {
            this.$emit('input', suggestion.label)
            this.$emit('select', suggestion)
            this.close()
        },
        highlightFirstItem () {
            if (this.showSuggestions) {
                const firstEnabledIndex = this.suggestions.findIndex(item => !item.disabled)
                if (firstEnabledIndex > -1 ) {
                    this.$refs.list.highlightItem(firstEnabledIndex)
                }
            }
        },
        selectFirstItem () {
            if (this.showSuggestions) {
                const firstEnabledIndex = this.suggestions.findIndex(item => !item.disabled)
                if (firstEnabledIndex > -1 ) {
                    this.onSelect(this.suggestions[0])
                }
            }
        },
        onDown (ev) {
            if (this.$refs.list) {
                this.$refs.list.focus()
                this.$refs.list.onKeyDown(ev)
            }
        },
    },
}
</script>

<style lang="less" scoped>
@import (reference) './less/common';
@import './less/typography';

.typeahead {
    .regular-font();
    position: relative;

    &__suggestions {
        position: absolute;
        margin-top: -7px;
        background-color: white;
        box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.3);
        padding: 15px 0;
        z-index: @z-heaven;
    }

    &__no-items-text {
        position: absolute;
        margin-top: -7px;
        box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.3);
        padding: 20px;
        width: 100%;
        box-sizing: border-box;
        background: white;
        font-size: 16px;
    }
}
</style>
