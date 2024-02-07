package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	go doWork(ctx)
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func doWork(ctx context.Context) {
	for {
		select {
		default:
			fmt.Println("working...")
			time.Sleep(time.Second)
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		}
	}
}
