EasyJSON 采用代码生成，而非反射

# 安装
```
go get -u github.com/mailru/easyjson/...
```
# 使用
## 结构定义
struct_def.go
## 编译
```
➜  easyjson git:(main) ✗ ~/go/bin/easyjson -all struct_def.go
```
在当前目录下生成一个文件: struct_def_easyjson.go
# 测试
```
➜  easyjson git:(main) ✗ go test -bench=.
{{Mike 30} {[Java Go C]}}
{"basic_info":{"name":"Mike","age":30},"job_info":{"skills":["Java","Go","C"]}}
{{Mike 30} {[Java Go C]}}
{"basic_info":{"name":"Mike","age":30},"job_info":{"skills":["Java","Go","C"]}}
goos: linux
goarch: amd64
pkg: geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVICommonTasks/easyjson
cpu: Intel(R) Core(TM) i9-10900K CPU @ 3.70GHz
BenchmarkEmbeddedJson-20          647100              1783 ns/op
BenchmarkEasyJson-20             2034846               612.3 ns/op
PASS
ok      geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVICommonTasks/easyjson  3.019s
```
easyjson  的效率差不多是内置 json 的三倍
