fmt:
	go fmt ./...

lint:
	golangci-lint run ./...

aim:
	./scripts/aim.sh assets/mobydick.txt

run:
	go run cmd/main.go assets/mobydick.txt

# test on preferred platform
test-build:
	docker build -t test-word-counter -f test/platform/Dockerfile .

test-run:
	docker run --rm test-word-counter

test-clear:
	docker rmi test-word-counter
