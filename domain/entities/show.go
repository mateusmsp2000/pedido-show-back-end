package entities

type Show struct {
	ID   string `gorm:"primaryKey"`
	Name string
}
