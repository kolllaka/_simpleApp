package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var tmpl *template.Template

type Todo struct {
	Item string `json:"item"`
	Done bool   `json:"done"`
}

type PageData struct {
	Title string `json:"-"`
	Todos []Todo `json:"todos"`
}

func main() {
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("./templates/index.go.html"))

	fs := http.FileServer(http.Dir("./static"))

	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", indexHandle)

	fmt.Println("server start on port: 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	data := PageData{Title: "Todo List"}

	databyte, err := ioutil.ReadFile("./todos.json")
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(databyte, &data)

	// buffer := new(bytes.Buffer)
	// tmpl.Execute(buffer, data)
	// if err := ioutil.WriteFile("./indexGenerated.html", buffer.Bytes(), 0644); err != nil {
	// 	log.Fatalf("Error while write to file: %v", err)
	// }

	tmpl.Execute(w, data)
}
