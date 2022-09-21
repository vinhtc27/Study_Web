run:
	go run cmd/main/*.go

build:
	go build cmd/main/main.go

clean:
	make clean-dist
	rm -rf vendor