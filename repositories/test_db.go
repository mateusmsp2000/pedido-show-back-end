package repositories

import (
	"PedidoShow/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitTestDB() (*gorm.DB, error) {
	// Criando um banco de dados em memória
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Rodando as migrações para garantir que as tabelas existam
	err = db.AutoMigrate(&models.Pedido{}, &models.Usuario{}, &models.Show{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
