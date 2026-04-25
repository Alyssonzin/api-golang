# API simples em Go + PostgreSQL (passo a passo)

Este guia cria um ambiente local para você aprender Go construindo uma API simples conectada ao PostgreSQL.

## 1) Pré-requisitos

- Windows 10/11
- Terminal (PowerShell)
- Editor (VS Code recomendado)

## 2) Instalar o Go

### Opção A (recomendada no Windows): winget

```powershell
winget install GoLang.Go
```

### Opção B: instalador oficial

- Baixe em: https://go.dev/dl/
- Execute o instalador e aceite as opções padrão

### Validar instalação

```powershell
go version
```

Você deve ver algo como `go version go1.xx.x windows/amd64`.

---

## 3) Instalar o PostgreSQL

### Opção A: winget

```powershell
winget install PostgreSQL.PostgreSQL
```

### Opção B: instalador oficial

- Baixe em: https://www.postgresql.org/download/windows/
- Durante a instalação:
  - Defina uma senha para o usuário `postgres`
  - Mantenha a porta padrão `5432`

### Validar instalação

Abra o terminal SQL (psql) e teste:

```powershell
psql --version
```

Se o comando não for reconhecido, adicione o diretório `bin` do PostgreSQL ao PATH.

---

## 4) Criar projeto Go

Na pasta do projeto:

```powershell
go mod init api-go
```

Isso cria o arquivo `go.mod`.

---

## 5) Instalar dependências da API

Para uma API simples com conexão ao Postgres, estas libs já são suficientes:

```powershell
go get github.com/jackc/pgx/v5
go get github.com/joho/godotenv
```

Dependências instaladas:

- `pgx/v5`: driver e client PostgreSQL para Go
- `godotenv`: carrega variáveis de ambiente de um arquivo `.env`

---

## 6) Configurar banco local

Entre no Postgres como usuário postgres:

```powershell
psql -U postgres -h localhost
```

Crie banco e usuário para o projeto:

```sql
CREATE DATABASE api_go_db;
CREATE USER api_go_user WITH PASSWORD 'api_go_pass';
GRANT ALL PRIVILEGES ON DATABASE api_go_db TO api_go_user;
```

Saia com:

```sql
\q
```

---

## 7) Configurar variáveis de ambiente

Crie um arquivo `.env` na raiz do projeto com:

```env
APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=api_go_user
DB_PASSWORD=api_go_pass
DB_NAME=api_go_db
DB_SSLMODE=disable
```

String de conexão equivalente:

```text
postgres://api_go_user:api_go_pass@localhost:5432/api_go_db?sslmode=disable
```

---

## 8) Estrutura inicial sugerida

```text
api-go/
  cmd/
    api/
      main.go
  internal/
    db/
      postgres.go
  .env
  go.mod
  go.sum
```

---

## 9) Criar uma tabela de teste

Conecte no banco criado:

```powershell
psql -U api_go_user -h localhost -d api_go_db
```

Execute:

```sql
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  email TEXT UNIQUE NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);
```

---

## 10) Comandos úteis no dia a dia

Baixar dependências do módulo:

```powershell
go mod tidy
```

Executar API (quando tiver `main.go`):

```powershell
go run ./cmd/api
```

Compilar:

```powershell
go build ./...
```

---

## 11) Problemas comuns

1. `psql` não encontrado
- Adicione o caminho do PostgreSQL ao PATH, algo como:
  - `C:\Program Files\PostgreSQL\<versao>\bin`

2. Erro de autenticação no banco
- Verifique usuário/senha em `.env`
- Confirme se o banco foi criado e se o usuário recebeu privilégios

3. Porta 5432 ocupada
- Descubra o processo usando a porta e altere a porta do PostgreSQL ou encerre o processo conflitante

4. Erro de SSL local
- Use `DB_SSLMODE=disable` em ambiente local

---

## 12) Próximo passo recomendado

Depois desse setup, você já pode implementar:

- Endpoint `GET /health` para testar API no ar
- Conexão ao banco na inicialização da aplicação
- CRUD básico de `users`

Se quiser, no próximo passo eu já posso gerar os arquivos iniciais da API com conexão ao Postgres e o endpoint de health check.
