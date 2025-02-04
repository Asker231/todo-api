package user

import "gorm.io/gorm"

type UserRepo struct{
	DataBase *gorm.DB
}

func NewUserRepo(DataBase *gorm.DB)*UserRepo{
	return &UserRepo{
		DataBase: DataBase,
	}
}


func(userRepo *UserRepo)Register(user *User)(){}