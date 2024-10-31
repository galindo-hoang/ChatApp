import axios from "axios";

const instance_v1 = axios.create({
    baseURL: 'http://localhost:8081/v1',
    timeout: 20000,
})

export default instance_v1