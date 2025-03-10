package models

type Show struct {
	ID   string `gorm:"primaryKey"`
	Name string
}
