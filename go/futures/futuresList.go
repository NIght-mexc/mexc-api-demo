package futuresList

import (
	"demo/config"
	"demo/utils"
	"fmt"
)

// 30-Day Fee Statistics
func FuturesAccountAssetBookOrderDealFeeTotal(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/account/asset_book/order_deal_fee/total")
	response := utils.FuturesPrivateGet("/api/v1/private/account/asset_book/order_deal_fee/total", jsonParams)
	return response
}

// Change Risk Level
func FuturesAccountChangeRiskLevelPost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/account/change_risk_level")
	response := utils.FuturesPrivatePost("/api/v1/private/account/change_risk_level", jsonParams)
	return response
}

// Create STP Group
func FuturesMarketMakerSelfTradeBlacklistCreatePost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/market_maker/self_trade/blacklist/create")
	response := utils.FuturesPrivatePost("/api/v1/private/market_maker/self_trade/blacklist/create", jsonParams)
	return response
}

// Deduction Configuration
func FuturesAccountFeeDeductConfigs(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/account/feeDeductConfigs")
	response := utils.FuturesPrivateGet("/api/v1/private/account/feeDeductConfigs", jsonParams)
	return response
}

// Delete STP Group
func FuturesMarketMakerSelfTradeBlacklistDeletePost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/market_maker/self_trade/blacklist/delete")
	response := utils.FuturesPrivatePost("/api/v1/private/market_maker/self_trade/blacklist/delete", jsonParams)
	return response
}

// Enable or Disable Auto-Add Margin
func FuturesPositionChangeAutoAddImPost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/position/change_auto_add_im")
	response := utils.FuturesPrivatePost("/api/v1/private/position/change_auto_add_im", jsonParams)
	return response
}

// Get All Account Assets
func FuturesAccountAssets(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/account/assets")
	response := utils.FuturesPrivateGet("/api/v1/private/account/assets", jsonParams)
	return response
}

// Get Asset Transfer Records
func FuturesAccountTransferRecord(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/account/transfer_record")
	response := utils.FuturesPrivateGet("/api/v1/private/account/transfer_record", jsonParams)
	return response
}

// Get Current User STP Group
func FuturesMarketMakerSelfTradeBlacklistSearch(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/market_maker/self_trade/blacklist/search")
	response := utils.FuturesPrivateGet("/api/v1/private/market_maker/self_trade/blacklist/search", jsonParams)
	return response
}

// Get Fee Details
func FuturesAccountTieredFeeRateV2(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/account/tiered_fee_rate/v2")
	response := utils.FuturesPrivateGet("/api/v1/private/account/tiered_fee_rate/v2", jsonParams)
	return response
}

// Get Funding Fee Details
func FuturesPositionFundingRecords(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/position/funding_records")
	response := utils.FuturesPrivateGet("/api/v1/private/position/funding_records", jsonParams)
	return response
}

// Get Historical Positions
func FuturesPositionListHistoryPositions(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/position/list/history_positions")
	response := utils.FuturesPrivateGet("/api/v1/private/position/list/history_positions", jsonParams)
	return response
}

// Get Open Positions
func FuturesPositionOpenPositions(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/position/open_positions")
	response := utils.FuturesPrivateGet("/api/v1/private/position/open_positions", jsonParams)
	return response
}

// Get Position Leverage Multipliers
func FuturesPositionLeverage(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/position/leverage")
	response := utils.FuturesPrivateGet("/api/v1/private/position/leverage", jsonParams)
	return response
}

// Get Risk Limits
func FuturesAccountRiskLimit(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/account/risk_limit")
	response := utils.FuturesPrivateGet("/api/v1/private/account/risk_limit", jsonParams)
	return response
}

// Get Single Currency Asset Information
func FuturesAccountAsset(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/account/asset/{currency}")
	response := utils.FuturesPrivateGet("/api/v1/private/account/asset/{currency}", jsonParams)
	return response
}

// Get User Position Mode
func FuturesPositionPositionMode(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/position/position_mode")
	response := utils.FuturesPrivateGet("/api/v1/private/position/position_mode", jsonParams)
	return response
}

// Modify Leverage
func FuturesPositionChangeLeveragePost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/position/change_leverage")
	response := utils.FuturesPrivatePost("/api/v1/private/position/change_leverage", jsonParams)
	return response
}

// Modify Position Margin
func FuturesPositionChangeMarginPost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/position/change_margin")
	response := utils.FuturesPrivatePost("/api/v1/private/position/change_margin", jsonParams)
	return response
}

