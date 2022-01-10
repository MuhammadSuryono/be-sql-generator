package main

import (
	"github.com/be-sql-generator/configs"
	"github.com/be-sql-generator/routes"
	"github.com/be-sql-generator/server"
)

func main() {
	configs.InitConfig().InitConnectionDatabase()
	_= routes.CreateRouter(server.Start()).Run(":8000")
}