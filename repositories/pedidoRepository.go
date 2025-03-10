package repositories

import (
	"PedidoShow/models"
	"fmt"
	"gorm.io/gorm"
)

type IPedidoRepository interface {
	Criar(pedido models.Pedido) error
	ObterTodos() ([]models.Pedido, error)
}

type PedidoRepository struct {
	db *gorm.DB
}

func NewPedidoRepository(db *gorm.DB) IPedidoRepository {
	return &PedidoRepository{db: db}
}

func (repo *PedidoRepository) Criar(pedido models.Pedido) error {
	if err := repo.db.Create(&pedido).Error; err != nil {
		return fmt.Errorf("erro ao salvar pedido: %w", err)
	}
	return nil
}

func (repo *PedidoRepository) ObterTodos() ([]models.Pedido, error) {
	var pedidos []models.Pedido
	if err := repo.db.Preload("Show").Preload("Usuario").Find(&pedidos).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar pedidos: %w", err)
	}
	return pedidos, nil
}
