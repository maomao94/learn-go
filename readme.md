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
    * 接口变量 包含 实现者的类型 实现者的值（指针）自带指针
    * 接口变量同样采用值传递，几乎不需要使用接口的指针
    * 指针接收者实现只能以指针方式使用，值接受者都可
* 查看接口变量
    * 标识任何类型：interface{}
    * Type Assertion
    * Type Switch