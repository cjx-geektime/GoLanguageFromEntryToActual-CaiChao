# sync.Map
• 适合读多写少，且 Key 相对稳定的环境
• 采⽤了空间换时间的⽅案，并且采⽤指针的⽅式间接实现值的映射，所以存储空间会较 built-in map ⼤
# Concurrent Map
• 适⽤于读写都很频繁的情况
# 别让性能被“锁”住
• 减少锁的影响范围
• 减少发⽣锁冲突的概率
  • sync.Map
  • ConcurrentMap
• 避免锁的使⽤
  • LAMX Disruptor：https://martinfowler.com/articles/lmax.html
# 分析
```
➜  maps git:(main) ✗ go test -bench=.
goos: linux
goarch: amd64
pkg: geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/maps
cpu: Intel(R) Core(TM) i9-10900K CPU @ 3.70GHz
BenchmarkSyncmap/map_with_RWLock-20         	    1132	   1007869 ns/op
BenchmarkSyncmap/sync.map-20                	    1486	    802110 ns/op
BenchmarkSyncmap/concurrent_map-20          	    2066	    574189 ns/op
PASS
ok  	geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/maps	3.771s
```
修改 map_benchmark_test.go 的代码，测试各种读写情况下的性能
```
const (
	NumOfReader = 100
	NumOfWriter = 10
)
```
