package services_test

import (
	"PedidoShow/dtos"
	"PedidoShow/models"
	"PedidoShow/services"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// Criando mocks para os repositórios e o serviço de fila
type MockPedidoRepository struct {
	mock.Mock
}

func (m *MockPedidoRepository) Criar(pedido models.Pedido) error {
	args := m.Called(pedido)
	return args.Error(0)
}

func (m *MockPedidoRepository) ObterTodos() ([]models.Pedido, error) {
	args := m.Called()
	return args.Get(0).([]models.Pedido), args.Error(1)
}

type MockUsuarioRepository struct {
	mock.Mock
}

func (m *MockUsuarioRepository) ObterPorID(id uint) (models.Usuario, error) {
	args := m.Called(id)
	return args.Get(0).(models.Usuario), args.Error(1)
}

func (m *MockUsuarioRepository) Criar(usuario models.Usuario) error {
	args := m.Called(usuario)
	return args.Error(0)
}

func (m *MockUsuarioRepository) Remover(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

type MockShowRepository struct {
	mock.Mock
}

func (m *MockShowRepository) Criar(show models.Show) error {
	args := m.Called(show)
	return args.Error(0)
}

func (m *MockShowRepository) Remover(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockShowRepository) ObterPorID(id string) (models.Show, error) {
	args := m.Called(id)
	return args.Get(0).(models.Show), args.Error(1)
}

type MockFilaPedidosService struct {
	mock.Mock
}

func (m *MockFilaPedidosService) Publicar(pedido dtos.PedidoDTO) {
	m.Called(pedido)
}

func (m *MockFilaPedidosService) Processar() {
	m.Called()
}

func (m *MockFilaPedidosService) ObterFilaPedidos() chan dtos.PedidoDTO {
	args := m.Called()
	return args.Get(0).(chan dtos.PedidoDTO)
}

func TestPedidoService_Criar(t *testing.T) {
	// Criando os mocks
	pedidoRepo := new(MockPedidoRepository)
	usuarioRepo := new(MockUsuarioRepository)
	showRepo := new(MockShowRepository)
	filaPedidosService := new(MockFilaPedidosService)

	// Criando o serviço
	pedidoService := services.NewPedidoService(pedidoRepo, usuarioRepo, showRepo, filaPedidosService)

	// Criando o DTO de pedido
	pedidoDTO := dtos.PedidoDTO{
		UserID: 1,
		ShowID: "123456",
	}

	// Configurando as expectativas dos mocks
	usuarioRepo.On("ObterPorID", uint(1)).Return(models.Usuario{ID: 1}, nil)
	showRepo.On("ObterPorID", "123456").Return(models.Show{ID: "123456"}, nil)
	pedidoRepo.On("Criar", mock.Anything).Return(nil)
	filaPedidosService.On("Publicar", mock.Anything).Return()

	// Chamando o método Criar
	err := pedidoService.Criar(pedidoDTO)

	// Verificando se não houve erro
	assert.Nil(t, err)

	// Verificando se os métodos dos mocks foram chamados corretamente
	usuarioRepo.AssertExpectations(t)
	showRepo.AssertExpectations(t)
	pedidoRepo.AssertExpectations(t)
	filaPedidosService.AssertExpectations(t)
}

func TestPedidoService_Criar_UsuarioNaoEncontrado(t *testing.T) {
	// Criando os mocks
	pedidoRepo := new(MockPedidoRepository)
	usuarioRepo := new(MockUsuarioRepository)
	showRepo := new(MockShowRepository)
	filaPedidosService := new(MockFilaPedidosService)

	// Criando o serviço
	pedidoService := services.NewPedidoService(pedidoRepo, usuarioRepo, showRepo, filaPedidosService)

	// Criando o DTO de pedido
	pedidoDTO := dtos.PedidoDTO{
		UserID: 1,
		ShowID: "123456",
	}

	// Configurando as expectativas dos mocks
	usuarioRepo.On("ObterPorID", uint(1)).Return(models.Usuario{}, errors.New("usuário não encontrado"))

	// Chamando o método Criar e verificando se o erro esperado ocorre
	err := pedidoService.Criar(pedidoDTO)

	// Verificando se houve erro
	assert.NotNil(t, err)
	assert.Equal(t, "usuário não encontrado", err.Error())

	// Verificando se os métodos dos mocks foram chamados corretamente
	usuarioRepo.AssertExpectations(t)
	showRepo.AssertExpectations(t)
	pedidoRepo.AssertExpectations(t)
	filaPedidosService.AssertExpectations(t)
}

func TestPedidoService_Criar_ShowNaoEncontrado(t *testing.T) {
	// Criando os mocks
	pedidoRepo := new(MockPedidoRepository)
	usuarioRepo := new(MockUsuarioRepository)
	showRepo := new(MockShowRepository)
	filaPedidosService := new(MockFilaPedidosService)

	// Criando o serviço
	pedidoService := services.NewPedidoService(pedidoRepo, usuarioRepo, showRepo, filaPedidosService)

	// Criando o DTO de pedido
	pedidoDTO := dtos.PedidoDTO{
		UserID: 1,
		ShowID: "123456",
	}

	// Configurando as expectativas dos mocks
	usuarioRepo.On("ObterPorID", uint(1)).Return(models.Usuario{ID: 1}, nil)
	showRepo.On("ObterPorID", "123456").Return(models.Show{}, errors.New("show não encontrado"))

	// Chamando o método Criar e verificando se o erro esperado ocorre
	err := pedidoService.Criar(pedidoDTO)

	// Verificando se houve erro
	assert.NotNil(t, err)
	assert.Equal(t, "show não encontrado", err.Error())

	// Verificando se os métodos dos mocks foram chamados corretamente
	usuarioRepo.AssertExpectations(t)
	showRepo.AssertExpectations(t)
	pedidoRepo.AssertExpectations(t)
	filaPedidosService.AssertExpectations(t)
}
