# 构建可复用的模块
需要模块的文件在 GOPATH 指定的目录下，引用时从目录的 src 的下一级开始写

1. 以首字母大写表明可以被包外代码访问
2. 代码的 package 可以和所在目录不一致
3. 同一目录里的 Go 代码的 package 要保持一致

引用远程包：go get -u github.com/programing-language/geek-GoLanguageFromEntryToActualCombat
