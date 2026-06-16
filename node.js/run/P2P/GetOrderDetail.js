const P2P = require('../../src/p2p')
const client = new P2P()

client.p2POrderDetail().then(response => client.logger.log(response.data))
  .catch(error => client.logger.error(error))