// Modify User Position Mode
func FuturesPositionChangePositionModePost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/position/change_position_mode")
	response := utils.FuturesPrivatePost("/api/v1/private/position/change_position_mode", jsonParams)
	return response
}

// Query All Spot Discount Configuration Information
func FuturesAccountConfigContractFeeDiscountConfig(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/account/config/contractFeeDiscountConfig")
	response := utils.FuturesPrivateGet("/api/v1/private/account/config/contractFeeDiscountConfig", jsonParams)
	return response
}

// Query Contract Fee Deduction Details
func FuturesOrderFeeDetails(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/order/fee_details")
	response := utils.FuturesPrivateGet("/api/v1/private/order/fee_details", jsonParams)
	return response
}

// Query STP Groups and Group Members
func FuturesMarketMakerSelfTradeBlacklist(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/market_maker/self_trade/blacklist")
	response := utils.FuturesPrivateGet("/api/v1/private/market_maker/self_trade/blacklist", jsonParams)
	return response
}

// Query User Discount Usage
func FuturesAccountDiscountType(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/account/discountType")
	response := utils.FuturesPrivateGet("/api/v1/private/account/discountType", jsonParams)
	return response
}

// Update STP Group
func FuturesMarketMakerSelfTradeBlacklistUpdatePost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/market_maker/self_trade/blacklist/update")
	response := utils.FuturesPrivatePost("/api/v1/private/market_maker/self_trade/blacklist/update", jsonParams)
	return response
}

// View Personal Profit Rate
func FuturesAccountProfitRate(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/account/profit_rate/{type}")
	response := utils.FuturesPrivateGet("/api/v1/private/account/profit_rate/{type}", jsonParams)
	return response
}

// Get Candlestick Data
func FuturesKline(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/contract/kline/{symbol}")
	response := utils.FuturesPublicGet("/api/v1/contract/kline/{symbol}", jsonParams)
	return response
}

// Get Contract Info
func FuturesDetailCountry(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/contract/detail/country")
	response := utils.FuturesPublicGet("/api/v1/contract/detail/country", jsonParams)
	return response
}

// Get Contract Order Book Depth
func FuturesDepth(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/contract/depth/{symbol}")
	response := utils.FuturesPublicGet("/api/v1/contract/depth/{symbol}", jsonParams)
	return response
}

// Get Fair Price Candles
func FuturesKlineFairPrice(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/contract/kline/fair_price/{symbol}")
	response := utils.FuturesPublicGet("/api/v1/contract/kline/fair_price/{symbol}", jsonParams)
	return response
}

// Get Fair Price
func FuturesFairPrice(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/contract/fair_price/{symbol}")
	response := utils.FuturesPublicGet("/api/v1/contract/fair_price/{symbol}", jsonParams)
	return response
}

// Get Funding Rate History
func FuturesFundingRateHistory(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/contract/funding_rate/history")
	response := utils.FuturesPublicGet("/api/v1/contract/funding_rate/history", jsonParams)
	return response
}

// Get Funding Rate
func FuturesFundingRate(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/contract/funding_rate/{symbol}")
	response := utils.FuturesPublicGet("/api/v1/contract/funding_rate/{symbol}", jsonParams)
	return response
}

// Get Index Price Candles
func FuturesKlineIndexPrice(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/contract/kline/index_price/{symbol}")
	response := utils.FuturesPublicGet("/api/v1/contract/kline/index_price/{symbol}", jsonParams)
	return response
}

// Get Index Price
func FuturesIndexPrice(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/contract/index_price/{symbol}")
	response := utils.FuturesPublicGet("/api/v1/contract/index_price/{symbol}", jsonParams)
	return response
}

// Get Insurance Fund Balance History
func FuturesRiskReverseHistory(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/contract/risk_reverse/history")
	response := utils.FuturesPublicGet("/api/v1/contract/risk_reverse/history", jsonParams)
	return response
}

// Get Insurance Fund Balance
func FuturesRiskReverse(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/contract/risk_reverse/{symbol}")
	response := utils.FuturesPublicGet("/api/v1/contract/risk_reverse/{symbol}", jsonParams)
	return response
}

// Get Recent Trades
func FuturesDeals(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/contract/deals/{symbol}")
	response := utils.FuturesPublicGet("/api/v1/contract/deals/{symbol}", jsonParams)
	return response
}

// Get Server Time
func FuturesPing(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/contract/ping")
	response := utils.FuturesPublicGet("/api/v1/contract/ping", jsonParams)
	return response
}

