# 复杂对象尽量传递引⽤
```
➜  passing_ref git:(main) ✗ ll
总用量 4.0K
-rw-rw-r-- 1 cjx cjx 750 10月 13 22:48 pass_array_test.go
➜  passing_ref git:(main) ✗ go test -bench=.
goos: linux
goarch: amd64
pkg: geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/gc_friendly/passing_ref
cpu: Intel(R) Core(TM) i9-10900K CPU @ 3.70GHz
BenchmarkPassingArrayWithValue-20    	     100	  10089567 ns/op
BenchmarkPassingArrayWithRef-20      	1000000000	         0.4090 ns/op
PASS
ok  	geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/gc_friendly/passing_ref	1.521s
```
# 打开 GC ⽇志
## 值传递
只运行 BenchmarkPassingArrayWithValue
```
➜  passing_ref git:(main) ✗ GODEBUG=gctrace=1 go test -bench=BenchmarkPassingArrayWithValue
gc 1 @0.003s 2%: 0.047+0.25+0.021 ms clock, 0.94+0.16/0.29/0.009+0.43 ms cpu, 4->4->0 MB, 5 MB goal, 20 P
gc 2 @0.005s 2%: 0.022+0.22+0.017 ms clock, 0.45+0.12/0.44/0+0.35 ms cpu, 4->4->0 MB, 5 MB goal, 20 P
......
gc 57 @1.094s 0%: 0.014+0.21+0.005 ms clock, 0.29+0/0.15/0.071+0.11 ms cpu, 305->305->152 MB, 306 MB goal, 20 P
gc 58 @1.117s 0%: 0.015+0.18+0.007 ms clock, 0.30+0/0.16/0.084+0.14 ms cpu, 305->305->152 MB, 306 MB goal, 20 P
     100	  11003071 ns/op
PASS
ok  	geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/gc_friendly/passing_ref	1.136s
```
58次 GC ，每次操作用时 11003071 ns/op
## 引用传递
只运行 BenchmarkPassingArrayWithRef
```
➜  passing_ref git:(main) ✗ GODEBUG=gctrace=1 go test -bench=BenchmarkPassingArrayWithRef  
gc 1 @0.007s 6%: 0.44+0.29+0.056 ms clock, 8.9+0.10/0.32/0+1.1 ms cpu, 4->4->0 MB, 5 MB goal, 20 P
gc 2 @0.008s 6%: 0.065+0.19+0.013 ms clock, 1.3+0.13/0.32/0+0.27 ms cpu, 4->4->0 MB, 5 MB goal, 20 P
......
gc 14 @0.077s 0%: 0.040+0.12+0.001 ms clock, 0.81+0/0.099/0.018+0.038 ms cpu, 0->0->0 MB, 4 MB goal, 20 P (forced)
gc 15 @0.082s 0%: 0.006+0.19+0.005 ms clock, 0.12+0/0.10/0.041+0.10 ms cpu, 76->76->0 MB, 77 MB goal, 20 P
1000000000	         0.4061 ns/op
PASS
ok  	geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/gc_friendly/passing_ref	0.489s
```
15次 GC ，每次操作用时 0.4061 ns/op
# go tool trace
生成分析文件
- 值传递：trace_val.out
- 引用传递：trace_ref.out
```
➜  passing_ref git:(main) ✗ go test -bench=BenchmarkPassingArrayWithRef -trace=trace_ref.out
goos: linux
goarch: amd64
pkg: geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/gc_friendly/passing_ref
cpu: Intel(R) Core(TM) i9-10900K CPU @ 3.70GHz
BenchmarkPassingArrayWithRef-20    	1000000000	         0.4060 ns/op
PASS
ok  	geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/gc_friendly/passing_ref	0.491s
➜  passing_ref git:(main) ✗ ls
pass_array_test.go  README.md  trace_ref.out
➜  passing_ref git:(main) ✗ go test -bench=BenchmarkPassingArrayWithValue -trace=trace_val.out
goos: linux
goarch: amd64
pkg: geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/gc_friendly/passing_ref
cpu: Intel(R) Core(TM) i9-10900K CPU @ 3.70GHz
BenchmarkPassingArrayWithValue-20    	     120	  10097241 ns/op
PASS
ok  	geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/gc_friendly/passing_ref	2.249s
➜  passing_ref git:(main) ✗ ls
pass_array_test.go  README.md  trace_ref.out  trace_val.out
```
## 可视化分析
```
➜  passing_ref git:(main) ✗ go tool trace trace_ref.out
2021/10/13 23:02:06 Parsing trace...
2021/10/13 23:02:06 Splitting trace...
2021/10/13 23:02:06 Opening browser. Trace viewer is listening on http://127.0.0.1:43079
```
自动跳转到浏览器
- View trace：查看过程
