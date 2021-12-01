package utilities

import "github.com/JuanMira/todolist/models"

func Migrate() {
	db := Connect()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Level{})
	db.AutoMigrate(&models.ToDoList{})
	db.AutoMigrate(&models.HomeWork{})
}
