package auth

import (
	"net/http"

	"github.com/Asker231/todo-api.git/config"
	"github.com/Asker231/todo-api.git/pkg/req"
	"github.com/Asker231/todo-api.git/pkg/res"
)

type AuthUser struct{
	config *config.AuthConfig
}

func NewAuthUser(router *http.ServeMux,config *config.AuthConfig){
	handleAuth := AuthUser{
		config: config,
	}
	router.Handle("POST /auth/register",handleAuth.Register())
	router.Handle("POST /auth/login",handleAuth.Login())

}

func(a *AuthUser)Register()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		_ , err := req.HandleBody[RegisterRequest](w,r)
		if err != nil{
		   res.Response(w,err.Error(),401)	
		   return	
		}
		res.Response(w,a.config.SECRET,200)

	}
}
func(a *AuthUser)Login()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {}
}