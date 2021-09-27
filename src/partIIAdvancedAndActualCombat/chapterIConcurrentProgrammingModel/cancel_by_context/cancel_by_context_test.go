package cancel_by_context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancalled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	// 创建根 Context
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 10; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancalled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}
