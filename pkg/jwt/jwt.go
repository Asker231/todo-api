package jwt

import "github.com/golang-jwt/jwt/v5"

type JWTData struct{
	SecretKey string
}


func NewGenerateJWT(secret string)*JWTData{
	 return &JWTData{
		SecretKey: secret,
	 }
}

func(jdata *JWTData)GenJWT(email string)(string,error){
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email":email,
	})
	res ,err := t.SignedString([]byte(jdata.SecretKey))
	if err != nil{
		return err.Error(),nil
	}
	return res,nil
}