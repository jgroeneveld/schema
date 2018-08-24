default: test

test:
	go test -v ./...

nice:
	golint ./... && go vet ./... && gofmt -s -w .
