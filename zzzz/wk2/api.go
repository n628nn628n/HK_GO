package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("LOADING.....")
	//推送網址到首頁<<
	//類似laravel router 概念
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	err := http.ListenAndServe(":8000", nil)
	log.Println(err)
}
