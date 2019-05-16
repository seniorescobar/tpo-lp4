export const dialogPaths = [
    { 
        condition: { isEditingSizes: true },
        steps: ['edit_sizes']
    },
    { 
        condition: { isEditingSizes: false },
        steps: ['select_sizes', 'select_experience', 'content', 'output', 'finish']
    },
]

export const goals = [
    { id: 'brand-awareness', label: 'Brand Awareness' },
    { id: 'product-awareness', label: 'Product Awareness' },
    { id: 'user-aquisition', label: 'User Aquisition' },
    { id: 'user-reengagement', label: 'User Re-engagement' }
]

export const campaigns = [
    { id: 'campaign1', label: 'My Campaign' },
    { id: 'campaign2', label: 'My Campaign #2' },
    { id: 'campaign3', label: 'My Campaign #3' },
    { id: 'campaign4', label: 'My Campaign #4' },
]

export const feeds = {
    bg: [
        {
            label: 'Feed name 1', 
            items: [
                { label: 'Model', id: 'bg-model.png' },
            ]
        },
        {
            label: 'Feed name 2', 
            items: [
                { label: 'Coconut', id: 'bg-coconut.png' },
            ]
        },
    ],
    logo: [
        {
            label: 'Feed logos', 
            items: [
                { label: 'Logo black full', id: 'hannah-logo-black-full.png' },
                { label: 'Logo black', id: 'hannah-logo-black.png' },
                { label: 'Logo white full', id: 'hannah-logo-white-full.png' },
            ]
        },
        {
            label: 'Feed products', 
            items: [
                { label: 'Shampoo', id: 'hannah-shampoo.png' },
                { label: 'Conditioner', id: 'hannah-conditioner.png' },
            ]
        },
    ],
    headline: [
        {
            label: 'Feed name 1', 
            items: [
                { label: 'Headline 1.1', id: 'h1.1' },
                { label: 'Headline 1.2', id: 'h1.2' },
                { label: 'Headline 1.3', id: 'h1.3' },
                { label: 'Headline 1.4', id: 'h1.4' },
            ]
        },
        {
            label: 'Feed name 2', 
            items: [
                { label: 'Headline 2.1', id: 'h2.1' },
                { label: 'Headline 2.2', id: 'h2.2' },
                { label: 'Headline 2.3', id: 'h2.3' },
                { label: 'Headline 2.4', id: 'h2.4' },
            ]
        },
    ],
    button: [
        {
            label: 'Feed name 1', 
            items: [
                { label: 'Action 1.1', id: 'a1.1' },
                { label: 'Action 1.2', id: 'a1.2' },
                { label: 'Action 1.3', id: 'a1.3' },
                { label: 'Action 1.4', id: 'a1.4' },
            ]
        },
        {
            label: 'Feed name 2', 
            items: [
                { label: 'Action 2.1', id: 'a2.1' },
                { label: 'Action 2.2', id: 'a2.2' },
                { label: 'Action 2.3', id: 'a2.3' },
                { label: 'Action 2.4', id: 'a2.4' },
            ]
        },
    ]
}

