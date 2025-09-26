swag:
	@swag init --parseInternal --parseDependency

go: swag
	@go run main.go

dev: swag
	@air

migrate-up:
	@go run main.go migrate up

migrate-down:
	@go run main.go migrate down