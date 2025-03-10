package fila

import (
	"PedidoShow/dtos"
	"log"
	"time"
)

type IFilaPedidosService interface {
	Publicar(pedido dtos.PedidoDTO)
	Processar()
	ObterFilaPedidos() chan dtos.PedidoDTO // Método público para acessar o canal
}

type FilaPedidosService struct {
	filaPedidos chan dtos.PedidoDTO
}

func NewFilaPedidosService(tamanho int) *FilaPedidosService {
	return &FilaPedidosService{
		filaPedidos: make(chan dtos.PedidoDTO, tamanho),
	}
}

func (f *FilaPedidosService) Publicar(pedido dtos.PedidoDTO) {
	f.filaPedidos <- pedido
	log.Printf("Pedido publicado na fila: ShowID: %s, UserID: %d", pedido.ShowID, pedido.UserID)
}

func (f *FilaPedidosService) Processar() {
	for pedido := range f.filaPedidos {
		// Simula o processamento (ex: pagamento)
		log.Printf("Iniciando processamento do pedido: ShowID: %s, UserID: %d", pedido.ShowID, pedido.UserID)
		time.Sleep(2 * time.Second) // Simulando tempo de processamento
		log.Printf("Pedido processado e realizado pagamento com sucesso! ShowID: %s, UserID: %d", pedido.ShowID, pedido.UserID)
	}
}

func (f *FilaPedidosService) ObterFilaPedidos() chan dtos.PedidoDTO {
	return f.filaPedidos // Método público para acessar o canal
}
