package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Asker231/todo-api.git/config"
	"github.com/Asker231/todo-api.git/pkg/jwt"
	"github.com/Asker231/todo-api.git/pkg/res"
)



func IsLogined(next http.Handler,conf *config.AppConfig)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		str := r.URL.Query().Get("auth")

		token := strings.TrimPrefix(str,"Berier ")
		
		isLog , email := jwt.NewGenerateJWT(conf.AuthConfig.SECRET).Parse(token)

		if !isLog{
			res.Response(w,http.StatusUnauthorized,404)
			return
		}
		fmt.Println(email)
		next.ServeHTTP(w,r)
	})
}