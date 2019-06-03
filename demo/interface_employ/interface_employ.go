package main

import "fmt"

type Employer interface {
	CalcSalary() float32
}

type Programmer struct {
	name string
	base float32
	extra float32
}

func NewProgrammer(name string, base, extra float32) Programmer {
	return Programmer{
		name: name,
		base: base,
		extra: extra,
	}
}

func (p Programmer) CalcSalary() float32 {
	return p.base
}

type Sale struct {
	name string
	base float32
	extra float32
}

func NewSale(name string, base, extra float32) Sale {
	return Sale{
		name: name,
		base: base,
		extra: extra,
	}
}

func (s Sale) CalcSalary() float32 {
	return s.base + s.extra * s.base * 0.5
}

func main() {
	p1 := NewProgrammer("搬砖1", 1500.0, 0)
	p2 := NewProgrammer("搬砖2", 1500.0, 0)
	p3 := NewProgrammer("搬砖3", 1500.0, 0)

	s1 := NewSale("销售1", 800.0, 2.5)
	s2 := NewSale("销售2", 800.0, 2.5)
	s3 := NewSale("销售3", 800.0, 2.5)

	var employList []Employer
	employList = append(employList, p1)
	employList = append(employList, p2)
	employList = append(employList, p3)

	employList = append(employList, s1)
	employList = append(employList, s2)
	employList = append(employList, s3)

	cost := calcAll(employList)
	fmt.Printf("人力成本: %f\n", cost)
}

func calcAll(e []Employer) float32 {
	var cost float32
	for _, v := range e {
		cost += v.CalcSalary()
	}
	return cost
}