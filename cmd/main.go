package main

import (
	"fmt"
	"net/http"

	"github.com/Asker231/todo-api.git/config"
	"github.com/Asker231/todo-api.git/internal/auth"
	"github.com/Asker231/todo-api.git/internal/todo"
	"github.com/Asker231/todo-api.git/internal/user"
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
	repoUser := user.NewUserRepo(db)

	//services
	authService := auth.NewServiceAuth(repoUser)	
	todoService := todo.NewServiceTodo(repoTodo)

	//handlers
	todo.NewTodoHandler(router, todoService)
	auth.NewAuthUser(router,&config.AuthConfig,*authService)


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
