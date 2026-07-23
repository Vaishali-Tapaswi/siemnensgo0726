package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Emp struct {
	Empno  int     `json:"empno"`
	Ename  string  `json:"ename"`
	Salary float64 `json:"sal"`
}

func jsonhelper() string {
	emp := Emp{
		Empno:  11,
		Ename:  "AAA",
		Salary: 11000,
	}
	jsonformated, _ := json.Marshal(emp)
	// 1. % error
	// 2. how much -> client

	return string(jsonformated)
}

func main() {
	roothhandle := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<html><h1>Index Page</h1>")
	}
	pageHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		switch r.Method {
		case "GET":
			io.WriteString(w, jsonhelper())
		case "POST":
			body, _ := io.ReadAll(r.Body)
			emp := readhelper(body)
			fmt.Println("In post", emp)
		case "DELETE":
			fmt.Println(r.URL)
			fmt.Println(r.PathValue("id"))
		default:
			break
		}
	}
	http.HandleFunc("/", roothhandle)
	http.HandleFunc("/my", pageHandler)
	http.HandleFunc("/my/{id}", pageHandler)
	http.ListenAndServe(":8080", nil)
}

func readhelper(data []byte) Emp {
	var emp Emp
	json.Unmarshal(data, &emp)
	return emp
}
