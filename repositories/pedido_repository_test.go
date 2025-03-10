package repositories

import (
	"PedidoShow/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Teste de criação de pedido
func TestCriarPedido(t *testing.T) {
	// Inicializando o banco de dados em memória
	db, err := InitTestDB()
	assert.Nil(t, err)

	// Criando o repositório com o banco de dados em memória
	repo := NewPedidoRepository(db)

	// Criando um pedido
	pedido := models.Pedido{
		ShowID: "123456",
		UserID: 1,
	}

	// Chamando o método Criar
	err = repo.Criar(pedido)
	assert.Nil(t, err)

	// Verificando se o pedido foi inserido no banco
	var pedidos []models.Pedido
	err = db.Find(&pedidos).Error
	assert.Nil(t, err)
	assert.Len(t, pedidos, 1)
	assert.Equal(t, pedidos[0].ShowID, "123456")
	assert.Equal(t, pedidos[0].UserID, uint(1))
}

// Teste para obter todos os pedidos
func TestObterTodosPedidos(t *testing.T) {
	// Inicializando o banco de dados em memória
	db, err := InitTestDB()
	assert.Nil(t, err)

	// Criando o repositório com o banco de dados em memória
	repo := NewPedidoRepository(db)

	// Inserindo alguns pedidos para teste
	pedidos := []models.Pedido{
		{ShowID: "123456", UserID: 1},
		{ShowID: "789101", UserID: 2},
	}

	for _, p := range pedidos {
		err := repo.Criar(p)
		assert.Nil(t, err)
	}

	// Chamando o método ObterTodos
	resultados, err := repo.ObterTodos()
	assert.Nil(t, err)

	// Verificando se os pedidos foram retornados corretamente
	assert.Len(t, resultados, 2)
	assert.Equal(t, resultados[0].ShowID, "123456")
	assert.Equal(t, resultados[1].ShowID, "789101")
}
