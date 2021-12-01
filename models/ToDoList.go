package models

//to do list model, need a levelID and userID
type ToDoList struct {
	ID          uint   `json:"todolist_id"`
	Name        string `json:"todolist_name"`
	Description string `json:"todolist_description"`
	LevelID     uint   `json:"level_id"`
	UserID      uint   `json:"user_id"`
	Level       Level  `gorm:"foreignKey:LevelID;references:ID" json:"-"`
	User        User   `gorm:"foreignKey:UserID;references:ID" json:"-"`
}
