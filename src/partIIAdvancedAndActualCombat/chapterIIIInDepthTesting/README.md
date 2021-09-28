# 单元测试
## 内置单元测试框架
- Fail, Error: 该测试失败,该测试继续,其他测试继续执行行行
- FailNow, Fatal: 该测试失败,该测试中止止,其他测试继续执行行行

代码覆盖率
```
go test -v - cover
```
断言
```
https://github.com/stretchr/testify
go get -u github.com/stretchr/testify
```
