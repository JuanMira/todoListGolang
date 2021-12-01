package authrespositories

import (
	"github.com/JuanMira/todolist/models"
	"github.com/JuanMira/todolist/utilities"
)

func GetUser(username, password string) (bool, uint, error) {
	db := utilities.Connect()

	var user models.User

	row := db.Where("username = ?", username).First(&user)

	if row.RowsAffected <= 0 {
		return false, 0, nil
	}

	ok, err := utilities.CheckPassword(password, user.Password)

	if !ok {
		return false, 0, err
	}

	return true, user.ID, nil

}
