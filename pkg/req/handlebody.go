package req

import (
	"net/http"
	"github.com/Asker231/todo-api.git/pkg/res"
)


func HandleBody[T any](w http.ResponseWriter, r *http.Request)(*T,error){
	body ,err := DecodeJson[T](r.Body)
	if err != nil{
		res.Response(w,err.Error(),http.StatusBadRequest)
		return nil,err
	}
	// err  = IsValid(body)
	// if err != nil{
	// 	res.Response(w,err.Error(),402)
	// 	return nil,err
	// }
	return body,nil
}