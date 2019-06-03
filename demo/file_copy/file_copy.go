package main

import (
	"fmt"
	"io"
	"os"
)

var (
	srcFile = "./src/demo/file_write/file.txt"
	dstFile = "./src/demo/file_copy/file.txt"
)

func main() {
	_, err := CopyFile(dstFile, srcFile)
	if err != nil {
		fmt.Printf("copy file failed, err:%v\n", err)
		return
	}
	fmt.Println("Copy done!")
}

/**
 * 复制文件
 */
func CopyFile(dst_file, src_file string) (written int64, err error) {
	src, err := os.Open(src_file)
	if err != nil {
		fmt.Printf("open source file %s failed, err:%v\n", src_file, err)
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dst_file, os.O_WRONLY | os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open dest file %s failed, err:%v\n", dst_file, err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)

}
