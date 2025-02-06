package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct{
	SecretKey string
}
func NewGenerateJWT(secret string)*JWT{
	 return &JWT{
		SecretKey: secret,
	 }
}

func(jdata *JWT)GenJWT(email string)(string,error){
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email":email,
	})
	res ,err := t.SignedString([]byte(jdata.SecretKey))
	if err != nil{
		return err.Error(),nil
	}
	return res,nil
}

func(jdata *JWT)Parse(token string)(bool,string){
	t ,err := jwt.Parse(token,func(t *jwt.Token) (interface{}, error) {
		return []byte(jdata.SecretKey) , nil
	})
	if err != nil{
		fmt.Println(err.Error())
	}
	return t.Valid, t.Claims.(jwt.MapClaims)["email"].(string)
}