test:
	go test -count=80 ./...

test log:
	go test -v ./...

test cover:
	go test ./... -cover