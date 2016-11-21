build:
	go test -v -race ./...
	go build ./bin/couchref
	go build ./bin/scores
