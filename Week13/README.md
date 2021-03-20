### 前言
对于目前团长项目的一些想法,团长项目目前为一个巨石项目,牵扯的业务范围涉及 用户关系、积分体系、团长管理


### 理想架构


![](https://huijianwang.oss-cn-shanghai.aliyuncs.com/customer.png)

#### BFF层: 
关注业务逻辑, 如: 调用account服务进行用户角色校验、将团长信息聚合等逻辑

#### customer-service: 
专注与团长的基本API的实现, 关注系统的稳定性、可用性, 将复杂的业务逻辑剥离出去

#### customer-job: 
利用消息队列进行削峰, 同时进行一些积分类的业务逻辑异步处理

#### customer-admin:
后台管理项目, 目前是共享同一个数据库

#### account-service:
用户服务, 提供用户校验、用户信息查询等服务,再BFF层进行收敛

### 项目目录
kratos 的项目目录

项目内部实现风格能够大大的提高生产效率 大仓的优势在于能够推得动基础组件库和方便proto文件的管理
```
.
├── LICENSE
├── README.md
├── api             // 维护proto文件与派生文件
│   └── helloworld
│       ├── errors
│       │   ├── helloworld.pb.go
│       │   ├── helloworld.proto
│       │   └── helloworld_errors.pb.go
│       └── v1
│           ├── greeter.pb.go
│           ├── greeter.proto
│           ├── greeter_grpc.pb.go
│           └── greeter_http.pb.go
├── cmd  // 项目入口
│   └── demo  
│       ├── main.go
│       ├── wire.go      // wire注入依赖以及wire的派生文件
│       └── wire_gen.go
├── configs // 配置文件
│   └── config.yaml
├── generate.go
├── go.mod
├── go.sum
└── internal // internal 避免被错误引用
    ├── biz // 业务逻辑
    │   ├── README.md
    │   ├── biz.go
    │   └── greeter.go
    ├── conf // 内部使用的结构定义
    │   ├── conf.pb.go
    │   └── conf.proto
    ├── data // 业务数据访问
    │   ├── README.md
    │   ├── data.go
    │   └── greeter.go
    ├── server // grpc&&http的实例
    │   ├── grpc.go
    │   ├── http.go
    │   └── server.go
    └── service api定义的服务层
        ├── README.md
        ├── greeter.go
        └── service.go
```

#### error 处理
> github.com/pkg/errors 包实现对上下文的信息传递, 仅在API或吞掉error的方法中输出error日志

```
func UpdateModel() error {
	err := sql.ErrNoRows
	return errors.Wrap(err, "raise a sql err")
}
```

#### 错误码的规范
实现一个通用的 errcode来维护, 当错误码重复时,直接抛出panic, 在编译阶段即可发现问题所在,同时也能够使得错误码和错误信息能够匹配
```
type Error struct {
	// 错误码
	code int `json:"code"`
	// 错误消息
	msg string `json:"msg"`
	// 详细信息
	details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息:：%s", e.Code(), e.Msg())
}

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服务内部错误")
	InvalidParams             = NewError(10000001, "入参错误")
	NotFound                  = NewError(10000002, "找不到")
)

```

#### 并发请求
> sync.WaitGroup

协程比较重要的知道其生命周期, 通过waitgroup 来保证goroutine的关闭

```
import (
    "fmt"
    "sync"
)

func worker(x int, wg *sync.WaitGroup) {
    defer wg.Done()
    do something
}

func main() {
    var wg sync.WaitGroup

    wg.Add(2)
    go worker(1, &wg)
    go worker(2, &wg)

    wg.Wait()
}
```

总结

1. 每一个项目都有其诞生的目的, 确定这个目的坚持的走下去. 不要为了做而做
2. 单体服务还是微服务都各有优劣,不盲目的追求选择, 适合当下的才是最好的
3. 毛大的训练营不仅仅是Go这种语言的,给我收获更大的是架构方面的知识, 对于微服务有了整体上的理解,微服务的各个中间件也都有了概念
4. 收获了很多职场技巧和信心
5. 收获了一波毛大表情包

最后感谢毛大的倾囊相授,今天又是吸毛大的一天