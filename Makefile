base:
	docker run --name=short-link-bd -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres
test:
	go test -v ./...
migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
