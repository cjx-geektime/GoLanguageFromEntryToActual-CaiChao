# 构建可复用的模块
## 引用本地包
需要模块的文件在 GOPATH 指定的目录下，引用时从目录的 src 的下一级开始写

1. 以首字母大写表明可以被包外代码访问
2. 代码的 package 可以和所在目录不一致
3. 同一目录里的 Go 代码的 package 要保持一致
## 引用远程包
1. 通过 go get 来获取远程依赖
•go get -u 强制从网网络更更新远程依赖
2. 注意代码在 GitHub 上的组织形式,以适应 go get
•直接以代码路路径开始,不要有 src

引用远程包：go get -u github.com/easierway/concurrent_map

## 依赖管理
```
✗ sudo apt install golang-glide
✗ glide init
安装依赖
$ glide install
```
