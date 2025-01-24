package main

import (
	"fmt"
	"net/http"
	"github.com/Asker231/todo-api.git/config"
	"github.com/Asker231/todo-api.git/internal/todo"
	"github.com/Asker231/todo-api.git/pkg/db"
)



func main(){
	//config
	config := config.NewAppConfig()

	//database
	_ = db.ConnectDataBase(config)

	//create router
	router := http.NewServeMux()

	//handlers
	todo.NewTodoHandler(router)

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
