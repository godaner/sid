build:buildsid
buildsid:
	-rm sid.amd
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o sid.amd ./cmd/boot.go
	-upx -9 sid.amd
	-rm sid.exe
	CGO_ENABLED=0 GOOS=windows go build -o sid.exe ./cmd/boot.go
	-upx -9 sid.exe
	-rm sid.arm
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o sid.arm ./cmd/boot.go
	-upx -9 sid.arm
