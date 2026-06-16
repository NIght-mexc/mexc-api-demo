const Broker = require('../../src/broker')
const apiKey = ''
const apiSecret = ''
const client = new Broker(apiKey, apiSecret, { baseURL: 'https://api.mexc.com' })


client.brokerCapitalDepositSubHisrecGetall({})
  .then(response => client.logger.log(response.data))
  .catch(error => client.logger.error(error))
