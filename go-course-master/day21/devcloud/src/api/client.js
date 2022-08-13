import axios from 'axios'

const client = axios.create({
    baseURL: 'http://localhost:8050',
    timeout: 3000,
})

client.interceptors.request.use(
    request => {
        // 每一个请求 都添加 token的头
        return request
    },
    err => {
        console.log(err)
        return Promise.reject(err)
    }
)

client.interceptors.response.use(
    response =>  {
        const resp = response.data
        if (resp.code === 0) {
            return resp
        }
        console.log(resp)
        return resp
    },
    err => {
        console.log(err)
        return Promise.reject(err)
    }
)

export default client