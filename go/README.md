# API-Demo

## V3 Spot API
Description:
* **Public API**
    1. access to file: `run demo`，select a function you want to call in `spot/spotList.go`. eg:
       > `run demo/Market Data/Kline.go`
    2. input the params in json format if this request need any params,like：
       > `var params string = \`{"symbol":"BTCUSDT","limit":"200"}\``  
    3. if no params needed, just write 
       > `var params string =""`
    4. run the request with code 
       > `go run Kline.go`
* **Private API**
    1. input the `api_key` and `sec_key` in `config/config.go` first
    2. same as how to call Public api in next

## Futures API
* REST SDK is in `futures/mexcfutures/` (see `futures/README.md` and `futures/cmd/testrest`)
* Public market endpoints do not require signing
* Private endpoints use Futures header signing (`ApiKey`, `Request-Time`, `Signature`)
* Example:
    > `go run "run demo/Futures/Market/ContractPing.go"`

## Broker API
* REST endpoints are in `broker/brokerList.go`
* Example:
    > `go run "run demo/Broker/Endpoints/QuerySubAccountList.go"`

## P2P API
* REST endpoints are in `p2p/p2pList.go`
* Example:
    > `go run "run demo/P2P/Account/GetOrderDetail.go"`

## WebSocket For Spot
**Description:**
1. Spot **v3** market streams demo: `ws/v2/public/publicWs.go` — note: `v2` here is the **repository directory name only**, not the WebSocket protocol version; the example uses the v3 endpoint `wss://wbs-api.mexc.com/ws` and v3 `SUBSCRIPTION` messages per api-docs.
2. User data stream: `ws/v3/Ws.go` (requires listenKey from REST API; v3 user stream protocol)
3. run the request with code 
    > `go run publicWs.go`
