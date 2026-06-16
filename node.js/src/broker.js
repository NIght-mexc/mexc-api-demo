const APIBase = require('./APIbase')
const { Broker: BrokerModule } = require('./modules')
const { flowRight } = require('./helpers/utils')

class Broker extends flowRight(BrokerModule)(APIBase) {
  constructor (apiKey = '', apiSecret = '', options = {}) {
    options.baseURL = options.baseURL || 'https://api.mexc.com'
    super({
      apiKey,
      apiSecret,
      ...options
    })
  }
}

module.exports = Broker
