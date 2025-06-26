package main

import (
	"github.com/tonyseptiann12/go-crud/initializer"
	"github.com/tonyseptiann12/go-crud/models"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	initializer.DB.AutoMigrate(&models.Post{})
}
