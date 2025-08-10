package database

import (
	"fmt"
	"todoapp/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect Database ")
	}
	e := db.AutoMigrate(&models.Todo{})
	if e != nil {
		fmt.Println(e)
	}
	return db
}

func InsertTask(T *models.Todo) {
	db := ConnectDB()
	res := db.Create(T)
	if res.Error != nil {
		fmt.Printf("Failed to insert: %v\n", res.Error)
		panic("Data not inserted")
	}
	fmt.Println(res.Rows())
}

func GetTask() ([]models.Todo, error) {
	db := ConnectDB()
	var T []models.Todo
	res := db.Find(&T)
	if res.Error != nil {
		return nil, res.Error
	}
	return T, nil
}
