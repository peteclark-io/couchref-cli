build:
	go test -v -race ./...
	cd ./bin/couchref-admin && rice embed-go
	go build ./bin/couchref
	go build ./bin/scores
	go build ./bin/couchref-admin
	go install ./bin/couchref
	go install ./bin/scores
	go install ./bin/couchref-admin
