package todo

import "net/http"

type(
	TodoHandler struct{}
)

func NewTodoHandler(router *http.ServeMux){
		todohandler := &TodoHandler{}
		//paths
		router.HandleFunc("GET   /all",todohandler.GetAll())
		router.HandleFunc("GET   /todos/{id}",todohandler.GetById())
		router.HandleFunc("POST  /todos/delete/{id}",todohandler.Delete())
		router.HandleFunc("POST  /todos/create",todohandler.Create())
		router.HandleFunc("PATCH /todos/update/{id}",todohandler.Update())
}

func(todoHandler *TodoHandler)GetAll()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All todos"))
	}
}
func(todoHandler *TodoHandler)Create()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func(todoHandler *TodoHandler)Delete()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func(todoHandler *TodoHandler)Update()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func(todoHandler *TodoHandler)GetById()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}