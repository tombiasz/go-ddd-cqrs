run:
	go run ./src/main.go

build:
	go build -o go-contacts ./src/main.go

test:
	go test ./...
