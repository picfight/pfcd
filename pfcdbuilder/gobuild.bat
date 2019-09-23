cd D:\PICFIGHT\src\github.com\picfight\pfcwallet
set GO111MODULE=on
go build
go clean -testcache
go test -v ./...