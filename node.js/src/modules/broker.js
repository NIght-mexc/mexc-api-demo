const Broker = superclass => class extends superclass {

  // Create a Sub-account
  brokerSubAccountVirtualSubAccountPost(options = {}) {
    return this.signRequest('POST', '/api/v3/broker/sub-account/virtualSubAccount', options)
  }

  // Create an APIKey for a Sub-account
  brokerSubAccountApiKeyPost(options = {}) {
    return this.signRequest('POST', '/api/v3/broker/sub-account/apiKey', options)
  }

  // Delete the APIKey of a Sub-account
  brokerSubAccountApiKeyDelete(options = {}) {
    return this.signRequest('DELETE', '/api/v3/broker/sub-account/apiKey', options)
  }

  // Deposit Address of Sub-account
  brokerCapitalDepositSubAddress(options = {}) {
    return this.signRequest('GET', '/api/v3/broker/capital/deposit/subAddress', options)
  }

  // Enable Futures for Sub-account
  brokerSubAccountFuturesPost(options = {}) {
    return this.signRequest('POST', '/api/v3/broker/sub-account/futures', options)
  }

  // Generate Deposit Address of Sub-account
  brokerCapitalDepositSubAddressPost(options = {}) {
    return this.signRequest('POST', '/api/v3/broker/capital/deposit/subAddress', options)
  }

  // Get Broker Rebate History Records
  brokerRebateTaxQuery(options = {}) {
    return this.signRequest('GET', '/api/v3/broker/rebate/taxQuery', options)
  }

  // Query All Sub-account Deposit History
  brokerCapitalDepositSubHisrecGetall(options = {}) {
    return this.signRequest('GET', '/api/v3/broker/capital/deposit/subHisrec/getall', options)
  }

  // Query Sub-account Deposit History
  brokerCapitalDepositSubHisrec(options = {}) {
    return this.signRequest('GET', '/api/v3/broker/capital/deposit/subHisrec', options)
  }

  // Query Sub-account List
  brokerSubAccountList(options = {}) {
    return this.signRequest('GET', '/api/v3/broker/sub-account/list', options)
  }

  // Query Sub-account Status
  brokerSubAccountStatus(options = {}) {
    return this.signRequest('GET', '/api/v3/broker/sub-account/status', options)
  }

  // Query the APIKey of a Sub-account
  brokerSubAccountApiKey(options = {}) {
    return this.signRequest('GET', '/api/v3/broker/sub-account/apiKey', options)
  }

  // Query Universal Transfer History - broker user
  brokerSubAccountUniversalTransfer(options = {}) {
    return this.signRequest('GET', '/api/v3/broker/sub-account/universalTransfer', options)
  }

  // Universal Transfer
  brokerSubAccountUniversalTransferPost(options = {}) {
    return this.signRequest('POST', '/api/v3/broker/sub-account/universalTransfer', options)
  }

  // Withdraw
  brokerCapitalWithdrawApplyPost(options = {}) {
    return this.signRequest('POST', '/api/v3/broker/capital/withdraw/apply', options)
  }

}

module.exports = Broker