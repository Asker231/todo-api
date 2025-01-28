package main

import (
	"fmt"
	"net/http"

	"github.com/Asker231/todo-api.git/config"
	"github.com/Asker231/todo-api.git/internal/auth"
	"github.com/Asker231/todo-api.git/internal/todo"
	"github.com/Asker231/todo-api.git/pkg/db"
)



func main(){
	//config
	config := config.NewAppConfig()

	//database
	 db := db.ConnectDataBase(config)

	//create router
	router := http.NewServeMux()

	//repository
	repoTodo := todo.NewTodoRepository(*db)

	//handlers
	todo.NewTodoHandler(router,*repoTodo)
	auth.NewAuthUser(router,&config.AuthConfig)

	//create server
	server := http.Server{
		Addr: ":8080",
		Handler: router,
	}
	
	//server started
	err := server.ListenAndServe()
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println("Server started http://localhost:8080/")
}
