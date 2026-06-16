const P2P = require('../../src/p2p')
const apiKey = ''
const apiSecret = ''
const client = new P2P(apiKey, apiSecret, { baseURL: 'https://api.mexc.com' })


client.p2PMerchantOrderDealPost({ advNo: 'a1369623409544374272', amount: 100 })
  .then(response => client.logger.log(response.data))
  .catch(error => client.logger.error(error))
