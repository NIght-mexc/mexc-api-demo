const Broker = require('../../src/broker')
const client = new Broker()

client.brokerSubAccountList().then(response => client.logger.log(response.data))
  .catch(error => client.logger.error(error))
