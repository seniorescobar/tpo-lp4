<template>
    <div class="new-dialog__step-wrapper">
        <div class="new-dialog__content-wrapper">
            <div class="topbar">
                <div class="topbar__tabs">
                    <instance-tabs :instances="selectedSizes" :disabled="true"/>
                </div>
                <div v-if="isCompacted" class="topbar__goal">
                    <selectbox size="condensed" placeholder="SELECT GOAL" :options="goals" v-model="goal"></selectbox>
                </div>
            </div>

            <div class="experiences__wrapper">
                <div v-if="!isCompacted" class="experiences__goal">
                    <div class="experiences__goal-label">I'm looking for an experience to</div>
                    <selectbox class="experiences__goal-selection" placeholder="select goal" :options="goals" size="phat" v-model="goal"></selectbox>
                </div>

                <div class="experiences__container">
                    <div v-for="group in groupedExperiences" :key="group.title" class="experience-group">
                        <div v-if="group.title" class="experience-group__title">{{ group.title }}</div>
                        <div class="experience-group__content">
                            <div v-for="experience in group.experiences" :key="experience.id" class="experience" :class="getClass(experience.id) | prefix('experience--')" @mouseenter="focusedExperience = experience.id" @mouseleave="focusedExperience = ''" @click="selectAndContinue(experience)">
                                <img class="experience__thumbnail" :src="`images/${experience.imageSource}`">
                                <div class="experience__title">{{ experience.title }}</div>
                                <div class="experience__description">{{ experience.description }}</div>
                                <div class="experience__description">{{ experience.attributes.join(', ') }}</div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>    
</template>

<script>
import { mapActions, mapGetters } from 'vuex'
import { Selectbox } from 'design-system'
import InstanceTabs from '../InstanceTabs.vue'

export default {
    components: {
        Selectbox,
        InstanceTabs
    },
    data () {
        return {
            isCompacted: false,
            goal: '',
            focusedExperience: ''
        }
    },
    computed: {
        ...mapGetters('massProduction', {
            goals: 'getGoals',
            selectedSizes: 'getSelectedSizes',
            experiences: 'getExperiences'
        }),
        groupedExperiences () {
            if (this.goal) {
                return [{ experiences: this.experiences.filter(experience => experience.goals.includes(this.goal)) }]
            } else {
                return [
                    { title: 'Featured experiences', experiences: this.experiences.filter(experience => experience.isFeatured) },
                    { title: 'Custom experiences', experiences: this.experiences.filter(experience => experience.isCustom) },
                ]
            }
        }
    },
    methods: {
        ...mapActions('massProduction', [
            'setExperience',
            'nextStep',
        ]),
        selectAndContinue (experience) {
            this.setExperience(experience.id)
            this.nextStep()
        },
        getClass (experienceId) {
            return {
                darkened: this.focusedExperience && this.focusedExperience !== experienceId
            }
        }
    }
}
</script>

<style lang="less" scoped>
@import (reference) '../../less/common';
@import '../../less/dialogs_common';
@import '../../less/dialogs_columns_layout';

.new-dialog__content-wrapper {
    height: 100%;
    width: 100%;
    max-width: 1216px;
    margin: 0 auto;
}

.topbar {
    display: flex;
    padding: 0 24px;

    &__tabs { flex: 1 1 0; }

    &__goal { width: 300px; }
}

.experiences {
    // &__wrapper {}

    &__goal {
        height: 80px;
        display: flex;
        margin: 48px auto 0;
        width: fit-content;

        &-label {
            line-height: 71px;
            color: @bluish-gray;
            font-size: 22px;
            margin-right: 16px;
        }
        
        &-selection {
            width: 320px;
        }
    }

    &__container {
        margin-top: 40px;
    }
}

.experience-group {
    &__content {
        display: flex;
        flex-wrap: wrap;
        margin: 0 -8px;
    }

    &__title {
        margin: 8px 0;
        padding-left: 16px;
        font-size: 20px;
        color: white;
    }
}

.experience {
    width: 292px;
    height: 380px;
    margin: 8px;
    padding: 16px;
    box-sizing: border-box;
    position: relative;
    cursor: pointer;
    transition: opacity 0.2s;

    &--darkened { opacity: 0.4; }

    &__thumbnail {
        width: 100%;
        margin-bottom: 8px;
    }

    &__title {
        margin: 8px 0;
        font-size: 16px;
        color: @light-gray;
    }

    &__description {
        font-size: 11px;
        line-height: 16px;
        color: @pale-gray;
    }
}
</style>