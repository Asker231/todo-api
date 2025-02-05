package middleware

import (
	"net/http"

	"github.com/Asker231/todo-api.git/pkg/res"
)



func IsLogined(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("auth")
		if token == ""{
			res.Response(w,"Not authorized",404)
			return
		}
		next.ServeHTTP(w,r)
	})
}