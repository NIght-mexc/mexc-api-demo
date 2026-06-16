const P2P = require('../../src/p2p')
const apiKey = ''
const apiSecret = ''
const client = new P2P(apiKey, apiSecret, { baseURL: 'https://api.mexc.com' })


client.p2PMerchantOrderPaginationV2({ startTime: 1737879140000, endTime: 1737965540000 })
  .then(response => client.logger.log(response.data))
  .catch(error => client.logger.error(error))
