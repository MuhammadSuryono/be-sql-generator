package configs

import "github.com/MuhammadSuryono/module-golang-siakad/db"

func migrateDatabase() {
	
}

func (c configHandler) InitConnectionDatabase() {
	db.InitConnectionFromEnvirontment().CreateNewConnection()
	migrateDatabase()
}