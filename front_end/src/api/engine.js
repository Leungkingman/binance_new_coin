import service from '@/utils/request'

export const startEngine = (params) => {
  return service({
    url: "/engine/startEngine",
    method: 'get',
    params
  })
}

export const stopEngine = (params) => {
  return service({
    url: "/engine/stopEngine",
    method: 'get',
    params
  })
}

export const getEngineRunningStatus = (params) => {
  return service({
    url: "/engine/getEngineRunningStatus",
    method: 'get',
    params
  })
}

export const getEngineLog = (params) => {
  return service({
    url: "/engine/getEngineLog",
    method: 'get',
    params
  })
}

export const clearEngineLog = (params) => {
  return service({
    url: "/engine/clearEngineLog",
    method: 'get',
    params
  })
}

export const testBinanceApi = (params) => {
  return service({
    url: "/test/testGetApi",
    method: 'get',
    params
  })
}
