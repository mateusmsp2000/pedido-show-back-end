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
)

func initDB() (*gorm.DB, error) {
	// Conectando ao banco de dados SQLite
	db, err := gorm.Open(sqlite.Open("./infrastructure/db.sqlite"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Rodando migrações automaticamente para garantir que as tabelas sejam criadas
	err = db.AutoMigrate(&entities.Usuario{}, &entities.Show{}, &entities.Pedido{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	// Inicializando o banco de dados
	db, err := initDB()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Criando o repositório com GORM
	pedidoRepo := repositories2.NewPedidoRepository(db)
	userRepo := repositories2.NewUsuarioRepository(db)
	showRepo := repositories2.NewShowRepository(db)

	_ = userRepo.Remover(1)
	_ = userRepo.Criar(entities.Usuario{ID: 1, Name: "João Palesmano"})
	_ = showRepo.Remover("6b3ed050-11b0-42dc-b7b5-892aac8b97c7")
	_ = showRepo.Criar(entities.Show{ID: "6b3ed050-11b0-42dc-b7b5-892aac8b97c7", Name: "Gustavo Lima"})

	filaPedidos := application.NewFilaPedidosService(100)
	go filaPedidos.Processar()

	// Criando o serviço com o repositório
	pedidoService := application.NewPedidoService(pedidoRepo, userRepo, showRepo, filaPedidos)

	// Criando o controlador com o serviço
	pedidoController := api.NewPedidoController(pedidoService)

	// Inicializando o Gin e criando as rotas
	r := gin.Default()

	// Definindo rotas para a API
	r.POST("/pedidos", pedidoController.Criar)     // Criar pedido
	r.GET("/pedidos", pedidoController.ObterTodos) // Listar pedidos

	// Iniciando o servidor na porta 8080
	if err := r.Run(":5001"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
