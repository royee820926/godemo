package main

/**
 * 反射总结：
 * 在运行时动态获取一个变量的类型信息和值信息
 *
 * 应用场景：
 * A.序列化和反序列化，比如：json，protobuf等各种数据协议
 * B.各种数据库的ORM，比如：gorm，sqlx等数据库的中间件
 * C.配置文件解析相关的库，比如：yaml、ini等
 */

import (
    "fmt"
    "reflect"
)

type Student struct {
    Name string `json:"name"`
    Sex int
    Age int
    Score float32
}

func (s *Student) SetName(name string) {
    s.Name = name
}

func (s *Student) Print() {
    fmt.Printf("反射调用:%#v\n", s)
}

func main() {
    var s Student
    s.SetName("xxx")

    v := reflect.ValueOf(&s)
    t := v.Type()

    field0 := t.Elem().Field(0)
    jsonStr := field0.Tag.Get("json")
    fmt.Printf("tag json=%s\n", jsonStr)
}
