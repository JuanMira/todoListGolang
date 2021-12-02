package taskrepositories

import (
	"github.com/JuanMira/todolist/models"
	"github.com/JuanMira/todolist/utilities"
)

func CreateToDoList(newTodolist *models.ToDoList) error {
	db := utilities.Connect() //connection to database
	uID := utilities.UserID   //user id

	newTodolist.UserID = uint(uID)

	row := db.Create(&newTodolist)

	if row.RowsAffected <= 0 {
		return row.Error
	}

	return nil
}
