release:
	cd ./code && GO111MODULE=on go mod tidy && GO111MODULE=on GOOS=linux CGO_ENABLED=1 GOARCH=amd64 CGO_LDFLAGS="-static" go build -o target/main main.go