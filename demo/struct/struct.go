package main

import (
	"encoding/json"
	"fmt"
)

//type User struct {
//	name string
//	age int
//	sex string
//}

type User struct {
	UserName 	string `json:"my_username"`
	Sex			string `json:"my_sex"`
	Score		float32
}

func (u *User) set(username string) {
	u.UserName = username
}

func (u *User) printName() {
	fmt.Println(u.UserName)
}

type Integer int

func (i Integer) print() {
	fmt.Println(i)
}

func main() {
	//testCreateUser()
	//testStructTag()
	//testBaseTypeFunction()
	updateMember()
}



func updateMember() {
	var user *User = &User{
		UserName: "people01",
		Sex: "女",
		Score: 99.8,
	}
	user.set("roy")
	user.printName()

}

/**
 * 基础类型添加方法
 */
func testBaseTypeWithFunction() {
	var a Integer
	a = 1000
	a.print()

	var b = 200
	a = Integer(b)
	a.print()
}

/**
 * 结构体tag
 */
func testStructTag() {
	user := &User{
		UserName:"user01",
		Sex:"男",
		Score:99.2,
	}

	data, _ := json.Marshal(user)
	fmt.Printf("json str:%s\n", string(data))
}

/**
 * 实例化结构体
 */
//func testCreateUser()  {
//	var user1 User
//	user1.name = "邓俊1"
//	user1.age = 34
//	user1.sex = "男"
//
//	var user2 = &User{}
//	user2.name = "邓俊2"
//	user2.age = 34
//	user2.sex = "男"
//
//
//	var user3 = new(User)
//	user3.name = "邓俊3"
//	user3.age = 34
//	user3.sex = "男"
//
//	fmt.Printf("%v , sizeof: %d\n", user1, unsafe.Sizeof(user1))
//	fmt.Printf("%v , sizeof: %d\n", user2, unsafe.Sizeof(user2))
//	fmt.Printf("%v , sizeof: %d\n", user3, unsafe.Sizeof(user3))
//}