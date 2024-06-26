import service from '@/utils/request'

export const getCoin = (params) => {
  return service({
    url: "/coin/getCoin",
    method: 'get',
    params
  })
}

export const updateCoin = (data) => {
  return service({
    url: "/coin/updateCoin",
    method: 'post',
    data
  })
}
