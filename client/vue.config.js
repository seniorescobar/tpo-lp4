const path = require('path')

module.exports = {
    chainWebpack: config => {
        config.resolve.alias
            .set('design-system', path.resolve(__dirname, 'src/components/design_system'))
            .set('client-api', path.resolve(__dirname, 'api/ClientApi.js'))
    }
}
