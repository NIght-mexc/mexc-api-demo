package brokerList

import (
	"demo/config"
	"demo/utils"
	"fmt"
)

// Create a Sub-account
func BrokerSubAccountVirtualSubAccountPost(jsonParams string) interface{} {
	caseUrl := "/broker/sub-account/virtualSubAccount"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivatePost(requestUrl, jsonParams)
	return response
}

// Create an APIKey for a Sub-account
func BrokerSubAccountApiKeyPost(jsonParams string) interface{} {
	caseUrl := "/broker/sub-account/apiKey"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivatePost(requestUrl, jsonParams)
	return response
}

// Delete the APIKey of a Sub-account
func BrokerSubAccountApiKeyDelete(jsonParams string) interface{} {
	caseUrl := "/broker/sub-account/apiKey"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivateDelete(requestUrl, jsonParams)
	return response
}

// Deposit Address of Sub-account
func BrokerCapitalDepositSubAddress(jsonParams string) interface{} {
	caseUrl := "/broker/capital/deposit/subAddress"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivateGet(requestUrl, jsonParams)
	return response
}

// Enable Futures for Sub-account
func BrokerSubAccountFuturesPost(jsonParams string) interface{} {
	caseUrl := "/broker/sub-account/futures"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivatePost(requestUrl, jsonParams)
	return response
}

// Generate Deposit Address of Sub-account
func BrokerCapitalDepositSubAddressPost(jsonParams string) interface{} {
	caseUrl := "/broker/capital/deposit/subAddress"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivatePost(requestUrl, jsonParams)
	return response
}

// Get Broker Rebate History Records
func BrokerRebateTaxQuery(jsonParams string) interface{} {
	caseUrl := "/broker/rebate/taxQuery"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivateGet(requestUrl, jsonParams)
	return response
}

// Query All Sub-account Deposit History
func BrokerCapitalDepositSubHisrecGetall(jsonParams string) interface{} {
	caseUrl := "/broker/capital/deposit/subHisrec/getall"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivateGet(requestUrl, jsonParams)
	return response
}

// Query Sub-account Deposit History
func BrokerCapitalDepositSubHisrec(jsonParams string) interface{} {
	caseUrl := "/broker/capital/deposit/subHisrec"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivateGet(requestUrl, jsonParams)
	return response
}

// Query Sub-account List
func BrokerSubAccountList(jsonParams string) interface{} {
	caseUrl := "/broker/sub-account/list"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivateGet(requestUrl, jsonParams)
	return response
}

// Query Sub-account Status
func BrokerSubAccountStatus(jsonParams string) interface{} {
	caseUrl := "/broker/sub-account/status"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivateGet(requestUrl, jsonParams)
	return response
}

// Query the APIKey of a Sub-account
func BrokerSubAccountApiKey(jsonParams string) interface{} {
	caseUrl := "/broker/sub-account/apiKey"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivateGet(requestUrl, jsonParams)
	return response
}

// Query Universal Transfer History - broker user
func BrokerSubAccountUniversalTransfer(jsonParams string) interface{} {
	caseUrl := "/broker/sub-account/universalTransfer"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivateGet(requestUrl, jsonParams)
	return response
}

// Universal Transfer
func BrokerSubAccountUniversalTransferPost(jsonParams string) interface{} {
	caseUrl := "/broker/sub-account/universalTransfer"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivatePost(requestUrl, jsonParams)
	return response
}

// Withdraw
func BrokerCapitalWithdrawApplyPost(jsonParams string) interface{} {
	caseUrl := "/broker/capital/withdraw/apply"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PrivatePost(requestUrl, jsonParams)
	return response
}
