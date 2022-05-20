build:
	docker-compose build shortLinks
run:
	docker-compose up shortLinks
test:
	go test -v ./...
migrate:
	migrate -path ./sсhema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
