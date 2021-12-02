package taskrepositories

import (
	"github.com/JuanMira/todolist/models"
	"github.com/JuanMira/todolist/utilities"
)

func GetLevelsPriority() ([]models.Level, error) {
	var level []models.Level
	db := utilities.Connect()

	row := db.Find(&level)

	if row.RowsAffected <= 0 {
		return level, row.Error
	}

	return level, nil
}
