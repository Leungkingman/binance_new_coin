import axios from 'axios'
import Lockr from 'lockr'
import { Message, Loading } from 'element-ui'
const service = axios.create({
    baseURL: process.env.VUE_APP_HOST,
    timeout: 99999,
    headers: {
      "Content-Type": "application/json;charset=utf-8",
      "x-token": Lockr.get("x-token") || ""
    }
})

service.resetBaseURL = (host) => {
  service.defaults.baseURL = host
}

service.interceptors.request.use(
  config => {
    const token = Lockr.get("x-token") || ""
    config.headers = {
      'Content-Type': 'application/json',
      'x-token': token
    }
    return config
  },
  error => {
    Message({
      showClose: true,
      message: error,
      type: 'error'
    })
    return Promise.reject(error);
  }
);


//http response 拦截器
service.interceptors.response.use(
  response => {
    // console.log("response = ", response)
    if (response.data.data && response.data.data.reload) {
      router.push('/')
    }
    if (response.data.code == 0) {
      return response.data
    } else {
      Message({
        showClose: true,
        message: response.data.msg,
        type: 'error',
      })
      return Promise.reject(response.data.msg)
    }
  },
  error => {
    Message({
      showClose: true,
      message: error,
      type: 'error'
    })
    return Promise.reject(error)
  }
)

export default service