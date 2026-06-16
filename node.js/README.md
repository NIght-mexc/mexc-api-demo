# Mexc in Nodejs

## Installation

```
npm install 
```

## SDK layout

| Package | Path | Description |
|---------|------|-------------|
| Spot SDK | `node.js/spot/` | Spot REST + Broker + P2P + thin Futures helpers |
| Futures SDK | `node.js/futures/` | Standalone Futures REST + WebSocket SDK (upstream) |

## RESTful APIs

### Spot (`node.js/spot/`)

```javascript
const Spot = require('./spot/src/spot')
const apiKey = ''
const apiSecret = ''
const client = new Spot(apiKey, apiSecret, { baseURL: 'https://api.mexc.com' })

client.Ping().then(response => client.logger.log(response.data))
  .catch(error => client.logger.error(error))
```

```javascript
const Spot = require('./spot/src/spot')
const apiKey = ''
const apiSecret = ''
const client = new Spot(apiKey, apiSecret, { baseURL: 'https://api.mexc.com' })

client.AccountInformation().then(response => client.logger.log(response.data))
  .catch(error => client.logger.error(error))
```

### Futures (`node.js/futures/`)

Use the standalone Futures SDK for contract REST and WebSocket examples:

```javascript
const { MexcFuturesRestClient } = require('./futures/src')
const apiKey = ''      // replace with your API key
const apiSecret = ''   // replace with your API secret
const client = new MexcFuturesRestClient({ apiKey, apiSecret })

client.ping().then(response => console.log(response))
  .catch(error => console.error(error))
```

`spot/src` also exposes a thin Futures mixin (`require('./spot/src').Futures`) for legacy demos; prefer `node.js/futures/` for full Futures coverage.

### Broker (`node.js/spot/`)

```javascript
const { Broker } = require('./spot/src')
const apiKey = ''      // replace with your API key
const apiSecret = ''   // replace with your API secret
const client = new Broker(apiKey, apiSecret, { baseURL: 'https://api.mexc.com' })

client.brokerSubAccountList().then(response => client.logger.log(response.data))
```

### P2P (`node.js/spot/`)

```javascript
const { P2P } = require('./spot/src')
const apiKey = ''      // replace with your API key
const apiSecret = ''   // replace with your API secret
const client = new P2P(apiKey, apiSecret, { baseURL: 'https://api.mexc.com' })

client.p2POrderDetail({ advOrderNo: '123' }).then(response => client.logger.log(response.data))
```

Please find `spot/src/modules` folder to check for more Spot / Broker / P2P endpoints.

### Base URL

* Spot / Broker / P2P: `https://api.mexc.com`
* Futures REST: `https://api.mexc.com`

### Optional Parameters

Optional parameters are encapsulated to a single object as the last function parameter.

```javascript
const Spot = require('./spot/src/spot')
const client = new Spot()
client.Depth({ symbol: 'BTCUSDT', limit: 5 }).then(response => client.logger.log(response.data))
```

## Websocket

### Spot WebSocket (`node.js/spot/websocket/`)

Environmental requirements: nodejs 12.22.3+

Demo: `spot/websocket/websocket_proto.js` — Spot WebSocket v3 protobuf streams.

```
cd spot/websocket
npm install
node websocket_proto.js
```

### Futures WebSocket (`node.js/futures/`)

See `futures/examples/ws-ticker.js`, `ws-depth.js`, `ws-private.js`, or run `node futures/run-ws.js`.

## Run REST demos

### Spot / Broker / P2P (`node.js/spot/run/`)

```
node spot/run/Market/Ping.js
node spot/run/Broker/QuerySubAccountList.js
node spot/run/P2P/GetOrderDetail.js
```

### Futures (`node.js/futures/`)

```
node futures/run.js
node futures/examples/rest-public.js
node futures/examples/rest-private.js
node futures/examples/ws-ticker.js
node futures/examples/ws-depth.js
node futures/examples/ws-private.js
```
