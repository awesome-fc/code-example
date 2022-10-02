release:
	cd ./code && GO111MODULE=on go mod tidy && GO111MODULE=on GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o target/main main.go