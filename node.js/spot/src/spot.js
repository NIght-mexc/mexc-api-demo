const APIBase = require('./APIbase')
const { Spot: SpotModule } = require('./modules')
const { flowRight } = require('./helpers/utils')

class Spot extends flowRight(SpotModule)(APIBase) {
  constructor (apiKey = '', apiSecret = '', options = {}) {
    options.baseURL = options.baseURL || 'https://api.mexc.com'
    super({
      apiKey,
      apiSecret,
      ...options
    })
  }
}

module.exports = Spot
