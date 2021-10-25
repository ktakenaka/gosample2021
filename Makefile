up:
	docker-compose up -d app database

down:
	docker-compose down

mod:
	docker-compose exec app go mod tidy
	docker-compose exec app go mod vendor

mysql:
	docker-compose exec database mysql -u gosample -p gosample_development

ent-init-%:
	$(eval ARG:= $*)
	go run entgo.io/ent/cmd/ent@latest init $(ARG)

ent-gen:
	docker-compose exec app go generate ./ent

migrate:
	docker-compose exec app go run cmd/migrate/main.go -c "environment/development/config.yml"
	docker-compose exec app go run cmd/migrate/main.go -c "environment/test/config.yml"

srv:
	docker-compose exec app go run cmd/srv/main.go

fmt:
	@go fmt ./...
