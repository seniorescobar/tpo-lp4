const path = require('path')

module.exports = {
    chainWebpack: config => {
        config.resolve.alias
            .set('design-system', path.resolve(__dirname, 'src/components/design_system'))
            .set('api-client', path.resolve(__dirname, 'src/helpers/apiClient.js'))
    },
    devServer: {
        port: 8000,
        proxy: 'http://localhost:8080',
    },
}
