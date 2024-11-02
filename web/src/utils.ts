class LocalStorageWrapper {
    getAccount() {
        return localStorage.getItem('account')
    }

    setAccount(data: string) {
        localStorage.setItem('account', data)
    }

    setToken(data: string) {
        localStorage.setItem('session_token', data)
    }

    getToken() {
        return localStorage.getItem('session_token')
    }
}

// eslint-disable-next-line import/no-anonymous-default-export
export default new LocalStorageWrapper()

