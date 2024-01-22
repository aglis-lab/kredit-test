migration:
	go run cmd/migration/main.go

seed:
	go run cmd/seed/main.go

run.transaction:
	go build -o tmp/transaction worker/transaction/main.go
	./tmp/transaction
