package main

import (
    "fmt"
    "unsafe"
)

const (
    a = "abc"
    b = len(a)
    c = unsafe.Sizeof(a)
)

/**
 * 不论字符串的len有多大，sizeof始终返回16。
 * 实际上字符串类型对应一个结构体，该结构体有两个域，
 * 第一个域是指向该字符串的指针，第二个域是字符串的长度，
 * 每个域占8个字节，但是并不包含指针指向的字符串的内容，
 * 这也就是为什么sizeof始终返回的是16。
 */
func main() {

    var aa = "abc"
    var bb = len(aa)
    var cc = unsafe.Sizeof(aa)

    fmt.Println(a, b, c)
    fmt.Println(aa, bb, cc)
}
