package todo

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Asker231/todo-api.git/pkg/middleware"
	"github.com/Asker231/todo-api.git/pkg/req"
	"github.com/Asker231/todo-api.git/pkg/res"
	"gorm.io/gorm"
)

type(
	TodoHandler struct{
		service *TodoService
	}
)

func NewTodoHandler(router *http.ServeMux,service *TodoService){
		todohandler := &TodoHandler{
			service: service,
		}
		//paths
		router.HandleFunc("GET   /all",todohandler.GetAll())
		router.HandleFunc("GET   /todo/{id}",todohandler.GetById())
		router.HandleFunc("DELETE  /todo/delete/{id}",todohandler.Delete())
		router.Handle("POST  /todo/create", middleware.IsLogined( todohandler.Create()))
		router.HandleFunc("PATCH /todo/update/{id}",todohandler.Update())
}

func(todoHandler *TodoHandler)GetAll()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		todos,err := todoHandler.service.ServiceGetAllTodos()
		if err != nil{
			 res.Response(w,err.Error(),404) 
			return
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
		result,err :=todoHandler.service.ServiceCreateTodo(&Todo{
			Text: body.Text,
			Complited: body.Complited,	
		})
		if err != nil{
			res.Response(w,err.Error(),404) 
			return
		}
		res.Response(w,result,200)

	}
}
func(todoHandler *TodoHandler)Delete()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id , _:= strconv.Atoi(idStr)

		err := todoHandler.service.ServiceDeleteTodo(id)
		if err != nil{
			res.Response(w,err.Error(),404)
			return
		}
		res.Response(w,"Success",200)
	}
}
func(todoHandler *TodoHandler)Update()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		body,err := req.HandleBody[TodoRequest](w,r)
		if err != nil{
			fmt.Println(err.Error())
		}
		idStr := r.PathValue("id")
		id , _:= strconv.Atoi(idStr)

		err = todoHandler.service.ServiceUpdateTodo(Todo{
			Model: &gorm.Model{
				ID: uint(id),
			},
			Text: body.Text,
			Complited: body.Complited,
		})

		if err != nil{
			res.Response(w,err.Error(),404)
			return
		}
		res.Response(w,"Success",200)

		
	}
}
func(todoHandler *TodoHandler)GetById()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r.PathValue("id")

	}
}