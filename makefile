run:
	go run cmd/server/main.go

run.watch:
	air

tidy:
	go mod tidy -v

test:
	go test  ./...

gosec:
	gosec ./...

doc:
	godoc

depend.update:
	go get -u

depend.patch:
	go get -u=patch
