package main

import (
	"demo/p2p"
	"fmt"
)

var params string = ""

func main() {
	resp := p2pList.P2POrderDetail(params)
	fmt.Println("返回信息:", resp)
}
