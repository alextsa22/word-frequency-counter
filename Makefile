fmt:
	go fmt ./...

lint:
	golangci-lint run ./...

aim:
	./scripts/aim.sh assets/mobydick.txt

run:
	go run cmd/main.go assets/mobydick.txt