export const sizes = [
    { id: 'largeRectangle', label: 'Large Rectangle', size: '336x280' },
    { id: 'mediumRectangle', label: 'Medium Rectangle', size: '300x250' },
    { id: 'square', label: 'Square', size: '250x250' },
    { id: 'smallSquare', label: 'Small Square', size: '200x200' },
    { id: 'smallRectangle', label: 'Small Rectangle', size: '180x150' },
    { id: 'billboard', label: 'Billboard', size: '970x250' },
    { id: 'leaderboard', label: 'Leaderboard', size: '728x90' },
    { id: 'largeLeaderboard', label: 'Large Leaderboard', size: '970x90' },
    { id: 'mainBanner', label: 'Main Banner', size: '468x60' },
    { id: 'mobile', label: 'Mobile', size: '320x50' },
    { id: 'largeMobileBanner', label: 'Large mobile banner', size: '320x100' },
    { id: 'halfBanner', label: 'Half Banner', size: '234x60' },
    { id: 'halfPage', label: 'Half Page', size: '300x600' },
    { id: 'wideSkyscraper', label: 'Wide Skyscraper', size: '160x600' },
    { id: 'skyscraper', label: 'Skyscraper', size: '120x600' },
    { id: 'portrait', label: 'Portrait', size: '300x1050' },
    { id: 'tbd', label: 'TBD', size: '320x480' },
    { id: 'instagramImage', label: 'Instagram image', size: '1080x1080' },
    { id: 'facebookImage', label: 'Facebook image', size: '1200x900' },
    { id: 'facebookAdImage', label: 'Facebook Ad image', size: '1200x628' },
    { id: 'twitterImage', label: 'Twitter image', size: '1024x512' },
    { id: 'facebookCoverImage', label: 'Facebook Cover image', size: '851x315' },
    { id: 'instagramVideo', label: 'Instagram video', size: '1080x1080' },
    { id: 'facebookVideo', label: 'Facebook video', size: '1920x1080' },
    { id: 'twitterVideo', label: 'Twitter video', size: '1920x1080' },
    { id: 'snapchatTopSnapVideo', label: 'Snapchat top snap video', size: '1080x1920' },
]

export const outputTypes = [
    { id: 'display', label: 'Display' },
    { id: 'image', label: 'Image' },
    { id: 'video', label: 'Video' },
    { id: 'fastAds', label: 'Fast ads' }
]

export const experiences = [
    {
        id: 'summerSale',
        title: 'Summer Sale',
        description: 'Template description comes here',
        attributes: ['Background', 'Logo', 'Headline', 'CTA'],
        imageSource: 'summerSale.png',
        isFeatured: true,
        goals: ['user-reengagement', 'product-awareness'],
    },
    {
        id: 'brandPackshotWithModel',
        title: 'Brand packshot with model',
        description: 'Template description comes here',
        attributes: ['Background', 'Logo', 'Headline', 'CTA'],
        imageSource: 'brandPackshotWithModel.png',
        isFeatured: true,
        goals: ['user-aquisition'],
    },
    {
        id: 'brandProductIngredient',
        title: 'Brand product ingredient',
        description: 'Template description comes here',
        attributes: ['Background', 'Logo', 'Headline', 'CTA'],
        imageSource: 'brandProductIngredient.png',
        isFeatured: true,
        goals: ['brand-awareness', 'product-awareness'],
    },
    {
        id: 'lightboxOverBackgroundImage',
        title: 'Lightbox over background image',
        description: 'Template description comes here',
        attributes: ['Background', 'Logo', 'Headline', 'CTA'],
        imageSource: 'lightboxOverBackgroundImage.png',
        isFeatured: true,
        goals: ['user-reengagement', 'user-aquisition'],
    },
    {
        id: 'brandPackshotGlorified',
        title: 'Brand packshot glorified',
        description: 'Template description comes here',
        attributes: ['Background', 'Logo', 'Headline', 'CTA'],
        imageSource: 'brandPackshotGlorified.png',
        isCustom: true,
        goals: ['product-awareness'],
    },
    {
        id: 'lightboxOverBackgroundImage2',
        title: 'Lightbox over background image 2',
        description: 'Template description comes here',
        attributes: ['Background', 'Logo', 'Headline', 'CTA'],
        imageSource: 'lightboxOverBackgroundImage2.png',
        isCustom: true,
        goals: ['user-reengagement'],
    },
    {
        id: 'brandPackshotRight',
        title: 'Brand packshot right',
        description: 'Template description comes here',
        attributes: ['Background', 'Logo', 'Headline', 'CTA'],
        imageSource: 'brandPackshotRight.png',
        isCustom: true,
        goals: ['brand-awareness'],
    },
    {
        id: 'brandPackshotWithIngredient',
        title: 'Brand packshot with ingredient',
        description: 'Template description comes here',
        attributes: ['Background', 'Logo', 'Headline', 'CTA'],
        imageSource: 'brandPackshotWithIngredient.png',
        isCustom: true,
        goals: ['user-aquisition'],
    },
]

