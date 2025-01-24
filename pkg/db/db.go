package db

import (
	"fmt"
	"github.com/Asker231/todo-api.git/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func ConnectDataBase(conf *config.AppConfig)*gorm.DB{
	db,err := gorm.Open(postgres.Open(conf.DataBaseConfig.DNS),&gorm.Config{})
	if err != nil{
		fmt.Println(err.Error())
	}
	return db
}