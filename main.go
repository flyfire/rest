package main

import (
    "fmt"
    "log"
    "time"

    . "server/rest"
)

var (
    RuleName = &StringRule{"name", 4, 20}
    RuleAge  = &IntRule{"age", 10, 60}
)

func test(params Value) Value {
    fmt.Println("in:", params)

    // 使用规则验证输入参数。
    if !RuleCheck(params, RuleName, RuleAge) {
        return NewValue().Error("invalid parameters")
    }

    // 返回数据。
    params["__time__"] = time.Now()
    return params
}

func main() {
    s := NewServer(":8080")

    s.StaticDir = "."  // 静态文件目录。
    s.Pattern = "/v2/" // REST 路径前缀匹配模式。

    s.HandleFunc("/test/{name}/{age}", GET, test) // http://localhost:8080/v2/test/user1/23

    if err := s.ListenAndServe(); err != nil {
        log.Println(err)
    }
}
