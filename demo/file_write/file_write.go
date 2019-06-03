package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	filename = "./src/demo/file_write/file.txt"
)

func main() {
	//testWrite()
	//testWriteWithBufio()
	testWriteWithIOutil()
}

/**
 * 使用ioutil写入文件
 */
func testWriteWithIOutil()  {
	str := "hello world"
	err := ioutil.WriteFile(filename, []byte(str), 0755)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}

}

/**
 * 使用bufio写入文件
 */
func testWriteWithBufio()  {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i:=0; i<10; i++ {
		writer.WriteString(fmt.Sprintf("hello world %d\n", i))
	}
	writer.Flush()
}

/**
 * 普通文件写入，每次清空
 */
func testWrite() {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "hello world 1"
	file.Write([]byte(str))
}