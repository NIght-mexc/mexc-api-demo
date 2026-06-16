const Broker = require('../../src/broker')
const apiKey = ''
const apiSecret = ''
const client = new Broker(apiKey, apiSecret, { baseURL: 'https://api.mexc.com' })


client.brokerSubAccountStatus({ subAccount: 'mysub001' })
  .then(response => client.logger.log(response.data))
  .catch(error => client.logger.error(error))
