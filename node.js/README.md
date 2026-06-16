# Mexc in Nodejs

## Installation

```
npm install 
```

## RESTful APIs

### Spot
```javascript
const { Spot } = require('./src')
const apiKey = ''
const apiSecret = ''
const client = new Spot(apiKey, apiSecret, { baseURL: 'https://api.mexc.com' })

client.Ping().then(response => client.logger.log(response.data))
  .catch(error => client.logger.error(error))
```

### Futures
```javascript
const { Futures } = require('./src')
const apiKey = ''      // replace with your API key
const apiSecret = ''   // replace with your API secret
const client = new Futures(apiKey, apiSecret, { baseURL: 'https://api.mexc.com' })

client.futuresPing().then(response => client.logger.log(response.data))
```

### Broker
```javascript
const { Broker } = require('./src')
const apiKey = ''      // replace with your API key
const apiSecret = ''   // replace with your API secret
const client = new Broker(apiKey, apiSecret, { baseURL: 'https://api.mexc.com' })

client.brokerSubAccountList().then(response => client.logger.log(response.data))
```

### P2P
```javascript
const { P2P } = require('./src')
const apiKey = ''      // replace with your API key
const apiSecret = ''   // replace with your API secret
const client = new P2P(apiKey, apiSecret, { baseURL: 'https://api.mexc.com' })

client.p2POrderDetail({ orderId: '123' }).then(response => client.logger.log(response.data))
```

Please find `src/modules` folder to check for more endpoints.

### Base URL
* Spot / Broker / P2P: `https://api.mexc.com`
* Futures REST: `https://api.mexc.com`

### Optional Parameters

Optional parameters are encapsulated to a single object as the last function parameter.

```javascript
const { Spot } = require('./src')
const client = new Spot()
client.Depth({ symbol: 'BTCUSDT', limit: 5 }).then(response => client.logger.log(response.data))
```

## Websocket
## Environmental requirements
nodejs 12.22.3+

## Demo Description
`websocket/websocket_proto.js`
Demo using spot websocket v3 protobuf streams

## User's Guide
```
cd websocket
npm install
node websocket_proto.js
```

## Run REST demos
```
node run/Market/Ping.js
node run/Futures/Ping.js
node run/Broker/QuerySubAccountList.js
node run/P2P/GetOrderDetail.js
```
