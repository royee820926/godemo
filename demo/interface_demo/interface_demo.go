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

func main() {
	var d Dog
	var a Animal
	a = d

	a.Eat()
	a.Talk()
	a.Name()
}

