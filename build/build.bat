windres -o resources.syso build/res.rc
go build -o http-win-notice.exe -ldflags="-extldflags -static -H windowsgui -s -w"
del resources.syso