.PHONY: build
build:
	docker build --force-rm --target local -f Dockerfile -t algorithms-complexity-golang:v1.0 .

.PHONY: rebuild
rebuild:
	make build && make up

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
