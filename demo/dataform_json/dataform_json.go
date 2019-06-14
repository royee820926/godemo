package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "math/rand"
)

type Person struct {
    Name string
    Age int
    Sex string
}

var (
    filename = "d:/logs/person.json"
)

func main() {
    err := testWriteJson(filename)
    if err != nil {
        fmt.Printf("write json failed, err:%v\n", err)
        return
    }

    err = testReadJson(filename)
    if err != nil {
        fmt.Printf("read json failed, err:%v\n", err)
        return
    }
}

func testReadJson(filename string) (err error) {
    var persons []*Person
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return
    }

    err = json.Unmarshal(data, &persons)
    if err != nil {
        return
    }

    for _, v := range persons {
        fmt.Printf("%#v\n", v)
    }
    return
}

func testWriteJson(filename string) (err error) {
    var persons []*Person
    for i := 0; i < 10; i++ {
        p := &Person{
            Name: fmt.Sprintf("name%d", i),
            Age: rand.Intn(100),
            Sex: "Man",
        }

        persons = append(persons, p)
    }
    data, err := json.Marshal(persons)
    if err != nil {
        fmt.Printf("=marshal failed, err:%v\n", err)
        return
    }

    err = ioutil.WriteFile(filename, data, 0755)
    if err != nil {
        fmt.Printf("write file failed, err:%v\n", err)
        return
    }
    return
}