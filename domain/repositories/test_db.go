package repositories

import (
	"PedidoShow/domain/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entities.Pedido{}, &entities.Usuario{}, &entities.Show{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
