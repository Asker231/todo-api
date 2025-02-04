package auth

import (
	"net/http"

	"github.com/Asker231/todo-api.git/config"
	"github.com/Asker231/todo-api.git/pkg/req"
	"github.com/Asker231/todo-api.git/pkg/res"
)

type AuthUser struct{
	config *config.AuthConfig
	servise *ServiceAuth
}

func NewAuthUser(router *http.ServeMux,config *config.AuthConfig,serviceAuth ServiceAuth){
	handleAuth := AuthUser{
		config: config,
		servise: &serviceAuth,
	}
	router.Handle("POST /auth/register",handleAuth.Register())
	router.Handle("POST /auth/login",handleAuth.Login())

}

func(a *AuthUser)Register()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		payload , err := req.HandleBody[RegisterRequest](w,r)
		if err != nil{
		   res.Response(w,err.Error(),401)	
		   return	
		}
		us ,err := a.servise.Register(payload.Email,payload.Password,payload.Name)
		if err != nil{
			res.Response(w,err.Error(),204)
			return
		}

		res.Response(w,us,200)

	}
}
func(a *AuthUser)Login()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		payload,err := req.HandleBody[LoginRequest](w,r)
		if err != nil{
			res.Response(w,err.Error(),401)	
			return	
		 }
		 _,err = a.servise.Login(payload.Email,payload.Password)

		 if err != nil{
			res.Response(w,err.Error(),204)
			return
		 }
		 res.Response(w,"Create",201)

	}
}