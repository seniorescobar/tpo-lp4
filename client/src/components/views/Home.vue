<template>
    <div class="column__wrapper">
        <div class="column">
            
        </div>

        <div class="column">
            <div class="todo__wrapper">
                <div class="todo__actions">
                    <chip class="todo__action" label="Add"/>
                    <chip class="todo__action" label="Remove"/>
                </div>
                <div class="todo__container">
                    <div v-for="todo in mockData.todos" :key="todo.id" class="todo">{{ todo.description }}</div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { Chip } from 'design-system'
import { mapGetters } from 'vuex'

export default {
    components: {
        Chip
    },
    computed: {
        ...mapGetters({
            mockData: 'mockData'
        }),
    }
}
</script>

<style lang="less" scoped>
@import '../../less/variables';
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
