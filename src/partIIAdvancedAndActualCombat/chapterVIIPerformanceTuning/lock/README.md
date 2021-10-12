# 分析
```
➜  lock git:(main) ✗ go test -bench=.
goos: linux
goarch: amd64
pkg: geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/lock
cpu: Intel(R) Core(TM) i9-10900K CPU @ 3.70GHz
BenchmarkLockFree-20    	     837	   1276458 ns/op
BenchmarkLock-20        	      10	 114363953 ns/op
PASS
ok  	geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/lock	2.462s
➜  lock git:(main) ✗ go test -bench=. -cpuprofile=cpu.prof
goos: linux
goarch: amd64
pkg: geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/lock
cpu: Intel(R) Core(TM) i9-10900K CPU @ 3.70GHz
BenchmarkLockFree-20    	     850	   1253038 ns/op
BenchmarkLock-20        	      10	 116450342 ns/op
PASS
ok  	geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/lock	2.654s
➜  lock git:(main) ✗ go tool pprof cpu.prof
File: lock.test
Type: cpu
Time: Oct 12, 2021 at 10:45pm (CST)
Duration: 2.65s, Total samples = 7.42s (280.00%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 7.39s, 99.60% of 7.42s total
Dropped 16 nodes (cum <= 0.04s)
      flat  flat%   sum%        cum   cum%
     2.15s 28.98% 28.98%      2.15s 28.98%  sync.(*RWMutex).RUnlock (inline)
     2.14s 28.84% 57.82%      2.14s 28.84%  sync.(*RWMutex).RLock (inline)
     1.95s 26.28% 84.10%      2.38s 32.08%  runtime.mapaccess2_faststr
     0.67s  9.03% 93.13%      3.03s 40.84%  geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/lock_test.lockFreeAccess.func1
     0.43s  5.80% 98.92%      0.43s  5.80%  runtime.add (inline)
     0.04s  0.54% 99.46%      0.04s  0.54%  sync.(*WaitGroup).Done (inline)
     0.01s  0.13% 99.60%      4.36s 58.76%  geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/lock_test.lockAccess.func1
         0     0% 99.60%      0.08s  1.08%  runtime.(*bmap).keys (inline)
(pprof) list lockAccess
Total: 7.42s
ROUTINE ======================== geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/lock_test.lockAccess.func1 in /bigdata/git-projects/programming-projects/geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/lock/lock_test.go
      10ms      4.36s (flat, cum) 58.76% of Total
         .          .     43:	m := new(sync.RWMutex)
         .          .     44:	for i := 0; i < NUM_OF_READER; i++ {
         .          .     45:		go func() {
         .          .     46:			for j := 0; j < READ_TIMES; j++ {
         .          .     47:
         .      2.14s     48:				m.RLock()
      10ms       70ms     49:				_, err := cache["a"]
         .          .     50:				if !err {
         .          .     51:					fmt.Println("Nothing")
         .          .     52:				}
         .      2.15s     53:				m.RUnlock()
         .          .     54:			}
         .          .     55:			wg.Done()
         .          .     56:		}()
         .          .     57:	}
         .          .     58:	wg.Wait()
```
         .      2.14s     48:				m.RLock()
