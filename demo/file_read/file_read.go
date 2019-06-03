package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//testFileRW()

	//testBufioRead()

	testIOutil()
}

/**
 * ioutil读取文件
 */
func testIOutil() {
	content, err := ioutil.ReadFile("./src/demo/file_rw/file.txt")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}

/**
 * bufio读取文件
 */
func testBufioRead() {
	file, err := os.Open("./src/demo/file_rw/file.txt")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("this is last line:\n", line)
			fmt.Println("\nlast line count:", len(line))
			break
		}
		if err != nil {
			fmt.Println("read file failed, err: ", err)
			return
		}
		fmt.Println(line, " , count is ", len(line))
	}
}

/**
 * 文件读取
 */
func testFileRead() {
	file, err := os.Open("./src/demo/file_rw/file.txt")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()

	var content []byte
	var buf [128]byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file:", err)
			return
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))
}