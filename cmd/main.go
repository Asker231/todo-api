package main

import (
	"fmt"
	"net/http"
)



func main(){
	//create router
	router := http.NewServeMux()

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
