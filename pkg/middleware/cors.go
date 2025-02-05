package middleware

import "net/http"


func CORS(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header
		header.Add("Access-Control-Allow-Origin","*")
		header.Add("Access-Control-Allow-Methods","POST, GET, OPTIONS, DELETE, PUT")

		next.ServeHTTP(w,r)
	})
}
