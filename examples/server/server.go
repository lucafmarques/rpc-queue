package main

import (
	"context"
	"fmt"
	"time"

	"github.com/lucafmarques/rpchan"
)

type T struct {
	A, B int
	C    string
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	for v := range channel.Listen[T](ctx, 1) {
		fmt.Printf("%+v\n", v)
	}
}
