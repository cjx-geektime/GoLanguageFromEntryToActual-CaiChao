# 退出返回值
- Go 中 main 函数不支持任何返回值
- 通过 os.Exit 来返回状态

# 返回值为 0 为正常，否则打印退出状态
```
➜  ch1 git:(main) ✗ go run exit/hello_world.go
Hello World
➜  ch1 git:(main) ✗ go run exit/hello_world.go
Hello World
exit status 255
```
