import service from '@/utils/request'

export const getOrders = (params) => {
  return service({
    url: "/order/getGateOrders",
    method: 'get',
    params
  })
}
