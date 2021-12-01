package models

//Task description model, need a todoListID
type HomeWork struct {
	ID          uint     `json:"id_homework"`
	Name        string   `json:"name_homework"`
	Description string   `json:"description_homework"`
	Status      bool     `gorm:"default:false" json:"status_homework"`
	ToDoListID  uint     `json:"-"`
	ToDoList    ToDoList `gorm:"foreignKey:ToDoListID;references:ID" json:"-"`
}
