// Celtra API Mock

import { } from './constants'

const state = {
    creatives: []
}

class CeltraApi {
    constructor () {
        this.state = state
    }

    get (req) {
        return {
            creatives: this.state.creatives,
        }[req]
    }

    post (req, payload) {
        return {
            'creative': () => this.state.creatives.push(payload),
        }[req]()
    }
}

export default new CeltraApi ()