package repositories

import (
	"PedidoShow/domain/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCriarPedido(t *testing.T) {
	db, err := InitTestDB()
	assert.Nil(t, err)

	repo := NewPedidoRepository(db)

	pedido := entities.Pedido{
		ShowID: "123456",
		UserID: 1,
	}

	err = repo.Criar(pedido)
	assert.Nil(t, err)

	var pedidos []entities.Pedido
	err = db.Find(&pedidos).Error
	assert.Nil(t, err)
	assert.Len(t, pedidos, 1)
	assert.Equal(t, pedidos[0].ShowID, "123456")
	assert.Equal(t, pedidos[0].UserID, uint(1))
}

func TestObterTodosPedidos(t *testing.T) {
	db, err := InitTestDB()
	assert.Nil(t, err)

	repo := NewPedidoRepository(db)

	pedidos := []entities.Pedido{
		{ShowID: "123456", UserID: 1},
		{ShowID: "789101", UserID: 2},
	}

	for _, p := range pedidos {
		err := repo.Criar(p)
		assert.Nil(t, err)
	}

	resultados, err := repo.ObterTodos()
	assert.Nil(t, err)

	assert.Len(t, resultados, 2)
	assert.Equal(t, resultados[0].ShowID, "123456")
	assert.Equal(t, resultados[1].ShowID, "789101")
}
