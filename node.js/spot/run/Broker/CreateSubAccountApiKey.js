const Broker = require('../../src/broker')
const apiKey = ''
const apiSecret = ''
const client = new Broker(apiKey, apiSecret, { baseURL: 'https://api.mexc.com' })


client.brokerSubAccountApiKeyPost({ subAccount: 'mysub001', permissions: 'SPOT_ACCOUNT_READ,SPOT_ACCOUNT_WRITE', note: 'demo key' })
  .then(response => client.logger.log(response.data))
  .catch(error => client.logger.error(error))
