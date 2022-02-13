package models

type ToDo struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Task      string `json:"title"`
	Completed bool   `json:"completed"`
}
