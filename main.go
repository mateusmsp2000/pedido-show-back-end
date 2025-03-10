package main

import (
	"PedidoShow/api"
	"PedidoShow/application"
	"PedidoShow/domain/entities"
	repositories2 "PedidoShow/domain/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

func initDB() (*gorm.DB, error) {

	dbPath := "./infrastructure/pedido.db"

	if _, err := os.Stat("./infrastructure"); os.IsNotExist(err) {
		err := os.MkdirAll("./infrastructure", os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entities.Usuario{}, &entities.Show{}, &entities.Pedido{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	pedidoRepo := repositories2.NewPedidoRepository(db)
	userRepo := repositories2.NewUsuarioRepository(db)
	showRepo := repositories2.NewShowRepository(db)

	_ = userRepo.Remover(1)
	_ = userRepo.Criar(entities.Usuario{ID: 1, Name: "João Palesmano"})
	_ = showRepo.Remover("6b3ed050-11b0-42dc-b7b5-892aac8b97c7")
	_ = showRepo.Criar(entities.Show{ID: "6b3ed050-11b0-42dc-b7b5-892aac8b97c7", Name: "Gustavo Lima"})

	filaPedidos := application.NewFilaPedidosService(100)
	go filaPedidos.Processar()

	pedidoService := application.NewPedidoService(pedidoRepo, userRepo, showRepo, filaPedidos)

	pedidoController := api.NewPedidoController(pedidoService)
	r := gin.Default()

	r.POST("/api/pedidos", pedidoController.Criar)
	r.GET("/api/pedidos", pedidoController.ObterTodos)

	if err := r.Run(":5001"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
