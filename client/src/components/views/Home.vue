<template>
    <div class="column__wrapper">
        <div class="column">
        </div>

        <div class="column">
            <div class="todo__wrapper">
                <div class="todo__actions">
                    <chip class="todo__action" label="Add" @click="openDialog"/>
                    <chip class="todo__action" label="Remove"/>
                </div>
                <div class="todo__container">
                    <div v-for="todo in mockData.todos" :key="todo.id" class="todo">{{ todo.description }}</div>
                </div>
            </div>
        </div>

        <md-dialog :md-active.sync="isDialogShown" class="">
            test
        </md-dialog>
    </div>
</template>

<script>
import { Chip, Input } from 'design-system'
import { mapGetters } from 'vuex'

export default {
    components: {
        Chip,
        InputElement: Input,
    },
    data () {
        return {
            isDialogShown: false,
            calendarData: {},
        }
    },
    computed: {
        ...mapGetters({
            mockData: 'mockData'
        }),
    },
    methods: {
        openDialog () {
            this.isDialogShown = true
        }
    }
}
</script>

<style lang="less" scoped>
@import '../../less/common';

.column {
    flex: 1 1 0;
    margin-right: 32px;

    &:last-child { margin-right: 0; }

    &__wrapper {
        display: flex;
        margin: 0 48px;
    }
}

.todo {
    height: 240px;
    width: 280px;
    background-color: @orange-yellow;
    overflow: scroll;
    color: black;
    margin: 16px;
    padding: 32px;
    box-sizing: border-box;

    &__wrapper {
        overflow: scroll;
    }
    &__actions {
        display: flex;
        border-bottom: 2px solid @dolphin;
    }
    &__action {
        cursor: pointer;
        margin: 12px 0;
    }
    &__container {
        display: flex;
        margin: 0 -16px;
    }
}
</style>
