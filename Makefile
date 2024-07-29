
run-client:
	go run cmd/client/main.go
	
run-server:
	go run cmd/server/main.go

build-client:
	go build -o bin/client cmd/client/main.go

build-server:
	go build -o bin/server cmd/server/main.go