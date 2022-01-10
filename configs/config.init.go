package configs

import (
	"log"

	"github.com/joho/godotenv"
)

type configs interface {
	InitConnectionDatabase()
}

type configHandler struct {

}

func InitConfig() configs {
	// Load env file
	errEnv := godotenv.Load("configs/.env")
	if errEnv != nil {
		log.Fatalln("Error when load file .env")
	}
	
	return configHandler{}
}