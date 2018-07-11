package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Demo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Grow int    `json:"-"`
}

func main() {

	http.HandleFunc("/api", apiController)

	err := http.ListenAndServe(":8000", nil)
	log.Fatalln(err)
}

func apiController(w http.ResponseWriter, r *http.Request) {
	s := map[string]interface{}{}

	data, _ := json.Marshal(s)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
