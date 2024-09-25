APP_NAME = iot-go
GO_FILES = ./cmd/server/main.go

build:
	go build -o $(APP_NAME) $(GO_FILES)

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(APP_NAME) $(GO_FILES)

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(APP_NAME).exe $(GO_FILES)

clean:
	rm -f $(APP_NAME) $(APP_NAME).exe