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

func DeleteTask() error {
	db := ConnectDB()
	result := db.Where("1==1").Delete(&models.Todo{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteTaskByID(id int) error {
	db := ConnectDB()
	result := db.Where("id=?", id).Delete(&models.Todo{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateTask(T *models.Todo) error {
	db := ConnectDB()
	result := db.Model(&models.Todo{}).Where("id =?", T.ID).Update("status", T.Status)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		fmt.Printf("No rows")
		return fmt.Errorf("no rows affected %d", T.ID)
	}
	return nil
}
