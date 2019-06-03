package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var s map[string][]int
	s = make(map[string][]int, 16)

	key := "stu01"
	value, ok := s[key]
	if !ok {
		s[key] = make([]int, 0, 16)
		value = s[key]
	}

	value = append(value, 100)
	fmt.Printf("value : %v\n", value)
	fmt.Printf("value : %v\n", s)
}
