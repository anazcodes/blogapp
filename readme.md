
# Blog App
A Blog Post API in Go powered by Fiber framework.

## Technologies

- Go Programming Language

- Fiber Framework

- Swagger

- Docker

## Run on Docker
```bash
docker run -p 3000:3000 --name blogapp anazibinurasheed/blogapp:latest --port=3000 --cache-capacity=30
```

## Run Locally

Clone the project

```bash
  git clone https://github.com/anazcodes/blogapp.git
```

Get into project directory

```bash
  cd blogapp
```

Install dependencies

```bash
  go mod tidy
```

Start the server

```bash
  go run cmd/blogapp/main.go --port=3000 --cache-capacity=30
```

## Access Live Swagger UI
`https://blogapp-wkxi.onrender.com/swagger/index.html`

[Click Here To Redirect](https://blogapp-wkxi.onrender.com/swagger/index.html)

## Access Swagger UI Locally
`http://localhost:3000/swagger/index.html`

[Click Here To Redirect](http://localhost:3000/swagger/index.html)

<!-- 
