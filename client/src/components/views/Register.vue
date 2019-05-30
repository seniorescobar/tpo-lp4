<template>
    <div class="register__wrapper">
        <div class="register__container">
            <div class="register__text">Sign up!</div>
            <input class="register__credential" placeholder="Email" type="email" v-model="email"/>
            <input class="register__credential" placeholder="Password" type="password" v-model="password"/>
            <input class="register__credential" placeholder="Repeat password" type="password" v-model="passwordRepeat"/>
            <dialog-button :disabled="!isRegisterEnabled" @click="register">Sign up</dialog-button>
            <div class="register__login" @click="goToLogin">Or login!</div>
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
            passwordRepeat: '',
            errorMessage: '',
            isRegistring: false,
        }
    },
    computed: {
        isRegisterEnabled () {
            return this.email 
                && this.password
                && /\S+@\S+\.\S+/.test(this.email)
                && this.password.length > 5
                && this.password === this.passwordRepeat
        },
    },
    methods: {
        goToLogin () {
            this.$router.push('/login')
        },
        register () {
            const self = this
            this.isRegistring = true
            this.errorMessage = ''
            const credentials = { email: this.email, password: this.password }

            return api
                .post('auth/register', credentials)
                .then(success)
                .catch(failure)

            function success (user) {
                localStorage.setItem('user', JSON.stringify(user))
                self.$router.push('/app/home')
            }
            function failure () {
                self.isRegistring = false
                self.errorMessage = 'Failed to log in :('
            }
        },
    }
}
</script>

<style lang="less" scoped>
@import '../../less/common';

.register {
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
    &__login {
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
