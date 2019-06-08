package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x float64 = 3.2
    //v := reflect.ValueOf(x)
    //v.Type()

    //reflectExample(x)
    //reflectValue(x)
    reflectSetValue(&x)
}

func reflectSetValue(a interface{}) {
    v := reflect.ValueOf(a)

    k := v.Kind()
    switch k {
    case reflect.Int64:
        v.SetInt(100)
        fmt.Printf("a is int64, store value is:%d\n", v.Int())
    case reflect.Float64:
        v.SetFloat(6.8)
        fmt.Printf("a is float64, store value is:%f\n", v.Float())
    case reflect.Ptr:
        fmt.Printf("set a to 6.8")
        // 等价于：
        // var b *int = new(int)
        // *b = 100
        v.Elem().SetFloat(6.8)
    default:
        fmt.Printf("default switch\n")
    }
}

func reflectValue(a interface{}) {
    v := reflect.ValueOf(a)
    //t := reflect.TypeOf(a)
    k := v.Kind()
    switch k {
    case reflect.Int64:
        fmt.Printf("a is int64, store value is:%d\n", v.Int())
    case reflect.Float64:
        fmt.Printf("a is float64, store value is:%f\n", v.Float())
    }
}

func reflectExample(a interface{}) {
    t := reflect.TypeOf(a)
    fmt.Printf("type of a is:%v\n", t)

    k := t.Kind()
    switch k {
    case reflect.Int64:
        fmt.Printf("a is int64\n")
    case reflect.String:
        fmt.Printf("a is string\n")

    }
}
