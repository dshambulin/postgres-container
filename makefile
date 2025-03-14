DB_DSN := "postgres://postgres:elpasodelgigante@localhost:5432/main?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	$(MIGRATE) create -ext sql -dir ./migrations $(NAME)

migrate:
	$(MIGRATE) up	

migrate-down:
	$(MIGRATE) down
	
run:
	go run cmd/app/main.go

gen-tasks:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks -o ./internal/web/tasks/api.gen.go openapi/openapi.yaml

gen-users:
	oapi-codegen -config openapi/.openapi -include-tags users -package users -o ./internal/web/users/api.gen.go openapi/openapi.yaml

	lint:
	golangci-lint	run	--out-format=colored-line-number

git:
	git add .
	git commit -m "$(commit)"
	git push
