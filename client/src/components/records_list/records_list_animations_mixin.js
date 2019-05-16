/*
    Contains logic for handling records list insertion and deletion animations and hover "ghost" animation.

    Requires refs set on RecordsListFooter component (ref="footer")
    and on transition-grop containing rows (ref="recordsListRows") for records list animations.

    Add it to list component and wrap records list rows in transition group:
        <transition-group :name="rowsTransitionGroupName" @after-enter="handleAnimationEnd" @after-leave="handleAnimationEnd">
            <records-list-row v-for="1 in 24"></records-list-row>
        <transition-group>

    Then call any functions that add or remove records like this:
        this.handleFooterAndAnimate(<number of records that will be added>, <update function>)

        - where number of records is 2 if function will add 2 records and -2 if 2 records will be deleted, while <update function>
        is the function that will actually remove or insert records when called

 */

export default {
    data () {
        return {
            rowsTransitionGroupName: null,
        }
    },
    created () {
        this.rowHeight = 60
    },
    mounted () {
        this.$nextTick(() => {
            this.observer = new MutationObserver(this.rowsMutationHandler)
            this.observer.observe(this.$refs.recordsListContent, { childList: true, subtree: true })
            this.rowsMutationHandler()
        })
    },
    destroyed () {
        this.observer.disconnect()
    },
    computed: {
        isIntersectionObserverAvailable () {
            // This doesn't work in Safari, so we have to position top buttons differently and lose some fancy animations
            return window.IntersectionObserver !== undefined
        },
    },
    methods: {
        rowsMutationHandler () {
            if (this.$refs.recordsListRows !== undefined) {
                this.rowsMutationCallback()
            }
        },
        rowsMutationCallback () {
        },
        handleFooterAndAnimate (recordCountDifference, updateRecords, shouldAnimateFooter = true) {
            this.rowsTransitionGroupName = recordCountDifference > 0 ? 'adding-items' : 'removing-items'

            this.$nextTick(() => {
                if (shouldAnimateFooter && this.$refs.footer) {
                    this.$refs.footer.$el.style.transition = ''
                    window.requestAnimationFrame(() => {
                        updateRecords()
                        if (this.$refs.footer) { // can be gone after updateRecords()
                            this.$refs.footer.$el.style.transform = `translateY(${-this.rowHeight * recordCountDifference}px)`
                            window.requestAnimationFrame(() => {
                                this.$refs.footer.$el.style.transition = 'transform 200ms'
                                this.$refs.footer.$el.style.transform = 'translateY(0px)'
                            })
                        }
                    })
                } else {
                    updateRecords()
                }
            })
        },
    },
}
