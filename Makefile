.PHONY: run, build, fmt, install
run:
	go run cmd/hitdir/main.go
build:
	go build -o bin/hitdir cmd/hitdir/main.go 
fmt:
	go fmt
install:
	go install