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
		_ , err = a.servise.Register(payload.Email,payload.Password,payload.Name)
		if err != nil{
			return
		}

		res.Response(w,a.config.SECRET,200)

	}
}
func(a *AuthUser)Login()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {}
}