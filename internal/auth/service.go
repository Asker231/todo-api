package auth

import (
	"github.com/Asker231/todo-api.git/internal/user"
)

type ServiceAuth struct{
	UserRepo *user.UserRepo
}

func NewServiceAuth(userRepo *user.UserRepo)*ServiceAuth{
	return &ServiceAuth{
		UserRepo: userRepo,
	}
}


