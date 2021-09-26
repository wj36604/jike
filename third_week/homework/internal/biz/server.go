package biz

import (
	"context"
	"fmt"
	"time"
)

func CheckTask(ctx context.Context) {
	tk := time.NewTicker(time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("exit check task")
			return
		case <-tk.C:
			// 相关业务处理
		}
	}
}