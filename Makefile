build: fmt test
	mkdir -p bin && GO111MODULE=on go build -o ./bin/goback ./cmd/goback

run: build
	./bin/goback "$(command)"

test:
	GO111MODULE=on go test ./...

fmt:
	GO111MODULE=on go fmt ./...

lint:
	GO111MODULE=on golint ./...
