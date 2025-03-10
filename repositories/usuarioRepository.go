package repositories

import (
	"PedidoShow/models"
	"fmt"
	"gorm.io/gorm"
)

type IUsuarioRepository interface {
	Criar(usuario models.Usuario) error
	Remover(id uint) error
	ObterPorID(id uint) (models.Usuario, error)
}

type UsuarioRepository struct {
	db *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) IUsuarioRepository {
	return &UsuarioRepository{db: db}
}

func (repo *UsuarioRepository) Criar(usuario models.Usuario) error {
	if err := repo.db.Create(&usuario).Error; err != nil {
		return fmt.Errorf("erro ao salvar usuario: %w", err)
	}
	return nil
}

func (repo *UsuarioRepository) Remover(id uint) error {
	if err := repo.db.Delete(&models.Usuario{}, id).Error; err != nil {
		return fmt.Errorf("erro ao remover usuário: %w", err)
	}
	return nil
}

func (repo *UsuarioRepository) ObterPorID(id uint) (models.Usuario, error) {
	var usuario models.Usuario
	if err := repo.db.First(&usuario, id).Error; err != nil {
		return usuario, err
	}
	return usuario, nil
}
