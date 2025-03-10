package dtos

type PedidoDTO struct {
	ShowID string `json:"show_id" binding:"required"`
	UserID uint   `json:"user_id" binding:"required"`
}