export const instances = [
    {
        sizeId: 'mediumRectangle',
        size: {
            width: 300,
            height: 250,
        },
        units: [
            {
                id: 'bg',
                label: 'Background image',
                position: { left: '0', top: '0', width: '100%', height: '100%' },
                type: 'asset',
                value: 'bg-model.png'
            },
            {
                id: 'logo',
                label: 'Logo',
                position: { left: '20px', top: '20px', width: 'fit-content', height: 'fit-content' },
                type: 'asset',
                value: 'hannah-logo-black-full.png',
                style: 'object-fit: contain'
            },
            {
                id: 'headline',
                label: 'Headline',
                position: { left: '20px', top: '150px', width: 'fit-content', height: 'fit-content' },
                type: 'text',
                value: 'It matters what\'s inside',
                style: 'font-family: \'Gotham Medium\'; max-width: 130px; text-transform: uppercase;'
            },
            {
                id: 'button',
                label: 'Call to action',
                position: { left: '20px', top: '200px', width: '135px', height: '35px' },
                type: 'button',
                value: 'Shop now'
            },
        ]
    },
    {
        sizeId: 'portrait',
        size: {
            width: 160,
            height: 600,
        },
        units: [
            {
                id: 'bg',
                label: 'Background image',
                position: { left: '0', top: '0', width: '100%', height: '100%' },
                type: 'asset',
                value: 'bg-model.png',
                style: 'height: 100%; object-position: right;',
            },
            {
                id: 'logo',
                label: 'Logo',
                position: { left: '20px', top: '20px', width: 'fit-content', height: 'fit-content' },
                type: 'asset',
                value: 'hannah-logo-white-full.png',
                style: 'object-fit: contain'
            },
            {
                id: 'headline',
                label: 'Headline',
                position: { left: '15px', top: '300px', width: 'fit-content', height: 'fit-content' },
                type: 'text',
                value: 'It matters what\'s inside',
                style: 'font-family: \'Gotham Medium\'; max-width: 130px; text-transform: uppercase; color: white;'
            },
            {
                id: 'button',
                label: 'Call to action',
                position: { left: '0', top: '530px', width: '100%', height: '70px' },
                type: 'button',
                value: 'Shop now'
            },
        ]
    },
    {
        sizeId: 'halfPage',
        size: {
            width: 300,
            height: 600,
        },
        units: [
            {
                id: 'bg',
                label: 'Background image',
                position: { left: '0', top: '0', width: '100%', height: '100%' },
                type: 'asset',
                value: 'bg-model.png',
                style: 'height: 100%; object-position: right;',
            },
            {
                id: 'logo',
                label: 'Logo',
                position: { left: '90px', top: '20px', width: 'fit-content', height: 'fit-content' },
                type: 'asset',
                value: 'hannah-logo-white-full.png',
                style: 'object-fit: contain'
            },
            {
                id: 'headline',
                label: 'Headline',
                position: { left: '130px', top: '280px', width: 'fit-content', height: 'fit-content' },
                type: 'text',
                value: 'It matters what\'s inside',
                style: 'font-family: \'Gotham Medium\'; max-width: 130px; text-transform: uppercase; color: white;'
            },
            {
                id: 'button',
                label: 'Call to action',
                position: { left: '0', top: '530px', width: '100%', height: '70px' },
                type: 'button',
                value: 'Shop now'
            },
        ]
    },
    {
        sizeId: 'facebookImage',
        size: {
            width: 1200,
            height: 900,
        },
        units: [
            {
                id: 'bg',
                label: 'Background image',
                position: { left: '0', top: '0', width: '100%', height: '100%' },
                type: 'asset',
                value: 'bg-model.png'
            },
            {
                id: 'logo',
                label: 'Logo',
                position: { left: '20px', top: '18px', width: '50px', height: 'fit-content' },
                type: 'asset',
                value: 'hannah-logo-black-full.png',
                style: 'object-fit: contain'
            },
            {
                id: 'button',
                label: 'Call to action',
                position: { left: '75px', top: '13px', width: '80px', height: '30px' },
                type: 'button',
                value: 'Shop now'
            },
        ]
    },
    {
        sizeId: 'mobile',
        size: {
            width: 320,
            height: 50,
        },
        units: [
            {
                id: 'bg',
                label: 'Background image',
                position: { left: '0', top: '0', width: '100%', height: '100%' },
                type: 'asset',
                value: 'bg-model.png'
            },
            {
                id: 'logo',
                label: 'Logo',
                position: { left: '20px', top: '18px', width: '50px', height: 'fit-content' },
                type: 'asset',
                value: 'hannah-logo-black-full.png',
                style: 'object-fit: contain'
            },
            {
                id: 'button',
                label: 'Call to action',
                position: { left: '75px', top: '13px', width: '80px', height: '30px' },
                type: 'button',
                value: 'Shop now'
            },
        ]
    },
    {
        sizeId: 'billboard',
        size: {
            width: 970,
            height: 250,
        },
        units: [
            {
                id: 'bg',
                label: 'Background image',
                position: { left: '0', top: '0', width: '100%', height: '100%' },
                type: 'asset',
                value: 'bg-model.png'
            },
            {
                id: 'logo',
                label: 'Logo',
                position: { left: '20px', top: '20px', width: 'fit-content', height: 'fit-content' },
                type: 'asset',
                value: 'hannah-logo-black-full.png',
                style: 'object-fit: contain'
            },
            {
                id: 'headline',
                label: 'Headline',
                position: { left: '20px', top: '150px', width: 'fit-content', height: 'fit-content' },
                type: 'text',
                value: 'It matters what\'s inside',
                style: 'font-family: \'Gotham Medium\'; max-width: 130px; text-transform: uppercase;'
            },
            {
                id: 'button',
                label: 'Call to action',
                position: { left: '20px', top: '200px', width: '135px', height: '35px' },
                type: 'button',
                value: 'Shop now'
            },
        ]
    },
    {
        sizeId: 'leaderboard',
        size: {
            width: 728,
            height: 90,
        },
        units: [
            {
                id: 'bg',
                label: 'Background image',
                position: { left: '0', top: '0', width: '100%', height: '100%' },
                type: 'asset',
                value: 'bg-model.png'
            },
            {
                id: 'logo',
                label: 'Logo',
                position: { left: '20px', top: '30px', width: 'fit-content', height: 'fit-content' },
                type: 'asset',
                value: 'hannah-logo-black-full.png',
                style: 'object-fit: contain'
            },
            {
                id: 'button',
                label: 'Call to action',
                position: { left: '180px', top: '30px', width: '135px', height: '35px' },
                type: 'button',
                value: 'Shop now'
            },
        ]
    },
    {
        sizeId: 'square',
        size: {
            width: 300,
            height: 300,
        },
        units: [
            {
                id: 'bg',
                label: 'Background image',
                position: { left: '0', top: '0', width: '100%', height: '100%' },
                type: 'asset',
                value: 'bg-model.png'
            },
            {
                id: 'logo',
                label: 'Logo',
                position: { left: '20px', top: '20px', width: 'fit-content', height: 'fit-content' },
                type: 'asset',
                value: 'hannah-logo-black-full.png',
                style: 'object-fit: contain'
            },
            {
                id: 'headline',
                label: 'Headline',
                position: { left: '20px', top: '150px', width: 'fit-content', height: 'fit-content' },
                type: 'text',
                value: 'It matters what\'s inside',
                style: 'font-family: \'Gotham Medium\'; max-width: 130px; text-transform: uppercase;'
            },
            {
                id: 'button',
                label: 'Call to action',
                position: { left: '20px', top: '200px', width: '135px', height: '35px' },
                type: 'button',
                value: 'Shop now'
            },
        ]
    }
]
