build:
	go build -o bin/app cmd/main.go

run:
	go run cmd/main.go

test:
	go test ./... -v

docker-build:
	docker build -t go-gin-layout .

