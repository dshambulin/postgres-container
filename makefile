DB_DSN := "postgres://postgres:your_new_password@localhost:5432/main?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	$(MIGRATE) create -ext sql -dir ./migrations $(NAME)

migrate:
	$(MIGRATE) up	

migrate-down:
	$(MIGRATE) down
	
run:
	go run cmd/app/main.go

gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks -o ./internal/web/tasks/api.gen.go openapi/openapi.yaml

lint:
	golangci-lint	run	--out-format=colored-line-number