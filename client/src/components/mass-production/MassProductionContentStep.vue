<template>
    <div class="new-dialog__step-wrapper">
        <div class="content-step__wrapper">  
            <div class="instance__tabs">
                <instance-tabs :instances="selectedSizes" :disabled="state.isInstanceTabsDisabled" v-model="activeTab"/>
            </div>

            <div class="instance">
                <div class="column__preview">
                    <div class="column__preview-panel">
                        <div v-if="state.isPreviewDisplayed" class="column__preview-wrapper">
                            <template-preview :template="activeInstance" :is-overflow-hidden="state.isOverflowHidden" :exposedId="state.exposedUnitId"/>
                            <!-- fuck all states though -->
                            <!-- add output and finish real quick -->
                        </div>
                        <div v-else class="column__preview-wrapper">
                            <img class="cropper__image" ref="imageCrop" @click="setHotspot"/>
                            <svg v-if="croppingOption === 'auto-crop' && areCropParametersSet" class="cropper__image-overlay" xmlns="http://www.w3.org/2000/svg" @click="setHotspot">
                                <defs>
                                    <mask id="mask" x="0" y="0" width="100%" height="100%">
                                        <rect x="0" y="0" width="100%" height="100%" fill="#fff"/>
                                        <circle :cx="x" :cy="y" :r="radius"/>
                                    </mask>
                                </defs>
                                <circle :cx="x" :cy="y" :r="radius - .8" class="cropper__hotspot" fill="none"/>
                                <circle :cx="x" :cy="y" :r="radius ? '8px' : 0" class="cropper__hotspot-point"/>
                                <rect x="0" y="0" width="100%" height="100%" fill-opacity="0.7" mask="url(#mask)"/>
                            </svg>
                        </div>

                        <svg v-if="state.isPreviewMaskDisplayed" class="cropper__image-overlay" xmlns="http://www.w3.org/2000/svg">
                            <defs>
                                <mask id="mask" x="0" y="0" width="100%" height="100%">
                                    <rect x="0" y="0" width="100%" height="200%" fill="#fff"/>
                                    <!-- <rect x="50%" y="50%" :width="activeInstance.size.width" :height="activeInstance.size.height" :transform="`translate(-${activeInstance.size.width / 2}, -${activeInstance.size.height / 2})`"/> -->
                                    <rect x="50%" y="100%" width="5" height="5"/>
                                </mask>
                            </defs>
                            <rect x="0" y="0" width="100%" height="100%" fill-opacity="0.7" mask="url(#mask)"/>
                        </svg>

                        <div class="column__preview-actions">
                            <div v-if="state.isZoomActionDisplayed" class="column__preview-actions__action">
                                <div class="column__preview-actions__action-icon">
                                    <svg width="16" height="16" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M12.7 11.2298c.9-1.1925 1.4-2.58384 1.4-4.1739C14.1 3.18012 11 0 7.1 0S0 3.18012 0 7.0559c0 3.8758 3.2 7.0559 7.1 7.0559 1.6 0 3.1-.4969 4.2-1.3913l3 2.9814c.2.1987.5.2981.7.2981.2 0 .5-.0994.7-.2981.4-.3976.4-.9938 0-1.3913l-3-3.0808zm-5.6.795c-2.8 0-5.1-2.18629-5.1-4.9689 0-2.78261 2.3-5.06832 5.1-5.06832s5.1 2.28571 5.1 5.06832-2.3 4.9689-5.1 4.9689z" fill="#B2B2B8"/></svg>
                                </div>
                                <div class="column__preview-actions__action-label">100%</div>
                            </div>
                            <div v-if="state.isReplayActionDisplayed" class="column__preview-actions__action">
                                <div class="column__preview-actions__action-icon">
                                    <svg width="16" height="16" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M8.71123 0l.69875 6.48649 2.09622-2.12838c1.8967 1.92567 1.8967 5.16892 0 7.09459-.8984 1.0135-2.19604 1.5203-3.49372 1.5203-1.29768 0-2.59537-.5068-3.49376-1.5203-1.89662-1.92567-1.89662-5.16892 0-7.09459.59893-.60811 1.3975-1.11487 2.2959-1.31757l-.59893-1.92568c-1.19786.30406-2.2959.91217-3.1943 1.82433-2.695187 2.73649-2.695187 7.19591 0 10.03381C4.31907 14.2905 6.11586 15 7.91266 15c1.89661 0 3.59354-.7095 4.89124-2.027 2.6952-2.7365 2.6952-7.19597 0-10.03381L15 .709459 8.71123 0z" fill="#B2B2B8"/></svg>
                                </div>
                                <div class="column__preview-actions__action-label">REPLAY</div>
                            </div>
                            <div v-if="state.isSyncActionDisplayed" class="column__preview-actions__action">
                                <checkbox v-model="isContentOverlaySimulated" size="condensed" :is-toggle="true">Simulate content overlay</checkbox>
                            </div>
                        </div>
                    </div>
                </div>

                <div v-if="state.areGlobalAttributesDisplayed" class="column__content">
                    <div class="column__options">
                        <div v-for="unit in activeInstance.units" :key="unit.id" class="attribute">
                            <component-switcher :options="options" default="manual">
                                <file-upload   slot="manual" v-if="unit.type === 'asset'" :file="customAttributes[unit.id]" @update="value => setGlobalAttribute({ id: unit.id, type: unit.type, value })" @open-image-crop="openImageCropDialog(unit.id)"/>
                                <input-element slot="manual" v-else :value="customAttributes[unit.id]" :label="unit.label" @input="value => setGlobalAttribute({ id: unit.id, value })"/>
                                <selectbox     slot="feed"   :options="feeds[unit.id]" :value="globalAttributes[unit.id]" @input="value => setGlobalAttribute({ id: unit.id, value })"/>
                            </component-switcher>
                        </div>
                    </div>

                    <div class="column__actions">
                        <div class="attribute__reset">
                            <svg width="14" height="14" xmlns="http://www.w3.org/2000/svg"><path d="M7 2.125c-2.42 0-4.466 1.708-4.865 4.06L.188 5.864C.748 2.568 3.612.175 7 .175a6.97 6.97 0 0 1 4.564 1.707L13.468 0 14 6.055 7.871 5.53l2.292-2.264c-.883-.728-2-1.14-3.163-1.14zm0 9.75c2.42 0 4.466-1.708 4.865-4.06l1.947.322c-.56 3.295-3.424 5.688-6.812 5.688a6.97 6.97 0 0 1-4.564-1.707L.532 14 0 7.945l6.129.525-2.292 2.264c.883.728 2 1.14 3.163 1.14z" fill="#FFF" fill-rule="nonzero" opacity=".25"/></svg>
                            <div class="attribute__reset-button">Reset</div>
                        </div>
                    </div>
                </div>
                <div v-else class="column__content">
                    <div class="column__options">
                        <radio-button v-model="croppingOption" value="no-crop">Don't crop</radio-button>

                        <radio-button v-model="croppingOption" value="auto-crop">Auto focus cropping</radio-button>
                        <div class="cropper__option">
                            <p v-if="croppingOption === 'auto-crop'">Click and dragg the parts of the image that are important. Click on APPLY button will automatically apply crop to all the selected sizes.</p>
                            <p v-else>A clear description of what auto focus cropping does and how is it applied to ALL the sizes.</p>
                            <template v-if="croppingOption === 'auto-crop'">
                                <slider :min="0" :max="10" :step="0.1" v-model="radiusValue" :disabled="!areCropParametersSet" :is-input-hidden="true" label="Radius"/>
                                <button class="cropper__crop-button" :class="{ 'disabled': !areCropParametersSet } | prefix('cropper__crop-button--')" @click="crop">Apply to all sizes</button>
                            </template>
                        </div>

                        <radio-button v-model="croppingOption" value="manual-crop">Specific size cropping</radio-button>
                        <div class="cropper__option">
                            <p v-if="croppingOption === 'manual-crop'">Resize, re-position and re-crop by using the handles and panning the image. Click save crop to confirm.</p>
                            <p v-else>A clear description of what specific size cropping does and how is it applied to ONLY ONE size.</p>
                            <template v-if="croppingOption === 'manual-crop'">
                                <slider :min="0" :max="10" :step="0.1" :value="scale" @input="value => setLocalImageHotspotScale({ instanceId: activeInstance.sizeId, assetId: imageToCropId, scale: value })" :is-input-hidden="true" label="Scale"/>
                                <slider :min="0" :max="360" minLabel="0°" maxLabel="360°" :step="10" :value="angle" @input="value => setLocalImageHotspotAngle({ instanceId: activeInstance.sizeId, assetId: imageToCropId, angle: value })" :is-input-hidden="true" label="Rotate"/>
                            </template>
                        </div>
                    </div>
                    <div class="column__actions">
                        <dialog-button :warning="true">Cancel</dialog-button>
                        <dialog-button @click="isCropDialogOpen = false">Save</dialog-button>
                    </div>
                </div>
            </div>
        </div>

        <div class="new-dialog__button-container">
            <dialog-button v-if="state.isNextStepButtonDisplayed" @click="nextStep">Next step</dialog-button>
        </div>
    </div>
