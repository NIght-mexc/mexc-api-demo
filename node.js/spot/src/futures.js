const APIBase = require('./APIbase')
const { Futures: FuturesModule } = require('./modules')
const { flowRight } = require('./helpers/utils')

class Futures extends flowRight(FuturesModule)(APIBase) {
  constructor (apiKey = '', apiSecret = '', options = {}) {
    options.baseURL = options.baseURL || 'https://api.mexc.com'
    super({
      apiKey,
      apiSecret,
      ...options
    })
  }
}

module.exports = Futures
