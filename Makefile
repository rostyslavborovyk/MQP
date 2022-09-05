build:
	env GOOS=windows GOARCH=386 go build -o bin/mqp-windows-386.exe cmd/mqp/main.go
	env GOOS=windows GOARCH=amd64 go build -o bin/mqp-windows-amd64.exe cmd/mqp/main.go
	env GOOS=darwin GOARCH=arm go build -o bin/mqp-darwin-arm cmd/mqp/main.go