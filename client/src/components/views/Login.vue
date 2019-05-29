<template>
    <div class="login__wrapper">
        <div class="login__container">
            <div class="login__text">Log in!</div>
            <input class="login__credential" placeholder="Email" type="email" v-model="email"/>
            <input class="login__credential" placeholder="Password" type="password" v-model="password"/>
            <dialog-button :disabled="!isLoginEnabled" @click="login">Log in</dialog-button>
            <div class="login__register" @click="goToRegister">Or sign up!</div>
            <div v-if="errorMessage.length" class="login__error">{{ errorMessage }}</div>
        </div>
    </div>
</template>

<script>
import api from 'api-client'
import { DialogButton } from 'design-system'

export default {
    components: {
        DialogButton,
    },
    data () {
        return {
            email: '',
            password: '',
            errorMessage: '',
            isLoggingIn: false,
        }
    },
    computed: {
        isLoginEnabled () {
            return this.email 
                && this.password
                && /\S+@\S+\.\S+/.test(this.email)
                && this.password.length > 5
        },
    },
    methods: {
        goToRegister () {
            this.$router.push('/register')
        },
        login () {
            const self = this
            this.isLoggingIn = true
            this.errorMessage = ''
            const credentials = { email: this.email, password: this.password }

            return api
                .post('auth/signin', credentials)
                .then(success)
                .catch(failure)

            function success (user) {
                localStorage.setItem('user', JSON.stringify(user))
                self.$router.push('/app/home')
            }
            function failure () {
                self.isLoggingIn = false
                self.errorMessage = 'Failed to log in :('
            }
        },
    }
}
</script>

<style lang="less">
@import '../../less/common';

.login {
    &__wrapper {
        min-height: 100vh;
        display: flex;
    }
    &__container{
        width: fit-content;
        margin: auto;
        padding: 30px 30px 60px;
    }
    &__text {
        font-size: 32px;
        text-align: center;
        margin-bottom: 32px;
    }
    &__credential {
        display: block;
        border: 2px solid @dolphin;
        border-radius: 3px;
        outline: none;
        background-color: transparent;
        padding: 6px;
        width: 260px;
        font-size: 16px;
        color: @white;
        margin-bottom: 12px;
    }
    &__register {
        text-align: center;
        color: @dolphin;
        font-size: 14px;
        text-decoration: underline;
        margin-top: 12px;
        cursor: pointer;
    }
    &__error {
        text-align: center;
        margin-top: 12px;
        color: @pink-red;
    }
}
</style>
