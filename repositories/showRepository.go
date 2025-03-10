package repositories

import (
	"PedidoShow/models"
	"fmt"
	"gorm.io/gorm"
)

type IShowRepository interface {
	Criar(show models.Show) error
	Remover(id string) error
	ObterPorID(id string) (models.Show, error)
}

type ShowRepository struct {
	db *gorm.DB
}

func NewShowRepository(db *gorm.DB) *ShowRepository {
	return &ShowRepository{db: db}
}

func (repo *ShowRepository) Criar(show models.Show) error {
	if err := repo.db.Create(&show).Error; err != nil {
		return fmt.Errorf("erro ao salvar show: %w", err)
	}
	return nil
}

func (repo *ShowRepository) Remover(id string) error {
	if err := repo.db.Delete(&models.Show{}, id).Error; err != nil {
		return fmt.Errorf("erro ao remover show: %w", err)
	}
	return nil
}

func (repo *ShowRepository) ObterPorID(id string) (models.Show, error) {
	var show models.Show
	if err := repo.db.First(&show, "id = ?", id).Error; err != nil {
		return show, err
	}
	return show, nil
}
