# 编写测试程序
1. 源码文件以 _test 结尾: xxx_test.go
2. 测试方法名以 Test 开头: func TestXXX(t *testing.T) {...}

# 测试
```
➜  src git:(main) ✗ go test -v partIGoLanguageFoundation/chapterIBasicProgramStructure/test/first_test.go
=== RUN   TestFirstTry
    first_test.go:6: My first try!
--- PASS: TestFirstTry (0.00s)
PASS
ok  	command-line-arguments	0.001s
```
# 保存自动执行测试
Packages -> Setting View -> Manage Packages -> go-plus -> 勾选Test.Run Tests With Verbose Flag Set(重启Atom)
```
➜  src git:(main) ✗ go mod init partIGoLanguageFoundation
go: creating new go.mod: module partIGoLanguageFoundation
go: to add module requirements and sums:
	go mod tidy
```
## 保存文件
```
=== RUN   TestFirstTry
    first_test.go:6: My first try!
--- PASS: TestFirstTry (0.00s)
PASS
coverage: [no statements]
ok  	partIGoLanguageFoundation/partIGoLanguageFoundation/chapterIBasicProgramStructure/test	0.001s

Done
```
