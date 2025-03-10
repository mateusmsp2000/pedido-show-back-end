package application_test

import (
	"PedidoShow/application"
	"PedidoShow/dtos"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPublicarPedido(t *testing.T) {
	filaService := application.NewFilaPedidosService(1)

	pedido := dtos.PedidoDTO{
		ShowID: "123456",
		UserID: 1,
	}

	filaService.Publicar(pedido)

	select {
	case p := <-filaService.ObterFilaPedidos():
		assert.Equal(t, p.ShowID, pedido.ShowID)
		assert.Equal(t, p.UserID, pedido.UserID)
	case <-time.After(1 * time.Second):
		t.Fatal("Pedido não foi colocado na fila no tempo esperado")
	}
}

func TestProcessarPedidos(t *testing.T) {
	filaService := application.NewFilaPedidosService(3)

	pedidos := []dtos.PedidoDTO{
		{ShowID: "123456", UserID: 1},
		{ShowID: "789101", UserID: 2},
		{ShowID: "112233", UserID: 3},
	}

	for _, pedido := range pedidos {
		filaService.Publicar(pedido)
	}

	go filaService.Processar()

	select {
	case <-time.After(7 * time.Second):
	}
}
