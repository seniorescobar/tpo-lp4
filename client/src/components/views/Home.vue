<template>
    <div class="column__wrapper">
        <div class="column">
            <div class="calendar">
                
            </div>
            <div class="schedule">
                <div v-if="!isStudentIdSet" class="schedule__input">
                    <input-element v-model="studentIdModel" label="Your student ID"/>
                    <dialog-button @click="setStudentIdSet">Confirm</dialog-button>
                </div>
                <iframe v-else class="iframe" :src="calendarSrc" />
            </div>
        </div>

        <div class="column">
            <div class="todo__wrapper">
                <div class="todo__actions">
                    <chip class="todo__action" label="Add" @click="openDialog"/>
                </div>
                <div class="todo__container">
                    <div v-for="todo in todos" :key="todo._id" class="todo">
                        <div class="todo__remove" @click="remove(todo._id)">X</div>
                        {{ todo.description }}
                    </div>
                </div>
            </div>
        </div>

        <md-dialog :md-active.sync="isDialogShown" class="">
            <textarea ref="input" class="todo__input" v-model="description"/>
            <button class="todo__post" @click="post">POST</button>
        </md-dialog>
    </div>
</template>

<script>
import { Chip, Input, DialogButton } from 'design-system'
import { mapState, mapActions } from 'vuex'

export default {
    components: {
        Chip,
        DialogButton,
        InputElement: Input,
    },
    data () {
        return {
            isDialogShown: false,
            description: '',
        }
    },
    computed: {
        ...mapState(['todos', 'studentId', 'isStudentIdSet']),
        studentIdModel: {
            get () {
                return this.studentId
            },
            set (id) {
                this.setStudentId(id)
            },
        },
        calendarSrc () {
            return `https://urnik.fri.uni-lj.si/timetable/fri-2018_2019-letni-1-13/allocations?student=${this.studentId}`
        }
    },
    methods: {
        ...mapActions(['fetchAndSetTodos', 'postTodo', 'deleteTodo', 'setStudentId', 'setStudentIdSet']),
        openDialog () {
            this.isDialogShown = true
            setTimeout(() => {
                this.$refs.input.focus()
            }, 250)
        },
        async post () {
            if (!this.description.length) return

            const payload = { description: this.description }
            const res = await this.postTodo(payload)
            if (res) {
                this.description = ''
                this.isDialogShown = false
            } else {
                alert('Error in posting todo')
            }
        },
        async remove (id) {
            return await this.deleteTodo(id)
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

.iframe {
    background-color: @white;
    width: 100%;
    margin-top: 64px;
    height: 460px;
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
    position: relative;

    &__remove{
        position: absolute;
        color: @dolphin;
        right: 8px;
        top: 8px;
        cursor: pointer;
    }
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
        display: grid;
        grid-template-columns: 1fr 1fr;
    }
    &__input {
        background-color: transparent;
        outline: none;
        margin: 12px;
        resize: none;
        border: none;
        font-size: 12px;
        color: @white;
    }
    &__post {
        background: @very-light-green;
        outline: none;
        border: none;
        color: @black;
        height: 28px;
    }
}
</style>
