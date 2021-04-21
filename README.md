## 第一节
#### 随机数

`rand.Intn () `函数是个伪随机函数，不管运行多少次都只会返回同样的随机数，因为它默认的资源就是单一值，所以必须调用 rand.Seed (), 并且传入一个变化的值作为参数，如 `time.Now().UnixNano() `, 就是可以生成时刻变化的值.

```go
package main

import ("fmt"
        "math/rand"
        "time")

func main() {
    // 初始化随机数的资源库, 如果不执行这行, 不管运行多少次都返回同样的值
    rand.Seed(time.Now().UnixNano())
    fmt.Println("A number from 1-100", rand.Intn(81))
}
```

#### gorm新版

##### 安装

```go
go get gorm.io/gorm
// **注意** GORM `v2.0.0` 发布的 git tag 是 `v1.20.0`
```

##### 使用示例

```go
import (
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
)

func init() {
  	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database, err: " + err.Error())
	}
}
```

