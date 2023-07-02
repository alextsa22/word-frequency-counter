fmt:
	go fmt ./...

lint:
	golangci-lint run ./...

aim:
	./scripts/aim.sh assets/mobydick.txt
