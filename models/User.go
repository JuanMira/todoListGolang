package models

import "time"

type User struct {
	ID        uint      `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `gorm:"unique" json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"create_date,omitempty"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
