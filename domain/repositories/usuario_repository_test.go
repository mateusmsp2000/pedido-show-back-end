package repositories

import (
	"PedidoShow/domain/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCriarUsuario(t *testing.T) {
	db, err := InitTestDB()
	assert.Nil(t, err)

	repo := NewUsuarioRepository(db)

	usuario := entities.Usuario{Name: "João Palesmano"}

	err = repo.Criar(usuario)
	assert.Nil(t, err)

	var result entities.Usuario
	err = db.First(&result, "name = ?", usuario.Name).Error
	assert.Nil(t, err)
	assert.Equal(t, result.Name, usuario.Name)
}

func TestRemoverUsuario(t *testing.T) {
	db, err := InitTestDB()
	assert.Nil(t, err)

	repo := NewUsuarioRepository(db)

	usuario := entities.Usuario{Name: "João Palesmano"}
	err = repo.Criar(usuario)
	assert.Nil(t, err)

	err = repo.Remover(usuario.ID)
	assert.Nil(t, err)

	var result entities.Usuario
	err = db.First(&result, "id = ?", usuario.ID).Error
	assert.Error(t, err)
}

func TestObterUsuarioPorID(t *testing.T) {
	db, err := InitTestDB()
	assert.Nil(t, err)

	repo := NewUsuarioRepository(db)

	usuario := entities.Usuario{ID: 1, Name: "João Palesmano"}
	err = repo.Criar(usuario)
	assert.Nil(t, err)

	result, err := repo.ObterPorID(usuario.ID)
	assert.Nil(t, err)
	assert.Equal(t, result.Name, usuario.Name)

	result, err = repo.ObterPorID(99999)
	assert.Error(t, err)
}
