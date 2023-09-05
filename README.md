# Agendamento

## Descrição

Esse projeto e uma agenda que visa criar reservas com base em contratos:
  Contrato:
     - Deve possuir atendimento
     - Usuarios

  Reserva:
     - Deve ser gerada com base em um contrato
     - Deve ser possivel ter reserva sem contrato


## Arquitetura

A aplicação segue o padrão de arquitetura DDD (Domain-Driven Design), que organiza o código em camadas bem definidas, incluindo:

- **Camada de Domínio**: Aqui, estão contidas as regras de negócio, entidades, agregados, objetos de valor e serviços do domínio.

- **Camada de Aplicação**: Esta camada é responsável por orquestrar a lógica de negócios, implementar os casos de uso e interagir com a camada de infraestrutura.

- **Camada de Infraestrutura**: Esta camada abstrai os detalhes técnicos da aplicação, incluindo acesso a bancos de dados, serviços externos e comunicação com a camada de aplicação.

## Pré-requisitos

Para executar esta aplicação, você precisará das seguintes dependências:

- Go (versão 1.21.0)
- Banco de dados PostgreSQL


## Instalação

Siga estas instruções para instalar e configurar a aplicação:

1. Clone o repositório:

   ```bash
  git clone https://github.com/seuusuario/seuprojeto.git

  cd seuprojeto

  go mod download
   
  cp config.example.yaml config.yaml
   
  nano config.yaml

  go run main.go

# Uso

Para usar a aplicação, siga estas instruções.

# Testes

- Testes unitários:

   ```bash
   go test ./...

## Criar mock para a interface

```bash
  mockgen -source=internal/domain/services/interface/user_interface.go -destination=internaldomain/services/mocks/user_mock.go 