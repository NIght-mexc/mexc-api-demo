const APIBase = require('./APIbase')
const { P2P: P2PModule } = require('./modules')
const { flowRight } = require('./helpers/utils')

class P2P extends flowRight(P2PModule)(APIBase) {
  constructor (apiKey = '', apiSecret = '', options = {}) {
    options.baseURL = options.baseURL || 'https://api.mexc.com'
    super({
      apiKey,
      apiSecret,
      ...options
    })
  }
}

module.exports = P2P
