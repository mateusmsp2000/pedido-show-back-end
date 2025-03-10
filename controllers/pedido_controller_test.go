package controllers_test

import (
	"PedidoShow/controllers"
	"PedidoShow/dtos"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock do IPedidoService
type MockPedidoService struct {
	mock.Mock
}

func (m *MockPedidoService) Criar(pedido dtos.PedidoDTO) error {
	args := m.Called(pedido)
	return args.Error(0)
}

func (m *MockPedidoService) ObterTodos() ([]dtos.PedidoDTO, error) {
	args := m.Called()
	return args.Get(0).([]dtos.PedidoDTO), args.Error(1)
}

// Teste do controlador CriarPedido
func TestCriarPedido(t *testing.T) {
	// Criando o mock do serviço
	mockService := new(MockPedidoService)

	// Criando o controlador com o serviço mockado
	controller := &controllers.PedidoController{service: mockService}

	// Simulando uma requisição HTTP
	router := gin.Default()
	router.POST("/pedidos", controller.Criar)

	// Definindo o payload da requisição
	pedidoDTO := dtos.PedidoDTO{
		UserID: 1,
		ShowID: "123456",
	}

	// Definindo o comportamento esperado do mock
	mockService.On("Criar", pedidoDTO).Return(nil)

	// Criando o corpo da requisição em JSON
	jsonPayload, _ := json.Marshal(pedidoDTO)

	// Criando a requisição HTTP com o corpo JSON correto
	req, _ := http.NewRequest(http.MethodPost, "/pedidos", bytes.NewReader(jsonPayload))

	// Gravando a resposta
	w := httptest.NewRecorder()

	// Enviando a requisição via Gin
	router.ServeHTTP(w, req)

	// Verificando a resposta HTTP
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, `{"message": "Pedido criado com sucesso"}`, w.Body.String())

	// Verificando se o método Criar foi chamado corretamente
	mockService.AssertExpectations(t)
}

// Teste do controlador ObterTodosPedidos
func TestObterTodosPedidos(t *testing.T) {
	// Criando o mock do serviço
	mockService := new(MockPedidoService)

	// Criando o controlador com o serviço mockado
	controller := &controllers.PedidoController{service: mockService}

	// Definindo os pedidos a serem retornados
	pedidosDTO := []dtos.PedidoDTO{
		{UserID: 1, ShowID: "123456"},
		{UserID: 2, ShowID: "789101"},
	}

	// Definindo o comportamento esperado do mock
	mockService.On("ObterTodos").Return(pedidosDTO, nil)

	// Simulando uma requisição HTTP
	router := gin.Default()
	router.GET("/pedidos", controller.ObterTodos)

	// Criando a requisição HTTP
	req, _ := http.NewRequest(http.MethodGet, "/pedidos", nil)

	// Gravando a resposta
	w := httptest.NewRecorder()

	// Enviando a requisição via Gin
	router.ServeHTTP(w, req)

	// Verificando a resposta HTTP
	assert.Equal(t, http.StatusOK, w.Code)

	// Verificando o conteúdo da resposta JSON
	expected := `[{"user_id":1,"show_id":"123456"},{"user_id":2,"show_id":"789101"}]`
	assert.JSONEq(t, expected, w.Body.String())

	// Verificando se o método ObterTodos foi chamado corretamente
	mockService.AssertExpectations(t)
}