</template>

<script>
import TemplatePreview from '../template/TemplatePreview.vue'
import ComponentSwitcher from '../helpers/ComponentSwitcher.vue'
import InstanceTabs from '../InstanceTabs.vue'
import { FileUpload, Selectbox, Input, DialogButton, Checkbox, RadioButton, Slider } from 'design-system'
import { mapGetters, mapState, mapActions } from 'vuex'

export default {
    components: {
        FileUpload,
        TemplatePreview,
        Selectbox,
        Checkbox,
        RadioButton,
        Slider,
        DialogButton,
        ComponentSwitcher,
        InstanceTabs,
        inputElement: Input
    },
    beforeCreate () {
        this.options = [
            { id: 'feed', label: 'Feed' },
            { id: 'manual', label: 'Manual' },
        ]
    },
    data () {
        return {
            activeTab: null,
            imageToCropId: null,
            isCropDialogOpen: false,
            areCropParametersSet: false,
            isCropApplied: false,
            isContentOverlaySimulated: false,
            croppingOption: 'no-crop',
            radiusValue: 0,
            x: null,
            y: null,
        }
    },
    computed: {
        ...mapState('massProduction', [
            'globalAttributes',
            'localImageHotspot',
        ]),
        ...mapGetters('massProduction', {
            selectedSizes: 'getSelectedSizes',
            instances: 'getPopulatedInstances',
            feeds: 'getFeeds',
        }),
        activeInstance () {
            return this.instances.find(instance => instance.sizeId === this.activeTab)
        },
        customAttributes () {
            const obj = {}
            this.activeInstance.units.forEach(unit => obj[unit.id] = unit.type === 'asset' ? {} : '')
            Object.keys(this.globalAttributes).forEach(key => obj[key] = this.globalAttributes[key])
            return obj
        },
        state () {
            return {
                isInstanceTabsDisabled: this.isCropDialogOpen && !this.areCropParametersSet && this.croppingOption !== 'manual-crop',
                isPreviewDisplayed: !this.isCropDialogOpen || this.areCropParametersSet && (this.croppingOption !== 'auto-crop' || this.isCropApplied) || this.croppingOption === 'manual-crop',
                isZoomActionDisplayed: !this.isCropDialogOpen,
                isReplayActionDisplayed: !this.isCropDialogOpen,
                isSyncActionDisplayed: this.isCropDialogOpen && this.areCropParametersSet && (this.croppingOption !== 'auto-crop' || this.isCropApplied) || this.croppingOption === 'manual-crop',
                isPreviewMaskDisplayed: false, // this.isCropDialogOpen && this.areCropParametersSet && (this.croppingOption !== 'auto-crop' || this.isCropApplied) || this.isCropDialogOpen && this.croppingOption === 'manual-crop',
                isNextStepButtonDisplayed: !this.isCropDialogOpen,
                areGlobalAttributesDisplayed: !this.isCropDialogOpen,
                isOverflowHidden: true, // !this.isCropDialogOpen,
                exposedUnitId: (this.isCropDialogOpen && !this.isContentOverlaySimulated) ? this.imageToCropId : ''
            }
        },
        radius () {
            // halve to get radius, decimate to get percentage
            return Math.min(this.clientWidth, this.clientHeight) * this.radiusValue / 2 / 10
        },
        scale () {
            return this.localImageHotspot[this.activeInstance.sizeId] && this.localImageHotspot[this.activeInstance.sizeId][this.imageToCropId] && this.localImageHotspot[this.activeInstance.sizeId][this.imageToCropId].scale || 0 
        },
        angle () {
            return this.localImageHotspot[this.activeInstance.sizeId] && this.localImageHotspot[this.activeInstance.sizeId][this.imageToCropId] && this.localImageHotspot[this.activeInstance.sizeId][this.imageToCropId].angle || 0
        }
    },
    created () {
        this.activeTab = this.selectedSizes[0].id
    },
    methods: {
        ...mapActions('massProduction', [
            'setGlobalAttribute',
            'setGlobalImageHotspot',
            'setLocalImageHotspotScale',
            'setLocalImageHotspotAngle',
            'resetGlobalAttribute',
            'nextStep',
        ]),
        openImageCropDialog (unitId) {
            this.imageToCropId = unitId
            this.isCropDialogOpen = true
            const fileReader = new FileReader()
            fileReader.onloadend = () => { this.$refs.imageCrop.src = fileReader.result }
            fileReader.readAsDataURL(this.globalAttributes[unitId].file)
        },
        setHotspot (e) {
            this.croppingOption = 'auto-crop'
            this.clientWidth = e.srcElement.clientWidth
            this.clientHeight = e.srcElement.clientHeight
            this.x = e.offsetX
            this.y = e.offsetY
            this.radiusValue = this.radiusValue || 5
            this.areCropParametersSet = true
        },
        crop () {
            this.isCropApplied = true
            this.setGlobalImageHotspot({ id: this.imageToCropId, x: this.x / this.clientWidth, y: this.y / this.clientHeight, r: this.radiusValue })
        }
    },
}
</script>

