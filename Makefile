build:
	@echo "Building"
	@go build -o bin/app app/main.go

run:
	@echo "Running"
	@go run app/main.go

test:
	@echo "Testing"
	@go test -p 1  ./... -cover