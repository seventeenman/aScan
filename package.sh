CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./release/aScan_linux64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ./release/aScan_win64.exe ./main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ./release/aScan_darwin ./main.go
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o ./release/aScan_linux32
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o ./release/aScan_win32.exe ./main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags="-s -w" -o ./release/aScan_linuxarm
CGO_ENABLED=0 GOOS=windows GOARCH=arm go build -ldflags="-s -w" -o ./release/aScan_winarm.exe ./main.go
