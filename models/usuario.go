package models

type Usuario struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}
