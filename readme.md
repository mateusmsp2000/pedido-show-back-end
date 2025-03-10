# Pedido-Show-App

### Prefácio

Esta aplicação permite criar pedidos de shows, gravando os pedidos em banco de dados sqlite e também disparando notificações por mensageria quando um novo pedido for efetivado com sucesso. Todo o projeto foi desenvolvido em golang.

> **Nota:** Algumas implementações, como comunicação com sistemas externos, foram abstraídas para fins de simplificação.


### Funcionalidades

- **Cadastro de pedidos de shows** (informando dados específicos).
- **Listagem de pedidos de shows**.
1. **Mensageria interna com Go Goroutines**:
    - Explicação de que o sistema usa goroutines e canais internos para gerenciar a comunicação e notificação de pedidos processados com sucesso.
    - Alteração no texto da funcionalidade para refletir o uso de goroutines para mensageria.

2. **Exemplo de código**:
    - Foi incluído um exemplo simples de como as goroutines e canais são usados para enviar notificações internas.


## Como executar a aplicação

> **Observação:** É necessário ter o Docker instalado na máquina para utilizar a abordagem abaixo.

```bash
docker-compose up --build
```

## Disponibilidade

A aplicação estará disponível em: [http://localhost:5001/](http://localhost:5001/)

## Frameworks utilizados

### **Gorm**
O **Gorm** é um ORM (Object Relational Mapping) utilizado para gerenciar a criação e manutenção de tabelas por meio de migrações, além de simplificar operações de leitura e escrita no banco de dados. O Gorm permite criar consultas avançadas, buscar registros e realizar operações CRUD de forma eficiente e integrada com a API, facilitando o desenvolvimento e a manutenção dos dados armazenados.

### **Gin**
O **Gin** é utilizado para possibilitar requisições REST na API. Ele é um framework web em Go que oferece alta performance, facilitando a criação de endpoints e manipulação de rotas de maneira rápida e eficiente.

## Utilizações

### **Interfaces**
Interfaces são amplamente utilizadas para garantir a separação de preocupações no projeto. Elas permitem que as dependências sejam passadas de forma flexível, tornando o código mais modular e testável.

### **Injeção de dependência**
A **injeção de dependência** é uma técnica utilizada para melhorar o desacoplamento das classes e objetos do sistema. No projeto, utilizamos a injeção de dependência para fornecer as dependências necessárias para cada componente de forma controlada, o que facilita a manutenção, os testes e a escalabilidade da aplicação.

### Testes Unitários

Os **testes unitários** são essenciais para garantir que as partes do sistema funcionem corretamente de maneira isolada. Eles permitem verificar se uma unidade de código (como uma função ou método) funciona como esperado.

## Tabelas criadas no banco de dados

| Tabela    | Descrição                                                    |
|-----------|--------------------------------------------------------------|
| `Usuario` | Armazena informações dos usuários do sistema.                |
| `Pedido`  | Armazena as informações dos pedidos realizados pelo usuário. |
| `Show`    | Armazena as informações do show atrelado ao pedido.          |

## Endpoints

| Tipo | Rota           | Descrição                                 |
|------|----------------|-------------------------------------------|
| GET  | `/api/pedidos` | Lista todos os pedidos de shows.          |
| POST | `/api/pedidos` | Cadastra pedido de show informando dados. | |

## Como testar os endpoints via Postman ou Insomnia
#### rota /api/pedidos

> **Observação:** Para abstração do serviço, favor utilizar esse show_id e id de usuário pré cadastrado!

``` json
{
  "show_id": "6b3ed050-11b0-42dc-b7b5-892aac8b97c7", 
  "user_id": 1                             
}
```

Após a criação do pedido com sucesso será emitido uma mensageria para que outros services como o de faturamento tenham conhecimento que o pedido foi criado com sucesso.

Pedido processado e realizado pagamento com sucesso! ShowID: %s, UserID: %d

## Validações
* usuário não encontrado -> ocorre quando não encontra o usuário através do id informado.
* show não encontrado -> ocorre quando não encontra o show através do id informeado.

## Objetivo

Criar uma aplicação simples, escalável e fácil de manter, que seja extensível para suportar novas demandas e alterações em regras de negócio.

## Possíveis melhorias para a versão 2.0

- Integrar a API com Prometheus e Grafana para monitoramento de métricas.
- Monitorar o banco de dados com Prometheus e Grafana.
- Utilizar RabbitMq para mensageria e comunicação entre micro serviços.
- Implementar outras funcionalidades como delete, put.

## Diagrama Entidade-Relacionamento (DER)
## Entidade: Pedido
| Campo    | Tipo  | Restrição              |
|----------|-------|------------------------|
| ID       | uint  | PRIMARY KEY            |
| ShowID   | string| FOREIGN KEY (ShowID)   |
| UserID   | uint  | FOREIGN KEY (UserID)   |
| Show     | Show  | Relacionamento 1:N     |
| Usuario  | Usuario | Relacionamento 1:N |

## Entidade: Show
| Campo  | Tipo   | Restrição     |
|--------|--------|---------------|
| ID     | string | PRIMARY KEY   |
| Name   | string |               |

## Entidade: Usuario
| Campo  | Tipo   | Restrição     |
|--------|--------|---------------|
| ID     | uint   | PRIMARY KEY   |
| Name   | string |               |

## Relacionamentos

1. **Pedido → Show**: Relacionamento de **muitos para um**. Um `Pedido` está associado a um único `Show` por meio de `ShowID`, mas um `Show` pode ter múltiplos `Pedidos`.
2. **Pedido → Usuario**: Relacionamento de **muitos para um**. Um `Pedido` está associado a um único `Usuario` por meio de `UserID`, mas um `Usuario` pode ter múltiplos `Pedidos`.



