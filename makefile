run:
	go run "./cmd/hydralb"

build: 
	go build -o hydralb "./cmd/hydralb"

test: 
	go test -v ./...

fmt: 
	go fmt ./...

vet:
	go vet ./...

clean:
	rm -f hydralb