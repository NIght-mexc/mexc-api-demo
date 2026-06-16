const Broker = require('../../src/broker')
const apiKey = ''
const apiSecret = ''
const client = new Broker(apiKey, apiSecret, { baseURL: 'https://api.mexc.com' })


client.brokerCapitalWithdrawApplyPost({ coin: 'USDT', network: 'TRC-20', address: 'TDemoAddress123456789', amount: '1.00' })
  .then(response => client.logger.log(response.data))
  .catch(error => client.logger.error(error))
