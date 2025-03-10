package repositories

import (
	"PedidoShow/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Teste para criar um usuário
func TestCriarUsuario(t *testing.T) {
	// Usando o initTestDB() existente no seu código de testes
	db, err := InitTestDB() // Aqui está utilizando a função que já existe
	assert.Nil(t, err)

	// Criando o repositório com o banco de dados em memória
	repo := NewUsuarioRepository(db)

	// Criando um usuário
	usuario := models.Usuario{Name: "João Palesmano"}

	// Chamando o método Criar
	err = repo.Criar(usuario)
	assert.Nil(t, err)

	// Verificando se o usuário foi inserido no banco
	var result models.Usuario
	err = db.First(&result, "name = ?", usuario.Name).Error
	assert.Nil(t, err)
	assert.Equal(t, result.Name, usuario.Name)
}

// Teste para remover um usuário
func TestRemoverUsuario(t *testing.T) {
	// Usando o initTestDB() existente no seu código de testes
	db, err := InitTestDB() // Aqui está utilizando a função que já existe
	assert.Nil(t, err)

	// Criando o repositório com o banco de dados em memória
	repo := NewUsuarioRepository(db)

	// Criando um usuário
	usuario := models.Usuario{Name: "João Palesmano"}
	err = repo.Criar(usuario)
	assert.Nil(t, err)

	// Chamando o método Remover
	err = repo.Remover(usuario.ID)
	assert.Nil(t, err)

	// Verificando se o usuário foi removido do banco
	var result models.Usuario
	err = db.First(&result, "id = ?", usuario.ID).Error
	assert.Error(t, err) // Espera-se que o usuário não seja encontrado
}

// Teste para obter um usuário por ID
func TestObterUsuarioPorID(t *testing.T) {
	// Usando o initTestDB() existente no seu código de testes
	db, err := InitTestDB() // Aqui está utilizando a função que já existe
	assert.Nil(t, err)

	// Criando o repositório com o banco de dados em memória
	repo := NewUsuarioRepository(db)

	// Criando um usuário
	usuario := models.Usuario{ID: 1, Name: "João Palesmano"}
	err = repo.Criar(usuario)
	assert.Nil(t, err)

	// Chamando o método ObterPorID
	result, err := repo.ObterPorID(usuario.ID)
	assert.Nil(t, err)
	assert.Equal(t, result.Name, usuario.Name)

	// Testando um ID inexistente
	result, err = repo.ObterPorID(99999) // ID inexistente
	assert.Error(t, err)                 // Espera-se erro pois o usuário não existe
}
