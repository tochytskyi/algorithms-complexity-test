.PHONY: build
build:
	docker build \
		--force-rm \
		--target local \
		-f Dockerfile \
		-t algorithms-complexity-golang:v1.0 \
		.

.PHONY: rebuild
rebuild:
	make down && make build && make up

.PHONY: rebuild
restart:
	make down && docker-compose up --build --force-recreate --no-deps

.PHONY: up
up:
	docker-compose up

.PHONY: down
down:
	docker-compose down

.PHONY: ps
ps:
	docker-compose ps -a

.PHONY: bash
bash:
	docker exec -it algorithms-complexity-test bash

.PHONY: test
test:
	docker exec algorithms-complexity-test go test ./...

.PHONY: copy-go-sum
copy-go-sum:
	docker cp algorithms-complexity-test:/go/src/app/go.sum ./go.sum
