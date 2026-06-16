const CryptoJS = require('crypto-js')
const { removeEmptyValue, buildQueryString, buildSortedQueryString, substitutePathParams, createRequest, createFuturesRequest, defaultLogger } = require('./helpers/utils')

class APIBase {
  constructor(options) {
    const { apiKey, apiSecret, baseURL, logger } = options
    this.apiKey = apiKey
    this.apiSecret = apiSecret
    this.baseURL = baseURL
    this.logger = logger || defaultLogger
  }
  
  publicRequest(method, path, params = {}) {
    params = removeEmptyValue(params)
    params = buildQueryString(params)
    if (params !== '') {
      path = `${path}?${params}`
    }
    return createRequest({
      method: method,
      baseURL: this.baseURL,
      url: path,
      apiKey: this.apiKey
    })
  }

  signRequest(method, path, params = {}) {
    params = removeEmptyValue(params)
    const timestamp = Date.now()
    let queryString = buildQueryString({ ...params, timestamp })
    queryString = queryString.replace(/\(/g, '%28').replace(/\)/g, '%29');
    const signature = CryptoJS.enc.Hex.stringify(CryptoJS.HmacSHA256(queryString, this.apiSecret))
    return createRequest({
      method: method,
      baseURL: this.baseURL,
      url: `${path}?${queryString}&signature=${signature}`,
      apiKey: this.apiKey
    })
  }

  futuresPublicRequest(method, path, params = {}) {
    const resolved = substitutePathParams(path, removeEmptyValue(params))
    let url = resolved.path
    const queryString = buildSortedQueryString(resolved.params)
    if (queryString !== '') {
      url = `${url}?${queryString}`
    }
    return createFuturesRequest({
      method,
      baseURL: this.baseURL,
      url
    })
  }

  futuresSignRequest(method, path, params = {}) {
    const timestamp = Date.now().toString()
    if (method === 'POST') {
      const body = Object.keys(params).length ? JSON.stringify(params) : ''
      const signTarget = `${this.apiKey}${timestamp}${body}`
      const signature = CryptoJS.enc.Hex.stringify(CryptoJS.HmacSHA256(signTarget, this.apiSecret))
      return createFuturesRequest({
        method,
        baseURL: this.baseURL,
        url: path,
        data: body || undefined,
        headers: {
          ApiKey: this.apiKey,
          'Request-Time': timestamp,
          Signature: signature
        }
      })
    }

    const resolved = substitutePathParams(path, removeEmptyValue(params))
    let url = resolved.path
    const queryString = buildSortedQueryString(resolved.params)
    const signTarget = `${this.apiKey}${timestamp}${queryString}`
    const signature = CryptoJS.enc.Hex.stringify(CryptoJS.HmacSHA256(signTarget, this.apiSecret))
    if (queryString !== '') {
      url = `${url}?${queryString}`
    }
    return createFuturesRequest({
      method,
      baseURL: this.baseURL,
      url,
      headers: {
        ApiKey: this.apiKey,
        'Request-Time': timestamp,
        Signature: signature
      }
    })
  }
}

module.exports = APIBase
