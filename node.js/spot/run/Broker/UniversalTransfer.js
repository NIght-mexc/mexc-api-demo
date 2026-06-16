const Broker = require('../../src/broker')
const apiKey = ''
const apiSecret = ''
const client = new Broker(apiKey, apiSecret, { baseURL: 'https://api.mexc.com' })


client.brokerSubAccountUniversalTransferPost({ fromAccountType: 'SPOT', toAccountType: 'SPOT', asset: 'USDT', amount: '10.00' })
  .then(response => client.logger.log(response.data))
  .catch(error => client.logger.error(error))
