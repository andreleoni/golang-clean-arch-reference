# This repository still in progress

# Golang Clean Arch Reference

This repository is an experiment of how to use the clean arch concepts on a Golang project.

The purposal is a study case about a customer, product and an order application with domain and layers isolation.

# Links

[Customers contract](CONTRACTS.md)

[Development utils](DEVELOPMENT.md)

# Development

Run compose with: `docker compose up -d`

Run tests with: `go test ./...`

Generate mock with: `make mocks`

# Sobre o clean arch

* Camada de Interface do Usuário: esta camada é responsável por lidar com todas as interações do usuário com o sistema. Ela pode ser implementada usando pacotes Go como "http" ou "grpc" e deve ser independente do restante da aplicação. Esta camada também pode fornecer APIs externas que os clientes podem usar para interagir com o sistema.

* Camada de Casos de Uso: esta camada contém a lógica de negócios da aplicação. Ela é responsável por receber solicitações da camada de interface do usuário e coordenar ações entre as outras camadas. Esta camada também deve ser independente do restante da aplicação e pode ser testada sem depender de outras camadas.

* Camada de Repositórios: esta camada é responsável pelo acesso e armazenamento de dados. Ela deve ser independente das outras camadas e pode ser implementada usando pacotes Go como "sql" ou "nosql". Os repositórios fornecem uma interface para as camadas superiores usarem para acessar e modificar dados.

* Camada de Infraestrutura: esta camada é responsável por implementar as partes específicas da plataforma, como as camadas de interface do usuário, bancos de dados e outras bibliotecas de terceiros. Esta camada não deve depender das outras camadas, mas deve ser responsável por integrá-las.
