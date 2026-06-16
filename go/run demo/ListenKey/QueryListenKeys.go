package main

import (
	"demo/spot"
	"fmt"
)

var params string = ""

func main() {
	resp := spotList.QueryListenKeys(params)
	fmt.Println("返回信息:", resp)
}
