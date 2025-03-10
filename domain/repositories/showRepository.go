package repositories

import (
	"PedidoShow/domain/entities"
	"fmt"
	"gorm.io/gorm"
)

type IShowRepository interface {
	Criar(show entities.Show) error
	Remover(id string) error
	ObterPorID(id string) (entities.Show, error)
}

type ShowRepository struct {
	db *gorm.DB
}

func NewShowRepository(db *gorm.DB) *ShowRepository {
	return &ShowRepository{db: db}
}

func (repo *ShowRepository) Criar(show entities.Show) error {
	if err := repo.db.Create(&show).Error; err != nil {
		return fmt.Errorf("erro ao salvar show: %w", err)
	}
	return nil
}

func (repo *ShowRepository) Remover(id string) error {
	if err := repo.db.Delete(&entities.Show{}, id).Error; err != nil {
		return fmt.Errorf("erro ao remover show: %w", err)
	}
	return nil
}

func (repo *ShowRepository) ObterPorID(id string) (entities.Show, error) {
	var show entities.Show
	if err := repo.db.First(&show, "id = ?", id).Error; err != nil {
		return show, err
	}
	return show, nil
}
