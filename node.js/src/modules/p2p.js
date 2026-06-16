const P2P = superclass => class extends superclass {

  // Get Market Ads List
  p2PMarketAdsPagination(options = {}) {
    return this.signRequest('GET', '/api/v3/fiat/market/ads/pagination', options)
  }

  // Get All Order List
  p2PMarketOrderPaginationV2(options = {}) {
    return this.signRequest('GET', '/api/v3/fiat/market/order/paginationV2', options)
  }

  // Get My Order List (Maker View)
  p2PMerchantOrderPaginationV2(options = {}) {
    return this.signRequest('GET', '/api/v3/fiat/merchant/order/paginationV2', options)
  }

  // Get My Ads List
  p2PMerchantAdsPagination(options = {}) {
    return this.signRequest('GET', '/api/v3/fiat/merchant/ads/pagination', options)
  }

  // Get Order Detail
  p2POrderDetail(options = {}) {
    return this.signRequest('GET', '/api/v3/fiat/order/detail', options)
  }

  // Toggle Business Status
  p2PMerchantServiceSwitchPost(options = {}) {
    return this.signRequest('POST', '/api/v3/fiat/merchant/service/switch', options)
  }

  // Create/Update Ad
  p2PMerchantAdsSaveOrUpdatePost(options = {}) {
    return this.signRequest('POST', '/api/v3/fiat/merchant/ads/save_or_update', options)
  }

  // Download File
  p2PDownloadFile(options = {}) {
    return this.signRequest('GET', '/api/v3/fiat/downloadFile', options)
  }

  // Generate listenKey
  p2PUserDataStreamPost(options = {}) {
    return this.signRequest('POST', '/api/v3/userDataStream', options)
  }

  // Get listenKey
  p2PUserDataStream(options = {}) {
    return this.signRequest('GET', '/api/v3/userDataStream', options)
  }

  // Get Chat Conversation Info
  p2PRetrieveChatConversation(options = {}) {
    return this.signRequest('GET', '/api/v3/fiat/retrieveChatConversation', options)
  }

  // Get Chat Messages
  p2PRetrieveChatMessageWithPagination(options = {}) {
    return this.signRequest('GET', '/api/v3/fiat/retrieveChatMessageWithPagination', options)
  }

  // Upload File
  p2PUploadFilePost(options = {}) {
    return this.signRequest('POST', '/api/v3/fiat/uploadFile', options)
  }

  // Mark Order as Paid
  p2PConfirmPaidPost(options = {}) {
    return this.signRequest('POST', '/api/v3/fiat/confirm_paid', options)
  }

  // Create Order
  p2PMerchantOrderDealPost(options = {}) {
    return this.signRequest('POST', '/api/v3/fiat/merchant/order/deal', options)
  }

  // Release Cryptocurrency
  p2PReleaseCoinPost(options = {}) {
    return this.signRequest('POST', '/api/v3/fiat/release_coin', options)
  }

}

module.exports = P2P