package application

import (
	"PedidoShow/domain/entities"
	repositories2 "PedidoShow/domain/repositories"
	"PedidoShow/dtos"
	"errors"
)

type IPedidoService interface {
	Criar(pedido dtos.PedidoDTO) error
	ObterTodos() ([]dtos.PedidoDTO, error)
}

type PedidoService struct {
	pedidoRepo  repositories2.IPedidoRepository
	userRepo    repositories2.IUsuarioRepository
	showRepo    repositories2.IShowRepository
	filaPedidos IFilaPedidosService
}

func NewPedidoService(
	pedidoRepo repositories2.IPedidoRepository,
	userRepo repositories2.IUsuarioRepository,
	showRepo repositories2.IShowRepository,
	filaPedidos IFilaPedidosService) IPedidoService {
	return &PedidoService{
		pedidoRepo:  pedidoRepo,
		userRepo:    userRepo,
		showRepo:    showRepo,
		filaPedidos: filaPedidos}
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

	err = s.pedidoRepo.Criar(entities.Pedido{
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
