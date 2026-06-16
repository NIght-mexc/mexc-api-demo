package main

import (
	"demo/futures"
	"fmt"
)

var params string = ""

func main() {
	resp := futuresList.FuturesPing(params)
	fmt.Println("返回信息:", resp)
}
