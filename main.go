package main

import "github.com/be-sql-generator/configs"

func main() {
	configs.InitConfig().InitConnectionDatabase()
}