env GOOS=linux GOARCH=amd64 go build -o scanstatus-linux
env GOOS=darwin GOARCH=amd64 go build -o scanstatus-macos
env GOOS=windows GOARCH=amd64 go build -o scanstatus-windows.exe