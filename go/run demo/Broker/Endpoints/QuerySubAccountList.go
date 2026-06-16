package main

import (
	"demo/broker"
	"fmt"
)

var params string = ""

func main() {
	resp := brokerList.BrokerSubAccountList(params)
	fmt.Println("返回信息:", resp)
}
