package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal("can't parse Request")
	}
	fmt.Println(r.Form)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println(r.Form["url_ long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value: ", strings.Join(v, ""))
	}
	_, err = fmt.Fprint(w, "Hello GoWeb!")
	if err != nil {
		log.Fatal("failed to write into ResponseWrite")
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
