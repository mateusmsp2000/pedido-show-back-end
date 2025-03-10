package application_test

import (
	"PedidoShow/application"
	"PedidoShow/dtos"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPublicarPedido(t *testing.T) {
	// Inicializa a fila com tamanho 1 para o teste
	filaService := application.NewFilaPedidosService(1)

	// Criar um pedido de exemplo
	pedido := dtos.PedidoDTO{
		ShowID: "123456",
		UserID: 1,
	}

	// Publica o pedido na fila
	filaService.Publicar(pedido)

	// Verifica se o pedido foi colocado na fila
	select {
	case p := <-filaService.ObterFilaPedidos(): // Obtém o pedido da fila
		assert.Equal(t, p.ShowID, pedido.ShowID)
		assert.Equal(t, p.UserID, pedido.UserID)
	case <-time.After(1 * time.Second):
		t.Fatal("Pedido não foi colocado na fila no tempo esperado")
	}
}

func TestProcessarPedidos(t *testing.T) {
	// Inicializa a fila com capacidade 3
	filaService := application.NewFilaPedidosService(3)

	// Criar alguns pedidos de exemplo
	pedidos := []dtos.PedidoDTO{
		{ShowID: "123456", UserID: 1},
		{ShowID: "789101", UserID: 2},
		{ShowID: "112233", UserID: 3},
	}

	// Publica os pedidos na fila
	for _, pedido := range pedidos {
		filaService.Publicar(pedido)
	}

	// Inicia o processamento dos pedidos
	go filaService.Processar()

	// Espera um pouco para garantir que o processamento ocorra
	select {
	case <-time.After(7 * time.Second): // Tempo maior para garantir que todos os pedidos foram processados
		// Teste de sucesso, não há falhas
	}
}
