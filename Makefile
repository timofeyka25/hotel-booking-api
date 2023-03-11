.PHONY: start
start:
	go run cmd/main.go

.PHONY: gen_docs
gen_docs:
	swag init -g ./cmd/main.go -o ./docs --parseDependency --parseInternal --quiet