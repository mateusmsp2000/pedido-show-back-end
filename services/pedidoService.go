package services

import (
	"PedidoShow/dtos"
	"PedidoShow/fila"
	"PedidoShow/models"
	"PedidoShow/repositories"
	"errors"
)

type IPedidoService interface {
	Criar(pedido dtos.PedidoDTO) error
	ObterTodos() ([]dtos.PedidoDTO, error)
}

type PedidoService struct {
	pedidoRepo  repositories.IPedidoRepository
	userRepo    repositories.IUsuarioRepository
	showRepo    repositories.IShowRepository
	filaPedidos fila.IFilaPedidosService
}

func NewPedidoService(pedidoRepo repositories.IPedidoRepository, userRepo repositories.IUsuarioRepository, showRepo repositories.IShowRepository, filaPedidos fila.IFilaPedidosService) IPedidoService {
	return &PedidoService{pedidoRepo: pedidoRepo, userRepo: userRepo, showRepo: showRepo, filaPedidos: filaPedidos}
}

func (s *PedidoService) Criar(pedido dtos.PedidoDTO) error {
	_, err := s.userRepo.ObterPorID(pedido.UserID)
	if err != nil {
		return errors.New("usuário não encontrado")
	}

	_, err = s.showRepo.ObterPorID(pedido.ShowID)
	if err != nil {
		return errors.New("show não encontrado")
	}

	err = s.pedidoRepo.Criar(models.Pedido{
		UserID: pedido.UserID,
		ShowID: pedido.ShowID,
	})

	if err != nil {
		return err
	}

	s.filaPedidos.Publicar(pedido)

	return nil
}

func (s *PedidoService) ObterTodos() ([]dtos.PedidoDTO, error) {
	pedidos, err := s.pedidoRepo.ObterTodos()
	if err != nil {
		return nil, err
	}

	var pedidosDTO []dtos.PedidoDTO
	for _, pedido := range pedidos {
		pedidosDTO = append(pedidosDTO, dtos.PedidoDTO{
			UserID: pedido.UserID,
			ShowID: pedido.ShowID,
		})
	}

	return pedidosDTO, nil
}
