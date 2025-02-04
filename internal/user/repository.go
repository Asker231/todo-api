package user

import (
	"fmt"

	"gorm.io/gorm"
)

type UserRepo struct{
	DataBase *gorm.DB
}

func NewUserRepo(DataBase *gorm.DB)*UserRepo{
	return &UserRepo{
		DataBase: DataBase,
	}
}

func(userRepo *UserRepo)CreateUser(user *User)(*User,error){
	result := userRepo.DataBase.Create(user)
	if result.Error != nil{
		fmt.Println(result.Error.Error())
		return nil,result.Error
	}
	return user,nil
}

func(userRepo *UserRepo)FindByEmail(email string)(*User){
	var payload User
	result := userRepo.DataBase.First(&payload,"email = ?",email)
	if result.Error != nil{
		fmt.Println(result.Error.Error())
		return nil
	}
	return &payload
}