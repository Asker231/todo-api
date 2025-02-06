package main

import (
	"fmt"
	"net/http"

	"github.com/Asker231/todo-api.git/config"
	"github.com/Asker231/todo-api.git/internal/auth"
	"github.com/Asker231/todo-api.git/internal/todo"
	"github.com/Asker231/todo-api.git/internal/user"
	"github.com/Asker231/todo-api.git/pkg/db"
	"github.com/Asker231/todo-api.git/pkg/middleware"
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
	todo.NewTodoHandler(router, todoService,config)
	auth.NewAuthUser(router,config,*authService)


	//create server
	server := http.Server{
		Addr: ":8080",
		Handler: middleware.CORS(router),
	}
	
	//server started
	err := server.ListenAndServe()
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println("Server started http://localhost:8080/")
}
