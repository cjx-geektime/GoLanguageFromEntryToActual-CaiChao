# 运行
```
➜  geek-GoLanguageFromEntryToActualCombat git:(main) ✗ cd src/ch1/main
➜  main git:(main) ✗ ll
总用量 4.0K
-rw-rw-r-- 1 cjx cjx 72 8月  18 11:59 hello_world.go
➜  main git:(main) ✗ go run hello_world.go
Hello World
```
# 编译运行
```
➜  main git:(main) ✗ go build hello_world.go
➜  main git:(main) ✗ ll
总用量 1.9M
-rwxrwxr-x 1 cjx cjx 1.9M 8月  18 12:01 hello_world
-rw-rw-r-- 1 cjx cjx   72 8月  18 11:59 hello_world.go
-rw-rw-r-- 1 cjx cjx  263 8月  18 12:00 README.md
➜  main git:(main) ✗ ./hello_world
Hello World
```
# 注意（应用程序入口）
1. 必须是 main 包: package main(不一定要放在 main 包下)
2. 必须是 main 方法: func main()
3. 文件名不一定是 main.go
