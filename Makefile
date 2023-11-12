.PHONY: build
build:
	go build --o server cmd/server/main.go

.PHONY: t
t:
	go generate ./...
	tern migrate --config migrations/tern.conf --migrations migrations --database todoer_test
	go test --v ./... --coverpkg=./internal/... --coverprofile=coverage.out --test.short=true

.PHONY: test
test:
	go generate ./...
	tern migrate --config migrations/tern.conf --migrations migrations --database todoer_test
	go test --v ./... --coverpkg=./internal/... --coverprofile=coverage.out

.PHONY: c
c:
	go tool cover --func coverage.out

.PHONY: gen
gen:
	swag fmt
	go generate ./...

.PHONY: migrate
migrate:
	tern migrate -c migrations/tern.conf -m migrations

.PHONY: tm
tm:
	tern migrate -c migrations/tern.conf -m migrations -database todoer_test

.PHONY: dock
dock:
	docker build . --file=infra/server.dockerfile --tag="vladmarlo/todoer_backend:latest"
	docker build . --file=infra/migrator.dockerfile --tag="vladmarlo/todoer_migrator:latest"

.PHONY: dock/push
dock/push: dock
	docker push vladmarlo/todoer_backend:latest
	docker push vladmarlo/todoer_migrator:latest

.PHONY: dock/run
dock/run:
	docker-compose up --d

.PHONY: lines
lines:
	git ls-files | xargs wc -l

.DEFAULT_GOAL := build