// Get the Last N Depth Snapshots
func FuturesDepthCommits(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/contract/depth_commits/{symbol}/{limit}")
	response := utils.FuturesPublicGet("/api/v1/contract/depth_commits/{symbol}/{limit}", jsonParams)
	return response
}

// Get Ticker (Contract Market Data)
func FuturesTicker(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/contract/ticker")
	response := utils.FuturesPublicGet("/api/v1/contract/ticker", jsonParams)
	return response
}

// Get Transferable Currencies
func FuturesSupportCurrencies(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/contract/support_currencies")
	response := utils.FuturesPublicGet("/api/v1/contract/support_currencies", jsonParams)
	return response
}

// Batch Cancel by External Order ID
func FuturesOrderBatchCancelWithExternalPost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/order/batch_cancel_with_external")
	response := utils.FuturesPrivatePost("/api/v1/private/order/batch_cancel_with_external", jsonParams)
	return response
}

// Batch Place Order
func FuturesOrderSubmitBatchPost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/order/submit_batch")
	response := utils.FuturesPrivatePost("/api/v1/private/order/submit_batch", jsonParams)
	return response
}

// Batch Query Orders by External Order ID
func FuturesOrderBatchQueryWithExternalPost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/order/batch_query_with_external")
	response := utils.FuturesPrivatePost("/api/v1/private/order/batch_query_with_external", jsonParams)
	return response
}

// Batch Query Orders by Order ID
func FuturesOrderBatchQuery(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/order/batch_query")
	response := utils.FuturesPrivateGet("/api/v1/private/order/batch_query", jsonParams)
	return response
}

// Cancel All Orders Under a Contract
func FuturesOrderCancelAllPost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/order/cancel_all")
	response := utils.FuturesPrivatePost("/api/v1/private/order/cancel_all", jsonParams)
	return response
}

// Cancel All Planned Orders
func FuturesPlanorderCancelAllPost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/planorder/cancel_all")
	response := utils.FuturesPrivatePost("/api/v1/private/planorder/cancel_all", jsonParams)
	return response
}

// Cancel All TP/SL Planned Orders
func FuturesStoporderCancelAllPost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/stoporder/cancel_all")
	response := utils.FuturesPrivatePost("/api/v1/private/stoporder/cancel_all", jsonParams)
	return response
}

// Cancel by External Order ID
func FuturesOrderCancelWithExternalPost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/order/cancel_with_external")
	response := utils.FuturesPrivatePost("/api/v1/private/order/cancel_with_external", jsonParams)
	return response
}

// Cancel Orders
func FuturesOrderCancelPost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/order/cancel")
	response := utils.FuturesPrivatePost("/api/v1/private/order/cancel", jsonParams)
	return response
}

// Cancel Planned Orders
func FuturesPlanorderCancelPost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/planorder/cancel")
	response := utils.FuturesPrivatePost("/api/v1/private/planorder/cancel", jsonParams)
	return response
}

// Cancel TP/SL Planned Orders
func FuturesStoporderCancelPost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/stoporder/cancel")
	response := utils.FuturesPrivatePost("/api/v1/private/stoporder/cancel", jsonParams)
	return response
}

// Cancel Trailing Order
func FuturesTrackorderCancelPost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/trackorder/cancel")
	response := utils.FuturesPrivatePost("/api/v1/private/trackorder/cancel", jsonParams)
	return response
}

// Chase Order
func FuturesOrderChaseLimitOrderPost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/order/chase_limit_order")
	response := utils.FuturesPrivatePost("/api/v1/private/order/chase_limit_order", jsonParams)
	return response
}

// Close All
func FuturesPositionCloseAllPost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/position/close_all")
	response := utils.FuturesPrivatePost("/api/v1/private/position/close_all", jsonParams)
	return response
}

// Get All Historical Orders
func FuturesOrderListHistoryOrders(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/order/list/history_orders")
	response := utils.FuturesPrivateGet("/api/v1/private/order/list/history_orders", jsonParams)
	return response
}

// Get Current Orders
func FuturesOrderListOpenOrders(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/order/list/open_orders")
	response := utils.FuturesPrivateGet("/api/v1/private/order/list/open_orders", jsonParams)
	return response
}

// Get Current Take-Profit/Stop-Loss Order List
func FuturesStoporderOpenOrders(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/stoporder/open_orders")
	response := utils.FuturesPrivateGet("/api/v1/private/stoporder/open_orders", jsonParams)
	return response
}

