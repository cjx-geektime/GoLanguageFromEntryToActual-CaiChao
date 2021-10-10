⾯向资源的架构（Resource Oriented Architecture）

参考: <RESTful Web Services>

# 运行
```
➜  roa git:(main) ✗ go run resource_oriented_arc.go

➜  ~ curl localhost:8080          
Welcome!
➜  ~ curl localhost:8080/employee/Mike
{"id":"e-1","name":"Mike","age":35}%                                                                                                                                                                               
➜  ~ curl localhost:8080/employee/Rose
{"id":"e-2","name":"Rose","age":45}%                                                                                                                                                                               
➜  ~ curl localhost:8080/employee/cjx
{"error":"Not Found"}%    
```
