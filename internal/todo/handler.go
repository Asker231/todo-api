package todo

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Asker231/todo-api.git/pkg/req"
	"github.com/Asker231/todo-api.git/pkg/res"
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
		router.HandleFunc("DELETE  /todo/delete/{id}",todohandler.Delete())
		router.HandleFunc("POST  /todo/create",todohandler.Create())
		router.HandleFunc("PATCH /todo/update/{id}",todohandler.Update())
}

func(todoHandler *TodoHandler)GetAll()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		todos,err := todoHandler.repo.GetAll()	
		if err != nil{
			fmt.Println(err.Error())
		}
		res.Response(w,todos,200)
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
			Complited: false,
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
		idStr := r.PathValue("id")
		id , _:= strconv.Atoi(idStr)

		err := todoHandler.repo.DeleteTodo(id)
		if err != nil{
			fmt.Println(err.Error())
		}
		todos ,err := todoHandler.repo.GetAll()
		if err != nil{
			fmt.Println(err)
		}
		res.Response(w,todos,200)

	}
}
func(todoHandler *TodoHandler)Update()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id , _:= strconv.Atoi(idStr)

		result,err := todoHandler.repo.UpdateTodo(id)
		if err != nil{
			fmt.Println(err.Error())
		}
		res.Response(w,result,200)


	}
}
func(todoHandler *TodoHandler)GetById()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r.PathValue("id")

	}
}