<style lang="less" scoped>
@import (reference) '../../less/common';
@import '../../less/dialogs_common';
@import '../../less/dialogs_columns_layout';

.content-step {
    &__wrapper {
        max-width: 1184px;
        width: 100%;
        margin: 0 auto;
        height: 100%;
        display: flex;
        flex-direction: column;
    }
}

.instance {
    display: flex;
    justify-content: center;
    width: 100%;
    flex: 1 1 0;

    &__tabs {
        margin-bottom: 16px;
    }
}

.column {
    &__content {
        display: flex;
        flex-direction: column;
        height: 100%;
    }

    &__actions {
        display: flex;
    }

    &__options {
        min-width: 190px;
        max-width: 320px;
        width: 100%;
        flex: 1 1 0;
        overflow: auto;
    }

    &__preview {
        margin-right: 64px;
        max-width: 800px;
        flex: auto;
    }

    &__preview-panel {
        display: flex;
        width: 100%;
        height: 100%;
        background-color: @extremely-dark-gray;
        position: relative;
        overflow: hidden;
    }

    &__preview-actions {
        position: absolute;
        top: 0;
        right: 0;
        display: flex;

        &__action {
            display: flex;
            color: @dolphin;
            height: 64px;
            margin-right: 24px;

            &-icon {
                height: fit-content;
                margin: auto 8px auto 0;
            }

            &-label { margin: auto; }
        }
    }

    &__preview-wrapper {
        margin: auto;
        position: relative;
    }
}

