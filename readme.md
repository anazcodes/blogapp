
# Blog Post CRUD 
A Blog Post CRUD API in Go powered by Fiber framework.

## Technologies

- Go Programming Language

- Fiber Framework

- Swagger

## Run on Docker
```bash
docker run -p 3000:3000 --name blogapp anazibinurasheed/blog-crud-api:latest --port=3000 --cache-capacity=30
```

## Run Locally

Clone the project

```bash
  git clone https://github.com/anazcodes/blog-crud-api.git
```

Get into project directory

```bash
  cd blog-crud-api
```

Install dependencies

```bash
  go mod tidy
```

Start the server

```bash
  go run cmd/blogapp/main.go --port=3000 --cache-capacity=30
```

## Access Swagger UI
`http://localhost:3000/swagger/index.html`
<!-- 