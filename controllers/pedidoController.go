package controllers

import (
	"PedidoShow/dtos"
	"PedidoShow/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PedidoController struct {
	service services.IPedidoService
}

func NewPedidoController(service services.IPedidoService) *PedidoController {
	return &PedidoController{service: service}
}

func (c *PedidoController) Criar(ctx *gin.Context) {
	var pedido dtos.PedidoDTO
	if err := ctx.ShouldBindJSON(&pedido); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	err := c.service.Criar(pedido)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Pedido criado com sucesso"})
}

func (c *PedidoController) ObterTodos(ctx *gin.Context) {
	pedidos, err := c.service.ObterTodos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar pedidos"})
		return
	}

	ctx.JSON(http.StatusOK, pedidos)
}
