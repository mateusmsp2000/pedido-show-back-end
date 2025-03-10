package repositories

import (
	"PedidoShow/domain/entities"
	"fmt"
	"gorm.io/gorm"
)

type IPedidoRepository interface {
	Criar(pedido entities.Pedido) error
	ObterTodos() ([]entities.Pedido, error)
}

type PedidoRepository struct {
	db *gorm.DB
}

func NewPedidoRepository(db *gorm.DB) IPedidoRepository {
	return &PedidoRepository{db: db}
}

func (repo *PedidoRepository) Criar(pedido entities.Pedido) error {
	if err := repo.db.Create(&pedido).Error; err != nil {
		return fmt.Errorf("erro ao salvar pedido: %w", err)
	}
	return nil
}

func (repo *PedidoRepository) ObterTodos() ([]entities.Pedido, error) {
	var pedidos []entities.Pedido
	if err := repo.db.Preload("Show").Preload("Usuario").Find(&pedidos).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar pedidos: %w", err)
	}
	return pedidos, nil
}
