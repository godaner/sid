build:buildsid
buildsid:
	-rm sid.amd
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o sid.amd ./cmd/boot.go
	-upx -9 sid.amd