// Get Historical Order Deal Details
func FuturesOrderListOrderDeals(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/order/list/order_deals/v3")
	response := utils.FuturesPrivateGet("/api/v1/private/order/list/order_deals/v3", jsonParams)
	return response
}

// Get Order by External ID
func FuturesOrderExternal(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/order/external/{symbol}/{external_oid}")
	response := utils.FuturesPrivateGet("/api/v1/private/order/external/{symbol}/{external_oid}", jsonParams)
	return response
}

// Get Order Information by Order ID
func FuturesOrderGet(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/order/get/{orderId}")
	response := utils.FuturesPrivateGet("/api/v1/private/order/get/{orderId}", jsonParams)
	return response
}

// Get Plan Order List
func FuturesPlanorderListOrders(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/planorder/list/orders")
	response := utils.FuturesPrivateGet("/api/v1/private/planorder/list/orders", jsonParams)
	return response
}

// Get Take-Profit/Stop-Loss Order List
func FuturesStoporderListOrders(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/stoporder/list/orders")
	response := utils.FuturesPrivateGet("/api/v1/private/stoporder/list/orders", jsonParams)
	return response
}

// Get Trade Records by Order ID
func FuturesOrderDealDetails(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/order/deal_details/{orderId}")
	response := utils.FuturesPrivateGet("/api/v1/private/order/deal_details/{orderId}", jsonParams)
	return response
}

// Modify Order Price & Quantity
func FuturesOrderChangeLimitOrderPost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/order/change_limit_order")
	response := utils.FuturesPrivatePost("/api/v1/private/order/change_limit_order", jsonParams)
	return response
}

// Modify Plan Order
func FuturesPlanorderChangePricePost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/planorder/change_price")
	response := utils.FuturesPrivatePost("/api/v1/private/planorder/change_price", jsonParams)
	return response
}

// Modify Take-Profit/Stop-Loss on Plan Order
func FuturesPlanorderChangeStopOrderPost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/planorder/change_stop_order")
	response := utils.FuturesPrivatePost("/api/v1/private/planorder/change_stop_order", jsonParams)
	return response
}

// Modify TP/SL Prices on a Limit Order
func FuturesStoporderChangePricePost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/stoporder/change_price")
	response := utils.FuturesPrivatePost("/api/v1/private/stoporder/change_price", jsonParams)
	return response
}

// Modify TP/SL Prices on a TP/SL Planned Order
func FuturesStoporderChangePlanPricePost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/stoporder/change_plan_price")
	response := utils.FuturesPrivatePost("/api/v1/private/stoporder/change_plan_price", jsonParams)
	return response
}

// Modify Trailing Order
func FuturesTrackorderChangeOrderPost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/trackorder/change_order")
	response := utils.FuturesPrivatePost("/api/v1/private/trackorder/change_order", jsonParams)
	return response
}

// Place Order
func FuturesOrderCreatePost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/order/create")
	response := utils.FuturesPrivatePost("/api/v1/private/order/create", jsonParams)
	return response
}

// Place Plan Order
func FuturesPlanorderPlaceV2Post(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/planorder/place/v2")
	response := utils.FuturesPrivatePost("/api/v1/private/planorder/place/v2", jsonParams)
	return response
}

// Place TP/SL Order by Position
func FuturesStoporderPlacePost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/stoporder/place")
	response := utils.FuturesPrivatePost("/api/v1/private/stoporder/place", jsonParams)
	return response
}

// Place Trailing Order
func FuturesTrackorderPlacePost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/trackorder/place")
	response := utils.FuturesPrivatePost("/api/v1/private/trackorder/place", jsonParams)
	return response
}

// Query Historical Orders
func FuturesOrderListCloseOrders(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/order/list/close_orders")
	response := utils.FuturesPrivateGet("/api/v1/private/order/list/close_orders", jsonParams)
	return response
}

// Query In-Flight Order Counts
func FuturesOrderOpenOrderTotalCountPost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/order/open_order_total_count")
	response := utils.FuturesPrivatePost("/api/v1/private/order/open_order_total_count", jsonParams)
	return response
}

// Query Trailing Orders
func FuturesTrackorderListOrders(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/trackorder/list/orders")
	response := utils.FuturesPrivateGet("/api/v1/private/trackorder/list/orders", jsonParams)
	return response
}

// Reverse Open Position
func FuturesPositionReversePost(jsonParams string) interface{} {
	fmt.Println("path:", "/api/v1/private/position/reverse")
	response := utils.FuturesPrivatePost("/api/v1/private/position/reverse", jsonParams)
	return response
}
