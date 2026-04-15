# go-k8s-service

Minimal Go monorepo with two microservices for learning Kubernetes deployment patterns.

## Services

| Service | Port | Endpoints |
|---------|------|-----------|
| order-service | 8080 | GET /health, GET /orders, POST /orders, GET /orders/{id} |
| payment-service | 8081 | GET /health, POST /payments, GET /payments/{id} |

## Run locally

```bash
# single service
cd order-service && go run main.go

# both
docker-compose up
```

## Test

```bash
cd order-service && go test ./...
cd payment-service && go test ./...
```

Import `postman/go-k8s-service.postman_collection.json` into Postman for manual testing.

## Deploy to Kubernetes

```bash
kubectl apply -f k8s/
```
