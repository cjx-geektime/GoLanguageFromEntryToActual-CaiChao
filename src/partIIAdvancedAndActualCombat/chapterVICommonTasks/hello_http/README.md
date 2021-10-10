# 路路由规则
- URL 分为两种,末尾是 /:表示一一个子子树,后面面可以跟其他子子路路径; 末尾不不
是 /,表示一一个叶子子,固定的路路径
  - 以/ 结尾的 URL 可以匹配它的任何子子路路径,比比如 /images 会匹配 /images/
cute-cat.jpg
- 它采用用最⻓长匹配原则,如果有多个匹配,一一定采用用匹配路路径最⻓长的那个进行行行处
理理
- 如果没有找到任何匹配项,会返回 404 错误
# 运行
```
➜  hello_http git:(main) ✗ go run hello_http.go

➜  ~ curl localhost:8080
Hello World!%                                                 
➜  ~ curl localhost:8080/time
{"time":"2021-10-10 22:21:21.942856221 +0800 CST m=+60.918022321"}%                                                                                                                                                
➜  ~ curl localhost:8080/time/1
Hello World!%  
```
## 修改路由
/time -> /time/

curl -i : 查看详细的 Response 头
```
➜  ~ curl -i localhost:8080       
HTTP/1.1 200 OK
Date: Sun, 10 Oct 2021 14:23:33 GMT
Content-Length: 12
Content-Type: text/plain; charset=utf-8

Hello World!%                                                                                                                                                                                                      ➜  ~ curl -i localhost:8080/time
HTTP/1.1 301 Moved Permanently
Content-Type: text/html; charset=utf-8
Location: /time/
Date: Sun, 10 Oct 2021 14:23:38 GMT
Content-Length: 41

<a href="/time/">Moved Permanently</a>.

➜  ~ curl -i localhost:8080/time/
HTTP/1.1 200 OK
Date: Sun, 10 Oct 2021 14:23:45 GMT
Content-Length: 66
Content-Type: text/plain; charset=utf-8

{"time":"2021-10-10 22:23:45.035495975 +0800 CST m=+60.317428248"}%                                                                                                                                                ➜  ~ curl -i localhost:8080/time/1
HTTP/1.1 200 OK
Date: Sun, 10 Oct 2021 14:23:50 GMT
Content-Length: 66
Content-Type: text/plain; charset=utf-8

{"time":"2021-10-10 22:23:50.094617588 +0800 CST m=+65.376549866"}%  
```
