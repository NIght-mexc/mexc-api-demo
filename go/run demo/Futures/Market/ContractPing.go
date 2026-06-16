package main

import (
	"context"
	"fmt"
	"time"

	"mexc-futures-go/mexcfutures"
)

func main() {
	cfg := mexcfutures.DefaultConfig()
	client := mexcfutures.NewFuturesRestClient(cfg)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	resp, err := client.Ping(ctx)
	if err != nil {
		fmt.Println("错误:", err)
		return
	}
	fmt.Println("返回信息:", resp.Text)
}
