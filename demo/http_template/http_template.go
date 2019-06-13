package main

import (
    "fmt"
    "html/template"
    "net/http"
)

type UserInfo struct {
    Name string
    Sex string
    Age int
    Address Address
    Score []Score
}

type Address struct {
    City string
    Province string
}

type Score struct {
    Chinese int
    Math int
    English int
}

func main() {
    http.HandleFunc("/login", login)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        fmt.Printf("listen server failed, err:%v\n", err)
        return
    }
}

func login(w http.ResponseWriter, r *http.Request) {
    method := r.Method
    if method == "GET" {
        t, err := template.ParseFiles("./src/demo/http_template/login.html")
        if err != nil {
            fmt.Fprintf(w, "load login.html failed")
            return
        }
        // 访问结构体
        scores := []Score{
            {Chinese:99, Math:98, English:97},
            {Chinese:88, Math:87, English:86},
        }

        user := &UserInfo{
           Name: "Tom",
           Sex: "男",
           Age: 10,
           Address: Address{
               City: "上海",
               Province: "上海市",
           },
           Score: scores,

        }
        t.Execute(w, user)

        // 访问map
        //mp := make(map[string]interface{})
        //mp["username"] = "张三"
        //mp["sex"] = "男"
        //mp["age"] = 12
        //t.Execute(w, mp)

        // 输出到终端
        //t.Execute(os.Stdout, mp)

    }else if method == "POST" {
        r.ParseForm()
        username := r.FormValue("username")
        password := r.FormValue("password")
        fmt.Printf("username:%s\n", username)
        fmt.Printf("password:%s\n", password)

        if username == "admin" && password == "admin" {
            fmt.Fprintf(w, "username:%s login success\n", username)
        }else {
            fmt.Fprintf(w, "username:%s login failed\n", username)
        }

    }
}