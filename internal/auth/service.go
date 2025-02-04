package auth

import (
	"errors"

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

	us := authService.UserRepo.FindByEmail(email)
	if us != nil{
		return nil,errors.New("Пользователь уже существует")
	}

	pass,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil{
		return nil,err
	}
	user := user.User{
		Email: email,
		Password: string(pass),
		Name: name,
	}
	
	u,err := authService.UserRepo.CreateUser(&user)
	if err != nil{
		return nil,err
	}
	return u,nil

}

//Login Method

func(authService *ServiceAuth)Login(email,password string)(*user.User,error){
	u:= authService.UserRepo.FindByEmail(email)
	if u == nil{
		return nil,errors.New("Нет такого пользователя")
	}
	err := bcrypt.CompareHashAndPassword([]byte(password),[]byte(u.Password))
	if err != nil{
		return nil,err
	}
	return u,nil
}

