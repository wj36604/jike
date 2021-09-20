package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancle := context.WithCancel(context.Background())
	ctx = context.WithValue(context.Background(), "wangj", "30")
	wg := sync.WaitGroup{}
	go checkMember(ctx)
	wg.Add(1)
	go func() {
		timeDelay(ctx)
		wg.Done()
	}()
	cancle()

	wg.Wait()
	fmt.Println("wait finished")
	time.Sleep(time.Second * 3)
	fmt.Println("time.sleep finished")

	time.Sleep(time.Second)
}

func timeDelay(ctx context.Context) {
	fmt.Println("start time count")
	defer fmt.Println("stop time count")
	ctx, _ = context.WithTimeout(ctx, time.Second*3)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("time out")
			return
		}

	}
}

func checkMember(ctx context.Context) {
	fmt.Println("start number")
	defer fmt.Println("stop number")
	tk := time.NewTicker(time.Second)
	ctx = context.WithValue(ctx, "louj", "29")
	lessAge(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx close")
			return
		case <-tk.C:
			fmt.Println("ticker time on")
		}

	}
}

func lessAge(ctx context.Context) {
	ageBig := ctx.Value("wangj")
	ageSml := ctx.Value("louj")
	ageMil := ctx.Value("wj")
	fmt.Println("*********age: ",ageBig, ageSml, ageMil)
}
