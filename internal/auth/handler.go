package auth

import (
	"net/http"

	"github.com/Asker231/todo-api.git/config"
	"github.com/Asker231/todo-api.git/pkg/jwt"
	"github.com/Asker231/todo-api.git/pkg/middleware"
	"github.com/Asker231/todo-api.git/pkg/req"
	"github.com/Asker231/todo-api.git/pkg/res"
)

type AuthUser struct{
	config *config.AppConfig
	servise *ServiceAuth
}

func NewAuthUser(router *http.ServeMux,config *config.AppConfig,serviceAuth ServiceAuth){
	handleAuth := AuthUser{
		config: config,
		servise: &serviceAuth,
	}
	router.HandleFunc("POST /auth/register",handleAuth.Register())
	router.HandleFunc("POST /auth/login",handleAuth.Login())

	router.Handle("/home",middleware.IsLogined(handleAuth.HomePage(),config))

}
func(a *AuthUser)HomePage()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from home page"))
	}
}
func(a *AuthUser)Register()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		payload , err := req.HandleBody[RegisterRequest](w,r)
		if err != nil{
		   res.Response(w,err.Error(),401)	
		   return	
		}
		_ ,err = a.servise.Register(payload.Email,payload.Password,payload.Name)
		if err != nil{
			res.Response(w,err.Error(),204)
			return
		}
		t ,err := jwt.NewGenerateJWT(a.config.AuthConfig.SECRET).GenJWT(payload.Email)

		if err != nil{
		   res.Response(w,err.Error(),404)
		   return
		}
		res.Response(w,t,201)
		http.Redirect(w,r,"/home",301)


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
		 http.Redirect(w,r,"/home",301)
	}
}