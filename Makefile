.PHONY: dev
dev: 
	@go run cmd/rest/main.go

.PHONY: console
console: 
	@go run cmd/console/main.go