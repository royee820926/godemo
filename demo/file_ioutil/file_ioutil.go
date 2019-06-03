package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    inputFile := "./src/demo/file_ioutil/file.txt"
    //outputFile := "./src/demo/file_ioutil/file_copy.txt"
    content, err := ioutil.ReadFile(inputFile)
    if err != nil {
       fmt.Fprintf(os.Stderr, "File Error:%s\n", err)
       return
    }
    fmt.Println(string(content))
}
