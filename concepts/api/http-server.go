package api

import (
	"html/template"
	"io"
	"net/http"
)

func MainApi() {

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/tasks", taskHandler)
	http.ListenAndServe(":8888", nil)

}

type Task struct {
	Name string
	Done bool
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Hello, world!\n")
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	task := Task{
		Name: "Aprendendo GoLang",
		Done: true,
	}

	t := template.Must(template.ParseFiles("task.html"))

	t.Execute(w, task)
}
