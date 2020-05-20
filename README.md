# cuten
HTTP web framework in go

## 快速上手

获取最新脚手架
[main.go](./test/main.go)

```go
package main

import (
	"fmt"
	"net/http"
	"github.com/ranxx/cuten"
)
func main() {
    e := cuten.New()
    e.GET("/", func(ctx *cuten.Context) {
        ctx.String(http.StatusOK, "Welcome use cuten")
    })
    fmt.Println(e.Run(":9999"))
}
```

### 中间件的使用 方式1

```go
package main

import (
	"fmt"
	"net/http"
	"github.com/ranxx/cuten"
)
func main() {
    e := cuten.New()
    e.Use(
        func(ctx *cuten.Context) {
		    fmt.Printf("我是%s路由中间件_1\n", ctx.Path)
        },
        func(ctx *cuten.Context) {
		    fmt.Printf("我是%s路由中间件_2\n", ctx.Path)
        })
    e.GET("/", func(ctx *cuten.Context) {
        ctx.String(http.StatusOK, "Welcome use cuten")
    })
    fmt.Println(e.Run(":9999"))
}
```


### 中间件的使用 方式2

```go
package main

import (
	"fmt"
	"net/http"
	"github.com/ranxx/cuten"
)
func main() {
    e := cuten.New()
    e.GET("/", 
        func(ctx *cuten.Context) {
            fmt.Printf("我是%s路由中间件_1\n", ctx.Path)
        },
        func(ctx *cuten.Context) {
            ctx.String(http.StatusOK, "Welcome use cuten")
        })
    fmt.Println(e.Run(":9999"))
}
```