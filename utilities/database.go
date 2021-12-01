package utilities

import (
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var once sync.Once
var db *gorm.DB

type DB *gorm.DB

var err error

const dsn = "host=localhost user=postgres password=root dbname=todolist port=5432"

func Connect() *gorm.DB {
	once.Do(func() {
		db, err = gorm.Open(postgres.Open(dsn))
	})
	if err != nil {
		panic("Could not connect to the database")
	}

	return db
}
