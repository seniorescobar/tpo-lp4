import axios from 'axios'

class ApiClient {
    constructor () {
    }

    get (path) {
        return this._makeRequest('GET', path)
    }

    post (path, data) {
        return this._makeRequest('POST', path, data)
    }

    put (path, data) {
        return this._makeRequest('PUT', path, data)
    }

    delete (path) {
        return this._makeRequest('DELETE', path)
    }

    _makeRequest (method, path, data) {
        return new Promise((resolve, reject) => {
            const headers = this._headers()
            const url = path
            axios({ method, headers, url, data })
                .then(response => resolve(response.data))
                .catch(error => reject(error))
        })
    }

    _headers () {
        const user = JSON.parse(localStorage.getItem('user'))
        const headers = { 'Content-Type': 'application/json' }

        if (user && user.token) {
            Object.assign(headers, { 'Authorization': 'Bearer ' + user.token })
        }

        return headers
    }
}

export default new ApiClient()
