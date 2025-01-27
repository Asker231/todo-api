package auth

import (
	"net/http"

	"github.com/Asker231/todo-api.git/config"
)

type AuthUser struct{
	config *config.AuthConfig
}

func NewAuthUser(router http.ServeMux,config *config.AuthConfig){
	handleAuth := AuthUser{
		config: config,
	}
	router.Handle("POST /auth/register",handleAuth.Register())
	router.Handle("POST /auth/login",handleAuth.Login())

}

func(a *AuthUser)Register()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {}
}
func(a *AuthUser)Login()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {}
}