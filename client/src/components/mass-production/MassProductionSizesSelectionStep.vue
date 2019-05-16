<template>
    <div class="new-dialog__step-wrapper">
        <div v-if="!isAddingCustomSize" class="sizes__content-wrapper">
            <div class="sizes__tab-container">
                <instance-tabs :instances="selectedSizes" :disabled="true"/>
            </div>

            <div class="sizes__container">
                <div class="size" @click="isAddingCustomSize = true">
                    <div class="size__add"><svg width="48" height="48" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M24 4v40M44 24H4" stroke="#D4D4D9" stroke-width="2" stroke-miterlimit="10" stroke-linecap="square"/></svg></div>
                    <div class="size__text">Custom</div>
                </div>

                <div v-for="size in availableSizes" :key="size.id" class="size" :class="getSizeClass(size) | prefix('size--')" @click="toggleSelection(size.id)">
                    <div class="size__box"></div>
                    <div class="size__text size__text--bold">{{ size.label }}</div>
                    <div class="size__text">{{ size.size }}</div>
                </div>
            </div>
        </div>

        <div v-else class="custom-size__wrapper">
            <div class="column">
                <div class="column-section">
                    <div class="common-dialog__column-title">Outputs</div>
                    <div class="output__container">
                         <div v-for="output in outputTypes" :key="output.id" class="output" :class="getOutputClass(output) | prefix('output--')" @click="customSizeType = output.id">
                            <div class="output__icon">
                                <svg v-if="output.id === 'display'" width="24" height="24" fill="none" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M6.40557 18.3717L5 4h15l-1.4076 14.3717L12.5 20l-6.09443-1.6283zM13.837 9.85467l-1.0288-2.038c-.1009-.2-.3924-.2-.4937 0l-1.0288 2.038-2.2998.32663c-.22568.032-.31568.3034-.15273.459l1.66463 1.586-.393 2.24c-.0382.22.1977.3877.3992.284l2.0574-1.0576 2.0573 1.0573c.2015.1037.4374-.0637.3989-.2837l-.3927-2.24 1.6643-1.586c.1633-.1556.0733-.4266-.1524-.459l-2.2998-.32663z" fill="#B2B2B8"/></svg>
                                <svg v-if="output.id === 'image'"   width="24" height="24" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M19 5H5c-.6 0-1 .4-1 1v12c0 .6.4 1 1 1h14c.6 0 1-.4 1-1V6c0-.6-.4-1-1-1zm-9 3c.6 0 1 .4 1 1s-.4 1-1 1-1-.4-1-1 .4-1 1-1zm-4 8l2-4 3 2 3-4 4 6H6z" fill="#B2B2B8"/></svg>
                                <svg v-if="output.id === 'video'"   width="24" height="24" fill="none" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M9.879 16V8l6 4.0023-6 3.9977zM12 4c-4.41833 0-8 3.58167-8 8 0 4.4183 3.58167 8 8 8 4.4183 0 8-3.5817 8-8 0-4.41833-3.5817-8-8-8z" fill="#B2B2B8"/></svg>
                                <svg v-if="output.id === 'fastAds'" width="24" height="24" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M9.86006 19.9967l1.07644-6.039s-2.29717.0006-3.24289 0c-.67144-.0007-.87287-.2957-.52751-.8264 1.3593-2.0866 2.72646-4.16897 4.0908-6.25263.5535-.84534 1.1168-1.68534 1.6546-2.539.1418-.225.2543-.33967.5693-.33967h.6375c-.1114.63333-.2 1.272-.3096 1.864-.2413 1.30606-.487 2.6114-.7372 3.916-.0461.2393.0761.289.2947.2883.9453-.0033 1.8911-.003 2.8364-.0033.0957 0 .1918-.0053.2868.0027.4689.0386.6443.3466.4018.719-.7263 1.1144-1.4549 2.2275-2.1857 3.3393-1.1965 1.8233-2.4018 3.6417-3.5858 5.472-.1739.2687-.2928.385-.7043.402l-.55534-.0033z" fill="#B2B2B8"/></svg>
                            </div>
                            <div class="output__label">{{ output.label }}</div>
                        </div>
                    </div>
                </div>

                <div class="column-section">
                    <div class="common-dialog__column-title">Unit size</div>
                    <div class="size__container">
                        <div class="size__input">
                            <input-element label="Width" type="number" v-model="customWidth">
                                <div slot="right">px</div>
                            </input-element>
                        </div>
                        <div class="size__input">
                            <input-element label="Height" type="number" v-model="customHeight">
                                <div slot="right">px</div>
                            </input-element>
                        </div>
                    </div>
                </div>
            </div>

            <div class="column">
                <svg v-if="customSizeType === 'display'" width="319" height="244" fill="none" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M28.6567 181.493V12.7363H292.935v79.602h11.145V6.36816C304.08 2.72239 301.349 0 297.711 0H22.2886c-3.6379 0-6.3682 2.72239-6.3682 6.36816V187.861c0 3.645 2.7303 6.368 6.3682 6.368h73.2338v-12.736H28.6567zm289.7383-76.418c.285-5.3162-3.921-9.5526-9.552-9.5526h-55.709c-5.645 0-11.142 4.2364-11.142 9.5526V234.03c0 5.618 4.206 9.552 9.55 9.552h57.301c5.631 0 9.536-4.236 9.552-9.552V105.075zm-79.589-54.1297V237.214c0 3.633-2.722 6.368-6.368 6.368H105.075c-3.646 0-6.3685-2.735-6.3685-6.368V50.9453c0-3.6331 2.7225-6.3682 6.3685-6.3682h127.363c3.646 0 6.368 2.7351 6.368 6.3682zM3.18408 200.597C.912239 200.597 0 200.903 0 202.189v3.184c0 .43.302488.736 1.59204 1.592 6.61493.979 13.60716 1.592 23.88056 1.592h70.0498v-7.96H3.18408z" fill="#535353"/><path fill-rule="evenodd" clip-rule="evenodd" d="M292.935 92.3383H249.95c-4.321.7578-7.772 3.45-7.96 4.7761V47.7612c-.602-3.407-4.01-6.0959-6.368-6.3682H103.483c-6.3686 0-7.9606 3.1841-7.9606 6.3682V181.493H28.6567V12.7363H292.935v79.602zM106.667 73.2338V62.0896h122.587V224.478H106.667V73.2338zm206.965 36.6172h-66.866v120.995h66.866V109.851z" fill="#fff"/><path fill-rule="evenodd" clip-rule="evenodd" d="M35.0249 19.1045V175.124h60.4975V49.3533c.0159-2.9183 2.2034-7.2868 6.3686-7.9602h12.736V19.1045H35.0249zM136.915 79.602h62.09v54.129h-62.09V79.602zm109.851 46.169h66.866v12.736h-66.866v-12.736z" fill="#0DE3FF"/><path fill-rule="evenodd" clip-rule="evenodd" d="M241.99 58.9055h44.577v-39.801h-163.98V41.393h113.035c2.415 0 5.946 2.7702 6.368 6.3682v11.1443zm-135.323 3.1841h122.587v11.1442H106.667V62.0896zM157.612 163.98h63.681v3.184h-63.681v-3.184zm63.681-23.88h-63.681v3.184h63.681V140.1zm-106.666 0h38.209v31.84h-38.209V140.1zm100.298 7.96h-57.313v3.184h57.313v-3.184zm-100.298 65.273h108.258v3.184H114.627v-3.184zm103.482-57.313h-60.497v3.184h60.497v-3.184zm-103.482 49.353h105.074v3.184H114.627v-3.184zm108.258-28.657H114.627v23.881h108.258v-23.881zm23.881-66.865h66.866v11.144h-66.866v-11.144zm62.09 33.433h-57.314v30.248h57.314v-30.248zM241.99 63.6816h44.577v23.8806H241.99V63.6816zm60.497 113.0344h-30.248v1.592h30.248v-1.592zm-30.248 4.777h31.84v1.592h-31.84v-1.592zm-3.184 9.552v-14.329h-17.513v14.329h17.513zm3.184-4.776h22.288v1.592h-22.288v-1.592zm30.248 7.96h-30.248v1.592h30.248v-1.592zm-30.248 4.776h31.84v1.592h-31.84v-1.592zm-3.184 9.552v-14.328h-17.513v14.328h17.513zm3.184-4.776h22.288v1.592h-22.288v-1.592zm30.248 7.96h-30.248v1.592h30.248v-1.592zm-30.248 4.776h31.84v1.592h-31.84v-1.592zm-3.184 9.553v-14.329h-17.513v14.329h17.513zm3.184-4.776h22.288v1.592h-22.288v-1.592z" fill="#E7E7E7"/></svg>
                <svg v-if="customSizeType === 'image'"   width="320" height="244" fill="none" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M28.6567 181.493V12.7363H292.935v79.602h11.145V6.36816C304.08 2.72239 301.349 0 297.711 0H22.2886c-3.6379 0-6.3682 2.72239-6.3682 6.36816V187.861c0 3.645 2.7303 6.368 6.3682 6.368h73.2338v-12.736H28.6567zm289.7383-76.418c.285-5.3162-3.921-9.5526-9.552-9.5526h-55.709c-5.645 0-11.142 4.2364-11.142 9.5526V234.03c0 5.618 4.206 9.552 9.55 9.552h57.301c5.631 0 9.536-4.236 9.552-9.552V105.075zm-79.589-54.1297V237.214c0 3.633-2.722 6.368-6.368 6.368H105.075c-3.646 0-6.3685-2.735-6.3685-6.368V50.9453c0-3.6331 2.7225-6.3682 6.3685-6.3682h127.363c3.646 0 6.368 2.7351 6.368 6.3682zM3.18408 200.597C.912239 200.597 0 200.903 0 202.189v3.184c0 .43.302488.736 1.59204 1.592 6.61493.979 13.60716 1.592 23.88056 1.592h70.0498v-7.96H3.18408z" fill="#535353"/><path fill-rule="evenodd" clip-rule="evenodd" d="M292.935 92.3383H249.95c-4.321.7578-7.772 3.45-7.96 4.7761V47.7612c-.602-3.407-4.01-6.0959-6.368-6.3682H103.483c-6.3686 0-7.9606 3.1841-7.9606 6.3682V181.493H28.6567V12.7363H292.935v79.602zM106.667 73.2338V62.0896h122.587V224.478H106.667V73.2338zm206.965 36.6172h-66.866v120.995h66.866V109.851z" fill="#fff"/><path fill-rule="evenodd" clip-rule="evenodd" d="M35.0249 19.1045V175.124h60.4975V49.3533c.0159-2.9183 2.2034-7.2868 6.3686-7.9602h12.736V19.1045H35.0249zM136.915 79.602h62.09v54.129h-62.09V79.602zm109.851 46.169h66.866v12.736h-66.866v-12.736z" fill="#0DE3FF"/><path fill-rule="evenodd" clip-rule="evenodd" d="M241.99 58.9055h44.577v-39.801h-163.98V41.393h113.035c2.415 0 5.946 2.7702 6.368 6.3682v11.1443zm-135.323 3.1841h122.587v11.1442H106.667V62.0896zM157.612 163.98h63.681v3.184h-63.681v-3.184zm63.681-23.88h-63.681v3.184h63.681V140.1zm-106.666 0h38.209v31.84h-38.209V140.1zm100.298 7.96h-57.313v3.184h57.313v-3.184zm-100.298 65.273h108.258v3.184H114.627v-3.184zm103.482-57.313h-60.497v3.184h60.497v-3.184zm-103.482 49.353h105.074v3.184H114.627v-3.184zm108.258-28.657H114.627v23.881h108.258v-23.881zm23.881-66.865h66.866v11.144h-66.866v-11.144zm62.09 33.433h-57.314v30.248h57.314v-30.248zM241.99 63.6816h44.577v23.8806H241.99V63.6816zm60.497 113.0344h-30.248v1.592h30.248v-1.592zm-30.248 4.777h31.84v1.592h-31.84v-1.592zm-3.184 9.552v-14.329h-17.513v14.329h17.513zm3.184-4.776h22.288v1.592h-22.288v-1.592zm30.248 7.96h-30.248v1.592h30.248v-1.592zm-30.248 4.776h31.84v1.592h-31.84v-1.592zm-3.184 9.552v-14.328h-17.513v14.328h17.513zm3.184-4.776h22.288v1.592h-22.288v-1.592zm30.248 7.96h-30.248v1.592h30.248v-1.592zm-30.248 4.776h31.84v1.592h-31.84v-1.592zm-3.184 9.553v-14.329h-17.513v14.329h17.513zm3.184-4.776h22.288v1.592h-22.288v-1.592z" fill="#E7E7E7"/><path d="M82 106V92c0-1.1-.9-2-2-2H66c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2zm-12.5-5.5l2.5 3.01L75.5 99l4.5 6H66l3.5-4.5zM178 114v-14c0-1.1-.9-2-2-2h-14c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2zm-12.5-5.5l2.5 3.01 3.5-4.51 4.5 6h-14l3.5-4.5z" fill="#fff"/></svg>
                <svg v-if="customSizeType === 'video'"   width="319" height="244" fill="none" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M28.6567 181.493V12.7363H291.343v79.602h12.737V6.36816C304.08 2.72239 301.349 0 297.711 0H22.2886c-3.6379 0-6.3682 2.72239-6.3682 6.36816V187.861c0 3.645 2.7303 6.368 6.3682 6.368h73.2338v-12.736H28.6567zm289.7383-76.418c.285-5.3162-3.921-9.5526-9.552-9.5526h-55.709c-5.645 0-11.142 4.2364-11.142 9.5526V234.03c0 5.618 4.206 9.552 9.55 9.552h57.301c5.631 0 9.536-4.236 9.552-9.552V105.075zm-79.589-54.1297V237.214c0 3.633-2.722 6.368-6.368 6.368H105.075c-3.646 0-6.3685-2.735-6.3685-6.368V50.9453c0-3.6331 2.7225-6.3682 6.3685-6.3682h127.363c3.646 0 6.368 2.7351 6.368 6.3682zM3.18408 200.597C.912239 200.597 0 200.903 0 202.189v3.184c0 .43.302488.736 1.59204 1.592 6.61493.979 13.60716 1.592 23.88056 1.592h70.0498v-7.96H3.18408z" fill="#535353"/><path fill-rule="evenodd" clip-rule="evenodd" d="M292.935 92.3383H249.95c-4.321.7578-7.772 3.45-7.96 4.7761V47.7612c-.602-3.407-4.01-6.0959-6.368-6.3682H103.483c-6.3686 0-7.9606 3.1841-7.9606 6.3682V181.493H28.6567V12.7363H292.935v79.602zM106.667 73.2338V62.0896h122.587V224.478H106.667V73.2338zm206.965 36.6172h-66.866v120.995h66.866V109.851z" fill="#fff"/><path fill-rule="evenodd" clip-rule="evenodd" d="M221.294 41.3931H103.483c-6.3685 0-7.9606 3.184-7.9606 6.3681v93.9308H35.0249V35.0249H221.294v6.3682zm-114.627 39.801h122.587V148.06H106.667V81.1941zm206.965 44.5769h-66.866v38.209h66.866v-38.209zm-138.906-10.348l-7.761 5.174v-10.348l7.761 5.174zm-5.174 10.348c-5.722 0-10.348-4.626-10.348-10.348-.002-2.745 1.087-5.379 3.029-7.32 1.941-1.941 4.574-3.03 7.319-3.028 5.722 0 10.349 4.626 10.349 10.348s-4.627 10.348-10.349 10.348zm107.264 14.727l7.761 5.174-7.761 5.174v-10.348zm2.587 15.522c-5.722 0-10.348-4.627-10.348-10.348-.002-2.746 1.087-5.379 3.028-7.32 1.941-1.941 4.575-3.031 7.32-3.029 5.722 0 10.348 4.627 10.348 10.349 0 5.721-4.626 10.348-10.348 10.348z" fill="#0DE3FF"/><path fill-rule="evenodd" clip-rule="evenodd" d="M28.6567 28.6567V12.7363H292.935v15.9204H28.6567zM229.254 62.0896H106.667v11.1442h122.587V62.0896zm-7.96 109.8504h-63.682v3.184h63.682v-3.184zm-106.667-15.92h38.209v27.065h-38.209V156.02zm100.298 0h-57.313v3.184h57.313v-3.184zm-57.313 7.96h60.497v3.184h-60.497v-3.184zm60.497 15.921h-60.497v3.184h60.497v-3.184zm-60.497 25.472h63.682v3.184h-63.682v-3.184zm-4.776-15.92h-38.209v27.064h38.209v-27.064zm4.776 0h57.313v3.184h-57.313v-3.184zm60.497 7.96h-60.497v3.184h60.497v-3.184zm-60.497 15.92h60.497v3.184h-60.497v-3.184zM286.567 63.6816H241.99v23.8806h44.577V63.6816zM308.856 226.07h-57.314v-19.105h57.314v19.105zm4.776-116.219h-66.866v11.144h66.866v-11.144zm-19.105 58.905h7.961v1.592h-30.249v-1.592h22.288zm9.553 4.776h-31.841v1.592h31.841v-1.592zm-35.025-4.776v14.329h-17.513v-14.329h17.513zm25.472 9.552h-22.288v1.593h22.288v-1.593zm0 9.553h7.961v1.592h-30.249v-1.592h22.288zm9.553 4.776h-31.841v1.592h31.841v-1.592zm-35.025-4.776v14.328h-17.513v-14.328h17.513zm25.472 9.552h-22.288v1.592h22.288v-1.592zM79.602 151.244v-3.184h15.9204v3.184H79.602zm0 15.92v3.184h15.9204v-3.184H79.602zm0 9.552v-3.184h15.9204v3.184H79.602zm0-22.288v3.184h15.9204v-3.184H79.602zm0 9.552v-3.184h15.9204v3.184H79.602zm-44.5771-15.92v28.656h38.2089V148.06H35.0249zM241.99 58.9055h44.577V35.0249h-60.539v6.3681h9.594c2.415 0 5.946 2.7702 6.368 6.3682v11.1443z" fill="#E7E7E7"/></svg>
                <svg v-if="customSizeType === 'fastAds'" width="319" height="244" fill="none" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M28.6567 181.493V12.7363H291.343v79.602h12.737V6.36816C304.08 2.72239 301.349 0 297.711 0H22.2886c-3.6379 0-6.3682 2.72239-6.3682 6.36816V187.861c0 3.645 2.7303 6.368 6.3682 6.368h73.2338v-12.736H28.6567zm289.7383-76.418c.285-5.3162-3.921-9.5526-9.552-9.5526h-55.709c-5.645 0-11.142 4.2364-11.142 9.5526V234.03c0 5.618 4.206 9.552 9.55 9.552h57.301c5.631 0 9.536-4.236 9.552-9.552V105.075zm-79.589-54.1297V237.214c0 3.633-2.722 6.368-6.368 6.368H105.075c-3.646 0-6.3685-2.735-6.3685-6.368V50.9453c0-3.6331 2.7225-6.3682 6.3685-6.3682h127.363c3.646 0 6.368 2.7351 6.368 6.3682zM3.18408 200.597C.912239 200.597 0 200.903 0 202.189v3.184c0 .43.302488.736 1.59204 1.592 6.61493.979 13.60716 1.592 23.88056 1.592h70.0498v-7.96H3.18408z" fill="#535353"/><path fill-rule="evenodd" clip-rule="evenodd" d="M292.935 92.3383H249.95c-4.321.7578-7.772 3.45-7.96 4.7761V47.7612c-.602-3.407-4.01-6.0959-6.368-6.3682H103.483c-6.3686 0-7.9606 3.1841-7.9606 6.3682V181.493H28.6567V12.7363H292.935v79.602zM106.667 73.2338V62.0896h122.587V224.478H106.667V73.2338zm206.965 36.6172h-66.866v120.995h66.866V109.851z" fill="#fff"/><path fill-rule="evenodd" clip-rule="evenodd" d="M103.483 41.3931h117.811v-6.3682H35.0249V141.692h60.4975V47.7612c0-3.1841 1.5921-6.3681 7.9606-6.3681zm3.184 39.801h122.587V148.06H106.667V81.1941zm57.32 52.5289l2.399-14.421s-5.12.002-7.228 0c-1.497-.002-1.946-.706-1.176-1.973 2.243-3.69 4.496-7.373 6.748-11.056l.008-.013c.787-1.288 1.575-2.575 2.362-3.863.341-.559.685-1.117 1.028-1.675.896-1.457 1.793-2.9142 2.66-4.3884.316-.5374.567-.8112 1.269-.8112h1.421c-.113.6869-.216 1.3765-.317 2.057-.121.8179-.24 1.6227-.374 2.3944-.537 3.1192-1.085 6.2362-1.643 9.3512-.102.572.17.691.657.689 1.779-.007 3.559-.007 5.339-.008h.983c.044 0 .087-.001.131-.001l.052-.001c.153-.003.305-.005.456.008 1.045.093 1.436.828.896 1.717-1.619 2.662-3.243 5.32-4.872 7.975-.784 1.281-1.571 2.561-2.357 3.841l-.001.002-.001.001-.001.001c-1.885 3.07-3.77 6.139-5.632 9.222-.388.642-.653.92-1.57.96l-1.237-.008zm82.779-7.952h66.866v38.209h-66.866v-38.209zm31.322 22.221l-1.599 9.615.825.005c.611-.027.788-.212 1.046-.64 1.242-2.056 2.499-4.103 3.756-6.15.525-.854 1.049-1.707 1.572-2.562 1.086-1.77 2.169-3.542 3.248-5.316.36-.593.1-1.083-.597-1.145-.101-.009-.202-.007-.304-.005-.041 0-.081.001-.122.001h-.645c-1.19.001-2.38.001-3.57.005-.325.001-.506-.078-.438-.459.372-2.077.737-4.155 1.096-6.234.089-.515.168-1.051.249-1.596.067-.454.136-.914.211-1.372h-.947c-.468 0-.636.183-.846.541-.578.983-1.176 1.954-1.773 2.925-.229.373-.458.745-.686 1.117l-1.575 2.577c-1.503 2.458-3.006 4.916-4.503 7.378-.514.845-.214 1.314.783 1.315 1.406.001 4.819 0 4.819 0z" fill="#0DE3FF"/><path fill-rule="evenodd" clip-rule="evenodd" d="M28.6567 28.6567V12.7363H292.935v15.9204H28.6567zM229.254 62.0896H106.667v11.1442h122.587V62.0896zm-7.96 109.8504h-63.682v3.184h63.682v-3.184zm-106.667-15.92h38.209v27.065h-38.209V156.02zm100.298 0h-57.313v3.184h57.313v-3.184zm-57.313 7.96h60.497v3.184h-60.497v-3.184zm60.497 15.921h-60.497v3.184h60.497v-3.184zm-60.497 25.472h63.682v3.184h-63.682v-3.184zm-4.776-15.92h-38.209v27.064h38.209v-27.064zm4.776 0h57.313v3.184h-57.313v-3.184zm60.497 7.96h-60.497v3.184h60.497v-3.184zm-60.497 15.92h60.497v3.184h-60.497v-3.184zM286.567 63.6816H241.99v23.8806h44.577V63.6816zM308.856 226.07h-57.314v-19.105h57.314v19.105zm4.776-116.219h-66.866v11.144h66.866v-11.144zm-19.105 58.905h7.961v1.592h-30.249v-1.592h22.288zm9.553 4.776h-31.841v1.592h31.841v-1.592zm-35.025-4.776v14.329h-17.513v-14.329h17.513zm25.472 9.552h-22.288v1.593h22.288v-1.593zm0 9.553h7.961v1.592h-30.249v-1.592h22.288zm9.553 4.776h-31.841v1.592h31.841v-1.592zm-35.025-4.776v14.328h-17.513v-14.328h17.513zm25.472 9.552h-22.288v1.592h22.288v-1.592zM79.602 151.244v-3.184h15.9204v3.184H79.602zm0 15.92v3.184h15.9204v-3.184H79.602zm0 9.552v-3.184h15.9204v3.184H79.602zm0-22.288v3.184h15.9204v-3.184H79.602zm0 9.552v-3.184h15.9204v3.184H79.602zm-44.5771-15.92v28.656h38.2089V148.06H35.0249zM241.99 58.9055h44.577V35.0249h-60.539v6.3681h9.594c2.415 0 5.946 2.7702 6.368 6.3682v11.1443z" fill="#E7E7E7"/></svg>
            </div>
        </div>

        <div class="new-dialog__button-container">
            <dialog-button v-if="!isAddingCustomSize" :disabled="!selectedSizesIds.length" @click="nextStep">Next step</dialog-button>
            <dialog-button v-else :disabled="isNil(customWidth) || isNil(customHeight)" @click="saveCustomSize">Save custom size</dialog-button>
        </div>
    </div>
