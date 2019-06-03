package main

import (
	"fmt"
	"time"

	//"time"
)

func main() {
	//var a int8 = 127
	//var b int  = 128
	//a = int8(b)
	//fmt.Printf("hello world : %d\n", a)
	//
	//var result = common.MyAdd(4, 66)
	//fmt.Printf("result: %d\n", result)
	//
	//var ru rune = '好'
	//fmt.Printf("ru = %c\n", ru)
	//
	//var str = "你好123。"
	//var runeSlice []rune = []rune(str)
	//fmt.Printf("str length: %d, len(str):%d", len(runeSlice), len(str))
    //fmt.Printf("time: %v", time.Now())

    //var now = time.Now()
    //var year = now.Year()
    //var month = now.Month()
    //var day  = now.Day()
    //var hour = now.Hour()
    //var minute = now.Minute()
    //var second = now.Second()
    //fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

    //var timeObj = time.Unix(1558943263, 0)
    //fmt.Printf("The time of timeObj is %d\n", timeObj.Year())

    //testTicker()

    //now = time.Now()
    //timeStr := now.Format("2006-01-02 15:04:05")
    //fmt.Printf("time is %s\n", timeStr)

	//testSwitch()

	testType()
}

func testType() {
	a := func() {
		fmt.Printf("abc hhh")
	}
	a()
}

func testSwitch(){
	var i int = 10
	switch i {
	case 10:
		fmt.Printf("i: ab\n")
		fallthrough
	case 11:
		fmt.Printf("i: ac\n")
		fallthrough
	default:
		fmt.Printf("default\n")
	}
}

func testTicker(){
    ticker := time.Tick(1 * time.Second)
	for i := range ticker {
		fmt.Printf("%v\n", i)
		processTask()
	}
}

func processTask(){
	fmt.Printf("do task\n")
}
