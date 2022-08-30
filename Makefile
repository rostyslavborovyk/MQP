build-mqp:
	go build -o bin/mqp cmd/mqp/main.go
build-ui:
	go build -o bin/ui-server cmd/ui/main.go
build:
	build-mqp build-ui
