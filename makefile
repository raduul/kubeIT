test_count:
	go test -count=80 ./...

test_cover:
	go test ./... -cover

test:
	go test ./...