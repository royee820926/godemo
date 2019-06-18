package main

import (
    "demo/dataform_protobuf/address"
    "fmt"
    "github.com/golang/protobuf/proto"
    "io/ioutil"
)

func main() {
    var filename = "d:/logs/contactbook.dat"

    // 序列化
    //err := testMarshal(filename)
    //if err != nil {
    //    fmt.Printf("write proto failed, err:%v\n", err)
    //}
    
    // 反序列化
    err := testUnmarshal(filename)
    if err != nil {
       fmt.Printf("read proto failed, err:%v\n", err)
    }
    
}

func testUnmarshal(filename string) (err error) {
    var contactBook address.ContactBook
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return
    }
    err = proto.Unmarshal(data, &contactBook)
    if err != nil {
        return
    }

    fmt.Printf("proto:%#v\n", contactBook)
    return
}

func testMarshal(filename string) (err error) {
    var contactBook address.ContactBook
    for i := 0; i < 64; i++ {
        p := &address.Person{
            Id: int32(i),
            Name: fmt.Sprintf("陈%d", i),
        }

        phone := &address.Phone{
            Type: address.PhoneType_HOME,
            Number: "15910624165",
        }

        p.Phones = append(p.Phones, phone)
        contactBook.Persons = append(contactBook.Persons, p)
    }

    // 序列化
    data, err := proto.Marshal(&contactBook)
    if err != nil {
        fmt.Printf("marshal proto buf failed, err:%v\n", err)
        return
    }

    err = ioutil.WriteFile(filename, data, 0755)
    if err != nil {
        fmt.Printf("marshal proto buf failed, err:%v\n", err)
        return
    }
    return
}