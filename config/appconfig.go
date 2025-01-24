package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

type(
	AppConfig struct{
		DataBaseConfig DataBaseConfig
		AuthConfig AuthConfig
	}
	DataBaseConfig struct{
		DNS string
	}
	AuthConfig struct{
		SECRET string
	}
)

func NewAppConfig()*AppConfig{
	err := godotenv.Load("../.env")
	if err != nil{
		fmt.Println(err.Error())
	}
	return &AppConfig{
		DataBaseConfig{
			DNS: os.Getenv("DNS"),
		},
		AuthConfig{
			SECRET:os.Getenv("TOKEN"),
		},
	}
}