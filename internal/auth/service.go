package auth

import (
	"github.com/Asker231/todo-api.git/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type ServiceAuth struct{
	UserRepo *user.UserRepo
}

func NewServiceAuth(userRepo *user.UserRepo)*ServiceAuth{
	return &ServiceAuth{
		UserRepo: userRepo,
	}
}

///Register Method
func(authService *ServiceAuth)Register(email,password,name string)(*user.User,error){

	pass,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil{
		return nil,err
	}
	user := user.User{
		Email: email,
		Password: string(pass),
		Name: name,
	}
	//bll
	u,err := authService.UserRepo.CreateUser(&user)
	if err != nil{
		return nil,err
	}
	return u,nil

}

