package repositories

import (
	"PedidoShow/domain/entities"
	"fmt"
	"gorm.io/gorm"
)

type IUsuarioRepository interface {
	Criar(usuario entities.Usuario) error
	Remover(id uint) error
	ObterPorID(id uint) (entities.Usuario, error)
}

type UsuarioRepository struct {
	db *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) IUsuarioRepository {
	return &UsuarioRepository{db: db}
}

func (repo *UsuarioRepository) Criar(usuario entities.Usuario) error {
	if err := repo.db.Create(&usuario).Error; err != nil {
		return fmt.Errorf("erro ao salvar usuario: %w", err)
	}
	return nil
}

func (repo *UsuarioRepository) Remover(id uint) error {
	if err := repo.db.Delete(&entities.Usuario{}, id).Error; err != nil {
		return fmt.Errorf("erro ao remover usuário: %w", err)
	}
	return nil
}

func (repo *UsuarioRepository) ObterPorID(id uint) (entities.Usuario, error) {
	var usuario entities.Usuario
	if err := repo.db.First(&usuario, id).Error; err != nil {
		return usuario, err
	}
	return usuario, nil
}
