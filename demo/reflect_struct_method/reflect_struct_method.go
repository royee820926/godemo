package main

import (
    "fmt"
    "reflect"
)

type Student struct {
    Name string
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
    //testMethodList()
    testMethodCall()
}

func testMethodCall() {
    var s Student
    v := reflect.ValueOf(&s)
    //t := v.Type()

    // 通过 reflect.Value获取对应方法并调用
    m := v.MethodByName("SetName")
    var args []reflect.Value
    name := "stu01"
    nameVal := reflect.ValueOf(name)
    args = append(args, nameVal)
    m.Call(args)

    m = v.MethodByName("Print")
    var args2 []reflect.Value
    m.Call(args2)
}

func testMethodList()  {
    var s Student
    v := reflect.ValueOf(&s)
    t := v.Type()

    // 注意：
    // 当值类型和指针类型的使用一致时，调用t.NumMethod()方法才能显示正确方法数量。
    // func (s Student) -> reflect.ValueOf(s)
    // func (s *Student) -> reflect.ValueOf(&s)
    fmt.Printf("struct student have %d methods\n", t.NumMethod())
    for i:=0; i<t.NumMethod(); i++ {
        method := t.Method(i)
        fmt.Printf("struct %d method, name:%s type:%v func:%v\n", i, method.Name, method.Type, method.Func)
    }
}