</template>

<script>
import { mapState, mapGetters, mapActions } from 'vuex'
import { DialogButton, Input } from 'design-system'
import { isNil } from 'lodash'
import InstanceTabs from '../InstanceTabs.vue'

export default {
    components: {
        DialogButton,
        InstanceTabs,
        inputElement: Input
    },
    data () {
        return {
            isAddingCustomSize: false,
            customSizeType: 'display',
            customWidth: null,
            customHeight: null
        }
    },
    computed: {
        ...mapState('massProduction', [
            'selectedSizesIds',
        ]),
        ...mapGetters('massProduction', {
            availableSizes: 'getAvailableSizes',
            selectedSizes: 'getSelectedSizes',
            outputTypes: 'getOutputTypes'
        }),
    },
    methods: {
        isNil: (v) => isNil(v),
        ...mapActions('massProduction', [
            'nextStep',
            'setSelectedSizesIds',
            'addCustomSize',
        ]),
        saveCustomSize () {
            const customSize = {
                id: `custom-${Math.random()}`,
                label: 'Custom',
                size: `${this.customWidth}x${this.customHeight}`,
                isCustom: true,
            }

            this.addCustomSize(customSize)
            this.toggleSelection(customSize.id)

            this.customWidth = null
            this.customHeight = null
            this.isAddingCustomSize = false
        },
        toggleSelection (sizeId) {
            if (this.selectedSizesIds.includes(sizeId)) {
                this.setSelectedSizesIds(this.selectedSizesIds.filter(size => size !== sizeId))
            } else {
                this.setSelectedSizesIds(this.selectedSizesIds.concat(sizeId))
            }
        },
        getSizeClass (size) {
            return {
                selected: this.selectedSizesIds.includes(size.id),
                custom: size.isCustom
            }
        },
        getOutputClass (output) {
            return {
                selected: output.id === this.customSizeType
            }
        }
    }
}
</script>

