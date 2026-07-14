# Climate by ZIP Code Exercise in GoLang(Exercício MBA)

Aplicação Go que consulta a temperatura e clima de uma localidade através do CEP.

## Pré-requisitos

### 1. Criar conta na WeatherAPI

- Acesse [https://www.weatherapi.com/](https://www.weatherapi.com/)
- Crie uma conta gratuita
- Gere uma chave de API

### 2. Configurar chave de API

1. Crie o arquivo `.env` em `cmd/.env`
2. Insira sua chave no formato:

```env
KEY_TEMP=sua_chave_aqui
```


## Como executar

### Localmente

```bash
cd cmd
go run main.go
```

## Via Docker

```bash
docker build -f DockerFile -t climate-app .
docker run -p 8080:8080 climate-app
```

## Testes

Na raiz do projeto, execute:

```bash
go test ./...
```

## Endpoint

| Ambiente | URL |
|----------|-----|
| Local | `http://localhost:8080/climate/{cep}` |


**Exemplo:**

```
GET http://localhost:8080/climate/01310100
```

