const P2P = require('../../src/p2p')
const apiKey = ''
const apiSecret = ''
const client = new P2P(apiKey, apiSecret, { baseURL: 'https://api.mexc.com' })

// API requires multipart/form-data; signRequest uses query signing only.
// Use a multipart HTTP client or extend APIbase for file upload.
client.p2PUploadFilePost({})
  .then(response => client.logger.log(response.data))
  .catch(error => client.logger.error(error))
