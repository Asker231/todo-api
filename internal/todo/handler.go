package todo

import (
	"fmt"
	"net/http"

	"github.com/Asker231/todo-api.git/pkg/req"
)

type(
	TodoHandler struct{
		repo RepositoryTodo
	}
)

func NewTodoHandler(router *http.ServeMux,repo RepositoryTodo){
		todohandler := &TodoHandler{
			repo:  repo,
		}
		//paths
		router.HandleFunc("GET   /all",todohandler.GetAll())
		router.HandleFunc("GET   /todo/{id}",todohandler.GetById())
		router.HandleFunc("POST  /todo/delete/{id}",todohandler.Delete())
		router.HandleFunc("POST  /todo/create",todohandler.Create())
		router.HandleFunc("PATCH /todo/update/{id}",todohandler.Update())
}

func(todoHandler *TodoHandler)GetAll()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All todos"))
	}
}
func(todoHandler *TodoHandler)Create()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		body,err := req.HandleBody[TodoRequest](w,r)
		if err != nil{
			fmt.Println(err.Error())
		}
		todo := &Todo{
			Text: body.Text,
			Complited: body.Complited,
		}
		err = todoHandler.repo.Create(todo)
		if err != nil{
			fmt.Println(err.Error())
			return
		}

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