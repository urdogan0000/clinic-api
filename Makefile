build:
	go build -ldflags "-s -w" -o service

buildw:
	go build -ldflags "-s -w" -o service.exe
	./service.exe

run:
	swag init
	go run main.go