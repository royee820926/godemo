package main

import (
    "fmt"
    "log"
    "path/filepath"
)

func main() {

    //dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    dir, err := filepath.Abs(".")
    if err != nil {
      log.Fatal(err)
    }
    fmt.Printf("%v", dir)

}
