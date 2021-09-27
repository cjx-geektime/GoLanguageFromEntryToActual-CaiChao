package shared_memory_concurrency_mechanism

import (
	"sync"
	"testing"
	"time"
)

// 最终结果不是 10000，因为各个协程之间是 线程不安全的
func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 10000; i++ {
		go func() {
			counter++
		}()
	}
	// 休眠一秒，保证主进程会在协程之后退出
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

// 使用线程安全的方式计数，最终结果正确
func TestCounterThreadSafe(t *testing.T) {
	// 定义 同步互斥锁
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 10000; i++ {
		// 运行一个协程
		go func() {
			// 协程运行结束时打开锁
			defer func() {
				mut.Unlock()
			}()
			// 上锁
			mut.Lock()
			// 计数
			counter++
		}()
	}
	// 休眠一秒，保证主进程会在协程之后退出
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

// 使用等待组，保证协程执行完再退出主程序
func TestCounterWaitGroup(t *testing.T) {
	// 定义 同步互斥锁
	var mut sync.Mutex
	// 定义 同步等待组
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 10000; i++ {
		// 添加一个需要等待的协程
		wg.Add(1)
		// 运行一个协程
		go func() {
			// 协程运行结束时打开锁
			defer func() {
				mut.Unlock()
			}()
			// 上锁
			mut.Lock()
			// 计数
			counter++
			// 取消一个需要等待的协程
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter = %d", counter)
}
