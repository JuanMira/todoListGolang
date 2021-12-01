package models

type Error struct {
	Message string      `json:"message"`
	Err     interface{} `json:"error"`
	Status  bool        `json:"status"`
}

type Success struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}
