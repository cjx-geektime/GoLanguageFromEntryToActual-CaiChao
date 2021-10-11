分析运行中的项目的性能
# 通过 HTTP ⽅式输出 Profile
- 简单，适合于持续性运⾏的应⽤
- 在应⽤程序中导⼊ import _ "net/http/pprof"，并启动 http server 即可
- http://<host>:<port>/debug/pprof/
- go tool pprof http://<host>:<port>/debug/pprof/profile?seconds=10 （默认值为30秒）
- go-torch -seconds 10 http://<host>:<port>/debug/pprof/profile
# 分析
## 运行
```
➜  http git:(main) ✗ ll
总用量 8.0K
-rw-rw-r-- 1 cjx cjx 636 10月 11 22:19 fb_server.go
-rw-rw-r-- 1 cjx cjx 390 10月 11 22:19 README.md
➜  http git:(main) ✗ go run fb_server.go

➜  http git:(main) ✗ curl localhost:8081/fb
[1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181 6765 10946 17711 28657 46368 75025 121393 196418 317811 514229 832040 1346269 2178309 3524578 5702887 9227465 14930352 24157817 39088169 63245986 102334155 165580141 267914296 433494437 701408733 1134903170 1836311903 2971215073 4807526976 7778742049 12586269025]
```
浏览器打开：http://localhost:8081/debug/pprof/
## 使用命令行分析性能
默认是 30s 的采样时间
```
➜  http git:(main) ✗ go tool pprof http://127.0.0.1:8081/debug/pprof/profile
Fetching profile over HTTP from http://127.0.0.1:8081/debug/pprof/profile
Saved profile in /home/cjx/pprof/pprof.fb_server.samples.cpu.001.pb.gz
File: fb_server
Type: cpu
Time: Oct 11, 2021 at 10:25pm (CST)
Duration: 30.03s, Total samples = 0
No samples were found with the default sample value type.
Try "sample_index" command to analyze different sample values.
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top

```
