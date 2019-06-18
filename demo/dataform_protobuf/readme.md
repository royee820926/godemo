github地址：
=============
https://github.com/protocolbuffers/protobuf/releases

使用步骤：
=============
1. 下载protobuf，将解压包中的 protoc 放在$GOPATH/bin目录中

2. go get -u github.com/golang/protobuf/protoc-gen-go

3. 编写protobuf数据结构，如下所示：
```cassandraql
// person.proto
// 指定版本
// 注意proto3与proto2的写法有些不同
syntax = "proto3";

// 包名，通过protoc生成时go文件时
package address;

// 手机类型
// 枚举类型第一个字段必须为0
enum PhoneType {
    HOME = 0;
    WORK = 1;
}

// 手机
message Phone {
    PhoneType type = 1;
    string number = 2;
}

// 人
message Person {
    // 后面的数字表示标识号
    int32 id = 1;
    string name = 2;
    // repeated表示可重复
    // 可以有多个手机
    repeated Phone phones = 3;
}

// 联系簿
message ContactBook {
    repeated Person persons = 1;
}
```

4.生成protobuf的go语言代码
```cassandraql
protoc --go_out=./address/ .\person.proto
```