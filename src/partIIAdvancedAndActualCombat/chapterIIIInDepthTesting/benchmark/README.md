# Benchmark
```
func BenchmarkConcatStringByAdd(b *testing.B) {
//与性能测试无无关的代码
b.ResetTimer()
for i := 0; i < b.N; i++ {
//测试代码
}
b.StopTimer()
//与性能测试无无关的代码
}
```
## 使用
-bench=<相关benchmark测试>
```
. 表示目录下的全部
go test -bench=.
查看更多信息
go test -bench=. -benchmem
```
## 测试
```
✗ go test -bench=.
goos: linux
goarch: amd64
pkg: geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterIIIInDepthTesting/benchmark
cpu: Intel(R) Core(TM) i9-10900K CPU @ 3.70GHz
BenchmarkConcatStringByAdd-20            	14025376	        82.59 ns/op
BenchmarkConcatStringByBytesBuffer-20    	31166922	        77.18 ns/op
PASS
ok  	geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterIIIInDepthTesting/benchmark	3.694s
```
```
✗ go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterIIIInDepthTesting/benchmark
cpu: Intel(R) Core(TM) i9-10900K CPU @ 3.70GHz
BenchmarkConcatStringByAdd-20            	14129676	        82.16 ns/op	      16 B/op	       4 allocs/op
BenchmarkConcatStringByBytesBuffer-20    	28282765	        68.07 ns/op	      64 B/op	       1 allocs/op
PASS
ok  	geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterIIIInDepthTesting/benchmark	3.220s
```
