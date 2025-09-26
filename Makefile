swag:
	@swag init --parseInternal --parseDependency

go: swag
	@go run main.go

dev: swag
	@air