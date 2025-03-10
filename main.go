package main

import (
	"PedidoShow/controllers"
	"PedidoShow/fila"
	"PedidoShow/models"
	"PedidoShow/repositories"
	"PedidoShow/services"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func initDB() (*gorm.DB, error) {
	// Conectando ao banco de dados SQLite
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Rodando migrações automaticamente para garantir que as tabelas sejam criadas
	err = db.AutoMigrate(&models.Usuario{}, &models.Show{}, &models.Pedido{})
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
	pedidoRepo := repositories.NewPedidoRepository(db)
	userRepo := repositories.NewUsuarioRepository(db)
	showRepo := repositories.NewShowRepository(db)

	_ = userRepo.Remover(1)
	_ = userRepo.Criar(models.Usuario{ID: 1, Name: "João Palesmano"})
	_ = showRepo.Remover("6b3ed050-11b0-42dc-b7b5-892aac8b97c7")
	_ = showRepo.Criar(models.Show{ID: "6b3ed050-11b0-42dc-b7b5-892aac8b97c7", Name: "Gustavo Lima"})

	filaPedidos := fila.NewFilaPedidosService(100)
	go filaPedidos.Processar()

	// Criando o serviço com o repositório
	pedidoService := services.NewPedidoService(pedidoRepo, userRepo, showRepo, filaPedidos)

	// Criando o controlador com o serviço
	pedidoController := controllers.NewPedidoController(pedidoService)

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

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
