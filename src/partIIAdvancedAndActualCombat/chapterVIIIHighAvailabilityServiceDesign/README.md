# ⾼可⽤性架构设计
## ⾯向错误的设计
" Once you accept that failures will happen, you have the ability to design your system’s reaction to the failures."
### 隔离
#### 隔离错误 — 设计
#### 隔离错误 — 部署
#### 重⽤ vs 隔离
逻辑结构的重⽤ vs 部署结构的隔离
### 冗余
#### 冗余
#### 单点失效
#### 慢响应
#### 不要⽆休⽌的等待
给阻塞操作都加上⼀个期限
#### 错误传递
##### 断路器
## ⾯向恢复的设计
"A priori prediction of all failure modes is not possible."
### 健康检查
• 注意僵⼫进程
  • 池化资源耗尽
  • 死锁
### 构建可恢复的系统
• 拒绝单体系统
• ⾯向错误和恢复的设计
  • 在依赖服务不可⽤时，可以继续存活
  • 快速启动
  • ⽆状态
### 与客户端协商
## Chaos Engineering
"If something hurts, do it more often!"
• 如果问题经常发⽣⼈们就会学习和思考解决它的⽅法
### Chaos Engineering 原则
• Build a Hypothesis around Steady State Behavior
• Vary Real-world Events
• Run Experiments in Production
• Automate Experiments to Run Continuously
• Minimize Blast Radius

http://principlesofchaos.org
#### 相关开源项⽬
- https://github.com/Netflix/chaosmonkey ()
- https://github.com/easierway/service_decorators/blob/master/README.md (用GO开发分布式服务做微服务架构)