.attribute {
    width: 320px;
    margin: 32px 0;
    min-height: 80px;
    display: flex;
    flex-direction: column;

    &__value {
        width: 100%;
    }

    &__source {
        display: flex;
        justify-content: space-between;
        margin-bottom: 8px;
        color: @bluish-gray;
        font-size: 14px;
    }

    &__reset {
        display: flex;
        height: 48px;
        cursor: pointer;

        &-button {
            color: @bluish-gray;
            margin: auto auto auto 0;
        }
        
        svg {
            margin: auto 9px auto auto;
        }
    }
}


.cropper {
    &__option {
        margin: -15px 0 20px 35px;
        font-size: 12px;
        color: @dolphin;
    }

    &__crop-button {
        width: 165px;
        height: 32px;
        text-transform: uppercase;
        background-color: @light-smoke;
        border: transparent;
        outline: none;
        cursor: pointer;

        &--disabled { background-color: @dolphin; }
    }

    &__image {
        max-width: 100%;
        max-height: 100%;
        height: auto;
        width: auto;
    }

    &__image-overlay {
        position: absolute;
        width: 100%;
        height: 100%;
        left: 0;
        cursor: crosshair;

        & * {
            pointer-events: none;
        }
    }

    &__hotspot {
        pointer-events: none;
        stroke: @royal-blue;
        stroke-width: 2px;
    }

    &__hotspot-point {
        pointer-events: none;
        fill: @royal-blue;
    }
}
</style>