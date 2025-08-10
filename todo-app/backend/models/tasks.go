package models

type Todo struct {
	ID     uint   `gorm:"primaryKey;autoincrement" json:"id"`
	Task   string `json:"task"`
	Status string `json:"status"`
}
