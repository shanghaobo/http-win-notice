windres -o resources.syso build/res.rc
go env -w CGO_ENABLED=1 GOOS=windows GOARCH=amd64
go build -o dist/http-win-notice.exe -ldflags="-extldflags -static -H windowsgui -s -w"
del resources.syso

go env -w CGO_ENABLED=0 GOOS=linux GOARCH=amd64
cd forward/server
go build -o "../../dist/forward/forward-server"

go env -w CGO_ENABLED=0 GOOS=windows GOARCH=amd64
go build -o "../../dist/forward/forward-server.exe"

go env -w CGO_ENABLED=1