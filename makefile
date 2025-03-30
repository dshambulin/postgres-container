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
	oapi-codegen	-generate	"types,server,spec"	-package users	-include-tags users	-o ./internal/web/users/api.gen.go	openapi/openapi.yaml

git:
	git add .
	git commit -m "$(commit)"
	git push

lint:
	golangci-lint run ./...
