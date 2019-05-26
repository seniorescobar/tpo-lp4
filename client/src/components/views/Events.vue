<template>
    <div class="event__wrapper">
        <div class="event__container">
            <div v-for="event in events" :key="event.id" class="event">
                <img :src="event.thumbnailUrl" class="event__thumbnail">
                <div>
                    <div class="event__title">{{ event.title }}</div>
                    <div class="event__description">{{ event.description }}</div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    beforeCreate () {
        this.events = [
            { id: 1, title: 'Wooo1', description: 'lorem ipsum dolor aawijaisjd', thumbnailUrl: 'https://partyflock.nl/ov/images/organization/9883_404x404_501008/Afterlife.jpg' },
            { id: 2, title: 'Wooo2', description: 'lorem ipsum dolor aawijaisjd', thumbnailUrl: 'https://partyflock.nl/ov/images/organization/9883_404x404_501008/Afterlife.jpg' },
            { id: 3, title: 'Wooo3', description: 'lorem ipsum dolor aawijaisjd', thumbnailUrl: 'https://partyflock.nl/ov/images/organization/9883_404x404_501008/Afterlife.jpg' },
            { id: 4, title: 'Wooo4', description: 'lorem ipsum dolor aawijaisjd', thumbnailUrl: 'https://partyflock.nl/ov/images/organization/9883_404x404_501008/Afterlife.jpg' },
        ]
    }
}
</script>

<style lang="less" scoped>
@import '../../less/common';

.event {
    display: flex;
    margin-bottom: 16px;
    padding: 24px;
    background-color: lighten(@purple, 5%);

    &__container {
        margin: 64px auto;
        width: 100%;
        max-width: 600px;
    }
    &__thumbnail {
        object-fit: cover;
        width: 240px;
        height: 180px;
        margin-right: 24px;
    }
    &__title {
        font-size: 32px;
        margin-bottom: 12px;
    }
}
</style>
