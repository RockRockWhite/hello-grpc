package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//withCancle()
	//withTimeout()
	withValue()
}

func withCancle() {
	pctx, pcancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithCancel(pctx)
	defer cancel()
	defer pcancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return
			default:
				fmt.Println("hello")
			}
		}
	}()

	go func() {
		for {
			select {
			case <-pctx.Done():
				fmt.Println(pctx.Err())
				return
			default:
				fmt.Println("phello")
			}
		}
	}()

	cancel()
	pcancel()
	time.Sleep(5 * time.Second)
}

func withTimeout() {
	pctx, pcancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(pctx, 2*time.Second)
	defer cancel()
	defer pcancel()

	time.Sleep(5 * time.Second)

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return
			default:
				fmt.Println(ctx.Value("key"))
				fmt.Println("hello")
			}
		}
	}()

	go func() {
		for {
			select {
			case <-pctx.Done():
				fmt.Println(pctx.Err())
				return
			default:
				//fmt.Println("phello")
			}
		}
	}()
	time.Sleep(5 * time.Second)
}

func withValue() {
	pctx, pcancel := context.WithCancel(context.Background())
	ctx := context.WithValue(pctx, "key", "value")
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	defer pcancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return
			default:
				fmt.Println(ctx.Value("key"))
				fmt.Println("hello")
			}
		}
	}()

	go func() {
		for {
			select {
			case <-pctx.Done():
				fmt.Println(pctx.Err())
				return
			default:
				//fmt.Println("phello")
			}
		}
	}()
	time.Sleep(5 * time.Second)
}
