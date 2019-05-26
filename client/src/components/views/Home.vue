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
                    <chip class="todo__action" label="Remove"/>
                </div>
                <div class="todo__container">
                    <div v-for="todo in todos" :key="todo.id" class="todo">{{ todo.description }}</div>
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
import { mapGetters, mapState, mapActions } from 'vuex'
import api from 'api-client'

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
        ...mapActions(['fetchAndSetTodos', 'setStudentId', 'setStudentIdSet']),
        openDialog () {
            this.isDialogShown = true
            setTimeout(() => {
                this.$refs.input.focus()
            }, 250)
        },
        post () {
            if (!this.description.length) return

            const payload = { description: this.description }
            return api
                .post('todo/', payload)
                .then(res => {
                    this.fetchAndSetTodos()
                    this.description = ''
                    this.isDialogShown = false
                })
                .catch((err) => console.log(err))
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
