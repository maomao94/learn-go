## GO 语言的依赖管理
* 依赖的概念
* 依赖管理的三个阶段 GOPATH，GOVENDOR，go mod
    * 每个项目有一个自己的vendor目录，存放第三方库
    * 大量第三方依赖管理工具：glide，dep，go dep，...
    * 初始化：go mode init
    * 增加依赖：go get
    * 更新依赖：go get[@v...],go mode tidy
    * 项目迁移到go mode：go mode init，go build ./...
## 接口
* duck typing 
    * 描述事物的外部行为而非内部结构
    * go属于结构化类型系统，类似 duck typing