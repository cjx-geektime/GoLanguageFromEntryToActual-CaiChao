# 性能分析⼯具
## 准备⼯作
### 安装 graphviz
```
✗ sudo apt-get install graphviz
```
### 安装  go-torch
```
✗ go get -v github.com/uber/go-torch
```
下载并复制 flamegraph.pl （https://github.com/brendangregg/FlameGraph）⾄ $GOPATH/bin 路径下

将 $GOPATH/bin 加⼊ $PATH
#### 下载或复制 flamegraph.pl 到当前目录下
赋予执行权限
```
➜  file git:(main) ✗ chmod +x flamegraph.pl
```
# 性能分析
```
➜  chapterVIIPerformanceTuning git:(main) ✗ cd tools/file
➜  file git:(main) ✗ ll
总用量 4.0K
-rw-rw-r-- 1 cjx cjx 1.5K 10月 11 21:32 prof.go
➜  file git:(main) ✗ ls
prof.go
➜  file git:(main) ✗ go build prof.go
➜  file git:(main) ✗ ls
prof  prof.go
➜  file git:(main) ✗ ./prof
➜  file git:(main) ✗ ls
cpu.prof  goroutine.prof  mem.prof  prof  prof.go
➜  file git:(main) ✗ go build prof.go
➜  file git:(main) ✗ ./prof
➜  file git:(main) ✗ ls
cpu.prof  goroutine.prof  mem.prof  prof  prof.go
```
## pprof
分析 CPU 性能
```
➜  file git:(main) ✗ go tool pprof prof cpu.prof
File: prof
Type: cpu
Time: Oct 11, 2021 at 9:33pm (CST)
Duration: 903.75ms, Total samples = 730ms (80.77%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 730ms, 100% of 730ms total
      flat  flat%   sum%        cum   cum%
     320ms 43.84% 43.84%      520ms 71.23%  math/rand.(*Rand).Int31n
     130ms 17.81% 61.64%      650ms 89.04%  math/rand.(*Rand).Intn
      90ms 12.33% 73.97%       90ms 12.33%  math/rand.(*rngSource).Uint64 (inline)
      80ms 10.96% 84.93%      170ms 23.29%  math/rand.(*rngSource).Int63
      60ms  8.22% 93.15%      710ms 97.26%  main.fillMatrix
      20ms  2.74% 95.89%       20ms  2.74%  main.calculate (inline)
      20ms  2.74% 98.63%      190ms 26.03%  math/rand.(*Rand).Int63 (inline)
      10ms  1.37%   100%      200ms 27.40%  math/rand.(*Rand).Int31
         0     0%   100%      730ms   100%  main.main
         0     0%   100%      730ms   100%  runtime.main
```
- cum：消耗时间
- cum%：时间占比

分析具体操作的性能
```
(pprof) list fillMatrix
Total: 730ms
ROUTINE ======================== main.fillMatrix in /bigdata/git-projects/programming-projects/geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/tools/file/prof.go
      60ms      710ms (flat, cum) 97.26% of Total
         .          .     16:
         .          .     17:func fillMatrix(m *[row][col]int) {
         .          .     18:	s := rand.New(rand.NewSource(time.Now().UnixNano()))
         .          .     19:
         .          .     20:	for i := 0; i < row; i++ {
      10ms       10ms     21:		for j := 0; j < col; j++ {
      50ms      700ms     22:			m[i][j] = s.Intn(100000)
         .          .     23:		}
         .          .     24:	}
         .          .     25:}
         .          .     26:
         .          .     27:func calculate(m *[row][col]int) {
```
发现赋值语句：m[i][j] = s.Intn(100000) 占用时间最长

