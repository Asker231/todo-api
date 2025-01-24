package req

import (
	"encoding/json"
	"fmt"
	"io"
)


func DecodeJson[T any](body io.ReadCloser)(*T,error){
	var payload T
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil{
		fmt.Println(err.Error())
		return nil,err
	}
	return &payload,nil
}