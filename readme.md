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
* 特殊接口
    * Stringer
    * Reader/Writer
  

## 函数与闭包
#### 函数式编程VS函数指针
* 函数是一等公民：参数，变量，返回值都可以是函数
* 高阶函数
* 函数 -> 闭包
#### "正统"函数式编程
* 不可变性：不能有状态，只有常量和函数
* 函数只能有一个参数
* go语言闭包的应用
  * 更为自然，不需要修饰如何访问自由变量
  * 没有lambda表达式，但是有匿名函数
  

## 资源管理和出错处理
#### defer 调用
* 确保调用在函数结束时发生
* 参数在defer语句时计算
* defer列表为后进先出
* 何时使用defer 调用
  * Open/Close
  * Lock/Unlock
  * PrintHeader/PrintFooter
#### 错误处理
#### panic
* 停止当前函数执行
* 一直向上返回，执行每一层的defer
* 如果没有遇见recover，程序就退出
#### recover
* 仅在defer调用中使用
* 获取panic的值
* 如果无法处理，可重新panic
#### error vs panic
* 意料之中：使用error。如文件打不开
* 意料之外的：使用panic。如：数组越界

## goroutine
#### 协程Coroutine
* 轻量级"线程"
* 非抢占式多任务处理，有协程主动交出控制权（1.4后优化成可以抢占式）
* 编译器/解释器/虚拟机层面的多任务
* 多个协程可能在一个或多个线程上运行
#### go语言的调度器
##### goroutine可能的切换点
    * I/O,select
    * channel
    * 等待锁
    * 函数调用（有时）
    * runtime.Gosched()
##### channel
* channel
* buffered channel
* range


  
  