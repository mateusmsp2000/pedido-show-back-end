package entities

type Pedido struct {
	ID      uint    `gorm:"primaryKey"`
	ShowID  string  `gorm:"not null;index"`
	UserID  uint    `gorm:"not null;index"`
	Show    Show    `gorm:"foreignKey:ShowID"`
	Usuario Usuario `gorm:"foreignKey:UserID"`
}
