# 性能调优示例
## 常⻅分析指标
• Wall Time
• CPU Time
• Block Time
• Memory allocation
• GC times/time spent
### work_time
通常是wall time和cpu time。
- wall time指程序运行的总体时间，包含程序不耗用cpu，被阻塞时的时间
- cpu time，指程序占用cpu的时间
# 分析
## 设定优化⽬标
## 分析系统瓶颈点
CPU
```
➜  optmization git:(main) ✗ go test -bench=.
goos: linux
goarch: amd64
pkg: geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization
cpu: Intel(R) Core(TM) i9-10900K CPU @ 3.70GHz
BenchmarkProcessRequest-20       	  284455	      5836 ns/op
BenchmarkProcessRequestOld-20    	   96697	     15265 ns/op
PASS
ok  	geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization	3.308s
➜  optmization git:(main) ✗ go test -bench=. -cpuprofile=cpu.prof
goos: linux
goarch: amd64
pkg: geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization
cpu: Intel(R) Core(TM) i9-10900K CPU @ 3.70GHz
BenchmarkProcessRequest-20       	  280094	      6241 ns/op
BenchmarkProcessRequestOld-20    	   89142	     13555 ns/op
PASS
ok  	geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization	3.315s
➜  optmization git:(main) ✗ ll
总用量 4.7M
-rw-rw-r-- 1 cjx cjx  14K 10月 12 22:21 cpu.prof
-rw-rw-r-- 1 cjx cjx  661 10月 12 22:18 optimization_test.go
-rw-rw-r-- 1 cjx cjx 1.3K 10月 12 22:15 optmization.go
-rwxrwxr-x 1 cjx cjx 4.7M 10月 12 22:21 optmization.test
-rw-rw-r-- 1 cjx cjx  321 10月 12 22:20 README.md
-rw-rw-r-- 1 cjx cjx 4.5K 10月 12 22:17 structs_easyjson.go
-rw-rw-r-- 1 cjx cjx  235 10月 12 22:16 structs.go
➜  optmization git:(main) ✗ go tool pprof cpu.prof
File: optmization.test
Type: cpu
Time: Oct 12, 2021 at 10:21pm (CST)
Duration: 3.31s, Total samples = 3.62s (109.32%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 1650ms, 45.58% of 3620ms total
Dropped 79 nodes (cum <= 18.10ms)
Showing top 10 nodes out of 146
      flat  flat%   sum%        cum   cum%
     240ms  6.63%  6.63%      270ms  7.46%  github.com/mailru/easyjson/jlexer.(*Lexer).fetchNumber
     240ms  6.63% 13.26%      240ms  6.63%  strconv.ParseUint
     210ms  5.80% 19.06%      520ms 14.36%  github.com/mailru/easyjson/jlexer.(*Lexer).FetchToken
     200ms  5.52% 24.59%      440ms 12.15%  strconv.ParseInt
     180ms  4.97% 29.56%      660ms 18.23%  runtime.mallocgc
     140ms  3.87% 33.43%      230ms  6.35%  github.com/mailru/easyjson/jwriter.(*Writer).String
     120ms  3.31% 36.74%      120ms  3.31%  runtime.memclrNoHeapPointers
     110ms  3.04% 39.78%      420ms 11.60%  runtime.concatstrings
     110ms  3.04% 42.82%      110ms  3.04%  runtime.memmove
     100ms  2.76% 45.58%      100ms  2.76%  runtime.nextFreeFast (inline)
(pprof) top -cum
Showing nodes accounting for 0.31s, 8.56% of 3.62s total
Dropped 79 nodes (cum <= 0.02s)
Showing top 10 nodes out of 146
      flat  flat%   sum%        cum   cum%
         0     0%     0%      3.07s 84.81%  testing.(*B).launch
         0     0%     0%      3.07s 84.81%  testing.(*B).runN
         0     0%     0%      1.77s 48.90%  geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization.BenchmarkProcessRequest
     0.04s  1.10%  1.10%      1.77s 48.90%  geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization.processRequest
         0     0%  1.10%      1.59s 43.92%  geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization.(*Request).UnmarshalJSON (partial-inline)
     0.08s  2.21%  3.31%      1.59s 43.92%  geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization.easyjson6a975c40DecodeCh471
         0     0%  3.31%      1.30s 35.91%  geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization.BenchmarkProcessRequestOld
     0.01s  0.28%  3.59%      1.30s 35.91%  geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization.processRequestOld
         0     0%  3.59%      0.67s 18.51%  encoding/json.Unmarshal
     0.18s  4.97%  8.56%      0.66s 18.23%  runtime.mallocgc

     (pprof) list processRequest
     Total: 3.62s
     ROUTINE ======================== geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization.processRequest in /bigdata/git-projects/programming-projects/geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization/optmization.go
           40ms      1.77s (flat, cum) 48.90% of Total
              .          .     21:
              .          .     22:func processRequest(reqs []string) []string {
              .          .     23:	reps := []string{}
              .          .     24:	for _, req := range reqs {
              .          .     25:		reqObj := &Request{}
              .      1.18s     26:		reqObj.UnmarshalJSON([]byte(req))
              .          .     27:		//	json.Unmarshal([]byte(req), reqObj)
              .          .     28:
              .          .     29:		var buf strings.Builder
           30ms       30ms     30:		for _, e := range reqObj.PayLoad {
              .      220ms     31:			buf.WriteString(strconv.Itoa(e))
              .       50ms     32:			buf.WriteString(",")
              .          .     33:		}
              .          .     34:		repObj := &Response{reqObj.TransactionID, buf.String()}
              .      210ms     35:		repJson, err := repObj.MarshalJSON()
              .          .     36:		//repJson, err := json.Marshal(&repObj)
              .          .     37:		if err != nil {
              .          .     38:			panic(err)
              .          .     39:		}
              .       70ms     40:		reps = append(reps, string(repJson))
              .          .     41:	}
           10ms       10ms     42:	return reps
              .          .     43:}
              .          .     44:
              .          .     45:func processRequestOld(reqs []string) []string {
              .          .     46:	reps := []string{}
              .          .     47:	for _, req := range reqs {
     ROUTINE ======================== geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization.processRequestOld in /bigdata/git-projects/programming-projects/geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization/optmization.go
           10ms      1.30s (flat, cum) 35.91% of Total
              .          .     43:}
              .          .     44:
              .          .     45:func processRequestOld(reqs []string) []string {
              .          .     46:	reps := []string{}
              .          .     47:	for _, req := range reqs {
              .       10ms     48:		reqObj := &Request{}
              .      680ms     49:		json.Unmarshal([]byte(req), reqObj)
              .          .     50:		ret := ""
              .          .     51:		for _, e := range reqObj.PayLoad {
           10ms      440ms     52:			ret += strconv.Itoa(e) + ","
              .          .     53:		}
              .       10ms     54:		repObj := &Response{reqObj.TransactionID, ret}
              .      160ms     55:		repJson, err := json.Marshal(&repObj)
              .          .     56:		if err != nil {
              .          .     57:			panic(err)
              .          .     58:		}
              .          .     59:		reps = append(reps, string(repJson))
              .          .     60:	}
```
## 优化瓶颈点
使用 easyjson
```
➜  optmization git:(main) ✗ ls
cpu.prof  optimization_test.go  optmization.go  optmization.test  README.md  structs_easyjson.go  structs.go
➜  optmization git:(main) ✗ easyjson -all structs.go
zsh: command not found: easyjson
➜  optmization git:(main) ✗  ~/go/bin/easyjson -all structs.go   
➜  optmization git:(main) ✗ ll
总用量 4.7M
-rw-rw-r-- 1 cjx cjx  14K 10月 12 22:21 cpu.prof
-rw-rw-r-- 1 cjx cjx  661 10月 12 22:18 optimization_test.go
-rw-rw-r-- 1 cjx cjx 1.3K 10月 12 22:15 optmization.go
-rwxrwxr-x 1 cjx cjx 4.7M 10月 12 22:21 optmization.test
-rw-rw-r-- 1 cjx cjx 7.8K 10月 12 22:26 README.md
-rw-rw-r-- 1 cjx cjx 5.4K 10月 12 22:29 structs_easyjson.go
-rw-rw-r-- 1 cjx cjx  235 10月 12 22:16 structs.go
```
分析CPU
```
➜  optmization git:(main) ✗ go test -bench=.
goos: linux
goarch: amd64
pkg: geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization
cpu: Intel(R) Core(TM) i9-10900K CPU @ 3.70GHz
BenchmarkProcessRequest-20       	  285826	      6368 ns/op
BenchmarkProcessRequestOld-20    	   92131	     14180 ns/op
PASS
ok  	geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization	3.305s
➜  optmization git:(main) ✗ go test -bench=. -cpuprofile=cpu.prof
goos: linux
goarch: amd64
pkg: geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization
cpu: Intel(R) Core(TM) i9-10900K CPU @ 3.70GHz
BenchmarkProcessRequest-20       	  264986	      7533 ns/op
BenchmarkProcessRequestOld-20    	   83464	     16830 ns/op
PASS
ok  	geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization	3.727s
➜  optmization git:(main) ✗ go tool pprof cpu.prof
File: optmization.test
Type: cpu
Time: Oct 12, 2021 at 10:31pm (CST)
Duration: 3.72s, Total samples = 4.07s (109.53%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top -cum
Showing nodes accounting for 0.18s, 4.42% of 4.07s total
Dropped 113 nodes (cum <= 0.02s)
Showing top 10 nodes out of 121
      flat  flat%   sum%        cum   cum%
         0     0%     0%      3.40s 83.54%  testing.(*B).launch
         0     0%     0%      3.40s 83.54%  testing.(*B).runN
         0     0%     0%      1.95s 47.91%  geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization.BenchmarkProcessRequest
     0.02s  0.49%  0.49%      1.95s 47.91%  geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization.processRequest
         0     0%  0.49%      1.68s 41.28%  geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization.(*Request).UnmarshalJSON (partial-inline)
     0.09s  2.21%  2.70%      1.68s 41.28%  geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization.easyjson6a975c40DecodeGeekGoLanguageFromEntryToActualCombatSrcPartIIAdvancedAndActualCombatChapterVIIPerformanceTuningOptmization1
         0     0%  2.70%      1.45s 35.63%  geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization.BenchmarkProcessRequestOld
         0     0%  2.70%      1.45s 35.63%  geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization.processRequestOld
     0.01s  0.25%  2.95%      0.74s 18.18%  github.com/mailru/easyjson/jlexer.(*Lexer).Int (inline)
     0.06s  1.47%  4.42%      0.73s 17.94%  github.com/mailru/easyjson/jlexer.(*Lexer).Int64
(pprof) list processRequest
Total: 4.07s
ROUTINE ======================== geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization.processRequest in /bigdata/git-projects/programming-projects/geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization/optmization.go
      20ms      1.95s (flat, cum) 47.91% of Total
         .          .     21:
         .          .     22:func processRequest(reqs []string) []string {
         .          .     23:	reps := []string{}
         .          .     24:	for _, req := range reqs {
         .          .     25:		reqObj := &Request{}
         .      1.43s     26:		reqObj.UnmarshalJSON([]byte(req))
         .          .     27:		//	json.Unmarshal([]byte(req), reqObj)
         .          .     28:
         .          .     29:		var buf strings.Builder
      10ms       10ms     30:		for _, e := range reqObj.PayLoad {
         .      130ms     31:			buf.WriteString(strconv.Itoa(e))
      10ms       90ms     32:			buf.WriteString(",")
         .          .     33:		}
         .          .     34:		repObj := &Response{reqObj.TransactionID, buf.String()}
         .      260ms     35:		repJson, err := repObj.MarshalJSON()
         .          .     36:		//repJson, err := json.Marshal(&repObj)
         .          .     37:		if err != nil {
         .          .     38:			panic(err)
         .          .     39:		}
         .       30ms     40:		reps = append(reps, string(repJson))
         .          .     41:	}
         .          .     42:	return reps
         .          .     43:}
         .          .     44:
         .          .     45:func processRequestOld(reqs []string) []string {
ROUTINE ======================== geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization.processRequestOld in /bigdata/git-projects/programming-projects/geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization/optmization.go
         0      1.45s (flat, cum) 35.63% of Total
         .          .     44:
         .          .     45:func processRequestOld(reqs []string) []string {
         .          .     46:	reps := []string{}
         .          .     47:	for _, req := range reqs {
         .          .     48:		reqObj := &Request{}
         .      620ms     49:		json.Unmarshal([]byte(req), reqObj)
         .          .     50:		ret := ""
         .          .     51:		for _, e := range reqObj.PayLoad {
         .      550ms     52:			ret += strconv.Itoa(e) + ","
         .          .     53:		}
         .          .     54:		repObj := &Response{reqObj.TransactionID, ret}
         .      280ms     55:		repJson, err := json.Marshal(&repObj)
         .          .     56:		if err != nil {
         .          .     57:			panic(err)
         .          .     58:		}
         .          .     59:		reps = append(reps, string(repJson))
         .          .     60:	}
```
分析内存
```
➜  optmization git:(main) ✗ go test -bench=. -memprofile=mem.prof
goos: linux
goarch: amd64
pkg: geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization
cpu: Intel(R) Core(TM) i9-10900K CPU @ 3.70GHz
BenchmarkProcessRequest-20       	  277652	      5842 ns/op
BenchmarkProcessRequestOld-20    	   91297	     14769 ns/op
PASS
ok  	geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization	3.152s
➜  optmization git:(main) ✗ go tool pprof mem.prof
File: optmization.test
Type: alloc_space
Time: Oct 12, 2021 at 10:35pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 3.09GB, 99.46% of 3.10GB total
Dropped 30 nodes (cum <= 0.02GB)
Showing top 10 nodes out of 29
      flat  flat%   sum%        cum   cum%
    1.52GB 49.08% 49.08%     1.84GB 59.32%  geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization.processRequestOld
    0.70GB 22.68% 71.76%     0.71GB 22.87%  geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization.easyjson6a975c40DecodeGeekGoLanguageFromEntryToActualCombatSrcPartIIAdvancedAndActualCombatChapterVIIPerformanceTuningOptmization1
    0.27GB  8.79% 80.55%     0.27GB  8.79%  strings.(*Builder).WriteString (inline)
    0.18GB  5.68% 86.24%     1.26GB 40.54%  geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization.processRequest
    0.15GB  4.68% 90.91%     0.15GB  4.68%  github.com/mailru/easyjson/buffer.getBuf
    0.13GB  4.27% 95.18%     0.13GB  4.27%  github.com/mailru/easyjson/buffer.(*Buffer).BuildBytes
    0.08GB  2.71% 97.89%     0.23GB  7.38%  github.com/mailru/easyjson/buffer.(*Buffer).ensureSpaceSlow
    0.03GB  1.12% 99.01%     0.13GB  4.28%  encoding/json.Marshal
    0.01GB  0.46% 99.46%     0.18GB  5.95%  encoding/json.Unmarshal
         0     0% 99.46%     0.17GB  5.39%  encoding/json.(*decodeState).object
(pprof) list processRequest
Total: 3.10GB
ROUTINE ======================== geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization.processRequest in /bigdata/git-projects/programming-projects/geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization/optmization.go
  180.56MB     1.26GB (flat, cum) 40.54% of Total
         .          .     21:
         .          .     22:func processRequest(reqs []string) []string {
         .          .     23:	reps := []string{}
         .          .     24:	for _, req := range reqs {
         .          .     25:		reqObj := &Request{}
   88.03MB   643.40MB     26:		reqObj.UnmarshalJSON([]byte(req))
         .          .     27:		//	json.Unmarshal([]byte(req), reqObj)
         .          .     28:
         .          .     29:		var buf strings.Builder
         .          .     30:		for _, e := range reqObj.PayLoad {
         .   100.52MB     31:			buf.WriteString(strconv.Itoa(e))
         .   178.57MB     32:			buf.WriteString(",")
         .          .     33:		}
         .          .     34:		repObj := &Response{reqObj.TransactionID, buf.String()}
         .   272.57MB     35:		repJson, err := repObj.MarshalJSON()
         .          .     36:		//repJson, err := json.Marshal(&repObj)
         .          .     37:		if err != nil {
         .          .     38:			panic(err)
         .          .     39:		}
   92.53MB    92.53MB     40:		reps = append(reps, string(repJson))
         .          .     41:	}
         .          .     42:	return reps
         .          .     43:}
         .          .     44:
         .          .     45:func processRequestOld(reqs []string) []string {
ROUTINE ======================== geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization.processRequestOld in /bigdata/git-projects/programming-projects/geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/optmization/optmization.go
    1.52GB     1.84GB (flat, cum) 59.32% of Total
         .          .     43:}
         .          .     44:
         .          .     45:func processRequestOld(reqs []string) []string {
         .          .     46:	reps := []string{}
         .          .     47:	for _, req := range reqs {
    4.50MB     4.50MB     48:		reqObj := &Request{}
   30.51MB   219.63MB     49:		json.Unmarshal([]byte(req), reqObj)
         .          .     50:		ret := ""
         .          .     51:		for _, e := range reqObj.PayLoad {
    1.45GB     1.45GB     52:			ret += strconv.Itoa(e) + ","
         .          .     53:		}
    3.50MB     3.50MB     54:		repObj := &Response{reqObj.TransactionID, ret}
         .   136.04MB     55:		repJson, err := json.Marshal(&repObj)
         .          .     56:		if err != nil {
         .          .     57:			panic(err)
         .          .     58:		}
   40.51MB    40.51MB     59:		reps = append(reps, string(repJson))
         .          .     60:	}
         .          .     61:	return reps
         .          .     62:}
```
