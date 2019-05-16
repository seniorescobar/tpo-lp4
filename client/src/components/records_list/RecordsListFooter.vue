<template>
    <div class="records-list-footer">
        <div class="records-list-footer__left">
            {{ footerLeftText }}
        </div>

        <div v-if="numPages > 1" class="records-list-footer__center">
            <div v-for="link in pageLinks" :class="{'active': link.isActive, 'disabled': link.isEllipsis} | prefix('records-list-footer__link--')"
                 :key="link.page"
                 class="records-list-footer__link" @click="goToPage(link.page)" v-html="link.label">
            </div>
        </div>

        <div v-if="numPages > 1" class="records-list-footer__right">
            <div :class="{'records-list-footer__link--disabled': currentPage === 0}" class="records-list-footer__link"
                 @click="goToPage(currentPage - 1)">
                Previous page
            </div>
            <div :class="{'records-list-footer__link--disabled': currentPage === numPages - 1 || totalRecords === 0}"
                 class="records-list-footer__link"
                 @click="goToPage(currentPage + 1)">
                Next page
            </div>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        recordNameOverride: String,
        perPage: { type: Number, required: true },
        currentPage: { type: Number, required: true },
        totalRecords: { type: Number, required: true },
        currentPageRecords: { type: Number, required: true },
        trackName: { type: String, default: 'goTo' },
    },
    computed: {
        numPages () {
            return Math.ceil(this.totalRecords/this.perPage)
        },
        footerLeftText () {
            if (this.currentPage === 0) {
                return `Showing ${Math.min(this.currentPageRecords)} out of ${this.totalRecords} total records`
            } else {
                const start = this.currentPage * this.perPage
                const end = start + this.currentPageRecords
                return `Showing ${start}-${end} out of ${this.totalRecords} total records`
            }
        },
        pageLinks () {
            if (this.numPages <= 9) {
                /* [1 2 3 4 5 6 7 8 9] */
                return Array.from(Array(this.numPages)).map((_, idx) => {
                    return { label: idx + 1, page: idx, isActive: idx === this.currentPage }
                })
            } else {
                let indexes = []

                if (this.currentPage < 3 || this.currentPage > this.numPages - 4) {
                    /* [1 2 3] 4 ... 17 [18 19 20] */
                    indexes = [0, 1, 2, 3, '...', this.numPages - 4, this.numPages - 3, this.numPages - 2, this.numPages - 1]
                } else if (this.currentPage === 3) {
                    /* 1 2 3 [4] 5 ... 18 19 20 */
                    indexes = [0, 1, 2, 3, 4, '...', this.numPages - 3, this.numPages - 2, this.numPages - 1]
                } else if (this.currentPage >= 4 && this.currentPage < this.numPages - 4) {
                    /* 1 ... 3 4 [5] 6 7 ... 20 */
                    indexes = [0, '...', this.currentPage - 2, this.currentPage - 1, this.currentPage, this.currentPage + 1, this.currentPage + 2, '...', this.numPages - 1]
                } else if (this.currentPage === this.numPages - 4) {
                    /* 1 2 3 ... 16 [17] 18 19 20 */
                    indexes = [0, 1, 2, '...', this.currentPage - 1, this.currentPage, this.currentPage + 1, this.currentPage + 2, this.numPages - 1]
                }

                return indexes.map((pageIndex, idx) => {
                    if (pageIndex === '...') {
                        return { label: '&hellip;', page: -idx, isActive: false, isEllipsis: true }
                    } else {
                        return { label: pageIndex + 1, page: pageIndex, isActive: this.currentPage === pageIndex }
                    }
                })
            }
        },
    },
    methods: {
        goToPage (page) {
            this.$emit('go-to-page', page)
            this.$root.$emit('tracking-event', { type: 'pagination', label: this.trackName, trigger: 'click', data: { currentPage: this.currentPage + 1, goToPage: page + 1 } }) // pages need to be padded with + 1 so we track human-readable pages
        },
    },
}
</script>

<style lang="less" scoped>
@import (reference) '~design-system/less/variables';

.records-list-footer {
    position: relative;
    z-index: 25;
    box-sizing: border-box;
    width: 100%;
    height: 60px;
    display: flex;
    border-top: 1px solid @extremely-light-gray;
    background-color: white;
    padding: 0 25px;
    justify-content: space-between;
    align-items: center;
    .regular-font();
    font-size: 14px;
    color: @bluish-gray;

    &__left, &__center, &__right {
        display: flex;
    }

    &__link {
        height: 30px;
        line-height: 30px;
        padding: 0 15px;
        color: @dark-gray;
        transition: background-color 200ms ease-in;
        cursor: pointer;

        &:hover:not(.records-list-footer__link--disabled):not(.records-list-footer__link--active) {
            color: @black;
        }

        &--active {
            color: @royal-blue;
        }

        &--disabled {
            color: @bluish-gray;
            cursor: default;
            pointer-events: none;
        }
    }

}
</style>
