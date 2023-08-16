
# Padawan Challenge Repository

This repository is about my implementation of the [Padawan Challenge Vacancy](https://github.com/genesisbankly/desafio-padawan-go)



### Requirements

Make sure you have the following tools installed on your system:

- Go (https://go.dev/doc/install)  
- Docker (https://docs.docker.com/engine/install/)


### Configuration and Execution

1. Clone this repository to your system:

```bash
git clone https://github.com/richardnfag/desafio-padawan-go.git
```

2. Navigate to the project directory:

```bash
cd desafio-padawan-go
```

3. Execute the following command to create and start the Docker containers:

```
docker-compose up -d
```

This will create and start the necessary containers for the project, including a container for the database.

4. After the database and application containers are running, execute the database migrations:

```bash
docker-compose run --rm web make migrate
```

5. After the containers are initialized and the migrations are applied, the API will be accessible at localhost:8000.  
Access: http://localhost:8000/exchange/10/BRL/USD/4.50 The response should follow the following format:
   ```json
   {
     "valorConvertido": 45,
     "simboloMoeda": "$"
   }

6. To run the tests, use the following command:

```bash
docker-compose run --rm web make test
```

## Challenge Resolution

The following conversions were requested:

- Convertion from Brazilian Real to US Dollar:
http://localhost:8000/exchange/10/BRL/USD/0.20

- Convertion from US Dollar to Brazilian Real:
http://localhost:8000/exchange/10/USD/BRL/5.00

- Convertion from Brazilian Real to Euro:
http://localhost:8000/exchange/10/BRL/EUR/0.18

- Convertion from Euro to Brazilian Real:
http://localhost:8000/exchange/10/EUR/BRL/5.40

- Convertion from Bitcoin to US Dollar:
http://localhost:8000/exchange/10/BTC/USD/30000.00

- Convertion from Bitcoin to Brazilian Real:
http://localhost:8000/exchange/5/BTC/BRL/145000.00


Another request was to retain the history for future queries. To view the history, use the following command:

```bash

QUERY="SELECT c.amount, f.code AS from_currency, t.code AS to_currency, c.rate, c.result
       FROM conversions c
       INNER JOIN currencies f ON c.from_currency_id = f.id
       INNER JOIN currencies t ON c.to_currency_id = t.id"

docker-compose run --rm db mysql -u exchange -h db -pexchange -Dexchange_db -e "$QUERY"

```


## Design and Architecture Decisions

### Technologies Used

[**Go**](https://go.dev/):  The chosen programming language for the project.

[**GORM**](https://gorm.io/index.html): An Object-Relational Mapping (ORM) framework for Go, used to simplify interaction with the database.

[**ATLAS**](https://atlasgo.io/): A migration manager that helps control changes in the database schema.

[**Docker**](https://www.docker.com/): A platform that enables packaging the application, its dependencies, and configurations in an isolated environment.

### Ports and Adapters (Hexagonal Architecture)

For this project, the Ports and Adapters architecture, also known as Hexagonal Architecture, was used. It aims to clearly separate core application concerns from external concerns, such as user interfaces, external services, and databases. This approach promotes a more cohesive organization and facilitates independent evolution of different parts of the system.

#### Testability and Evolution


The Ports and Adapters architecture promotes testability and continuous evolution of the system, allowing core parts to be tested independently and adapters to be replaced with new implementations without affecting the core of the application.


#### Flexibility and Maintenance

The clear separation between the core application and the adapters provides flexibility in choosing technologies and facilitates system maintenance over time. Changes in external components can be accommodated without affecting the internal functioning of the application.

### Directory Structure
```sh
.
├── Dockerfile # Dockerfile is the configuration file of the application
├── Makefile # Makefile is the automation script of the application
├── README.md 
├── atlas.hcl # Atlas is the configuration file of the application
├── bin # Bin is the output directory of the application
├── cmd # Cmd is the entrypoint of the application
│   └── main.go
├── coverage # Coverage is the directory of the coverage report
├── docker-compose.yml # Docker-compose is the configuration file of the application
├── go.mod # Go.mod is the dependency manager of the application
├── go.sum 
├── internal # Houses the core application logic, including use cases, domain models, and business rules.
│   ├── adapters # Adapters are the implementation of the ports
│   │   ├── database
│   │   │   ├── gorm_conversion_repository.go
│   │   │   ├── gorm_conversion_repository_test.go
│   │   │   ├── gorm_currency_repository.go
│   │   │   └── gorm_currency_repository_test.go
│   │   └── http
│   │       ├── http_handler.go
│   │       └── http_handler_test.go
│   ├── entities # Entities are the representation of the domain
│   │   ├── conversion_entity.go
│   │   └── currency_entity.go
│   ├── ports # Ports are the interfaces that define the contract of the adapters
│   │   ├── conversion_repository.go
│   │   └── currency_repository.go
│   └── services # Services are the use cases of the application
│       ├── conversion_service.go
│       └── conversion_service_test.go
└── migrations # Migrations are the database migrations
    ├── 20230816135044.sql
    ├── 20230816135153_add_currencies.sql
    └── atlas.sum # Atlas.sum is the summary of the migrations


# Note: The test files are organized adjacent to their respective targets with the _test suffix for improved maintainability.
```

