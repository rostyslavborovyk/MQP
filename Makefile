build:
	env GOOS=windows GOARCH=386 go build -o bin/mqp-windows-386.exe cmd/mqp/main.go
	env GOOS=windows GOARCH=amd64 go build -o bin/mqp-windows-amd64.exe cmd/mqp/main.go
	env GOOS=darwin GOARCH=amd64 go build -o bin/mqp-darwin-amd64 cmd/mqp/main.go
	env GOOS=linux GOARCH=amd64 go build -o bin/mqp-linux-amd64 cmd/mqp/main.go
	env GOOS=linux GOARCH=386 go build -o bin/mqp-linux-386 cmd/mqp/main.go
