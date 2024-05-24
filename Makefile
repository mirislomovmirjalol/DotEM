run:
	@go build main.go
	@./main

build:
	go build -o bin/main main.go