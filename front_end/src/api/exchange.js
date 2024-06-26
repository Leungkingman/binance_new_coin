import service from '@/utils/request'

export const initSystem = (params) => {
  return service({
    url: "/exchange/initialExchangeData",
    method: 'get',
    params
  })
}

export const getInitialInfo = (params) => {
  return service({
    url: "/exchange/getInitialInfo",
    method: 'get',
    params
  })
}

