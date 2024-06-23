# Comando para levantar el servidor
run-server:
	go run ./cmd/server/main.go

run-client-read-all-files:
	go run ./cmd/client/main.go

# Comando para correr los tests
test:
	go test ./... -v

# Comando para ejecutar golangci-lint
validate:
	golangci-lint run

# Comando para ejecutar golangci-lint con --fix
validate-fix:
	golangci-lint run --fix

# Comando para limpiar archivos binarios
clean:
	go clean
	rm -f server client

# Comando para compilar el binario del servidor
build-server:
	go build -o server ./cmd/server

# Comando para compilar el binario del cliente
build-client:
	go build -o client ./cmd/client

.PHONY: run-server test lint clean build-server build-client
