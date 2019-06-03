package main

import (
    "encoding/json"
    "fmt"
)

type Student struct {
    Name string
    Sex string
    Id string
}

type Class struct {
    Name string
    Count int
    Students []*Student
}

var rawJson = `
{
    "Name": "101",
    "Count": 0,
    "Students": [
        {
            "Name": "stu=0",
            "Sex": "man",
            "Id": "0"
        },
        {
            "Name": "stu=1",
            "Sex": "man",
            "Id": "1"
        },
        {
            "Name": "stu=2",
            "Sex": "man",
            "Id": "2"
        },
        {
            "Name": "stu=3",
            "Sex": "man",
            "Id": "3"
        },
        {
            "Name": "stu=4",
            "Sex": "man",
            "Id": "4"
        },
        {
            "Name": "stu=5",
            "Sex": "man",
            "Id": "5"
        },
        {
            "Name": "stu=6",
            "Sex": "man",
            "Id": "6"
        },
        {
            "Name": "stu=7",
            "Sex": "man",
            "Id": "7"
        },
        {
            "Name": "stu=8",
            "Sex": "man",
            "Id": "8"
        },
        {
            "Name": "stu=9",
            "Sex": "man",
            "Id": "9"
        }
    ]
}
`

func main() {
    //testStructMarshal()
    testUnstructMarshal()
}

/**
 * 结构体反序列化
 */
func testUnstructMarshal()  {
    fmt.Println("\n\nunmarshal result is \n\n")

    var c1 *Class = &Class{}
    err := json.Unmarshal([]byte(rawJson), c1)
    if err != nil {
        fmt.Println("unmarshal failed")
        return
    }
    fmt.Printf("c1:%#v\n", c1)

    for _, v := range c1.Students {
        fmt.Printf("stu:%#v\n", v)
    }
}

/**
 * 结构体序列化
 */
func testStructMarshal() {
    c := &Class {
        Name: "101",
        Count: 0,
    }

    for i:= 0; i<10; i++ {
        stu := &Student{
            Name: fmt.Sprintf("stu=%d", i),
            Sex: "man",
            Id: fmt.Sprintf("%d", i),
        }

        c.Students = append(c.Students, stu)
    }

    data, err := json.Marshal(c)
    if err != nil {
        fmt.Println("json marshal failed")
        return
    }

    fmt.Printf("json:%s\n", string(data))
}