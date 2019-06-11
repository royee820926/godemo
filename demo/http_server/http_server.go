package main

import "net/http"

func main() {
    http.HandleFunc("/", sayHelloName)
    err := http.ListenAndServe(":9090", nil)
}

func sayHelloName() {

}
