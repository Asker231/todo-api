package req

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)


func IsValid[T any](body T)error{
	v := validator.New()
	err := v.Struct(body)

	if err != nil{
		fmt.Println(err.Error())
		return err
	}
	return err
}