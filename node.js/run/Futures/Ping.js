const Futures = require('../../src/futures')
const client = new Futures()

client.futuresPing().then(response => client.logger.log(response.data))
  .catch(error => client.logger.error(error))