<style lang="less" scoped>
@import (reference) '../../less/common';
@import '../../less/dialogs_common';
@import '../../less/dialogs_columns_layout';

.sizes {
    &__content-wrapper {
        height: 100%;
        max-width: 1232px;
        display: flex;
        margin: 0 auto;
        flex-direction: column;
    }

    &__tab-container {
        height: 64px;
        padding: 0 24px;
    }

    &__container {
        flex: 1 1 0;
        display: flex;
        flex-wrap: wrap;
        margin: 0 auto;
        overflow: auto;
    }
}

.size {
    width: 160px;
    height: 208px;
    margin: 8px;
    padding: 16px;
    border: 1px solid transparent;
    box-sizing: border-box;
    cursor: pointer;
    display: flex;
    flex-direction: column;

    &__box {
        width: 128px;
        height: 128px;
        margin-bottom: 16px;
        background-color: #858594;
    }

    &__text {
        font-size: 12px;
        line-height: 16px;
        color: white;

        &--bold {
            font-family: @demi-font;
        }
    }

    &__add {
        width: 128px;
        height: 128px;
        margin-bottom: 16px;
        border: 2px solid #2E2E3E;
        display: flex;
        box-sizing: border-box;

        svg { margin: auto; }
    }

    &--selected {
        border: 1px solid @royal-blue;
        background-color: #1F1F30;

        .size__text { color: @royal-blue; }
    }

    &--custom .size__box { background: white; }
}

.custom-size {
    &__wrapper {
        width: 100%;
        display: flex;
        flex: auto;
        align-items: center;
        justify-content: center;
    }
}

.column {
    font-size: 11px;
    letter-spacing: 0.5px;
    width: 320px;
    margin-right: 40px;

    &:last-child {
        margin-right: 0;
        margin-left: 40px;
    }

    &-section {
        margin-bottom: 40px;

        &:last-child { margin-bottom: 0; }
    }
}

.output {
    height: 32px;
    display: flex;
    margin-bottom: 12px;
    cursor: pointer;

    // &__container {}

    &__icon {
        display: flex;
        margin-right: 8px;

        svg { margin: auto; }
    }

    &__label {
        font-family: SF Pro Text;
        font-size: 18px;
        line-height: 32px;
        color: #B2B2B8;
    }

    &:last-child { margin-bottom: 0; }

    &--selected {
        .output__icon svg path { fill: @royal-blue; }
        .output__label { color: @royal-blue; }
    }
}

.size {
    &__container {
        display: flex;
        justify-content: space-between;
    }
    
    &__input { width: 148px; }
}
</style>