run:
	go run cmd/main.go

docs:
	swag init -g cmd/main.go -o docs

mocks:
	mockery --all --recursive

test:
	go test -v ./...

db:
	goose up
