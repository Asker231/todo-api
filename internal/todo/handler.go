package todo

import "net/http"

type(
	TodoHandler struct{}
)

func NewTodoHandler(router *http.ServeMux){
		//paths
}

func(todoHandler *TodoHandler)GetAll()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {}
}
func(todoHandler *TodoHandler)Create()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {}
}
func(todoHandler *TodoHandler)Delete()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {}
}
func(todoHandler *TodoHandler)Update()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {}
}
func(todoHandler *TodoHandler)GetById()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {}
}