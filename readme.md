# go-scaffold-gk

Ferramenta pessoal para migrations PostgreSQL usando `pgxpool`, com um gerador simples de arquivos base em Go.

A lib foi feita para o meu fluxo pessoal. Ela não pretende substituir ferramentas maduras como `golang-migrate` ou `goose`.

## Instalação

```bash
go get github.com/gabrielki09/go-scaffold-gk@latest
```

## Migrations

Importe o pacote:

```go
import migrator "github.com/gabrielki09/go-scaffold-gk/pkg/migration"
```

Exemplo de uso:

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

	if err := migrator.Run(ctx, db, migrator.Options{
		Dir:     "database/migration",
		Command: migrator.CommandUp,
	}); err != nil {
		log.Fatal(err)
	}
}
```

Se `Dir` ficar vazio, o pacote usa `database/migration` a partir do diretório atual.

### Padrão dos arquivos

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

### Comandos disponiveis

```go
migrator.CommandUp
migrator.CommandDown
migrator.CommandFresh
migrator.CommandStatus
```

Equivalentes aos comandos:

- `up`: aplica migrations pendentes em ordem crescente.
- `down`: desfaz a ultima migration aplicada.
- `fresh`: desfaz todas as migrations aplicadas e depois roda `up`.
- `status`: imprime o status de cada migration como `applied` ou `pending`.

O pacote cria automaticamente a tabela `schema_migrations` quando necessário.

## Gerador de arquivos

O gerador fica em `cmd/model` e cria arquivos base para model, migration, controller, request, resource e seed.

Uso local no repositório:

```bash
go run ./cmd/model -model users
```

Uso via modulo:

```bash
go run github.com/gabrielki09/go-scaffold-gk/cmd/model@latest -model users
```

A flag `-model` e obrigatória. Por padrão, o gerador sempre cria o arquivo de model em `models/`.

### Flags reais

```txt
-model string  nome da model; obrigatória
-uuid          usa UUID no model/resource e gen_random_uuid() na migration
-id            usa ID int/BIGSERIAL explicitamente
-S             habilita a opção SeparateByFolder no código
-R             cria request em requests/
-r             cria resource/response em resource/
-s             cria seed em seed/
-m             cria migration em migration/
-c             cria controller em controller/
-a             cria todos os arquivos opcionais, exceto escolha de id
```

Observações do comportamento atual:

- Não use `-uuid` e `-id` juntos; o comando retorna erro.
- Se nem `-uuid` nem `-id` forem informados, o comando retorna erro.
- A flag `-a` habilita `model`, `requests`, `resource`, `seed`, `migration` e `controller`; ela não habilita `uuid`, `id` nem `separate_by_folder`.
- Os diretórios são criados automaticamente quando não existem.

Exemplos:

```bash
# cria somente model em models/
go run ./cmd/model -model users

# cria model e migration
go run ./cmd/model -model users -m

# cria model, migration, request, resource, seed e controller usando UUID
go run ./cmd/model -model users -uuid -a
```
