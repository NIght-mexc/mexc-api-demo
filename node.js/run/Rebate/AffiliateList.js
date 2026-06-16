const Spot = require('../../src/spot')
const client = new Spot('')

client.AffiliateList().then(response => client.logger.log(response.data))
  .catch(error => client.logger.error(error))
