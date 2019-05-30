import axios from 'axios'

class ApiClient {
    constructor () {}

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
            const url = `/api/${path}/`

            axios({ method, headers, url, data })
                .then(response => resolve(response.data))
                .catch(error => reject(error))
        })
    }

    _headers () {
        const headers = { 'Content-Type': 'application/json' }
        const user = localStorage.getItem('user')
        const token = user && JSON.parse(user).token

        if (token) {
            Object.assign(headers, { 'Authorization': `Bearer ${token}` })
        }

        return headers
    }
}

export default new ApiClient()