package p2pList

import (
	"demo/config"
	"demo/utils"
	"fmt"
)

// Get Market Ads List
func P2PMarketAdsPagination(jsonParams string) interface{} {
	caseUrl := "/fiat/market/ads/pagination"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivateGet(requestUrl, jsonParams)
	return response
}

// Get All Order List
func P2PMarketOrderPaginationV2(jsonParams string) interface{} {
	caseUrl := "/fiat/market/order/paginationV2"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivateGet(requestUrl, jsonParams)
	return response
}

// Get My Order List (Maker View)
func P2PMerchantOrderPaginationV2(jsonParams string) interface{} {
	caseUrl := "/fiat/merchant/order/paginationV2"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivateGet(requestUrl, jsonParams)
	return response
}

// Get My Ads List
func P2PMerchantAdsPagination(jsonParams string) interface{} {
	caseUrl := "/fiat/merchant/ads/pagination"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivateGet(requestUrl, jsonParams)
	return response
}

// Get Order Detail
func P2POrderDetail(jsonParams string) interface{} {
	caseUrl := "/fiat/order/detail"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivateGet(requestUrl, jsonParams)
	return response
}

// Toggle Business Status
func P2PMerchantServiceSwitchPost(jsonParams string) interface{} {
	caseUrl := "/fiat/merchant/service/switch"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivatePost(requestUrl, jsonParams)
	return response
}

// Create/Update Ad
func P2PMerchantAdsSaveOrUpdatePost(jsonParams string) interface{} {
	caseUrl := "/fiat/merchant/ads/save_or_update"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivatePost(requestUrl, jsonParams)
	return response
}

// Download File
func P2PDownloadFile(jsonParams string) interface{} {
	caseUrl := "/fiat/downloadFile"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivateGet(requestUrl, jsonParams)
	return response
}

// Generate listenKey
func P2PUserDataStreamPost(jsonParams string) interface{} {
	caseUrl := "/userDataStream"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivatePost(requestUrl, jsonParams)
	return response
}

// Get listenKey
func P2PUserDataStream(jsonParams string) interface{} {
	caseUrl := "/userDataStream"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivateGet(requestUrl, jsonParams)
	return response
}

// Get Chat Conversation Info
func P2PRetrieveChatConversation(jsonParams string) interface{} {
	caseUrl := "/fiat/retrieveChatConversation"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivateGet(requestUrl, jsonParams)
	return response
}

// Get Chat Messages
func P2PRetrieveChatMessageWithPagination(jsonParams string) interface{} {
	caseUrl := "/fiat/retrieveChatMessageWithPagination"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivateGet(requestUrl, jsonParams)
	return response
}

// Upload File
func P2PUploadFilePost(jsonParams string) interface{} {
	caseUrl := "/fiat/uploadFile"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivatePost(requestUrl, jsonParams)
	return response
}

// Mark Order as Paid
func P2PConfirmPaidPost(jsonParams string) interface{} {
	caseUrl := "/fiat/confirm_paid"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivatePost(requestUrl, jsonParams)
	return response
}

// Create Order
func P2PMerchantOrderDealPost(jsonParams string) interface{} {
	caseUrl := "/fiat/merchant/order/deal"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivatePost(requestUrl, jsonParams)
	return response
}

// Release Cryptocurrency
func P2PReleaseCoinPost(jsonParams string) interface{} {
	caseUrl := "/fiat/release_coin"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivatePost(requestUrl, jsonParams)
	return response
}