使用图形化分析
```
(pprof) svg
Generating report in profile001.svg（生成的图，可以用浏览器打开）
(pprof) exit
```
块越大，说明该操作的CPU占用时间越长
## go-torch
分析 CPU 性能
```
➜  file git:(main) ✗ /home/cjx/go/bin/go-torch cpu.prof

INFO[22:00:45] Run pprof command: go tool pprof -raw -seconds 30 cpu.prof
INFO[22:00:45] Writing svg to torch.svg
```
用浏览器打开 torch.svg ，点击可以查看具体的执行分析

# 内存分析(不使用 GC的情况)
## pprof
```
➜  file git:(main) ✗ go build prof.go
➜  file git:(main) ✗ ./prof
➜  file git:(main) ✗ go tool pprof mem.prof
File: prof
Type: inuse_space
Time: Oct 11, 2021 at 10:14pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 762.95MB, 99.45% of 767.17MB total
Dropped 16 nodes (cum <= 3.84MB)
      flat  flat%   sum%        cum   cum%
  762.95MB 99.45% 99.45%   764.67MB 99.67%  main.main
         0     0% 99.45%   764.67MB 99.67%  runtime.main
(pprof) list main.main
Total: 767.17MB
ROUTINE ======================== main.main in /bigdata/git-projects/programming-projects/geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/tools/file/prof.go
  762.95MB   764.67MB (flat, cum) 99.67% of Total
         .          .     38:	if err != nil {
         .          .     39:		log.Fatal("could not create CPU profile: ", err)
         .          .     40:	}
         .          .     41:
         .          .     42:	// 获取系统信息
         .     1.72MB     43:	if err := pprof.StartCPUProfile(f); err != nil { //监控cpu
         .          .     44:		log.Fatal("could not start CPU profile: ", err)
         .          .     45:	}
         .          .     46:	defer pprof.StopCPUProfile()
         .          .     47:
         .          .     48:	// 主逻辑区，进行一些简单的代码运算
  762.95MB   762.95MB     49:	x := [row][col]int{}
         .          .     50:	fillMatrix(&x)
         .          .     51:	calculate(&x)
         .          .     52:
         .          .     53:	f1, err := os.Create("mem.prof")
         .          .     54:	if err != nil {

```
# 内存分析(使用 GC的情况)
## pprof
```
➜  file git:(main) ✗ go tool pprof prof mem.prof
File: prof
Type: inuse_space
Time: Oct 11, 2021 at 9:33pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 3233.98kB, 100% of 3233.98kB total
Showing top 10 nodes out of 18
      flat  flat%   sum%        cum   cum%
 1537.50kB 47.54% 47.54%  1537.50kB 47.54%  runtime.allocm
 1184.27kB 36.62% 84.16%  1184.27kB 36.62%  runtime/pprof.StartCPUProfile
  512.20kB 15.84%   100%   512.20kB 15.84%  runtime.malg
         0     0%   100%  1184.27kB 36.62%  main.main
         0     0%   100%  1184.27kB 36.62%  runtime.main
         0     0%   100%     1025kB 31.69%  runtime.mcall
         0     0%   100%   512.50kB 15.85%  runtime.mstart
         0     0%   100%   512.50kB 15.85%  runtime.mstart0
         0     0%   100%   512.50kB 15.85%  runtime.mstart1
         0     0%   100%  1537.50kB 47.54%  runtime.newm
(pprof) list main.main
Total: 3.16MB
ROUTINE ======================== main.main in /bigdata/git-projects/programming-projects/geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIPerformanceTuning/tools/file/prof.go
         0     1.16MB (flat, cum) 36.62% of Total
         .          .     39:	if err != nil {
         .          .     40:		log.Fatal("could not create CPU profile: ", err)
         .          .     41:	}
         .          .     42:
         .          .     43:	// 获取系统信息
         .     1.16MB     44:	if err := pprof.StartCPUProfile(f); err != nil { //监控cpu
         .          .     45:		log.Fatal("could not start CPU profile: ", err)
         .          .     46:	}
         .          .     47:	defer pprof.StopCPUProfile()
         .          .     48:
         .          .     49:	// 主逻辑区，进行一些简单的代码运算

```
