const Futures = superclass => class extends superclass {

  // 30-Day Fee Statistics
  futuresAccountAssetBookOrderDealFeeTotal(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/account/asset_book/order_deal_fee/total', options)
  }

  // Change Risk Level
  futuresAccountChangeRiskLevelPost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/account/change_risk_level', options)
  }

  // Create STP Group
  futuresMarketMakerSelfTradeBlacklistCreatePost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/market_maker/self_trade/blacklist/create', options)
  }

  // Deduction Configuration
  futuresAccountFeeDeductConfigs(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/account/feeDeductConfigs', options)
  }

  // Delete STP Group
  futuresMarketMakerSelfTradeBlacklistDeletePost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/market_maker/self_trade/blacklist/delete', options)
  }

  // Enable or Disable Auto-Add Margin
  futuresPositionChangeAutoAddImPost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/position/change_auto_add_im', options)
  }

  // Get All Account Assets
  futuresAccountAssets(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/account/assets', options)
  }

  // Get Asset Transfer Records
  futuresAccountTransferRecord(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/account/transfer_record', options)
  }

  // Get Current User STP Group
  futuresMarketMakerSelfTradeBlacklistSearch(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/market_maker/self_trade/blacklist/search', options)
  }

  // Get Fee Details
  futuresAccountTieredFeeRateV2(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/account/tiered_fee_rate/v2', options)
  }

  // Get Funding Fee Details
  futuresPositionFundingRecords(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/position/funding_records', options)
  }

  // Get Historical Positions
  futuresPositionListHistoryPositions(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/position/list/history_positions', options)
  }

  // Get Open Positions
  futuresPositionOpenPositions(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/position/open_positions', options)
  }

  // Get Position Leverage Multipliers
  futuresPositionLeverage(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/position/leverage', options)
  }

  // Get Risk Limits
  futuresAccountRiskLimit(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/account/risk_limit', options)
  }

  // Get Single Currency Asset Information
  futuresAccountAsset(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/account/asset/{currency}', options)
  }

  // Get User Position Mode
  futuresPositionPositionMode(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/position/position_mode', options)
  }

  // Modify Leverage
  futuresPositionChangeLeveragePost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/position/change_leverage', options)
  }

  // Modify Position Margin
  futuresPositionChangeMarginPost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/position/change_margin', options)
  }

  // Modify User Position Mode
  futuresPositionChangePositionModePost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/position/change_position_mode', options)
  }

  // Query All Spot Discount Configuration Information
  futuresAccountConfigContractFeeDiscountConfig(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/account/config/contractFeeDiscountConfig', options)
  }

  // Query Contract Fee Deduction Details
  futuresOrderFeeDetails(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/order/fee_details', options)
  }

  // Query STP Groups and Group Members
  futuresMarketMakerSelfTradeBlacklist(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/market_maker/self_trade/blacklist', options)
  }

  // Query User Discount Usage
  futuresAccountDiscountType(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/account/discountType', options)
  }

  // Update STP Group
  futuresMarketMakerSelfTradeBlacklistUpdatePost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/market_maker/self_trade/blacklist/update', options)
  }

  // View Personal Profit Rate
  futuresAccountProfitRate(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/account/profit_rate/{type}', options)
  }

  // Get Candlestick Data
  futuresKline(options = {}) {
    return this.futuresPublicRequest('GET', '/api/v1/contract/kline/{symbol}', options)
  }

  // Get Contract Info
  futuresDetailCountry(options = {}) {
    return this.futuresPublicRequest('GET', '/api/v1/contract/detail/country', options)
  }

  // Get Contract Order Book Depth
  futuresDepth(options = {}) {
    return this.futuresPublicRequest('GET', '/api/v1/contract/depth/{symbol}', options)
  }

  // Get Fair Price Candles
  futuresKlineFairPrice(options = {}) {
    return this.futuresPublicRequest('GET', '/api/v1/contract/kline/fair_price/{symbol}', options)
  }

  // Get Fair Price
  futuresFairPrice(options = {}) {
    return this.futuresPublicRequest('GET', '/api/v1/contract/fair_price/{symbol}', options)
  }

  // Get Funding Rate History
  futuresFundingRateHistory(options = {}) {
    return this.futuresPublicRequest('GET', '/api/v1/contract/funding_rate/history', options)
  }

  // Get Funding Rate
  futuresFundingRate(options = {}) {
    return this.futuresPublicRequest('GET', '/api/v1/contract/funding_rate/{symbol}', options)
  }

  // Get Index Price Candles
  futuresKlineIndexPrice(options = {}) {
    return this.futuresPublicRequest('GET', '/api/v1/contract/kline/index_price/{symbol}', options)
  }

  // Get Index Price
  futuresIndexPrice(options = {}) {
    return this.futuresPublicRequest('GET', '/api/v1/contract/index_price/{symbol}', options)
  }

  // Get Insurance Fund Balance History
  futuresRiskReverseHistory(options = {}) {
    return this.futuresPublicRequest('GET', '/api/v1/contract/risk_reverse/history', options)
  }

  // Get Insurance Fund Balance
  futuresRiskReverse(options = {}) {
    return this.futuresPublicRequest('GET', '/api/v1/contract/risk_reverse/{symbol}', options)
  }

  // Get Recent Trades
  futuresDeals(options = {}) {
    return this.futuresPublicRequest('GET', '/api/v1/contract/deals/{symbol}', options)
  }

  // Get Server Time
  futuresPing(options = {}) {
    return this.futuresPublicRequest('GET', '/api/v1/contract/ping', options)
  }

  // Get the Last N Depth Snapshots
  futuresDepthCommits(options = {}) {
    return this.futuresPublicRequest('GET', '/api/v1/contract/depth_commits/{symbol}/{limit}', options)
  }

  // Get Ticker (Contract Market Data)
  futuresTicker(options = {}) {
    return this.futuresPublicRequest('GET', '/api/v1/contract/ticker', options)
  }

  // Get Transferable Currencies
  futuresSupportCurrencies(options = {}) {
    return this.futuresPublicRequest('GET', '/api/v1/contract/support_currencies', options)
  }

  // Batch Cancel by External Order ID
  futuresOrderBatchCancelWithExternalPost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/order/batch_cancel_with_external', options)
  }

  // Batch Place Order
  futuresOrderSubmitBatchPost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/order/submit_batch', options)
  }

  // Batch Query Orders by External Order ID
  futuresOrderBatchQueryWithExternalPost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/order/batch_query_with_external', options)
  }

  // Batch Query Orders by Order ID
  futuresOrderBatchQuery(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/order/batch_query', options)
  }

  // Cancel All Orders Under a Contract
  futuresOrderCancelAllPost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/order/cancel_all', options)
  }

  // Cancel All Planned Orders
  futuresPlanorderCancelAllPost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/planorder/cancel_all', options)
  }

  // Cancel All TP/SL Planned Orders
  futuresStoporderCancelAllPost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/stoporder/cancel_all', options)
  }

  // Cancel by External Order ID
  futuresOrderCancelWithExternalPost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/order/cancel_with_external', options)
  }

  // Cancel Orders
  futuresOrderCancelPost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/order/cancel', options)
  }

  // Cancel Planned Orders
  futuresPlanorderCancelPost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/planorder/cancel', options)
  }

  // Cancel TP/SL Planned Orders
  futuresStoporderCancelPost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/stoporder/cancel', options)
  }

  // Cancel Trailing Order
  futuresTrackorderCancelPost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/trackorder/cancel', options)
  }

  // Chase Order
  futuresOrderChaseLimitOrderPost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/order/chase_limit_order', options)
  }

  // Close All
  futuresPositionCloseAllPost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/position/close_all', options)
  }

  // Get All Historical Orders
  futuresOrderListHistoryOrders(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/order/list/history_orders', options)
  }

  // Get Current Orders
  futuresOrderListOpenOrders(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/order/list/open_orders', options)
  }

  // Get Current Take-Profit/Stop-Loss Order List
  futuresStoporderOpenOrders(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/stoporder/open_orders', options)
  }

  // Get Historical Order Deal Details
  futuresOrderListOrderDeals(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/order/list/order_deals/v3', options)
  }

  // Get Order by External ID
  futuresOrderExternal(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/order/external/{symbol}/{external_oid}', options)
  }

  // Get Order Information by Order ID
  futuresOrderGet(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/order/get/{orderId}', options)
  }

  // Get Plan Order List
  futuresPlanorderListOrders(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/planorder/list/orders', options)
  }

  // Get Take-Profit/Stop-Loss Order List
  futuresStoporderListOrders(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/stoporder/list/orders', options)
  }

  // Get Trade Records by Order ID
  futuresOrderDealDetails(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/order/deal_details/{orderId}', options)
  }

  // Modify Order Price & Quantity
  futuresOrderChangeLimitOrderPost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/order/change_limit_order', options)
  }

  // Modify Plan Order
  futuresPlanorderChangePricePost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/planorder/change_price', options)
  }

  // Modify Take-Profit/Stop-Loss on Plan Order
  futuresPlanorderChangeStopOrderPost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/planorder/change_stop_order', options)
  }

  // Modify TP/SL Prices on a Limit Order
  futuresStoporderChangePricePost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/stoporder/change_price', options)
  }

  // Modify TP/SL Prices on a TP/SL Planned Order
  futuresStoporderChangePlanPricePost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/stoporder/change_plan_price', options)
  }

  // Modify Trailing Order
  futuresTrackorderChangeOrderPost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/trackorder/change_order', options)
  }

  // Place Order
  futuresOrderCreatePost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/order/create', options)
  }

  // Place Plan Order
  futuresPlanorderPlaceV2Post(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/planorder/place/v2', options)
  }

  // Place TP/SL Order by Position
  futuresStoporderPlacePost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/stoporder/place', options)
  }

  // Place Trailing Order
  futuresTrackorderPlacePost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/trackorder/place', options)
  }

  // Query Historical Orders
  futuresOrderListCloseOrders(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/order/list/close_orders', options)
  }

  // Query In-Flight Order Counts
  futuresOrderOpenOrderTotalCountPost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/order/open_order_total_count', options)
  }

  // Query Trailing Orders
  futuresTrackorderListOrders(options = {}) {
    return this.futuresSignRequest('GET', '/api/v1/private/trackorder/list/orders', options)
  }

  // Reverse Open Position
  futuresPositionReversePost(options = {}) {
    return this.futuresSignRequest('POST', '/api/v1/private/position/reverse', options)
  }

}

module.exports = Futures