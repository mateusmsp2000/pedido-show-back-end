package repositories

import (
	"PedidoShow/domain/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCriarShow(t *testing.T) {
	db, err := InitTestDB()
	assert.Nil(t, err)

	repo := NewShowRepository(db)

	show := entities.Show{ID: "123456", Name: "Show do Gustavo Lima"}

	err = repo.Criar(show)
	assert.Nil(t, err)

	var result entities.Show
	err = db.First(&result, "id = ?", show.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, result.Name, show.Name)
}

func TestRemoverShow(t *testing.T) {
	db, err := InitTestDB()
	assert.Nil(t, err)

	repo := NewShowRepository(db)

	show := entities.Show{ID: "123456", Name: "Show do Gustavo Lima"}
	err = repo.Criar(show)
	assert.Nil(t, err)

	err = repo.Remover(show.ID)
	assert.Nil(t, err)

	var result entities.Show
	err = db.First(&result, "id = ?", show.ID).Error
	assert.Error(t, err)
}

func TestObterShowPorID(t *testing.T) {
	db, err := InitTestDB()
	assert.Nil(t, err)

	repo := NewShowRepository(db)

	show := entities.Show{ID: "123456", Name: "Show do Gustavo Lima"}
	err = repo.Criar(show)
	assert.Nil(t, err)

	result, err := repo.ObterPorID(show.ID)
	assert.Nil(t, err)
	assert.Equal(t, result.Name, show.Name)

	result, err = repo.ObterPorID("inexistente")
	assert.Error(t, err)
}
