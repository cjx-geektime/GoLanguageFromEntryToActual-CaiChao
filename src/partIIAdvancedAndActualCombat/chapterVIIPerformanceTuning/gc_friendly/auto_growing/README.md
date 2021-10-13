# 避免内存分配和复制
```
➜  auto_growing git:(main) ✗ go test -bench=.
goos: linux
goarch: amd64
pkg: geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/gc_friendly/auto_growing
cpu: Intel(R) Core(TM) i9-10900K CPU @ 3.70GHz
BenchmarkAutoGrow-20        	    3750	    324889 ns/op
BenchmarkProperInit-20      	   12381	     96547 ns/op
BenchmarkOverSizeInit-20    	    3430	    322145 ns/op
PASS
ok  	geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/gc_friendly/auto_growing	5.269s
```
- 自增长：BenchmarkAutoGrow-20        	    3750	    324889 ns/op
- 合适的大小：BenchmarkProperInit-20      	   12381	     96547 ns/op
- 初始化一个比自增长初始化大八倍的buffer：BenchmarkOverSizeInit-20    	    3430	    322145 ns/op

合适的大小的每次操作时间最小
