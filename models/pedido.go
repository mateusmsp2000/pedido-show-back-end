package models

type Pedido struct {
	ID      uint    `gorm:"primaryKey"`
	ShowID  string  `gorm:"not null;index"`    // Definindo ShowID como chave estrangeira
	UserID  uint    `gorm:"not null;index"`    // Definindo UserID como chave estrangeira
	Show    Show    `gorm:"foreignKey:ShowID"` // Relacionamento com Show
	Usuario Usuario `gorm:"foreignKey:UserID"` // Relacionamento com User
}
