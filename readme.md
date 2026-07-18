# go-scaffold-gk

Scaffold em Go para criar estruturas comuns de aplicações com PostgreSQL. O projeto gera arquivos base para model, migration, request, response, seed, controller e um esqueleto de repository pattern.

O pacote de migrations continua disponível, mas o foco principal do projeto é acelerar a criação da estrutura inicial de módulos Go.

## Instalação

```bash
go get github.com/gabrielki09/go-scaffold-gk@latest
```

## Gerador de Scaffold

Uso local:

```bash
go run ./cmd/main.go -m users -id
```

Uso com UUID:

```bash
go run ./cmd/main.go -m users -uuid
```

Por padrão, o scaffold sempre cria a model em `models/`.

## Flags

```txt
-m string       nome do recurso/model; obrigatória
-uuid           usa UUID no model, response e migration
-id             usa ID int/BIGSERIAL
-R              cria request
-r              cria response/resource
-s              cria seed
-M              cria migration
-c              cria controller
-a              cria todos os arquivos opcionais, exceto -repo, -uuid e -id
-repo           cria estrutura de repository pattern
-path string    caminho raiz para estrutura repository pattern
-create-path    cria o caminho informado em -path caso ele não exista
```

## Exemplos

Criar somente model com ID inteiro:

```bash
go run ./cmd/main.go -m users -id
```

Criar model e migration com UUID:

```bash
go run ./cmd/main.go -m users -uuid -M
```

Criar model, request, response, seed, migration e controller:

```bash
go run ./cmd/main.go -m users -uuid -a
```

Criar estrutura com repository pattern:

```bash
go run ./cmd/main.go -m users -uuid -repo -path=internal/modules/user -create-path
```

## Estrutura Gerada

Com:

```bash
go run ./cmd/main.go -m users -uuid -repo -path=internal/modules/user -create-path
```

A estrutura esperada é:

```txt
models/
  users/
    users_model.go

internal/modules/user/
  routes/
    users_routes.go
  controller/
    users_controller.go
  service/
    users_service.go
  repository/
    users_repository.go
```

## Regras de Validação

- `-m` é obrigatório.
- É obrigatório informar `-uuid` ou `-id`.
- `-uuid` e `-id` não podem ser usados juntos.
- Quando `-repo` é usado, `-path` deve apontar para o diretório raiz do módulo.
- Quando `-create-path` não é usado, o diretório informado em `-path` precisa existir.

## Migrations

Além do scaffold, o projeto fornece um runner simples de migrations PostgreSQL usando `pgxpool`.

Importe o pacote:

```go
import migrator "github.com/gabrielki09/go-scaffold-gk/pkg/migration"
```

Exemplo:

```go
package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	migrator "github.com/gabrielki09/go-scaffold-gk/pkg/migration"
)

func main() {
	ctx := context.Background()

	db, err := pgxpool.New(ctx, "postgres://user:password@localhost:5432/app")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = migrator.Run(ctx, db, migrator.Options{
		Dir:     "database/migration",
		Command: migrator.CommandUp,
	})
	if err != nil {
		log.Fatal(err)
	}
}
```

Se `Dir` ficar vazio, o pacote usa `database/migration` a partir do diretório atual.

### Formato das Migrations

Cada migration precisa ter um arquivo `.up.sql` e um `.down.sql` com a mesma versão e o mesmo nome:

```txt
database/migration/
  20260711235000_create_users.up.sql
  20260711235000_create_users.down.sql
```

Formato esperado:

```txt
{version}_{name}.up.sql
{version}_{name}.down.sql
```

Exemplo:

```sql
-- 20260711235000_create_users.up.sql
CREATE TABLE IF NOT EXISTS users (
	id BIGSERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	deleted_at TIMESTAMPTZ NULL
);
```

```sql
-- 20260711235000_create_users.down.sql
DROP TABLE IF EXISTS users;
```

### Comandos de Migration

```go
migrator.CommandUp
migrator.CommandDown
migrator.CommandFresh
migrator.CommandStatus
```

- `up`: aplica migrations pendentes em ordem crescente.
- `down`: desfaz a última migration aplicada.
- `fresh`: desfaz todas as migrations aplicadas e depois roda `up`.
- `status`: imprime o status de cada migration como `applied` ou `pending`.

O pacote cria automaticamente a tabela `schema_migrations` quando necessário.

## Status Atual

Este projeto está em evolução para ser um scaffold mais completo. Alguns templates gerados ainda são esqueletos iniciais e devem ser ajustados conforme o padrão final da aplicação.
