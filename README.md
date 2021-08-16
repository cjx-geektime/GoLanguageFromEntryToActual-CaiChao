# geek-GoLanguageFromEntryToActualCombat
# 安装 GO
## 下载
https://golang.org/doc/install<br>
https://golang.google.cn/dl
## 安装 Go 语言
### 解压
用 root 用户操作
```
bigdata# tar -C /usr/local -xzf go1.16.7.linux-amd64.tar.gz
```
### 添加环境变量
```
vim /etc/profile
export GO_HOME=/usr/local/go

export PATH=$PATH:$GO_HOME/bin

bigdata# source /etc/profile
```
### 测试
```
# go version
go version go1.16.7 linux/amd64
```
## 安装 IDE
Atom: https://atom.io + Package: go-plus
