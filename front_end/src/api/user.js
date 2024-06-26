import service from '@/utils/request'

export const login = (data) => {
  return service({
    url: "/user/login",
    method: 'post',
    data
  })
}

export const getConfig = (params) => {
  return service({
    url: "/user/getConfig",
    method: 'get',
    params
  })
}

export const updateConfig = (data) => {
  return service({
    url: "/user/updateConfig",
    method: 'post',
    data
  })
}

export const updateSecret = (data) => {
  return service({
    url: "/user/updateSecret",
    method: 'post',
    data
  })
}

export const updatePassword = (data) => {
  return service({
    url: "/user/updatePassword",
    method: 'post',
    data
  })
}
