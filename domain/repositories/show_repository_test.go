package repositories

import (
	"PedidoShow/domain/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Teste para criar um show
func TestCriarShow(t *testing.T) {
	// Inicializando o banco de dados em memória
	db, err := InitTestDB()
	assert.Nil(t, err)

	// Criando o repositório com o banco de dados em memória
	repo := NewShowRepository(db)

	// Criando um show
	show := entities.Show{ID: "123456", Name: "Show do Gustavo Lima"}

	// Chamando o método Criar
	err = repo.Criar(show)
	assert.Nil(t, err)

	// Verificando se o show foi inserido no banco
	var result entities.Show
	err = db.First(&result, "id = ?", show.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, result.Name, show.Name)
}

// Teste para remover um show
func TestRemoverShow(t *testing.T) {
	// Inicializando o banco de dados em memória
	db, err := InitTestDB()
	assert.Nil(t, err)

	// Criando o repositório com o banco de dados em memória
	repo := NewShowRepository(db)

	// Criando um show
	show := entities.Show{ID: "123456", Name: "Show do Gustavo Lima"}
	err = repo.Criar(show)
	assert.Nil(t, err)

	// Chamando o método Remover
	err = repo.Remover(show.ID)
	assert.Nil(t, err)

	// Verificando se o show foi removido do banco
	var result entities.Show
	err = db.First(&result, "id = ?", show.ID).Error
	assert.Error(t, err) // Espera-se que o show não seja encontrado, ou seja, erro
}

// Teste para obter um show por ID
func TestObterShowPorID(t *testing.T) {
	// Inicializando o banco de dados em memória
	db, err := InitTestDB()
	assert.Nil(t, err)

	// Criando o repositório com o banco de dados em memória
	repo := NewShowRepository(db)

	// Criando um show
	show := entities.Show{ID: "123456", Name: "Show do Gustavo Lima"}
	err = repo.Criar(show)
	assert.Nil(t, err)

	// Chamando o método ObterPorID
	result, err := repo.ObterPorID(show.ID)
	assert.Nil(t, err)
	assert.Equal(t, result.Name, show.Name)

	// Testando um ID inexistente
	result, err = repo.ObterPorID("inexistente")
	assert.Error(t, err) // Espera-se erro pois o show não existe
}
