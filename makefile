run:
	go run cmd/server/main.go

run.watch:
	air

tidy:
	go mod tidy -v

test:
	go test  ./...

test.cov:
	go test -v -race -covermode=atomic -coverprofile=coverage.out ./...

gosec:
	gosec ./...

doc:
	godoc

depend.update:
	go get -u

depend.patch:
	go get -u=patch
