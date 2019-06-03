package main

import "fmt"

type Animal interface {
	Talk()
	Eat()
	Name() string
}

type Dog struct {
	
}

func (d Dog) Talk() {
	fmt.Println("wang wang wang")
}

func (d Dog) Eat() {
	fmt.Println("I am eating bone")
}

func (d Dog) Name() string {
	fmt.Println("My name is wang cai")
	return "wang cai"
}

type Pig struct {

}

func (d Pig) Talk() {
	fmt.Println("ken ken ken")
}

func (d Pig) Eat() {
	fmt.Println("I am eating straw")
}

func (d Pig) Name() string {
	fmt.Println("My name is zhu ba jie")
	return "wang cai"
}

func main() {
	var d Dog
	var a Animal
	a = d

	a.Eat()
	a.Talk()
	a.Name()

	var pig Pig
	a = pig
	a.Eat()
	a.Talk()
	a.Name()
}
