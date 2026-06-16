const P2P = require('../../src/p2p')
const apiKey = ''
const apiSecret = ''
const client = new P2P(apiKey, apiSecret, { baseURL: 'https://api.mexc.com' })


client.p2PMerchantAdsSaveOrUpdatePost({ payTimeLimit: 15, initQuantity: 10, price: 11, coinId: '5989b56ba96a43599dbeeca5bb053f43', side: 'SELL', fiatUnit: 'GBP', payMethod: '1605', minSingleTransAmount: 1, maxSingleTransAmount: 110, userAllTradeCountMin: 0, userAllTradeCountMax: 0 })
  .then(response => client.logger.log(response.data))
  .catch(error => client.logger.error(error))
