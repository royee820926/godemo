package main

import (
    "flag"
    "fmt"
    "os"
    "os/exec"
    "time"
)

var (
    child *bool
)

func main() {
    if child != nil && *child == true {
        fmt.Printf("继承于父进程的文件句柄\n")
        readFromParent()
        return
    }
    // 父进程的逻辑
    file, err := os.OpenFile("/tmp/test_inherit.log", os.O_APPEND | os.O_CREATE | os.O_RDONLY, 0755)
    if err != nil {
        fmt.Printf("open file failed, err:%v\n", err)
        return
    }

    _, err = file.WriteString("parent write one line\n")
    if err != nil {
        fmt.Printf("parent write failed, err:%v\n", err)
        return
    }

    startChild(file)
    fmt.Printf("parent exited")
}

func init()  {
    child = flag.Bool("child", false, "继承于父进程(internal use only)")
    flag.Parse()
}

func startChild(file *os.File) {
    args := []string{"-child"}
    cmd := exec.Command(os.Args[0], args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    // put socket DF at the first entry
    /********************* VIP ******************/
    cmd.ExtraFiles = []*os.File{file}
    err := cmd.Start()
    if err != nil {
        fmt.Printf("start child failed, err:%v\n", err)
        return
    }
}

func readFromParent() {
    f := os.NewFile(3, "")
    count := 0
    for {
        str := fmt.Sprintf("hello, i'child process, write:%d line\n", count)
        count += 1
        _, err := f.WriteString(str)
        if err != nil {
            fmt.Printf("write string failed, err:%v\n", err)
            time.Sleep(time.Second)
            continue
        }
        time.Sleep(time.Second)
    }
}
